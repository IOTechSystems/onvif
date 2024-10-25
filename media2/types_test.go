package media2

import (
	"fmt"
	"testing"

	"github.com/nbio/xml"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/secure-passage/onvif/xsd"
	"github.com/secure-passage/onvif/xsd/onvif"
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
