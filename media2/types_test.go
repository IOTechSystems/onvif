package media2

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/IOTechSystems/onvif/xsd"
	"github.com/IOTechSystems/onvif/xsd/onvif"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnmarshalGetProfilesResponse(t *testing.T) {
	profile1Name := "H26x_L1S1"
	profile1Token := "profile_1"
	profile1Fixed := false
	profile2Name := "JPEG_L1S3"
	profile2Token := "profile_2"
	profile2Fixed := true
	GetProfilesResponseData := fmt.Sprintf(`
		<tms:GetProfilesResponse>
			<tms:Profiles token="%s" fixed="%t"><tms:Name>%s</tms:Name></tms:Profiles>
			<tms:Profiles token="%s" fixed="%t"><tms:Name>%s</tms:Name></tms:Profiles>
		</tms:GetProfilesResponse>
	`, profile1Token, profile1Fixed, profile1Name, profile2Token, profile2Fixed, profile2Name)

	getProfilesResponse := &GetProfilesResponse{}
	err := xml.Unmarshal([]byte(GetProfilesResponseData), getProfilesResponse)
	require.NoError(t, err)

	assert.Equal(t, getProfilesResponse.Profiles[0].Token, profile1Token)
	assert.Equal(t, getProfilesResponse.Profiles[0].Fixed, profile1Fixed)
	assert.Equal(t, getProfilesResponse.Profiles[0].Name, profile1Name)
	assert.Equal(t, getProfilesResponse.Profiles[1].Token, profile2Token)
	assert.Equal(t, getProfilesResponse.Profiles[1].Fixed, profile2Fixed)
	assert.Equal(t, getProfilesResponse.Profiles[1].Name, profile2Name)
}

func TestUnmarshalGetAnalyticsConfigurationsResponse(t *testing.T) {
	configToken := onvif.ReferenceToken("token_1")
	configName := onvif.Name("Analytics_1")
	useCount := 0
	analyticsModuleName := "Viproc"
	analyticsModuleType := "tt:Viproc"
	analyticsModuleItemName := "AnalysisType"
	analyticsModuleItemValue := "Intelligent Video Analytics"
	ruleName := "The Min ObjectHeight"
	ruleType := "tt:ObjectInField"
	ruleItemName := "MaxObjectHeight"
	ruleItemValue := "100"

	responseData := fmt.Sprintf(`
		<tms:GetAnalyticsConfigurationsResponse>
			<tms:Configurations token="%s">
				<tt:Name>%s</tt:Name>
				<tt:UseCount>%d</tt:UseCount>
				<tt:AnalyticsEngineConfiguration>
					<tt:AnalyticsModule Name="%s" Type="%s">
						<tt:Parameters>
							<tt:SimpleItem Name="%s" Value="%s"></tt:SimpleItem>
						</tt:Parameters>
					</tt:AnalyticsModule>
				</tt:AnalyticsEngineConfiguration>
				<tt:RuleEngineConfiguration>
					<tt:Rule Name="%s" Type="%s">
						<tt:Parameters>
							<tt:SimpleItem Name="%s" Value="%s"></tt:SimpleItem>
						</tt:Parameters>
					</tt:Rule>
				</tt:RuleEngineConfiguration>
			</tms:Configurations>
		</tms:GetAnalyticsConfigurationsResponse>
	`, configToken, configName, useCount, analyticsModuleName, analyticsModuleType, analyticsModuleItemName, analyticsModuleItemValue,
		ruleName, ruleType, ruleItemName, ruleItemValue)

	response := &GetAnalyticsConfigurationsResponse{}
	err := xml.Unmarshal([]byte(responseData), response)
	require.NoError(t, err)

	assert.Equal(t, response.Configurations[0].Token, configToken)
	assert.Equal(t, response.Configurations[0].Name, configName)
	assert.Equal(t, response.Configurations[0].AnalyticsEngineConfiguration.AnalyticsModule[0].Name, analyticsModuleName)
	assert.Equal(t, response.Configurations[0].AnalyticsEngineConfiguration.AnalyticsModule[0].Type, analyticsModuleType)
	assert.Equal(t, response.Configurations[0].AnalyticsEngineConfiguration.AnalyticsModule[0].Parameters.SimpleItem[0].Name, analyticsModuleItemName)
	assert.Equal(t, response.Configurations[0].AnalyticsEngineConfiguration.AnalyticsModule[0].Parameters.SimpleItem[0].Value, analyticsModuleItemValue)
	assert.Equal(t, response.Configurations[0].RuleEngineConfiguration.Rule[0].Name, ruleName)
	assert.Equal(t, response.Configurations[0].RuleEngineConfiguration.Rule[0].Type, ruleType)
	assert.Equal(t, response.Configurations[0].RuleEngineConfiguration.Rule[0].Parameters.SimpleItem[0].Name, ruleItemName)
	assert.Equal(t, response.Configurations[0].RuleEngineConfiguration.Rule[0].Parameters.SimpleItem[0].Value, ruleItemValue)
}

