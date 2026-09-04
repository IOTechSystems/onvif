package onvif

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// digestParam matches key=value pairs in a WWW-Authenticate challenge, value quoted or bare.
var digestParam = regexp.MustCompile(`(\w+)=(?:"([^"]*)"|([^,\s]*))`)

// DigestClient represents an HTTP client used for making requests authenticated
// with http digest authentication.
type DigestClient struct {
	client     *http.Client
	username   string
	password   string
	snonce     string
	realm      string
	qop        string
	opaque     string
	nonceCount uint32
}

// NewDigestClient returns a DigestClient that wraps a given standard library http Client with the given username and password
func NewDigestClient(stdClient *http.Client, username string, password string) *DigestClient {
	return &DigestClient{
		client:   stdClient,
		username: username,
		password: password,
	}
}

func (dc *DigestClient) Do(httpMethod string, endpoint string, soap string) (*http.Response, error) {
	req, err := createHttpRequest(httpMethod, endpoint, soap)
	if err != nil {
		return nil, err
	}
	if dc.snonce != "" {
		digestAuth, err := dc.getDigestAuth(req.Method, req.URL.RequestURI())
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", digestAuth)
	}

	// Attempt the request using the underlying client
	resp, err := dc.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusUnauthorized {
		return resp, nil
	}

	dc.getDigestParts(resp)
	// We will need to return the response from another request, so defer a close on this one
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Failed to close io reader, %s", err.Error())
		}
	}(resp.Body)

	req, err = createHttpRequest(httpMethod, endpoint, soap)
	if err != nil {
		return nil, err
	}
	digestAuth, err := dc.getDigestAuth(req.Method, req.URL.RequestURI())
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", digestAuth)

	authedResp, err := dc.client.Do(req)
	if err != nil {
		return nil, err
	}
	return authedResp, nil
}

func (dc *DigestClient) getDigestParts(resp *http.Response) {
	result := map[string]string{}
	authHeader := strings.TrimPrefix(resp.Header.Get("WWW-Authenticate"), "Digest ")
	for _, m := range digestParam.FindAllStringSubmatch(authHeader, -1) {
		result[m[1]] = m[2] + m[3] // exactly one of the two groups matched
	}
	dc.snonce = result["nonce"]
	dc.realm = result["realm"]
	dc.opaque = result["opaque"]
	dc.qop = result["qop"]
	// ponytail: only qop=auth is implemented; pick it when the server offers "auth,auth-int"
	for _, q := range strings.Split(dc.qop, ",") {
		if strings.TrimSpace(q) == "auth" {
			dc.qop = "auth"
			break
		}
	}
	dc.nonceCount = 0
}

func getMD5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func getCnonce() (string, error) {
	b := make([]byte, 8)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return "", fmt.Errorf("generate random number failed: %w", err)
	}
	return fmt.Sprintf("%x", b)[:16], nil
}

func (dc *DigestClient) getDigestAuth(method string, uri string) (string, error) {
	ha1 := getMD5(dc.username + ":" + dc.realm + ":" + dc.password)
	ha2 := getMD5(method + ":" + uri)
	cnonce, err := getCnonce()
	if err != nil {
		return "", fmt.Errorf("get DigestAuth failed: %w", err)
	}
	dc.nonceCount++
	nc := fmt.Sprintf("%08x", dc.nonceCount)
	response := getMD5(fmt.Sprintf("%s:%s:%s:%s:%s:%s", ha1, dc.snonce, nc, cnonce, dc.qop, ha2))
	authorization := fmt.Sprintf(`Digest username="%s", realm="%s", nonce="%s", uri="%s", cnonce="%s", nc=%s, qop=%s, response="%s"`,
		dc.username, dc.realm, dc.snonce, uri, cnonce, nc, dc.qop, response)
	if dc.opaque != "" {
		authorization += fmt.Sprintf(`, opaque="%s"`, dc.opaque)
	}
	return authorization, nil
}
