package imaging

import (
	"encoding/xml"
	"testing"

	"github.com/IOTechSystems/onvif/xsd"
	"github.com/IOTechSystems/onvif/xsd/onvif"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ---------------------------------------------------------------------------
// SetImagingSettings
// ---------------------------------------------------------------------------

// onvif.xsd is elementFormDefault="qualified", so every child of
// tt:ImagingSettings20 must carry the ver10 schema namespace, and all eleven
// members are minOccurs="0".
//
// SetImagingSettings marshaled the response-side ImagingSettings20 directly, so
// it emitted unqualified children AND every field of the zero value -- including
// <Mode></Mode> for tt:ExposureMode and tt:AutoFocusMode, enum restrictions with
// no empty member. Changing only brightness could therefore reset exposure and
// focus, or fault outright.
func TestMarshalSetImagingSettingsOnlySetsSuppliedFields(t *testing.T) {
	brightness := xsd.Float(55)

	request := SetImagingSettings{
		VideoSourceToken: "VS_1",
		ImagingSettings: onvif.ImagingSettings20Request{
			Brightness: &brightness,
		},
		ForcePersistence: xsd.Boolean(true),
	}

	data, err := xml.Marshal(request)
	require.NoError(t, err)

	assert.Equal(t,
		`<timg:SetImagingSettings>`+
			`<timg:VideoSourceToken>VS_1</timg:VideoSourceToken>`+
			`<timg:ImagingSettings><onvif:Brightness>55</onvif:Brightness></timg:ImagingSettings>`+
			`<timg:ForcePersistence>true</timg:ForcePersistence>`+
			`</timg:SetImagingSettings>`,
		string(data))
}

// The WSDL sequence is VideoSourceToken, ImagingSettings, ForcePersistence, and
// every nested child must be onvif:-qualified too.
func TestMarshalSetImagingSettingsNestedFieldsAreQualified(t *testing.T) {
	mode := onvif.ExposureMode("MANUAL")
	exposureTime := xsd.Float(20000)
	autoFocusMode := onvif.AutoFocusMode("MANUAL")

	request := SetImagingSettings{
		VideoSourceToken: "VS_1",
		ImagingSettings: onvif.ImagingSettings20Request{
			Exposure: &onvif.Exposure20Request{
				Mode:         &mode,
				ExposureTime: &exposureTime,
			},
			Focus: &onvif.FocusConfiguration20Request{
				AutoFocusMode: &autoFocusMode,
			},
		},
	}

	data, err := xml.Marshal(request)
	require.NoError(t, err)

	assert.Contains(t, string(data), `<onvif:Exposure><onvif:Mode>MANUAL</onvif:Mode><onvif:ExposureTime>20000</onvif:ExposureTime></onvif:Exposure>`)
	assert.Contains(t, string(data), `<onvif:Focus><onvif:AutoFocusMode>MANUAL</onvif:AutoFocusMode></onvif:Focus>`)

	// No unqualified leakage and no empty enum elements.
	assert.NotContains(t, string(data), `<Mode>`)
	assert.NotContains(t, string(data), `<onvif:Mode></onvif:Mode>`)
	assert.NotContains(t, string(data), `<onvif:AutoFocusMode></onvif:AutoFocusMode>`)
	assert.NotContains(t, string(data), `<onvif:BacklightCompensation>`)
	assert.NotContains(t, string(data), `<onvif:WhiteBalance>`)
}

// ---------------------------------------------------------------------------
// Move
// ---------------------------------------------------------------------------

// tt:FocusMove's Absolute/Relative/Continuous are the three mutually-exclusive
// move modes, all minOccurs="0". Non-pointer fields meant a continuous move also
// carried <Absolute><Position>0</Position></Absolute>, so the camera could rack
// focus to absolute position 0 instead of nudging it.
func TestMarshalMoveSendsOnlyOneFocusMode(t *testing.T) {
	speed := xsd.Float(0.5)

	request := Move{
		VideoSourceToken: "VS_1",
		Focus: onvif.FocusMoveRequest{
			Continuous: &onvif.ContinuousFocusRequest{Speed: &speed},
		},
	}

	data, err := xml.Marshal(request)
	require.NoError(t, err)

	assert.Equal(t,
		`<timg:Move>`+
			`<timg:VideoSourceToken>VS_1</timg:VideoSourceToken>`+
			`<timg:Focus><onvif:Continuous><onvif:Speed>0.5</onvif:Speed></onvif:Continuous></timg:Focus>`+
			`</timg:Move>`,
		string(data))
}

func TestMarshalMoveAbsoluteFocus(t *testing.T) {
	position := xsd.Float(0.25)
	speed := xsd.Float(0.5)

	request := Move{
		VideoSourceToken: "VS_1",
		Focus: onvif.FocusMoveRequest{
			Absolute: &onvif.AbsoluteFocusRequest{Position: &position, Speed: &speed},
		},
	}

	data, err := xml.Marshal(request)
	require.NoError(t, err)

	assert.Contains(t, string(data),
		`<onvif:Absolute><onvif:Position>0.25</onvif:Position><onvif:Speed>0.5</onvif:Speed></onvif:Absolute>`)
	assert.NotContains(t, string(data), `<onvif:Relative>`)
	assert.NotContains(t, string(data), `<onvif:Continuous>`)
}

// ---------------------------------------------------------------------------
// Read path
// ---------------------------------------------------------------------------

// tt:FocusConfiguration20 carries an AFMode attribute (tt:StringAttrList). It was
// missing entirely, so a get-modify-set silently dropped the camera's AF
// sub-mode list.
func TestUnmarshalGetImagingSettingsResponseAFMode(t *testing.T) {
	responseData := `
		<GetImagingSettingsResponse>
			<ImagingSettings>
				<Brightness>55</Brightness>
				<Focus AFMode="OnceAfterMove">
					<AutoFocusMode>AUTO</AutoFocusMode>
					<NearLimit>0.1</NearLimit>
				</Focus>
			</ImagingSettings>
		</GetImagingSettingsResponse>
	`

	response := &GetImagingSettingsResponse{}
	err := xml.Unmarshal([]byte(responseData), response)
	require.NoError(t, err)

	assert.Equal(t, float64(55), response.ImagingSettings.Brightness)
	assert.Equal(t, onvif.AutoFocusMode("AUTO"), response.ImagingSettings.Focus.AutoFocusMode)
	require.Len(t, response.ImagingSettings.Focus.AFMode, 1)
	assert.Equal(t, "OnceAfterMove", response.ImagingSettings.Focus.AFMode[0])
}

// tt:ImagingSettingsExtension202 declares IrCutFilterAutoAdjustment with
// maxOccurs="unbounded"; a scalar field silently kept only the last entry, so a
// subsequent Set wiped the other boundary.
func TestUnmarshalIrCutFilterAutoAdjustmentKeepsAll(t *testing.T) {
	responseData := `
		<GetImagingSettingsResponse>
			<ImagingSettings>
				<Extension>
					<Extension>
						<IrCutFilterAutoAdjustment>
							<BoundaryType>Common</BoundaryType>
							<BoundaryOffset>0.5</BoundaryOffset>
						</IrCutFilterAutoAdjustment>
						<IrCutFilterAutoAdjustment>
							<BoundaryType>ToTele</BoundaryType>
							<BoundaryOffset>0.25</BoundaryOffset>
						</IrCutFilterAutoAdjustment>
					</Extension>
				</Extension>
			</ImagingSettings>
		</GetImagingSettingsResponse>
	`

	response := &GetImagingSettingsResponse{}
	err := xml.Unmarshal([]byte(responseData), response)
	require.NoError(t, err)

	adjustments := response.ImagingSettings.Extension.Extension.IrCutFilterAutoAdjustment
	require.Len(t, adjustments, 2)
	assert.Equal(t, "Common", string(adjustments[0].BoundaryType))
	assert.Equal(t, "ToTele", string(adjustments[1].BoundaryType))
}

// Rectangle's bounds are attributes -- guard against a regression to elements.
func TestUnmarshalExposureWindowRectangleAttrs(t *testing.T) {
	responseData := `
		<GetImagingSettingsResponse>
			<ImagingSettings>
				<Exposure>
					<Mode>AUTO</Mode>
					<Window bottom="-1" top="1" right="1" left="-1"/>
				</Exposure>
			</ImagingSettings>
		</GetImagingSettingsResponse>
	`

	response := &GetImagingSettingsResponse{}
	err := xml.Unmarshal([]byte(responseData), response)
	require.NoError(t, err)

	window := response.ImagingSettings.Exposure.Window
	assert.Equal(t, float64(-1), window.Bottom)
	assert.Equal(t, float64(1), window.Top)
	assert.Equal(t, float64(1), window.Right)
	assert.Equal(t, float64(-1), window.Left)
}
