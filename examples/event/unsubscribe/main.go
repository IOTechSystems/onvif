package main

import (
	"io"
	"log"

	"github.com/nbio/xml"

	"github.com/secure-passage/onvif"
	"github.com/secure-passage/onvif/event"
)

// === Geovision ===
// Request:
//  <tev:Unsubscribe />
// Response:
// <SOAP-ENV:Fault>
//   <SOAP-ENV:Code>
//      <SOAP-ENV:Value>SOAP-ENV:Sender</SOAP-ENV:Value>
//   </SOAP-ENV:Code>
//   <SOAP-ENV:Reason>
//       <SOAP-ENV:Text xml:lang="en">
//         Method 'tev:Unsubscribe' not implemented: method name or namespace not recognized
//       </SOAP-ENV:Text>
//   </SOAP-ENV:Reason>
// </SOAP-ENV:Fault>
//
// Test Summary: Geovision might not support unsubscribe

// === BOSCH ===
// Request:
// <tev:Unsubscribe />
// Response:
// <SOAP-ENV:Fault>
//   <SOAP-ENV:Code><SOAP-ENV:Value>SOAP-ENV:Receiver</SOAP-ENV:Value><SOAP-ENV:Subcode><SOAP-ENV:Value>ter:Action</SOAP-ENV:Value></SOAP-ENV:Subcode></SOAP-ENV:Code>
//    <SOAP-ENV:Reason><SOAP-ENV:Text xml:lang="en">Action Failed</SOAP-ENV:Text></SOAP-ENV:Reason>
//  <SOAP-ENV:Node>http://www.w3.org/2003/05/soap-envelope/node/ultimateReceiver</SOAP-ENV:Node><SOAP-ENV:Role>http://www.w3.org/2003/05/soap-envelope/node/ultimateReceiver</SOAP-ENV:Role>
// </SOAP-ENV:Fault>
//
// Test Summary: BOSCH might not support unsubscribe

// === Hikvision
// Request:
//  <tev:Unsubscribe />
// Response:
//  <wsnt:UnsubscribeResponse/>

func main() {
	dev, err := onvif.NewDevice(onvif.DeviceParams{
		// Xaddr:    "192.168.12.148", // BOSCH
		// Xaddr:    "192.168.12.149", // Geovision
		Xaddr:    "192.168.12.123", // Hikvision
		Username: "administrator",
		Password: "Password1!",
		AuthMode: onvif.UsernameTokenAuth,
	})
	if err != nil {
		log.Fatalln("fail to new device:", err)
	}

	unsubscribe := event.Unsubscribe{}

	// endPoint:= "http://192.168.12.148/Web_Service?Idx=0" // BOSCH
	// endPoint := "http://192.168.12.149:80/onvif/events" // Geovision
	endPoint := "http://192.168.12.123/onvif/Events/PullSubManager_2021-12-02T06:13:45Z_0" // Hikvision
	requestBody, err := xml.Marshal(unsubscribe)
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
