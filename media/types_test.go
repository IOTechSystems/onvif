package media

import (
	"encoding/xml"
	"testing"

	"github.com/IOTechSystems/onvif/xsd"
	"github.com/IOTechSystems/onvif/xsd/onvif"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ---------------------------------------------------------------------------
// MetadataConfiguration
// ---------------------------------------------------------------------------

// SessionTimeout is its own element in tt:MetadataConfiguration:
//
//	<xs:element name="SessionTimeout" type="xs:duration"/>
//
// It was tagged onvif:CompressionType, so the duration went out under the wrong
// name and the required SessionTimeout element was never sent at all.
func TestMarshalSetMetadataConfigurationSessionTimeout(t *testing.T) {
	sessionTimeout := xsd.Duration("PT60S")

	request := SetMetadataConfiguration{
		Configuration: onvif.MetadataConfigurationRequest{
			SessionTimeout: &sessionTimeout,
		},
	}

	data, err := xml.Marshal(request)
	require.NoError(t, err)

	assert.Contains(t, string(data), `<onvif:SessionTimeout>PT60S</onvif:SessionTimeout>`)
	assert.NotContains(t, string(data), `<onvif:CompressionType>`)
}

// onvif.xsd declares no attributeFormDefault, so it defaults to "unqualified":
// attributes must be emitted without a namespace prefix. This is the same fix
// already applied to VideoEncoder2Configuration's GovLength/Profile.
func TestMarshalSetMetadataConfigurationCompressionTypeIsUnqualifiedAttr(t *testing.T) {
	request := SetMetadataConfiguration{
		Configuration: onvif.MetadataConfigurationRequest{
			CompressionType: "GZIP",
		},
	}

	data, err := xml.Marshal(request)
	require.NoError(t, err)

	assert.Contains(t, string(data), `CompressionType="GZIP"`)
	assert.NotContains(t, string(data), `onvif:CompressionType=`)
}

// tt:MetadataConfiguration extends tt:ConfigurationEntity, whose Name/UseCount
// live in the ver10 schema namespace. Embedding the response-side
// ConfigurationEntity emitted them unqualified (no namespace at all).
func TestMarshalSetMetadataConfigurationQualifiesEntityFields(t *testing.T) {
	request := SetMetadataConfiguration{
		Configuration: onvif.MetadataConfigurationRequest{
			ConfigurationEntityRequest: onvif.ConfigurationEntityRequest{
				Token: "meta0",
				Name:  "MetaCfg",
			},
		},
	}

	data, err := xml.Marshal(request)
	require.NoError(t, err)

	assert.Contains(t, string(data), `token="meta0"`)
	assert.Contains(t, string(data), `<onvif:Name>MetaCfg</onvif:Name>`)
	assert.NotContains(t, string(data), `<Name>`)
}

// tt:PTZFilter declares Status and Position without minOccurs, so both are
// required. With ",omitempty" a request that disables PTZ metadata marshaled an
// empty <PTZStatus></PTZStatus>, which is schema-invalid -- so the camera kept
// emitting PTZ metadata while the call reported success.
func TestMarshalSetMetadataConfigurationPTZFilterSendsFalse(t *testing.T) {
	request := SetMetadataConfiguration{
		Configuration: onvif.MetadataConfigurationRequest{
			PTZStatus: &onvif.PTZFilterRequest{Status: false, Position: false},
		},
	}

	data, err := xml.Marshal(request)
	require.NoError(t, err)

	assert.Contains(t, string(data), `<onvif:Status>false</onvif:Status>`)
	assert.Contains(t, string(data), `<onvif:Position>false</onvif:Position>`)
}

// ---------------------------------------------------------------------------
// AnalyticsEngineConfiguration
// ---------------------------------------------------------------------------

// tt:AnalyticsEngineConfiguration declares
//
//	<xs:element name="AnalyticsModule" type="tt:Config" minOccurs="0" maxOccurs="unbounded"/>
//
// The field was tagged with the name of its own Go struct, so the camera
// received an unknown child element and the module list was never applied.
func TestMarshalAnalyticsEngineConfigurationModuleName(t *testing.T) {
	viprocType := xsd.QName("tt:Viproc")
	otherType := xsd.QName("tt:Other")

	request := SetMetadataConfiguration{
		Configuration: onvif.MetadataConfigurationRequest{
			AnalyticsEngineConfiguration: &onvif.AnalyticsEngineConfigurationRequest{
				AnalyticsModule: []onvif.ConfigRequest{
					{Name: "Viproc", Type: &viprocType},
					{Name: "Second", Type: &otherType},
				},
			},
		},
	}

	data, err := xml.Marshal(request)
	require.NoError(t, err)

	assert.Contains(t, string(data), `<onvif:AnalyticsModule Name="Viproc" Type="tt:Viproc">`)
	assert.Contains(t, string(data), `<onvif:AnalyticsModule Name="Second" Type="tt:Other">`)
	assert.NotContains(t, string(data), `AnalyticsEngineConfigurationRequest`)
}

// ---------------------------------------------------------------------------
// PTZConfiguration (read path)
// ---------------------------------------------------------------------------

// PTZConfiguration is a tt: type, but Name/UseCount/NodeToken were tagged in the
// tptz: namespace and PTZSpeed's children in the onvif: namespace. Go enforces
// the namespace half of a tag on unmarshal, so all of these silently stayed at
// their zero value on every GetProfiles response.
func TestUnmarshalPTZConfiguration(t *testing.T) {
	responseData := `
		<GetProfileResponse>
			<Profile token="profile_1">
				<Name>MainProfile</Name>
				<PTZConfiguration token="ptz_0">
					<Name>MyPTZ</Name>
					<UseCount>2</UseCount>
					<NodeToken>node_0</NodeToken>
					<DefaultPTZSpeed>
						<PanTilt x="0.5" y="0.25"/>
						<Zoom x="0.75"/>
					</DefaultPTZSpeed>
				</PTZConfiguration>
			</Profile>
		</GetProfileResponse>
	`

	response := &GetProfileResponse{}
	err := xml.Unmarshal([]byte(responseData), response)
	require.NoError(t, err)

	ptz := response.Profile.PTZConfiguration
	require.NotNil(t, ptz)
	assert.Equal(t, onvif.ReferenceToken("ptz_0"), ptz.Token)
	assert.Equal(t, onvif.Name("MyPTZ"), ptz.Name)
	assert.Equal(t, 2, ptz.UseCount)

	require.NotNil(t, ptz.NodeToken)
	assert.Equal(t, onvif.ReferenceToken("node_0"), *ptz.NodeToken)

	require.NotNil(t, ptz.DefaultPTZSpeed)
	require.NotNil(t, ptz.DefaultPTZSpeed.PanTilt)
	assert.Equal(t, 0.5, ptz.DefaultPTZSpeed.PanTilt.X)
	assert.Equal(t, 0.25, ptz.DefaultPTZSpeed.PanTilt.Y)
	require.NotNil(t, ptz.DefaultPTZSpeed.Zoom)
	assert.Equal(t, 0.75, ptz.DefaultPTZSpeed.Zoom.X)
}

// ---------------------------------------------------------------------------
// Unbounded response collections
// ---------------------------------------------------------------------------

// media.wsdl declares Configurations as
// minOccurs="0" maxOccurs="unbounded" for every Get*ConfigurationsResponse.
// With a scalar field encoding/xml overwrites in place, so a caller saw only the
// LAST configuration and concluded the device had a single encoder.
func TestUnmarshalGetVideoEncoderConfigurationsResponseKeepsAll(t *testing.T) {
	responseData := `
		<GetVideoEncoderConfigurationsResponse>
			<Configurations token="enc_0"><Name>A</Name></Configurations>
			<Configurations token="enc_1"><Name>B</Name></Configurations>
			<Configurations token="enc_2"><Name>C</Name></Configurations>
		</GetVideoEncoderConfigurationsResponse>
	`

	response := &GetVideoEncoderConfigurationsResponse{}
	err := xml.Unmarshal([]byte(responseData), response)
	require.NoError(t, err)

	require.Len(t, response.Configurations, 3)
	assert.Equal(t, onvif.ReferenceToken("enc_0"), response.Configurations[0].Token)
	assert.Equal(t, onvif.ReferenceToken("enc_1"), response.Configurations[1].Token)
	assert.Equal(t, onvif.ReferenceToken("enc_2"), response.Configurations[2].Token)
}

func TestUnmarshalGetVideoSourcesResponseKeepsAll(t *testing.T) {
	responseData := `
		<GetVideoSourcesResponse>
			<VideoSources token="vs_0"><Framerate>30</Framerate></VideoSources>
			<VideoSources token="vs_1"><Framerate>25</Framerate></VideoSources>
		</GetVideoSourcesResponse>
	`

	response := &GetVideoSourcesResponse{}
	err := xml.Unmarshal([]byte(responseData), response)
	require.NoError(t, err)

	require.Len(t, response.VideoSources, 2)
	assert.Equal(t, onvif.ReferenceToken("vs_0"), response.VideoSources[0].Token)
	assert.Equal(t, onvif.ReferenceToken("vs_1"), response.VideoSources[1].Token)
}

func TestUnmarshalGetVideoSourceConfigurationsResponseKeepsAll(t *testing.T) {
	responseData := `
		<GetVideoSourceConfigurationsResponse>
			<Configurations token="vsc_0"><Name>A</Name></Configurations>
			<Configurations token="vsc_1"><Name>B</Name></Configurations>
		</GetVideoSourceConfigurationsResponse>
	`

	response := &GetVideoSourceConfigurationsResponse{}
	err := xml.Unmarshal([]byte(responseData), response)
	require.NoError(t, err)

	require.Len(t, response.Configurations, 2)
	assert.Equal(t, onvif.ReferenceToken("vsc_0"), response.Configurations[0].Token)
	assert.Equal(t, onvif.ReferenceToken("vsc_1"), response.Configurations[1].Token)
}

func TestUnmarshalGetOSDsResponseKeepsAll(t *testing.T) {
	responseData := `
		<GetOSDsResponse>
			<OSDs token="osd_0"><Type>Text</Type></OSDs>
			<OSDs token="osd_1"><Type>Image</Type></OSDs>
		</GetOSDsResponse>
	`

	response := &GetOSDsResponse{}
	err := xml.Unmarshal([]byte(responseData), response)
	require.NoError(t, err)

	require.Len(t, response.OSDs, 2)
}

// ---------------------------------------------------------------------------
// Request element ordering
// ---------------------------------------------------------------------------

// media.wsdl orders every Get*ConfigurationOptions request as
// <xs:sequence><ConfigurationToken/><ProfileToken/></xs:sequence>.
// Go marshals in field-declaration order, so a reversed struct produces a
// sequence that schema-validating cameras reject.
func TestMarshalGetVideoEncoderConfigurationOptionsOrder(t *testing.T) {
	configToken := onvif.ReferenceToken("enc_0")
	profileToken := onvif.ReferenceToken("profile_1")

	request := GetVideoEncoderConfigurationOptions{
		ConfigurationToken: &configToken,
		ProfileToken:       &profileToken,
	}

	data, err := xml.Marshal(request)
	require.NoError(t, err)

	assert.Equal(t,
		`<trt:GetVideoEncoderConfigurationOptions>`+
			`<trt:ConfigurationToken>enc_0</trt:ConfigurationToken>`+
			`<trt:ProfileToken>profile_1</trt:ProfileToken>`+
			`</trt:GetVideoEncoderConfigurationOptions>`,
		string(data))
}

// Both tokens are minOccurs="0". Querying options for a profile alone must not
// ship an empty <trt:ConfigurationToken/>, which cameras that validate token
// length reject as an invalid reference token.
func TestMarshalGetVideoEncoderConfigurationOptionsOmitsUnsetTokens(t *testing.T) {
	profileToken := onvif.ReferenceToken("profile_1")

	request := GetVideoEncoderConfigurationOptions{ProfileToken: &profileToken}

	data, err := xml.Marshal(request)
	require.NoError(t, err)

	assert.Equal(t,
		`<trt:GetVideoEncoderConfigurationOptions>`+
			`<trt:ProfileToken>profile_1</trt:ProfileToken>`+
			`</trt:GetVideoEncoderConfigurationOptions>`,
		string(data))
}

// CreateProfile's Token is minOccurs="0" -- omitting it asks the device to
// assign one. An empty <trt:Token/> draws an InvalidArgVal fault instead.
func TestMarshalCreateProfileOmitsUnsetToken(t *testing.T) {
	request := CreateProfile{Name: "NewProfile"}

	data, err := xml.Marshal(request)
	require.NoError(t, err)

	assert.Equal(t,
		`<trt:CreateProfile><trt:Name>NewProfile</trt:Name></trt:CreateProfile>`,
		string(data))
}

// GetOSDs.ConfigurationToken is minOccurs="0"; omitting it means "all OSDs".
func TestMarshalGetOSDsOmitsUnsetToken(t *testing.T) {
	request := GetOSDs{}

	data, err := xml.Marshal(request)
	require.NoError(t, err)

	assert.Equal(t, `<trt:GetOSDs></trt:GetOSDs>`, string(data))
}

// ---------------------------------------------------------------------------
// Set* requests must namespace-qualify the configuration body
// ---------------------------------------------------------------------------

// onvif.xsd is elementFormDefault="qualified", so every child of a tt: type must
// carry the ver10 schema namespace. Marshaling the response-side struct emitted
// them in no namespace, which lenient cameras accept and then ignore -- the
// silent no-op that makes a get-modify-set appear to succeed.
func TestMarshalSetAudioEncoderConfigurationQualifiesChildren(t *testing.T) {
	bitrate := xsd.Int(64)
	sampleRate := xsd.Int(8000)

	request := SetAudioEncoderConfiguration{
		Configuration: onvif.AudioEncoderConfigurationRequest{
			ConfigurationEntityRequest: onvif.ConfigurationEntityRequest{
				Token: "aenc_0",
				Name:  "AudioCfg",
			},
			Encoding:   "G711",
			Bitrate:    &bitrate,
			SampleRate: &sampleRate,
		},
	}

	data, err := xml.Marshal(request)
	require.NoError(t, err)

	assert.Contains(t, string(data), `<onvif:Encoding>G711</onvif:Encoding>`)
	assert.Contains(t, string(data), `<onvif:Bitrate>64</onvif:Bitrate>`)
	assert.Contains(t, string(data), `<onvif:SampleRate>8000</onvif:SampleRate>`)
	assert.NotContains(t, string(data), `<Encoding>`)
	assert.NotContains(t, string(data), `<Bitrate>`)
}

func TestMarshalSetAudioSourceConfigurationQualifiesChildren(t *testing.T) {
	request := SetAudioSourceConfiguration{
		Configuration: onvif.AudioSourceConfigurationRequest{
			ConfigurationEntityRequest: onvif.ConfigurationEntityRequest{
				Token: "asrc_0",
				Name:  "AudioSrc",
			},
			SourceToken: "mic_0",
		},
	}

	data, err := xml.Marshal(request)
	require.NoError(t, err)

	assert.Contains(t, string(data), `<onvif:SourceToken>mic_0</onvif:SourceToken>`)
	assert.NotContains(t, string(data), `<SourceToken>`)
}
