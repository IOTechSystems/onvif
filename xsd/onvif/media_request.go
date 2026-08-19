package onvif

import "github.com/IOTechSystems/onvif/xsd"

/*
Request-side mirrors of the tt: configuration types used by the ver10 media
service's Set* operations.

onvif.xsd is elementFormDefault="qualified", so every child of a tt: type must be
emitted in the ver10 schema namespace ("onvif:" per Device.go's prefix map).
Marshaling the read-side structs sent these children in no namespace, which a
lenient device accepts and then ignores -- the request returns success while
nothing changes. Attributes stay unqualified: onvif.xsd sets no
attributeFormDefault.

Optional (minOccurs="0") members are pointers with ",omitempty" so a partial
update does not transmit zero values the device would apply verbatim.
*/

type AudioEncoderConfigurationRequest struct {
	ConfigurationEntityRequest
	Encoding       AudioEncoding                  `xml:"onvif:Encoding"`
	Bitrate        *xsd.Int                       `xml:"onvif:Bitrate,omitempty"`
	SampleRate     *xsd.Int                       `xml:"onvif:SampleRate,omitempty"`
	Multicast      *MulticastConfigurationRequest `xml:"onvif:Multicast,omitempty"`
	SessionTimeout *xsd.Duration                  `xml:"onvif:SessionTimeout,omitempty"`
}

type AudioSourceConfigurationRequest struct {
	ConfigurationEntityRequest
	SourceToken ReferenceToken `xml:"onvif:SourceToken"`
}

type VideoAnalyticsConfigurationRequest struct {
	ConfigurationEntityRequest
	AnalyticsEngineConfiguration *AnalyticsEngineConfigurationRequest `xml:"onvif:AnalyticsEngineConfiguration,omitempty"`
	RuleEngineConfiguration      *RuleEngineConfigurationRequest      `xml:"onvif:RuleEngineConfiguration,omitempty"`
}

type RuleEngineConfigurationRequest struct {
	Rule      []ConfigRequest                   `xml:"onvif:Rule,omitempty"`
	Extension *RuleEngineConfigurationExtension `xml:"onvif:Extension,omitempty"`
}

type AudioOutputConfigurationRequest struct {
	ConfigurationEntityRequest
	OutputToken ReferenceToken `xml:"onvif:OutputToken"`
	SendPrimacy *xsd.AnyURI    `xml:"onvif:SendPrimacy,omitempty"`
	OutputLevel *xsd.Int       `xml:"onvif:OutputLevel,omitempty"`
}

type AudioDecoderConfigurationRequest struct {
	ConfigurationEntityRequest
}

// OSDConfigurationRequest mirrors tt:OSDConfiguration, which derives from
// tt:DeviceEntity (so its token is an attribute).
type OSDConfigurationRequest struct {
	Token                         ReferenceToken               `xml:"token,attr,omitempty"`
	VideoSourceConfigurationToken *OSDReference                `xml:"onvif:VideoSourceConfigurationToken,omitempty"`
	Type                          *OSDType                     `xml:"onvif:Type,omitempty"`
	Position                      *OSDPosConfigurationRequest  `xml:"onvif:Position,omitempty"`
	TextString                    *OSDTextConfigurationRequest `xml:"onvif:TextString,omitempty"`
	Image                         *OSDImgConfigurationRequest  `xml:"onvif:Image,omitempty"`
}

type OSDPosConfigurationRequest struct {
	Type *xsd.String      `xml:"onvif:Type,omitempty"`
	Pos  *Vector2DRequest `xml:"onvif:Pos,omitempty"`
}

// Vector2DRequest mirrors tt:Vector, whose coordinates are attributes.
type Vector2DRequest struct {
	X *xsd.Float `xml:"x,attr,omitempty"`
	Y *xsd.Float `xml:"y,attr,omitempty"`
}

type OSDTextConfigurationRequest struct {
	Type             *xsd.String      `xml:"onvif:Type,omitempty"`
	DateFormat       *xsd.String      `xml:"onvif:DateFormat,omitempty"`
	TimeFormat       *xsd.String      `xml:"onvif:TimeFormat,omitempty"`
	FontSize         *xsd.Int         `xml:"onvif:FontSize,omitempty"`
	FontColor        *OSDColorRequest `xml:"onvif:FontColor,omitempty"`
	BackgroundColor  *OSDColorRequest `xml:"onvif:BackgroundColor,omitempty"`
	PlainText        *xsd.String      `xml:"onvif:PlainText,omitempty"`
	IsPersistentText *xsd.Boolean     `xml:"IsPersistentText,attr,omitempty"`
}

type OSDColorRequest struct {
	Color       *ColorRequest `xml:"onvif:Color,omitempty"`
	Transparent *xsd.Int      `xml:"Transparent,attr,omitempty"`
}

// ColorRequest mirrors tt:Color, whose components are attributes.
type ColorRequest struct {
	X          *xsd.Float  `xml:"X,attr,omitempty"`
	Y          *xsd.Float  `xml:"Y,attr,omitempty"`
	Z          *xsd.Float  `xml:"Z,attr,omitempty"`
	Colorspace *xsd.AnyURI `xml:"Colorspace,attr,omitempty"`
}

type OSDImgConfigurationRequest struct {
	ImgPath *xsd.AnyURI `xml:"onvif:ImgPath,omitempty"`
}
