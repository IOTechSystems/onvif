package onvif

import "github.com/IOTechSystems/onvif/xsd"

/*
Request-side mirrors of the tt: PTZ types.

The PTZ service marshals tt:PTZSpeed and tt:PTZConfiguration outbound
(ContinuousMove/Velocity, GotoPreset/Speed, GotoHomePosition/Speed,
GeoMove/Speed, SetConfiguration/PTZConfiguration), and onvif.xsd is
elementFormDefault="qualified", so their children must carry the ver10 schema
namespace. The read-side structs deliberately leave these tags unprefixed --
Go enforces the namespace half of a tag on unmarshal, so a prefix there would
silently leave the fields at their zero value.

Attributes stay unqualified: onvif.xsd sets no attributeFormDefault.
*/

type PTZSpeedRequest struct {
	PanTilt *Vector2D `xml:"onvif:PanTilt,omitempty"`
	Zoom    *Vector1D `xml:"onvif:Zoom,omitempty"`
}

type PTZConfigurationRequest struct {
	Token          ReferenceToken `xml:"token,attr,omitempty"`
	MoveRamp       *xsd.Int       `xml:"MoveRamp,attr,omitempty"`
	PresetRamp     *xsd.Int       `xml:"PresetRamp,attr,omitempty"`
	PresetTourRamp *xsd.Int       `xml:"PresetTourRamp,attr,omitempty"`

	Name     Name `xml:"onvif:Name,omitempty"`
	UseCount int  `xml:"onvif:UseCount,omitempty"`

	NodeToken                              *ReferenceToken       `xml:"onvif:NodeToken,omitempty"`
	DefaultAbsolutePantTiltPositionSpace   *xsd.AnyURI           `xml:"onvif:DefaultAbsolutePantTiltPositionSpace,omitempty"`
	DefaultAbsoluteZoomPositionSpace       *xsd.AnyURI           `xml:"onvif:DefaultAbsoluteZoomPositionSpace,omitempty"`
	DefaultRelativePanTiltTranslationSpace *xsd.AnyURI           `xml:"onvif:DefaultRelativePanTiltTranslationSpace,omitempty"`
	DefaultRelativeZoomTranslationSpace    *xsd.AnyURI           `xml:"onvif:DefaultRelativeZoomTranslationSpace,omitempty"`
	DefaultContinuousPanTiltVelocitySpace  *xsd.AnyURI           `xml:"onvif:DefaultContinuousPanTiltVelocitySpace,omitempty"`
	DefaultContinuousZoomVelocitySpace     *xsd.AnyURI           `xml:"onvif:DefaultContinuousZoomVelocitySpace,omitempty"`
	DefaultPTZSpeed                        *PTZSpeedRequest      `xml:"onvif:DefaultPTZSpeed,omitempty"`
	DefaultPTZTimeout                      *xsd.Duration         `xml:"onvif:DefaultPTZTimeout,omitempty"`
	PanTiltLimits                          *PanTiltLimitsRequest `xml:"onvif:PanTiltLimits,omitempty"`
	ZoomLimits                             *ZoomLimitsRequest    `xml:"onvif:ZoomLimits,omitempty"`
}

type PanTiltLimitsRequest struct {
	Range *Space2DDescriptionRequest `xml:"onvif:Range,omitempty"`
}

type Space2DDescriptionRequest struct {
	URI    *xsd.AnyURI        `xml:"onvif:URI,omitempty"`
	XRange *FloatRangeRequest `xml:"onvif:XRange,omitempty"`
	YRange *FloatRangeRequest `xml:"onvif:YRange,omitempty"`
}

type ZoomLimitsRequest struct {
	Range *Space1DDescriptionRequest `xml:"onvif:Range,omitempty"`
}

type Space1DDescriptionRequest struct {
	URI    *xsd.AnyURI        `xml:"onvif:URI,omitempty"`
	XRange *FloatRangeRequest `xml:"onvif:XRange,omitempty"`
}

type FloatRangeRequest struct {
	Min *xsd.Float `xml:"onvif:Min,omitempty"`
	Max *xsd.Float `xml:"onvif:Max,omitempty"`
}
