package main

import (
	"io"
	"log"

	"github.com/nbio/xml"

	"github.com/secure-passage/onvif"
	"github.com/secure-passage/onvif/event"
	"github.com/secure-passage/onvif/xsd"
)

// === Geovision ===
// Request:
// <tev:PullMessages>
//	 <tev:Timeout>PT20S</tev:Timeout>
//	 <tev:MessageLimit>10</tev:MessageLimit>
// </tev:PullMessages>
// Response:
// <tev:PullMessagesResponse>
//   <tev:CurrentTime>2021-12-02T02:42:30Z</tev:CurrentTime>
//   <tev:TerminationTime>2021-12-02T02:42:50Z</tev:TerminationTime>
//  <wsnt:NotificationMessage> ... </wsnt:NotificationMessage>
// </tev:PullMessagesResponse>
//
// Test Summary:
//	1. the TerminationTime = CurrentTime+Timeout
//	2. even current time exceed the TerminationTime, the pull point still alive

// === BOSCH ===
// Request:
//  <wsnt:Renew>
//        <wsnt:TerminationTime>
//            2021-12-03T15:50:03Z
//        </wsnt:TerminationTime>
//  </wsnt:Renew>
// Response:
// <wsnt:RenewResponse>
//	  <wsnt:TerminationTime>
//	 	 2021-12-03T15:50:03Z
//	 </wsnt:TerminationTime>
//    <wsnt:CurrentTime>
//   	2021-12-02T03:36:04Z
//    </wsnt:CurrentTime>
//	</wsnt:RenewResponse>
//
// Test Summary:
//	1. the response's TerminationTime equal request's TerminationTime
//  2. But the subscription still live for one minute

// === Hikvision ===
// Request:
// <wsnt:Renew><wsnt:TerminationTime>2021-12-02T18:30:53Z</wsnt:TerminationTime></wsnt:Renew>
// Response:
//  <wsnt:RenewResponse>
// 		<wsnt:TerminationTime>2021-12-02T18:30:53Z</wsnt:TerminationTime>
//    	<wsnt:CurrentTime>2021-12-02T06:31:57Z</wsnt:CurrentTime>
//  </wsnt:RenewResponse>
//
// Test Summary:
// 1. the subscription's termination time will update if the request's TerminationTime greater than the curren time

func main() {
	dev, err := onvif.NewDevice(onvif.DeviceParams{
		Xaddr: "192.168.12.148", // BOSCH
		// Xaddr:    "192.168.12.149", // Geovision
		// Xaddr: "192.168.12.123", //Hikvision
		Username: "administrator",
		Password: "Password1!",
		AuthMode: onvif.UsernameTokenAuth,
	})
	if err != nil {
		log.Fatalln("fail to new device:", err)
	}

	terminationTime := xsd.String("PT120S")
	renew := event.Renew{
		TerminationTime: terminationTime,
	}

	endPoint := "http://192.168.12.148/Web_Service?Idx=0" // BOSCH
	// endPoint := "http://192.168.12.149:80/onvif/events" // Geovision
	// endPoint := "http://192.168.12.123:80/onvif/Events/SubManager__0" // Hikvision
	requestBody, err := xml.Marshal(renew)
	if err != nil {
		log.Fatalln(err)
	}
	res, err := dev.SendSoap(endPoint, string(requestBody))
	if err != nil {
		log.Fatalln("fail to CallMethod:", err)
	}
	bs, _ := io.ReadAll(res.Body)

	log.Printf(">> Result: %+v \n %s", res.StatusCode, bs)
}