func TestMarshalAddConfigurationRequest(t *testing.T) {
	analyticsType := xsd.String("Analytics")
	analyticsToken := xsd.String("AnalyticsToken")
	request := AddConfiguration{
		ProfileToken: "profile_1",
		Configuration: []Configuration{
			{
				Type:  analyticsType,
				Token: &analyticsToken,
			},
		},
	}
	expected := fmt.Sprintf("<tr2:AddConfiguration><tr2:ProfileToken>%s</tr2:ProfileToken><tr2:Configuration><tr2:Type>%s</tr2:Type><tr2:Token>%s</tr2:Token></tr2:Configuration></tr2:AddConfiguration>",
		request.ProfileToken, request.Configuration[0].Type, *request.Configuration[0].Token)

	data, err := xml.Marshal(request)
	require.NoError(t, err)

	assert.Equal(t, expected, string(data))

}

func TestMarshalGetVideoEncoderConfigurationsRequest(t *testing.T) {
	request := GetVideoEncoderConfigurations{
		ConfigurationToken: "enc_token_1",
	}
	expected := fmt.Sprintf("<tr2:GetVideoEncoderConfigurations><tr2:ConfigurationToken>%s</tr2:ConfigurationToken></tr2:GetVideoEncoderConfigurations>",
		request.ConfigurationToken)

	data, err := xml.Marshal(request)
	require.NoError(t, err)
	assert.Equal(t, expected, string(data))
}

func TestMarshalSetVideoEncoderConfigurationRequest(t *testing.T) {
	encoding := xsd.String("H265")
	govLength := xsd.Int(30)
	profile := xsd.String("Main")
	width := xsd.Int(1920)
	height := xsd.Int(1080)
	frameRate := xsd.Float(25)

	constantBitRate := xsd.Boolean(false)

	request := SetVideoEncoderConfiguration{
		Configuration: onvif.VideoEncoder2ConfigurationRequest{
			ConfigurationEntityRequest: onvif.ConfigurationEntityRequest{
				Token: "enc_token_1",
			},
			Encoding: &encoding,
			Resolution: &onvif.VideoResolutionRequest{
				Width:  &width,
				Height: &height,
			},
			RateControl: &onvif.VideoRateControl2Request{
				FrameRateLimit:  &frameRate,
				ConstantBitRate: &constantBitRate,
			},
			GovLength: &govLength,
			Profile:   &profile,
		},
	}
	expected := fmt.Sprintf(`<tr2:SetVideoEncoderConfiguration><tr2:Configuration token="%s" GovLength="%d" Profile="%s"><onvif:Encoding>%s</onvif:Encoding><onvif:Resolution><onvif:Width>%d</onvif:Width><onvif:Height>%d</onvif:Height></onvif:Resolution><onvif:RateControl ConstantBitRate="%t"><onvif:FrameRateLimit>%g</onvif:FrameRateLimit></onvif:RateControl></tr2:Configuration></tr2:SetVideoEncoderConfiguration>`,
		request.Configuration.Token,
		*request.Configuration.GovLength,
		*request.Configuration.Profile,
		*request.Configuration.Encoding,
		*request.Configuration.Resolution.Width,
		*request.Configuration.Resolution.Height,
		*request.Configuration.RateControl.ConstantBitRate,
		*request.Configuration.RateControl.FrameRateLimit)

	data, err := xml.Marshal(request)
	require.NoError(t, err)
	assert.Equal(t, expected, string(data))
}

