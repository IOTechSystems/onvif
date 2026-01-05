package main

import (
	"io"
	"log"

	"github.com/IOTechSystems/onvif"
	"github.com/IOTechSystems/onvif/media2"
)

func main() {
	dev, err := onvif.NewDevice(onvif.DeviceParams{
		Xaddr:    "192.168.12.148", // BOSCH
		Username: "administrator",
		Password: "Password1!",
		AuthMode: onvif.UsernameTokenAuth,
	})
	if err != nil {
		log.Fatalln("fail to new device:", err)
	}

	res, err := dev.CallMethod(media2.GetProfiles{})
	if err != nil {
		log.Fatalln("fail to CallMethod:", err)
	}
	bs, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("fail to read CallMethod response:", err)
	}

	log.Printf(">> Result: %+v \n %s", res.StatusCode, bs)
}
