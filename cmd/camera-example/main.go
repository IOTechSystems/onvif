package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nbio/xml"

	"github.com/secure-passage/onvif"
	"github.com/secure-passage/onvif/gosoap"
	"github.com/secure-passage/onvif/networking"
)

func main() {
	RunApi()
}

func RunApi() {
	router := gin.Default()

	router.POST("/:service/:function", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		// c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

		serviceName := c.Param("service")
		functionName := c.Param("function")
		username := c.GetHeader("username")
		pass := c.GetHeader("password")
		xaddr := c.GetHeader("xaddr")
		acceptedData, err := c.GetRawData()
		if err != nil {
			fmt.Println(err)
		}
		dev, _ := onvif.NewDevice(onvif.DeviceParams{
			Xaddr:    xaddr,
			Username: username,
			Password: pass,
		})
		message, err := CallOnvifFunction(dev, serviceName, functionName, string(acceptedData))
		if err != nil {
			c.XML(http.StatusBadRequest, err.Error())
		} else {
			c.String(http.StatusOK, message)
		}
	})

	_ = router.Run()
}

func CallOnvifFunction(dev *onvif.Device, serviceName, functionName, acceptedData string) (string, error) {
	function, err := functionByServiceAndFunctionName(serviceName, functionName)
	if err != nil {
		return "", err
	}
	request := function.Request()

	if len(acceptedData) > 0 {
		err = json.Unmarshal([]byte(acceptedData), request)
		if err != nil {
			return "", err
		}
	}

	requestBody, err := xml.Marshal(request)
	if err != nil {
		return "", err
	}

	endpoint, err := dev.GetEndpointByRequestStruct(request)
	if err != nil {
		return "", err
	}

	soap := gosoap.NewEmptySOAP()
	soap.AddStringBodyContent(string(requestBody))
	soap.AddRootNamespaces(onvif.Xlmns)
	soap.AddWSSecurity(dev.GetDeviceParams().Username, dev.GetDeviceParams().Password)

	servResp, err := networking.SendSoap(new(http.Client), endpoint, soap.String())
	if err != nil {
		return "", err
	}

	defer servResp.Body.Close()

	rsp, err := io.ReadAll(servResp.Body)
	if err != nil {
		return "", err
	}

	responseEnvelope := gosoap.NewSOAPEnvelope(function.Response())
	err = xml.Unmarshal(rsp, responseEnvelope)
	if err != nil {
		return "", err
	}

	var jsonData []byte

	if responseEnvelope.Body.Fault != nil {
		jsonData, err = json.Marshal(responseEnvelope.Body.Fault)
		if err != nil {
			return "", err
		}
		return string(jsonData), nil
	} else {
		jsonData, err = json.Marshal(responseEnvelope.Body.Content)
		if err != nil {
			return "", err
		}
		return string(jsonData), nil
	}
}

func functionByServiceAndFunctionName(serviceName, functionName string) (onvif.Function, error) {
	var function onvif.Function
	var exist bool
	switch serviceName {
	case onvif.DeviceWebService:
		function, exist = onvif.DeviceFunctionMap[functionName]
		if !exist {
			return nil, fmt.Errorf("the web service '%s'not support the function '%s'", serviceName, functionName)
		}
	case onvif.MediaWebService:
		function, exist = onvif.MediaFunctionMap[functionName]
		if !exist {
			return nil, fmt.Errorf("the web service '%s' not support the function '%s'", serviceName, functionName)
		}
	default:
		return nil, fmt.Errorf("not support the web service '%s'", serviceName)
	}
	return function, nil
}
