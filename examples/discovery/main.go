package main

import (
	"encoding/json"
	"io"
	"log"
	"path"
	"regexp"
	"strings"

	"github.com/IOTechSystems/onvif"
	"github.com/IOTechSystems/onvif/device"
	wsdiscovery "github.com/IOTechSystems/onvif/ws-discovery"
	"github.com/beevik/etree"
)

func main() {
	client()
	runDiscovery("en0")
	s, _ := wsdiscovery.GetAvailableDevicesAtSpecificEthernetInterface("en0")

	log.Printf("%v", s)
}

func client() {
	dev, err := onvif.NewDevice(onvif.DeviceParams{Xaddr: "192.168.3.10", Username: "admin", Password: "zsyy12345"})
	if err != nil {
		panic(err)
	}

	log.Printf("output %+v", dev.GetServices())

	res, _ := dev.CallMethod(device.GetUsers{})
	bs, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("fail to read CallMethod response:", err)
	}

	log.Printf("output %+v %s", res.StatusCode, bs)
}

// Host host
type Host struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

func runDiscovery(interfaceName string) {
	var hosts []*Host
	devices, _ := wsdiscovery.SendProbe(interfaceName, nil, []string{"dn:NetworkVideoTransmitter"}, map[string]string{"dn": "http://www.onvif.org/ver10/network/wsdl"})
	for _, j := range devices {
		doc := etree.NewDocument()
		if err := doc.ReadFromString(j); err != nil {
			log.Printf("error %s", err)
		} else {

			endpoints := doc.Root().FindElements("./Body/ProbeMatches/ProbeMatch/XAddrs")
			scopes := doc.Root().FindElements("./Body/ProbeMatches/ProbeMatch/Scopes")

			flag := false

			host := &Host{}

			for _, xaddr := range endpoints {
				xaddr := strings.Split(strings.Split(xaddr.Text(), " ")[0], "/")[2]
				host.URL = xaddr
			}
			if flag {
				break
			}
			for _, scope := range scopes {
				re := regexp.MustCompile(`onvif:\/\/www\.onvif\.org\/name\/[A-Za-z0-9-]+`)
				match := re.FindStringSubmatch(scope.Text())
				host.Name = path.Base(match[0])
			}

			hosts = append(hosts, host)

		}

	}

	bys, _ := json.Marshal(hosts)
	log.Printf("done %s", bys)
}
