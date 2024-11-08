package main

import (
	"fmt"
	"github.com/IOTechSystems/onvif/ws-discovery"
)

func main() {
	probeResponses := []string{
		`<env:Envelope>
			<env:Header></env:Header>
			<env:Body>
				<d:ProbeMatches>
					<d:ProbeMatch>
						<wsadis:EndpointReference>
							<wsadis:Address>urn:uuid:cea94000-fb96-11b3-8260-686dbc5cb15d</wsadis:Address>
						</wsadis:EndpointReference>
						<d:Types>dn:NetworkVideoTransmitter tds:Device</d:Types>
						<d:Scopes>onvif://www.onvif.org/type/video_encoder onvif://www.onvif.org/Profile/Streaming onvif://www.onvif.org/MAC/68:6d:bc:5c:b1:5d onvif://www.onvif.org/hardware/DFI6256TE http:123</d:Scopes>
						<d:XAddrs>http://192.168.12.123/onvif/device_service</d:XAddrs>
						<d:MetadataVersion>10</d:MetadataVersion>
					</d:ProbeMatch>
				</d:ProbeMatches>
			</env:Body>
		</env:Envelope>`,
		`<SOAP-ENV:Envelope>
		<SOAP-ENV:Header></SOAP-ENV:Header>
		<SOAP-ENV:Body>
			<wsdd:ProbeMatches>
				<wsdd:ProbeMatch>
					<wsa:EndpointReference>
						<wsa:Address>uuid:3fa1fe68-b915-4053-a3e1-c006c3afec0e</wsa:Address>
						<wsa:ReferenceProperties>
						</wsa:ReferenceProperties>
						<wsa:PortType>ttl</wsa:PortType>
					</wsa:EndpointReference>
					<wsdd:Types>tdn:NetworkVideoTransmitter</wsdd:Types>
					<wsdd:Scopes>onvif://www.onvif.org/name/TP-IPC onvif://www.onvif.org/hardware/MODEL onvif://www.onvif.org/Profile/Streaming onvif://www.onvif.org/location/ShenZhen onvif://www.onvif.org/type/NetworkVideoTransmitter </wsdd:Scopes>
					<wsdd:XAddrs>http://192.168.12.128:2020/onvif/device_service</wsdd:XAddrs>
					<wsdd:MetadataVersion>1</wsdd:MetadataVersion>
				</wsdd:ProbeMatch>
			</wsdd:ProbeMatches>
		</SOAP-ENV:Body>
		</SOAP-ENV:Envelope>`,
	}

	devices, err := wsdiscovery.DevicesFromProbeResponses(probeResponses)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, device := range devices {
		fmt.Printf("Device Xaddr %s", device.GetDeviceParams().Xaddr)
		fmt.Printf("Device EndpointRefAddress %s", device.GetDeviceParams().EndpointRefAddress)
	}
}
