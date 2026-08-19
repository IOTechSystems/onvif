package onvif

import "github.com/IOTechSystems/onvif/xsd"

/*
Request-side mirrors of the tt: imaging types.

Two things distinguish these from the read types above:

  - onvif.xsd is elementFormDefault="qualified", so every child element must be
    emitted in the ver10 schema namespace ("onvif:" per Device.go's prefix map).
    The read types carry unprefixed tags, which is fine for unmarshaling -- Go
    matches on local name -- but marshaling them sends children in no namespace,
    which a device may silently ignore.

  - Every member of tt:ImagingSettings20 and tt:FocusMove is minOccurs="0", so the
    fields are pointers with ",omitempty". Sending a zero value is not the same as
    not sending the field: tt:ExposureMode and tt:AutoFocusMode are enum
    restrictions with no empty member, so an empty <Mode></Mode> is schema-invalid
    and can make a device reject or mis-apply the whole SetImagingSettings call.

Attributes stay unqualified: onvif.xsd does not set attributeFormDefault.
*/

type ImagingSettings20Request struct {
	BacklightCompensation *BacklightCompensation20Request    `xml:"onvif:BacklightCompensation,omitempty"`
	Brightness            *xsd.Float                         `xml:"onvif:Brightness,omitempty"`
	ColorSaturation       *xsd.Float                         `xml:"onvif:ColorSaturation,omitempty"`
	Contrast              *xsd.Float                         `xml:"onvif:Contrast,omitempty"`
	Exposure              *Exposure20Request                 `xml:"onvif:Exposure,omitempty"`
	Focus                 *FocusConfiguration20Request       `xml:"onvif:Focus,omitempty"`
	IrCutFilter           *IrCutFilterMode                   `xml:"onvif:IrCutFilter,omitempty"`
	Sharpness             *xsd.Float                         `xml:"onvif:Sharpness,omitempty"`
	WideDynamicRange      *WideDynamicRange20Request         `xml:"onvif:WideDynamicRange,omitempty"`
	WhiteBalance          *WhiteBalance20Request             `xml:"onvif:WhiteBalance,omitempty"`
	Extension             *ImagingSettingsExtension20Request `xml:"onvif:Extension,omitempty"`
}

type BacklightCompensation20Request struct {
	Mode  *BacklightCompensationMode `xml:"onvif:Mode,omitempty"`
	Level *xsd.Float                 `xml:"onvif:Level,omitempty"`
}

type Exposure20Request struct {
	Mode            *ExposureMode     `xml:"onvif:Mode,omitempty"`
	Priority        *ExposurePriority `xml:"onvif:Priority,omitempty"`
	Window          *RectangleRequest `xml:"onvif:Window,omitempty"`
	MinExposureTime *xsd.Float        `xml:"onvif:MinExposureTime,omitempty"`
	MaxExposureTime *xsd.Float        `xml:"onvif:MaxExposureTime,omitempty"`
	MinGain         *xsd.Float        `xml:"onvif:MinGain,omitempty"`
	MaxGain         *xsd.Float        `xml:"onvif:MaxGain,omitempty"`
	MinIris         *xsd.Float        `xml:"onvif:MinIris,omitempty"`
	MaxIris         *xsd.Float        `xml:"onvif:MaxIris,omitempty"`
	ExposureTime    *xsd.Float        `xml:"onvif:ExposureTime,omitempty"`
	Gain            *xsd.Float        `xml:"onvif:Gain,omitempty"`
	Iris            *xsd.Float        `xml:"onvif:Iris,omitempty"`
}

// RectangleRequest mirrors tt:Rectangle, whose bounds are attributes.
type RectangleRequest struct {
	Bottom *xsd.Float `xml:"bottom,attr,omitempty"`
	Top    *xsd.Float `xml:"top,attr,omitempty"`
	Right  *xsd.Float `xml:"right,attr,omitempty"`
	Left   *xsd.Float `xml:"left,attr,omitempty"`
}

