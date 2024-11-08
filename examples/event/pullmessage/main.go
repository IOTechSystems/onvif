package main

import (
	"encoding/xml"
	"github.com/IOTechSystems/onvif"
	"github.com/IOTechSystems/onvif/event"
	"github.com/IOTechSystems/onvif/xsd"
	"io/ioutil"
	"log"
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
// <tev:PullMessages>
//	 <tev:Timeout>PT20S</tev:Timeout>
//	 <tev:MessageLimit>10</tev:MessageLimit>
// </tev:PullMessages>
// Response:
// <tev:PullMessagesResponse>
//	<tev:CurrentTime>2021-12-02T03:06:39Z</tev:CurrentTime>
//	<tev:TerminationTime>2021-12-02T03:07:39Z</tev:TerminationTime>
//	<wsnt:NotificationMessage> ... </wsnt:NotificationMessage>
// </tev:PullMessagesResponse>
//
// Test Summary:
//	the TerminationTime increase one minute

// === Hikvisiion ===
// Request:
//
// Response:
// <tev:PullMessagesResponse><tev:CurrentTime>2021-12-02T06:08:35Z</tev:CurrentTime>
//
//	 <tev:TerminationTime>2021-12-02T06:18:40Z</tev:TerminationTime>
//	  <wsnt:NotificationMessage><wsnt:Topic Dialect="http://www.onvif.org/ver10/tev/topicExpression/ConcreteSet">tns1:RuleEngine/CellMotionDetector/Motion</wsnt:Topic>
//	   ...
//	  </wsnt:NotificationMessage>
//	</tev:PullMessagesResponse>
//
// Test Summary:
//
//	the TerminationTime increase ten minutes
func main() {
	dev, err := onvif.NewDevice(onvif.DeviceParams{
		Xaddr: "192.168.12.148", // BOSCH
		//Xaddr:    "192.168.12.149", // Geovision
		//Xaddr: "192.168.12.123", //Hikvision
		Username: "administrator",
		Password: "Password1!",
		AuthMode: onvif.UsernameTokenAuth,
	})
	if err != nil {
		log.Fatalln("fail to new device:", err)
	}

	pullMessage := event.PullMessages{
		Timeout:      xsd.Duration("PT5S"),
		MessageLimit: 10,
	}

	endPoint := "http://192.168.12.148/Web_Service?Idx=0" // BOSCH
	//endPoint := "http://192.168.12.149:80/onvif/events" // Geovision
	//endPoint := "http://192.168.12.123/onvif/Events/PullSubManager_2021-12-02T06:07:58Z_0"
	requestBody, err := xml.Marshal(pullMessage)
	if err != nil {
		log.Fatalln(err)
	}
	res, err := dev.SendSoap(endPoint, string(requestBody))
	if err != nil {
		log.Fatalln("fail to CallMethod:", err)
	}
	bs, _ := ioutil.ReadAll(res.Body)

	log.Printf(">> Result: %+v \n %s", res.StatusCode, bs)
}
