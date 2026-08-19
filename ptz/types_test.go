package ptz

import (
	"encoding/xml"
	"testing"

	"github.com/IOTechSystems/onvif/xsd"
	"github.com/IOTechSystems/onvif/xsd/onvif"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// tt:PTZSpeed's PanTilt/Zoom live in the ver10 schema namespace, which
// onvif.xsd declares elementFormDefault="qualified". The PTZ service marshals
// PTZSpeed outbound as ContinuousMove/Velocity, GotoPreset/Speed,
// GotoHomePosition/Speed and GeoMove/Speed, so these children must be
// onvif:-qualified -- unqualified means a strict device faults, and a lenient
// one ignores the requested velocity and the camera simply does not move.
func TestMarshalContinuousMoveVelocityIsQualified(t *testing.T) {
	profileToken := onvif.ReferenceToken("profile_1")

	request := ContinuousMove{
		ProfileToken: &profileToken,
		Velocity: &onvif.PTZSpeedRequest{
			PanTilt: &onvif.Vector2D{X: 0.5, Y: -0.5},
		},
	}

	data, err := xml.Marshal(request)
	require.NoError(t, err)

	assert.Equal(t,
		`<tptz:ContinuousMove>`+
			`<tptz:ProfileToken>profile_1</tptz:ProfileToken>`+
			`<tptz:Velocity><onvif:PanTilt x="0.5" y="-0.5"></onvif:PanTilt></tptz:Velocity>`+
			`</tptz:ContinuousMove>`,
		string(data))
}

func TestMarshalGotoPresetSpeedIsQualified(t *testing.T) {
	profileToken := onvif.ReferenceToken("profile_1")
	presetToken := onvif.ReferenceToken("preset_1")

	request := GotoPreset{
		ProfileToken: &profileToken,
		PresetToken:  &presetToken,
		Speed: &onvif.PTZSpeedRequest{
			Zoom: &onvif.Vector1D{X: 0.25},
		},
	}

	data, err := xml.Marshal(request)
	require.NoError(t, err)

	assert.Contains(t, string(data), `<tptz:Speed><onvif:Zoom x="0.25"></onvif:Zoom></tptz:Speed>`)
	assert.NotContains(t, string(data), `<Zoom`)
}

func TestMarshalGotoHomePositionSpeedIsQualified(t *testing.T) {
	profileToken := onvif.ReferenceToken("profile_1")

	request := GotoHomePosition{
		ProfileToken: &profileToken,
		Speed: &onvif.PTZSpeedRequest{
			PanTilt: &onvif.Vector2D{X: 1, Y: 1},
		},
	}

	data, err := xml.Marshal(request)
	require.NoError(t, err)

	assert.Contains(t, string(data), `<onvif:PanTilt x="1" y="1">`)
	assert.NotContains(t, string(data), `<PanTilt`)
}

// tt:PTZConfiguration extends tt:ConfigurationEntity. SetConfiguration marshals
// it outbound, so Name/UseCount/NodeToken must be onvif:-qualified. They were
// tagged tptz:, which was wrong in the other direction -- the namespace half of
// a tag is enforced on unmarshal, so reads silently returned zero values.
func TestMarshalSetConfigurationIsQualified(t *testing.T) {
	nodeToken := onvif.ReferenceToken("node_0")

	request := SetConfiguration{
		PTZConfiguration: onvif.PTZConfigurationRequest{
			Token:     "ptz_0",
			Name:      "MyPTZ",
			UseCount:  1,
			NodeToken: &nodeToken,
			DefaultPTZSpeed: &onvif.PTZSpeedRequest{
				PanTilt: &onvif.Vector2D{X: 0.5, Y: 0.5},
			},
		},
		ForcePersistence: xsd.Boolean(true),
	}

	data, err := xml.Marshal(request)
	require.NoError(t, err)

	assert.Contains(t, string(data), `<tptz:PTZConfiguration token="ptz_0">`)
	assert.Contains(t, string(data), `<onvif:Name>MyPTZ</onvif:Name>`)
	assert.Contains(t, string(data), `<onvif:UseCount>1</onvif:UseCount>`)
	assert.Contains(t, string(data), `<onvif:NodeToken>node_0</onvif:NodeToken>`)
	assert.Contains(t, string(data), `<onvif:DefaultPTZSpeed><onvif:PanTilt x="0.5" y="0.5"></onvif:PanTilt></onvif:DefaultPTZSpeed>`)

	// No unqualified leakage.
	assert.NotContains(t, string(data), `<Name>`)
	assert.NotContains(t, string(data), `<UseCount>`)
	assert.NotContains(t, string(data), `<NodeToken>`)
}

// The read path must keep working: a tt:-namespaced response populates the
// response-side PTZConfiguration.
func TestUnmarshalGetConfigurationResponse(t *testing.T) {
	responseData := `
		<GetConfigurationResponse>
			<PTZConfiguration token="ptz_0">
				<Name>MyPTZ</Name>
				<UseCount>2</UseCount>
				<NodeToken>node_0</NodeToken>
				<DefaultPTZSpeed>
					<PanTilt x="0.5" y="0.25"/>
				</DefaultPTZSpeed>
			</PTZConfiguration>
		</GetConfigurationResponse>
	`

	response := &GetConfigurationResponse{}
	err := xml.Unmarshal([]byte(responseData), response)
	require.NoError(t, err)

	config := response.PTZConfiguration
	assert.Equal(t, onvif.ReferenceToken("ptz_0"), config.Token)
	assert.Equal(t, onvif.Name("MyPTZ"), config.Name)
	assert.Equal(t, 2, config.UseCount)
	require.NotNil(t, config.NodeToken)
	assert.Equal(t, onvif.ReferenceToken("node_0"), *config.NodeToken)
	require.NotNil(t, config.DefaultPTZSpeed)
	require.NotNil(t, config.DefaultPTZSpeed.PanTilt)
	assert.Equal(t, 0.5, config.DefaultPTZSpeed.PanTilt.X)
}
