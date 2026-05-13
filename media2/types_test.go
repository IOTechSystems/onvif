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
				Type:  &analyticsType,
				Token: &analyticsToken,
			},
		},
	}
	expected := fmt.Sprintf("<tr2:AddConfiguration><tr2:ProfileToken>%s</tr2:ProfileToken><tr2:Configuration><tr2:Type>%s</tr2:Type><tr2:Token>%s</tr2:Token></tr2:Configuration></tr2:AddConfiguration>",
		request.ProfileToken, *request.Configuration[0].Type, *request.Configuration[0].Token)

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
				FrameRateLimit: &frameRate,
			},
			GovLength: &govLength,
			Profile:   &profile,
		},
	}
	expected := fmt.Sprintf(`<tr2:SetVideoEncoderConfiguration><tr2:Configuration token="%s"><onvif:Encoding>%s</onvif:Encoding><onvif:Resolution><onvif:Width>%d</onvif:Width><onvif:Height>%d</onvif:Height></onvif:Resolution><onvif:RateControl><onvif:FrameRateLimit>%g</onvif:FrameRateLimit></onvif:RateControl><onvif:GovLength>%d</onvif:GovLength><onvif:Profile>%s</onvif:Profile></tr2:Configuration></tr2:SetVideoEncoderConfiguration>`,
		request.Configuration.Token,
		*request.Configuration.Encoding,
		*request.Configuration.Resolution.Width,
		*request.Configuration.Resolution.Height,
		*request.Configuration.RateControl.FrameRateLimit,
		*request.Configuration.GovLength,
		*request.Configuration.Profile)

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
	config2Ref := "enc_2"
	config2Encoding := xsd.String("H264")

	responseData := fmt.Sprintf(`
		<tr2:GetVideoEncoderConfigurationsResponse>
			<tr2:Configurations token="%s">
				<tt:Name>%s</tt:Name>
				<tt:Encoding>%s</tt:Encoding>
				<tt:Resolution>
					<tt:Width>%d</tt:Width>
					<tt:Height>%d</tt:Height>
				</tt:Resolution>
				<tt:GovLength>%d</tt:GovLength>
				<tt:Profile>%s</tt:Profile>
			</tr2:Configurations>
			<tr2:Configurations token="%s">
				<tt:Encoding>%s</tt:Encoding>
			</tr2:Configurations>
		</tr2:GetVideoEncoderConfigurationsResponse>
	`, config1Ref, config1Name, config1Encoding, config1Width, config1Height, config1GovLength, config1Profile,
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
	assert.Equal(t, response.Configurations[1].Token, onvif.ReferenceToken(config2Ref))
	assert.Equal(t, response.Configurations[1].Encoding, &config2Encoding)
}

func TestUnmarshalGetVideoEncoderConfigurationOptionsResponse(t *testing.T) {
	encoding := xsd.String("H265")
	width := xsd.Int(1920)
	height := xsd.Int(1080)

	responseData := fmt.Sprintf(`
		<tr2:GetVideoEncoderConfigurationOptionsResponse>
			<tr2:Options>
				<tt:Encoding>%s</tt:Encoding>
				<tt:ResolutionsAvailable>
					<tt:Width>%d</tt:Width>
					<tt:Height>%d</tt:Height>
				</tt:ResolutionsAvailable>
				<tt:GovLengthRange>
					<tt:Min>1</tt:Min>
					<tt:Max>300</tt:Max>
				</tt:GovLengthRange>
				<tt:ProfilesSupported>Main</tt:ProfilesSupported>
				<tt:ProfilesSupported>Main10</tt:ProfilesSupported>
			</tr2:Options>
		</tr2:GetVideoEncoderConfigurationOptionsResponse>
	`, encoding, width, height)

	response := &GetVideoEncoderConfigurationOptionsResponse{}
	err := xml.Unmarshal([]byte(responseData), response)
	require.NoError(t, err)

	require.Len(t, response.Options, 1)
	assert.Equal(t, response.Options[0].Encoding, &encoding)
	require.Len(t, response.Options[0].ResolutionsAvailable, 1)
	assert.Equal(t, response.Options[0].ResolutionsAvailable[0].Width, &width)
	assert.Equal(t, response.Options[0].ResolutionsAvailable[0].Height, &height)
	require.NotNil(t, response.Options[0].GovLengthRange)
	assert.Equal(t, response.Options[0].GovLengthRange.Min, 1)
	assert.Equal(t, response.Options[0].GovLengthRange.Max, 300)
	require.Len(t, response.Options[0].ProfilesSupported, 2)
	assert.Equal(t, response.Options[0].ProfilesSupported[0], xsd.String("Main"))
	assert.Equal(t, response.Options[0].ProfilesSupported[1], xsd.String("Main10"))
}

func TestMarshalRemoveConfigurationRequest(t *testing.T) {
	analyticsType := xsd.String("Analytics")
	analyticsToken := xsd.String("AnalyticsToken")
	request := RemoveConfiguration{
		ProfileToken: "profile_1",
		Configuration: []Configuration{
			{
				Type:  &analyticsType,
				Token: &analyticsToken,
			},
		},
	}
	expected := fmt.Sprintf("<tr2:RemoveConfiguration><tr2:ProfileToken>%s</tr2:ProfileToken><tr2:Configuration><tr2:Type>%s</tr2:Type><tr2:Token>%s</tr2:Token></tr2:Configuration></tr2:RemoveConfiguration>",
		request.ProfileToken, *request.Configuration[0].Type, *request.Configuration[0].Token)

	data, err := xml.Marshal(request)
	require.NoError(t, err)

	assert.Equal(t, expected, string(data))
}
