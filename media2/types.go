package media2

//go:generate python3 ../python/gen_commands.py

import (
	"github.com/IOTechSystems/onvif/xsd"
	"github.com/IOTechSystems/onvif/xsd/onvif"
)

type GetVideoEncoderConfigurations struct {
	XMLName            string               `xml:"tr2:GetVideoEncoderConfigurations"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken,omitempty"`
	ProfileToken       onvif.ReferenceToken `xml:"tr2:ProfileToken,omitempty"`
}

type GetVideoEncoderConfigurationsResponse struct {
	Configurations []onvif.VideoEncoder2Configuration
}

type SetVideoEncoderConfiguration struct {
	XMLName       string                                  `xml:"tr2:SetVideoEncoderConfiguration"`
	Configuration onvif.VideoEncoder2ConfigurationRequest `xml:"tr2:Configuration"`
}

type SetVideoEncoderConfigurationResponse struct{}

type GetVideoEncoderConfigurationOptions struct {
	XMLName            string               `xml:"tr2:GetVideoEncoderConfigurationOptions"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken,omitempty"`
	ProfileToken       onvif.ReferenceToken `xml:"tr2:ProfileToken,omitempty"`
}

type GetVideoEncoderConfigurationOptionsResponse struct {
	Options []onvif.VideoEncoder2ConfigurationOptions
}

type GetProfiles struct {
	XMLName string               `xml:"tr2:GetProfiles"`
	Token   onvif.ReferenceToken `xml:"tr2:Token,omitempty"`
	Type    []xsd.String         `xml:"tr2:Type,omitempty"`
}

type GetProfilesResponse struct {
	Profiles []Profile
}

type Profile struct {
	Token          string `xml:"token,attr"`
	Fixed          bool   `xml:"fixed,attr"`
	Name           string
	Configurations *ProfileConfigurations `xml:"Configurations,omitempty"`
}

// ProfileConfigurations is returned per-profile when GetProfiles is called with
// a Type filter; each element mirrors the corresponding GetXConfigurations
// response's Configurations element.
type ProfileConfigurations struct {
	VideoEncoder *onvif.VideoEncoder2Configuration `xml:"VideoEncoder,omitempty"`
}

type GetAnalyticsConfigurations struct {
	XMLName string `xml:"tr2:GetAnalyticsConfigurations"`
}

type GetAnalyticsConfigurationsResponse struct {
	Configurations []Configurations
}

type Configurations struct {
	onvif.ConfigurationEntity
	AnalyticsEngineConfiguration *AnalyticsEngineConfiguration `json:",omitempty"`
	RuleEngineConfiguration      *RuleEngineConfiguration      `json:",omitempty"`
}

type AnalyticsEngineConfiguration struct {
	AnalyticsModule []AnalyticsModule
}

type AnalyticsModule struct {
	Name       string `xml:",attr"`
	Type       string `xml:",attr"`
	Parameters Parameters
}

type RuleEngineConfiguration struct {
	Rule []Rule `json:",omitempty"`
}

type Rule struct {
	Name       string `xml:",attr"`
	Type       string `xml:",attr"`
	Parameters Parameters
}

type Parameters struct {
	SimpleItem  []SimpleItem  `json:",omitempty"`
	ElementItem []ElementItem `json:",omitempty"`
}

type SimpleItem struct {
	Name  string `xml:",attr"`
	Value string `xml:",attr"`
}

type ElementItem struct {
	Name string `xml:",attr"`
}

type AddConfiguration struct {
	XMLName       string `xml:"tr2:AddConfiguration"`
	ProfileToken  string `xml:"tr2:ProfileToken"`
	Name          string `xml:"tr2:Name,omitempty"`
	Configuration []Configuration
}

type AddConfigurationResponse struct{}

type RemoveConfiguration struct {
	XMLName       string `xml:"tr2:RemoveConfiguration"`
	ProfileToken  string `xml:"tr2:ProfileToken"`
	Configuration []Configuration
}

type RemoveConfigurationResponse struct{}

// Configuration mirrors tr2:ConfigurationRef, whose Type is required
// (minOccurs defaults to 1) while Token is minOccurs="0" -- omitting Token asks
// the device to pick a configuration of that type itself.
type Configuration struct {
	XMLName xsd.String  `xml:"tr2:Configuration"`
	Type    xsd.String  `xml:"tr2:Type"`
	Token   *xsd.String `xml:"tr2:Token,omitempty"`
}

// GetStreamUri and its properties are defined in the Onvif specification:
// https://www.onvif.org/ver20/media/wsdl/media.wsdl#op.GetStreamUri
type GetStreamUri struct {
	XMLName      string               `xml:"tr2:GetStreamUri"`
	Protocol     string               `xml:"tr2:Protocol"`
	ProfileToken onvif.ReferenceToken `xml:"tr2:ProfileToken"`
}

// GetStreamUriResponse holds the ver20 response. The Media2 WSDL returns a bare
// Uri element (unlike ver10, which wraps a tt:MediaUri); the remaining MediaUri
// fields are accepted as optional for devices that still send them.
type GetStreamUriResponse struct {
	Uri                 string
	InvalidAfterConnect bool
	InvalidAfterReboot  bool
	Timeout             string
}
