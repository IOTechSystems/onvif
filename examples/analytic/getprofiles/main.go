package main

import (
	"io"
	"log"

	"github.com/secure-passage/onvif"
	"github.com/secure-passage/onvif/media2"
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

	res, err := dev.CallMethod(media2.GetProfiles{}, nil)
	if err != nil {
		log.Fatalln("fail to CallMethod:", err)
	}
	bs, _ := io.ReadAll(res.Body)

	log.Printf(">> Result: %+v \n %s", res.StatusCode, bs)
}