func TestMarshalGetVideoEncoderConfigurationOptionsRequest(t *testing.T) {
	request := GetVideoEncoderConfigurationOptions{
		ProfileToken: "profile_1",
	}
	expected := fmt.Sprintf("<tr2:GetVideoEncoderConfigurationOptions><tr2:ProfileToken>%s</tr2:ProfileToken></tr2:GetVideoEncoderConfigurationOptions>",
		request.ProfileToken)

	data, err := xml.Marshal(request)
	require.NoError(t, err)
	assert.Equal(t, expected, string(data))
}

func TestUnmarshalGetVideoEncoderConfigurationsResponse(t *testing.T) {
	config1Ref := "enc_1"
	config1Name := onvif.Name("Video Encoder 1")
	config1Encoding := xsd.String("H265")
	config1Width := xsd.Int(1920)
	config1Height := xsd.Int(1080)
	config1GovLength := xsd.Int(30)
	config1Profile := xsd.String("Main")
	config1ConstantBitRate := xsd.Boolean(false)
	config2Ref := "enc_2"
	config2Encoding := xsd.String("H264")

	responseData := fmt.Sprintf(`
		<tr2:GetVideoEncoderConfigurationsResponse>
			<tr2:Configurations token="%s" GovLength="%d" Profile="%s">
				<tt:Name>%s</tt:Name>
				<tt:Encoding>%s</tt:Encoding>
				<tt:Resolution>
					<tt:Width>%d</tt:Width>
					<tt:Height>%d</tt:Height>
				</tt:Resolution>
				<tt:RateControl ConstantBitRate="%t"></tt:RateControl>
			</tr2:Configurations>
			<tr2:Configurations token="%s">
				<tt:Encoding>%s</tt:Encoding>
			</tr2:Configurations>
		</tr2:GetVideoEncoderConfigurationsResponse>
	`, config1Ref, config1GovLength, config1Profile, config1Name, config1Encoding, config1Width, config1Height, config1ConstantBitRate,
		config2Ref, config2Encoding)

	response := &GetVideoEncoderConfigurationsResponse{}
	err := xml.Unmarshal([]byte(responseData), response)
	require.NoError(t, err)

	require.Len(t, response.Configurations, 2)
	assert.Equal(t, response.Configurations[0].Token, onvif.ReferenceToken(config1Ref))
	assert.Equal(t, response.Configurations[0].Name, config1Name)
	assert.Equal(t, response.Configurations[0].Encoding, &config1Encoding)
	require.NotNil(t, response.Configurations[0].Resolution)
	assert.Equal(t, response.Configurations[0].Resolution.Width, &config1Width)
	assert.Equal(t, response.Configurations[0].Resolution.Height, &config1Height)
	assert.Equal(t, response.Configurations[0].GovLength, &config1GovLength)
	assert.Equal(t, response.Configurations[0].Profile, &config1Profile)
	require.NotNil(t, response.Configurations[0].RateControl)
	assert.Equal(t, response.Configurations[0].RateControl.ConstantBitRate, &config1ConstantBitRate)
	assert.Equal(t, response.Configurations[1].Token, onvif.ReferenceToken(config2Ref))
	assert.Equal(t, response.Configurations[1].Encoding, &config2Encoding)
}