type FocusConfiguration20Request struct {
	AutoFocusMode *AutoFocusMode `xml:"onvif:AutoFocusMode,omitempty"`
	DefaultSpeed  *xsd.Float     `xml:"onvif:DefaultSpeed,omitempty"`
	NearLimit     *xsd.Float     `xml:"onvif:NearLimit,omitempty"`
	FarLimit      *xsd.Float     `xml:"onvif:FarLimit,omitempty"`
	AFMode        StringAttrList `xml:"AFMode,attr,omitempty"`
}

type WideDynamicRange20Request struct {
	Mode  *WideDynamicMode `xml:"onvif:Mode,omitempty"`
	Level *xsd.Float       `xml:"onvif:Level,omitempty"`
}

type WhiteBalance20Request struct {
	Mode   *WhiteBalanceMode `xml:"onvif:Mode,omitempty"`
	CrGain *xsd.Float        `xml:"onvif:CrGain,omitempty"`
	CbGain *xsd.Float        `xml:"onvif:CbGain,omitempty"`
}

type ImagingSettingsExtension20Request struct {
	ImageStabilization *ImageStabilizationRequest          `xml:"onvif:ImageStabilization,omitempty"`
	Extension          *ImagingSettingsExtension202Request `xml:"onvif:Extension,omitempty"`
}

type ImageStabilizationRequest struct {
	Mode  *ImageStabilizationMode `xml:"onvif:Mode,omitempty"`
	Level *xsd.Float              `xml:"onvif:Level,omitempty"`
}

type ImagingSettingsExtension202Request struct {
	IrCutFilterAutoAdjustment []IrCutFilterAutoAdjustmentRequest  `xml:"onvif:IrCutFilterAutoAdjustment,omitempty"`
	Extension                 *ImagingSettingsExtension203Request `xml:"onvif:Extension,omitempty"`
}

type IrCutFilterAutoAdjustmentRequest struct {
	BoundaryType   *xsd.String   `xml:"onvif:BoundaryType,omitempty"`
	BoundaryOffset *xsd.Float    `xml:"onvif:BoundaryOffset,omitempty"`
	ResponseTime   *xsd.Duration `xml:"onvif:ResponseTime,omitempty"`
}

type ImagingSettingsExtension203Request struct {
	ToneCompensation *ToneCompensationRequest `xml:"onvif:ToneCompensation,omitempty"`
	Defogging        *DefoggingRequest        `xml:"onvif:Defogging,omitempty"`
	NoiseReduction   *NoiseReductionRequest   `xml:"onvif:NoiseReduction,omitempty"`
}

type ToneCompensationRequest struct {
	Mode  *xsd.String `xml:"onvif:Mode,omitempty"`
	Level *xsd.Float  `xml:"onvif:Level,omitempty"`
}

type DefoggingRequest struct {
	Mode  *xsd.String `xml:"onvif:Mode,omitempty"`
	Level *xsd.Float  `xml:"onvif:Level,omitempty"`
}

type NoiseReductionRequest struct {
	Level *xsd.Float `xml:"onvif:Level,omitempty"`
}

/*
FocusMoveRequest mirrors tt:FocusMove. Absolute, Relative and Continuous are the
three mutually-exclusive move modes, each minOccurs="0" -- exactly one should be
set per Move call.
*/
type FocusMoveRequest struct {
	Absolute   *AbsoluteFocusRequest   `xml:"onvif:Absolute,omitempty"`
	Relative   *RelativeFocusRequest   `xml:"onvif:Relative,omitempty"`
	Continuous *ContinuousFocusRequest `xml:"onvif:Continuous,omitempty"`
}

type AbsoluteFocusRequest struct {
	Position *xsd.Float `xml:"onvif:Position,omitempty"`
	Speed    *xsd.Float `xml:"onvif:Speed,omitempty"`
}

type RelativeFocusRequest struct {
	Distance *xsd.Float `xml:"onvif:Distance,omitempty"`
	Speed    *xsd.Float `xml:"onvif:Speed,omitempty"`
}

type ContinuousFocusRequest struct {
	Speed *xsd.Float `xml:"onvif:Speed,omitempty"`
}
