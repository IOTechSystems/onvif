package main

import (
	"io"
	"log"

	"github.com/IOTechSystems/onvif"
	"github.com/IOTechSystems/onvif/event"
	"github.com/IOTechSystems/onvif/xsd"
)

// Geovision
// Request:
// <tev:CreatePullPointSubscription>
//      <tev:InitialTerminationTime>PT120S</tev:InitialTerminationTime>
//  </tev:CreatePullPointSubscription>
// Response:
// <tev:CreatePullPointSubscriptionResponse>
//	 <tev:SubscriptionReference>
//	   <wsa5:Address>http://192.168.12.149:80/onvif/events</wsa5:Address>
//	 </tev:SubscriptionReference>
//	 <wsnt:CurrentTime>2021-12-02T02:02:15Z</wsnt:CurrentTime>
//	 <wsnt:TerminationTime>2021-12-02T02:04:15Z</wsnt:TerminationTime>
// </tev:CreatePullPointSubscriptionResponse>
//
// Test Summary:
//	1. TerminationTime = CurrentTime+InitialTerminationTime
//	2. but always return http://192.168.12.149:80/onvif/events
//  3. We can pulling the event without creating the pull point, and the unsubscribe is not supported

// BOSCH
// Request:
// <tev:CreatePullPointSubscription>
//      <tev:InitialTerminationTime>PT1H</tev:InitialTerminationTime>
//  </tev:CreatePullPointSubscription>
// Response:
// <tev:CreatePullPointSubscriptionResponse>
//    <tev:SubscriptionReference>
//       <wsa5:Address>http://192.168.12.148/Web_Service?Idx=1</wsa5:Address>
//    </tev:SubscriptionReference>
//   <wsnt:CurrentTime>2021-12-02T02:59:30Z</wsnt:CurrentTime>
//   <wsnt:TerminationTime>2021-12-02T03:00:30Z</wsnt:TerminationTime>
// </tev:CreatePullPointSubscriptionResponse>
//
// Test Summary
//  1. The TerminationTime always increase one minute
//  2. The SubscriptionReference Address will change when create another PullPoint

// === Hikvision ===
// Request:
//  <tev:CreatePullPointSubscription>
//      <tev:InitialTerminationTime>PT120S</tev:InitialTerminationTime>
//  </tev:CreatePullPointSubscription>
// Response:
// <tev:CreatePullPointSubscriptionResponse>
//   <tev:SubscriptionReference>
//  	<wsa:Address>http://192.168.12.123/onvif/Events/PullSubManager_2021-12-02T06:04:46Z_0</wsa:Address>
//   </tev:SubscriptionReference>
//   <wsnt:CurrentTime>2021-12-02T06:04:46Z</wsnt:CurrentTime>
//   <wsnt:TerminationTime>2021-12-02T06:06:46Z</wsnt:TerminationTime>
// </tev:CreatePullPointSubscriptionResponse>
//
// Test Summary:
//	 1. TerminationTime = CurrentTime+InitialTerminationTime
//   2. The pull point will drop after exceeding the termination time

func main() {
	dev, err := onvif.NewDevice(onvif.DeviceParams{
		//Xaddr:    "192.168.12.148", // BOSCH
		//Xaddr:    "192.168.12.149", // Geovision
		Xaddr:    "192.168.12.123", //Hikvision
		Username: "administrator",
		Password: "Password1!",
		AuthMode: onvif.UsernameTokenAuth,
	})
	if err != nil {
		log.Fatalln("fail to new device:", err)
	}

	initialTerminationTime := xsd.String("PT120S")
	res, err := dev.CallMethod(event.CreatePullPointSubscription{InitialTerminationTime: &initialTerminationTime})
	if err != nil {
		log.Fatalln("fail to CallMethod:", err)
	}
	bs, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("fail to read CallMethod response:", err)
	}

	log.Printf(">> Result: %+v \n %s", res.StatusCode, bs)
}