func TestUnmarshalGetVideoEncoderConfigurationOptionsResponse(t *testing.T) {
	encoding := xsd.String("H265")
	width := xsd.Int(1920)
	height := xsd.Int(1080)

	// tt:VideoEncoder2ConfigurationOptions declares only four child ELEMENTS
	// (Encoding, QualityRange, ResolutionsAvailable, BitrateRange). GovLengthRange,
	// ProfilesSupported, FrameRatesSupported and ConstantBitRateSupported are
	// ATTRIBUTES, the list-valued ones being whitespace-separated xs:list values.
	// This is the wire format a real device emits.
	responseData := fmt.Sprintf(`
		<tr2:GetVideoEncoderConfigurationOptionsResponse>
			<tr2:Options GovLengthRange="1 300" ProfilesSupported="Main Main10"
			             FrameRatesSupported="30 25 12.5" ConstantBitRateSupported="true">
				<tt:Encoding>%s</tt:Encoding>
				<tt:ResolutionsAvailable>
					<tt:Width>%d</tt:Width>
					<tt:Height>%d</tt:Height>
				</tt:ResolutionsAvailable>
			</tr2:Options>
		</tr2:GetVideoEncoderConfigurationOptionsResponse>
	`, encoding, width, height)

	response := &GetVideoEncoderConfigurationOptionsResponse{}
	err := xml.Unmarshal([]byte(responseData), response)
	require.NoError(t, err)

	require.Len(t, response.Options, 1)
	options := response.Options[0]
	assert.Equal(t, options.Encoding, &encoding)
	require.Len(t, options.ResolutionsAvailable, 1)
	assert.Equal(t, options.ResolutionsAvailable[0].Width, &width)
	assert.Equal(t, options.ResolutionsAvailable[0].Height, &height)

	assert.Equal(t, onvif.IntAttrList{1, 300}, options.GovLengthRange)
	assert.Equal(t, onvif.StringAttrList{"Main", "Main10"}, options.ProfilesSupported)
	assert.Equal(t, onvif.FloatAttrList{30, 25, 12.5}, options.FrameRatesSupported)

	require.NotNil(t, options.ConstantBitRateSupported)
	assert.Equal(t, xsd.Boolean(true), *options.ConstantBitRateSupported)
}

// tt:VideoRateControl2 has no EncodingInterval (that is the ver10
// VideoRateControl) and does have an optional AverageBitRate. Sending
// EncodingInterval draws a fault from validating cameras, and the missing
// AverageBitRate meant a get-modify-set silently wiped the configured value.
func TestUnmarshalVideoRateControl2AverageBitRate(t *testing.T) {
	responseData := `
		<tr2:GetVideoEncoderConfigurationsResponse>
			<tr2:Configurations token="enc_1">
				<tt:RateControl ConstantBitRate="false">
					<tt:FrameRateLimit>25</tt:FrameRateLimit>
					<tt:BitrateLimit>4096</tt:BitrateLimit>
					<tt:AverageBitRate>2048</tt:AverageBitRate>
				</tt:RateControl>
			</tr2:Configurations>
		</tr2:GetVideoEncoderConfigurationsResponse>
	`

	response := &GetVideoEncoderConfigurationsResponse{}
	err := xml.Unmarshal([]byte(responseData), response)
	require.NoError(t, err)

	require.Len(t, response.Configurations, 1)
	rateControl := response.Configurations[0].RateControl
	require.NotNil(t, rateControl)
	require.NotNil(t, rateControl.AverageBitRate)
	assert.Equal(t, xsd.Int(2048), *rateControl.AverageBitRate)
	require.NotNil(t, rateControl.BitrateLimit)
	assert.Equal(t, xsd.Int(4096), *rateControl.BitrateLimit)
}

