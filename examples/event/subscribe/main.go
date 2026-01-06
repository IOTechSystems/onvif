package main

import (
	"io"
	"log"

	"github.com/IOTechSystems/onvif"
	"github.com/IOTechSystems/onvif/event"
	"github.com/IOTechSystems/onvif/xsd"
)

// === Geovision ===
// Request:
// <wsnt:Subscribe>
//      <wsnt:TerminationTime>2021-12-02T10:10:16Z</wsnt:TerminationTime>
//  </wsnt:Subscribe>
// Response:
//   <SOAP-ENV:Fault>
//  	<SOAP-ENV:Code>
// 		    <SOAP-ENV:Value>
//		      SOAP-ENV:Sender
//		    </SOAP-ENV:Value>
//		</SOAP-ENV:Code>
//	 	<SOAP-ENV:Reason>
//		  <SOAP-ENV:Text xml:lang="en">
//			  Validation constraint violation: invalid value in element 'wsnt:TerminationTime'
//		  </SOAP-ENV:Text>
//	    </SOAP-ENV:Reason>
//	</SOAP-ENV:Fault>
//
// Test Summary:
//   1. The time pattern 2021-12-02T10:20:15Z and PT1H not work and always return error
//   2. If not provide the time, camera return the error: ' fail to CallMethod: Post "http://192.168.12.149/onvif/events": net/http: HTTP/1.x transport connection broken: unexpected EOF'
//   3. Geovision might not support the BaseNotification

// === BOSCH ===
// Request:
//	<wsnt:Subscribe>
//	   <wsnt:ConsumerReference>
//		   <wsa:Address>http://192.168.12.112:8080/ping</wsa:Address>
//	   </wsnt:ConsumerReference>
//	   <wsnt:TerminationTime>
//		   2021-12-01T15:50:03Z
//	   </wsnt:TerminationTime>
//   </wsnt:Subscribe>
// Response:
// <wsnt:SubscribeResponse>
//		<wsnt:SubscriptionReference>
//			<wsa5:Address>http://192.168.12.148/Web_Service?Idx=1</wsa5:Address>
//		</wsnt:SubscriptionReference>
//		<wsnt:CurrentTime>2021-12-02T03:21:13Z</wsnt:CurrentTime>
//		<wsnt:TerminationTime>2021-12-02T03:22:13Z</wsnt:TerminationTime>
//	</wsnt:SubscribeResponse>
//
// Test Summary
//  1. The TerminationTime always increase one minute
//  2. The SubscriptionReference Address will change when create another PullPoint

// Hikvision
// Request:
//  <wsnt:Subscribe>
//      <wsnt:ConsumerReference>
//          <wsa:Address>http://192.168.12.112:8080/ping</wsa:Address>
//      </wsnt:ConsumerReference>
//      <wsnt:TerminationTime>2021-12-03T15:50:03Z</wsnt:TerminationTime>
//  </wsnt:Subscribe>
// Response:
//  <wsnt:SubscribeResponse>
// 	  <wsnt:SubscriptionReference>
//		 <wsa:Address>http:///onvif/Events/SubManager_2021-12-02T06:21:11Z_1</wsa:Address>
//    </wsnt:SubscriptionReference>
//    <wsnt:CurrentTime>2021-12-02T06:21:11Z</wsnt:CurrentTime>
//    <wsnt:TerminationTime>2021-12-02T06:22:11Z</wsnt:TerminationTime>
//  </wsnt:SubscribeResponse>
//
// Test Summary
//  1. Regardless of the request's TerminationTime, the response TerminationTime always increase one minute
//  2. The SubscriptionReference Address will change when create another PullPoint
//  3. The response's reference address is invalid, renew request will fail

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

	consumerAddress := event.AttributedURIType("http://192.168.12.112:8080/ping")
	//terminationTime:= xsd.String("2021-12-03T15:50:03Z")
	terminationTime := xsd.String("PT180S")
	res, err := dev.CallMethod(event.Subscribe{
		ConsumerReference: &event.EndpointReferenceType{
			Address: consumerAddress,
		},
		TerminationTime: &terminationTime,
	},
	)
	if err != nil {
		log.Fatalln("fail to CallMethod:", err)
	}
	log.Printf(">> Status Code: %+v \n", res.StatusCode)
	bs, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("fail to read CallMethod response:", err)
	}

	log.Printf(">> Result: %+v \n %s", res.StatusCode, bs)
}
