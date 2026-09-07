package onvif

import (
	"net/http"

	"github.com/icholy/digest"
)

// DigestClient represents an HTTP client used for making requests authenticated
// with http digest authentication.
type DigestClient struct {
	client *http.Client
}

// NewDigestClient returns a DigestClient that uses a copy of the provided http.Client configured for digest authentication.
func NewDigestClient(stdClient *http.Client, username string, password string) *DigestClient {
	if stdClient == nil {
		stdClient = http.DefaultClient
	}
	c := *stdClient
	c.Transport = &digest.Transport{
		Username:  username,
		Password:  password,
		Transport: stdClient.Transport,
	}
	return &DigestClient{client: &c}
}

func (dc *DigestClient) Do(httpMethod string, endpoint string, soap string) (*http.Response, error) {
	req, err := createHttpRequest(httpMethod, endpoint, soap)
	if err != nil {
		return nil, err
	}
	return dc.client.Do(req)
}