func TestMarshalVideoRateControl2RequestAverageBitRate(t *testing.T) {
	frameRate := xsd.Float(25)
	bitrateLimit := xsd.Int(4096)
	averageBitRate := xsd.Int(2048)

	request := SetVideoEncoderConfiguration{
		Configuration: onvif.VideoEncoder2ConfigurationRequest{
			ConfigurationEntityRequest: onvif.ConfigurationEntityRequest{Token: "enc_1"},
			RateControl: &onvif.VideoRateControl2Request{
				FrameRateLimit: &frameRate,
				BitrateLimit:   &bitrateLimit,
				AverageBitRate: &averageBitRate,
			},
		},
	}

	data, err := xml.Marshal(request)
	require.NoError(t, err)

	assert.Contains(t, string(data),
		`<onvif:RateControl><onvif:FrameRateLimit>25</onvif:FrameRateLimit>`+
			`<onvif:BitrateLimit>4096</onvif:BitrateLimit>`+
			`<onvif:AverageBitRate>2048</onvif:AverageBitRate></onvif:RateControl>`)
	assert.NotContains(t, string(data), `EncodingInterval`)
}

// GuaranteedFrameRate and AnchorFrameDistance are attributes of
// tt:VideoEncoder2Configuration. Go drops unknown attributes silently, so
// without these fields a get-modify-set erased whatever the camera had set.
func TestUnmarshalVideoEncoder2ConfigurationExtraAttrs(t *testing.T) {
	responseData := `
		<tr2:GetVideoEncoderConfigurationsResponse>
			<tr2:Configurations token="enc_1" GovLength="30" Profile="Main"
			                    GuaranteedFrameRate="true" AnchorFrameDistance="5">
				<tt:Encoding>H264</tt:Encoding>
			</tr2:Configurations>
		</tr2:GetVideoEncoderConfigurationsResponse>
	`

	response := &GetVideoEncoderConfigurationsResponse{}
	err := xml.Unmarshal([]byte(responseData), response)
	require.NoError(t, err)

	require.Len(t, response.Configurations, 1)
	config := response.Configurations[0]

	require.NotNil(t, config.GuaranteedFrameRate)
	assert.Equal(t, xsd.Boolean(true), *config.GuaranteedFrameRate)
	require.NotNil(t, config.AnchorFrameDistance)
	assert.Equal(t, xsd.Int(5), *config.AnchorFrameDistance)
}

// tr2:ConfigurationRef declares Type as required (minOccurs defaults to 1) and
// Token as minOccurs="0". Type must therefore always be on the wire.
func TestMarshalAddConfigurationAlwaysSendsType(t *testing.T) {
	request := AddConfiguration{
		ProfileToken: "profile_1",
		Configuration: []Configuration{
			{Type: "VideoEncoder"},
		},
	}

	data, err := xml.Marshal(request)
	require.NoError(t, err)

	assert.Contains(t, string(data), `<tr2:Type>VideoEncoder</tr2:Type>`)
	assert.NotContains(t, string(data), `<tr2:Token>`)
}

func TestMarshalRemoveConfigurationRequest(t *testing.T) {
	analyticsType := xsd.String("Analytics")
	analyticsToken := xsd.String("AnalyticsToken")
	request := RemoveConfiguration{
		ProfileToken: "profile_1",
		Configuration: []Configuration{
			{
				Type:  analyticsType,
				Token: &analyticsToken,
			},
		},
	}
	expected := fmt.Sprintf("<tr2:RemoveConfiguration><tr2:ProfileToken>%s</tr2:ProfileToken><tr2:Configuration><tr2:Type>%s</tr2:Type><tr2:Token>%s</tr2:Token></tr2:Configuration></tr2:RemoveConfiguration>",
		request.ProfileToken, request.Configuration[0].Type, *request.Configuration[0].Token)

	data, err := xml.Marshal(request)
	require.NoError(t, err)

	assert.Equal(t, expected, string(data))
}

