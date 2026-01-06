package main

import (
	"io"
	"log"

	"github.com/IOTechSystems/onvif"
	"github.com/IOTechSystems/onvif/device"
)

func main() {
	dev, err := onvif.NewDevice(onvif.DeviceParams{
		Xaddr:    "192.168.12.149",
		Username: "administrator",
		Password: "Password1!",
	})
	if err != nil {
		log.Fatalln("fail to new device:", err)
	}
	// CreateUsers
	res, err := dev.CallMethod(device.GetUsers{})
	if err != nil {
		log.Fatalln("fail to CallMethod:", err)
	}
	bs, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("fail to read CallMethod response:", err)
	}

	log.Printf(">> Result: %+v \n %s", res.StatusCode, bs)
}
