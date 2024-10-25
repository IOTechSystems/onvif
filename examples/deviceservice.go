package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	goonvif "github.com/secure-passage/onvif"
	"github.com/secure-passage/onvif/device"
	"github.com/secure-passage/onvif/gosoap"
	"github.com/secure-passage/onvif/xsd/onvif"
)

const (
	login    = "admin"
	password = "Supervisor"
)

func readResponse(resp *http.Response) string {
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func main() {
	// Getting an camera instance
	dev, err := goonvif.NewDevice(goonvif.DeviceParams{
		Xaddr:      "192.168.13.14:80",
		Username:   login,
		Password:   password,
		HttpClient: new(http.Client),
	})
	if err != nil {
		panic(err)
	}

	// Preparing commands
	UserLevel := onvif.UserLevel("User")
	systemDateAndTyme := device.GetSystemDateAndTime{}
	getCapabilities := device.GetCapabilities{Category: []onvif.CapabilityCategory{"All"}}
	createUser := device.CreateUsers{
		User: []onvif.UserRequest{
			{
				Username:  "TestUser",
				Password:  "TestPassword",
				UserLevel: &UserLevel,
			},
		},
	}

	// Commands execution
	systemDateAndTymeResponse, err := dev.CallMethod(systemDateAndTyme, nil)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(readResponse(systemDateAndTymeResponse))
	}
	getCapabilitiesResponse, err := dev.CallMethod(getCapabilities, nil)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(readResponse(getCapabilitiesResponse))
	}
	createUserResponse, err := dev.CallMethod(createUser, nil)
	if err != nil {
		log.Println(err)
	} else {
		/*
			You could use https://github.com/secure-passage/onvif/gosoap for pretty printing response
		*/
		fmt.Println(gosoap.SoapMessage(readResponse(createUserResponse)).StringIndent())
	}

}