func TestUnmarshalGetStreamUriResponse(t *testing.T) {
	uri := "rtsp://192.168.1.64:554/live1s1.sdp"
	getStreamUriResponseData := fmt.Sprintf(`
		<tr2:GetStreamUriResponse xmlns:tr2="http://www.onvif.org/ver20/media/wsdl">
			<tr2:Uri>%s</tr2:Uri>
		</tr2:GetStreamUriResponse>
	`, uri)

	getStreamUriResponse := &GetStreamUriResponse{}
	err := xml.Unmarshal([]byte(getStreamUriResponseData), getStreamUriResponse)
	require.NoError(t, err)

	assert.Equal(t, uri, getStreamUriResponse.Uri)
}

func TestMarshalGetStreamUri(t *testing.T) {
	request := GetStreamUri{
		Protocol:     "RTSP",
		ProfileToken: onvif.ReferenceToken("Profilee8a6"),
	}

	output, err := xml.Marshal(request)
	require.NoError(t, err)

	assert.Equal(t, `<tr2:GetStreamUri><tr2:Protocol>RTSP</tr2:Protocol><tr2:ProfileToken>Profilee8a6</tr2:ProfileToken></tr2:GetStreamUri>`, string(output))
}

func TestMarshalGetProfiles(t *testing.T) {
	request := GetProfiles{}
	output, err := xml.Marshal(request)
	require.NoError(t, err)
	assert.Equal(t, `<tr2:GetProfiles></tr2:GetProfiles>`, string(output))
}

func TestMarshalGetProfilesWithTypeAndToken(t *testing.T) {
	request := GetProfiles{
		Token: onvif.ReferenceToken("Profilee8a6"),
		Type:  []xsd.String{"All", "VideoEncoder"},
	}
	output, err := xml.Marshal(request)
	require.NoError(t, err)
	assert.Equal(t, `<tr2:GetProfiles><tr2:Token>Profilee8a6</tr2:Token><tr2:Type>All</tr2:Type><tr2:Type>VideoEncoder</tr2:Type></tr2:GetProfiles>`, string(output))
}

func TestUnmarshalGetProfilesResponseWithConfigurations(t *testing.T) {
	profileToken := "Profilee8a6"
	encoderToken := "VideoEnc_1"
	width := 1920
	height := 1080
	GetProfilesResponseData := fmt.Sprintf(`
		<tr2:GetProfilesResponse>
			<tr2:Profiles token="%s" fixed="false">
				<tr2:Name>MainStream</tr2:Name>
				<tr2:Configurations>
					<tr2:VideoEncoder token="%s">
						<tt:Encoding>H264</tt:Encoding>
						<tt:Resolution>
							<tt:Width>%d</tt:Width>
							<tt:Height>%d</tt:Height>
						</tt:Resolution>
					</tr2:VideoEncoder>
				</tr2:Configurations>
			</tr2:Profiles>
		</tr2:GetProfilesResponse>
	`, profileToken, encoderToken, width, height)

	getProfilesResponse := &GetProfilesResponse{}
	err := xml.Unmarshal([]byte(GetProfilesResponseData), getProfilesResponse)
	require.NoError(t, err)

	require.Len(t, getProfilesResponse.Profiles, 1)
	profile := getProfilesResponse.Profiles[0]
	assert.Equal(t, profileToken, profile.Token)
	require.NotNil(t, profile.Configurations)
	require.NotNil(t, profile.Configurations.VideoEncoder)
	assert.Equal(t, onvif.ReferenceToken(encoderToken), profile.Configurations.VideoEncoder.Token)
	require.NotNil(t, profile.Configurations.VideoEncoder.Resolution)
	assert.Equal(t, width, int(*profile.Configurations.VideoEncoder.Resolution.Width))
	assert.Equal(t, height, int(*profile.Configurations.VideoEncoder.Resolution.Height))
}
