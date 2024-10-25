package event

import (
	"testing"

	"github.com/nbio/xml"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/secure-passage/onvif/xsd"
)

var eventPropertiesData = []byte(`
        <tev:GetEventPropertiesResponse>
            <tev:TopicNamespaceLocation>http://www.onvif.org/onvif/ver10/topics/topicns.xml</tev:TopicNamespaceLocation>
            <wsnt:FixedTopicSet>true</wsnt:FixedTopicSet>
            <wstop:TopicSet>
                <tns1:UserAlarm wstop:topic="true">
                    <tnshik:IllegalAccess wstop:topic="true"/>
                </tns1:UserAlarm>
                <tns1:RuleEngine wstop:topic="true">
                    <CellMotionDetector wstop:topic="true">
                        <Motion wstop:topic="true">
                            <tt:MessageDescription IsProperty="true">
                                <tt:Source>
                                    <tt:SimpleItemDescription Name="VideoSourceConfigurationToken" Type="tt:ReferenceToken"/>
                                    <tt:SimpleItemDescription Name="VideoAnalyticsConfigurationToken" Type="tt:ReferenceToken"/>
                                    <tt:SimpleItemDescription Name="Rule" Type="xs:string"/>
                                </tt:Source>
                                <tt:Data>
                                    <tt:SimpleItemDescription Name="IsMotion" Type="xs:boolean"/>
                                </tt:Data>
                            </tt:MessageDescription>
                        </Motion>
                    </CellMotionDetector>
                    <TamperDetector wstop:topic="true">
                        <Tamper wstop:topic="true">
                            <tt:MessageDescription IsProperty="true">
                                <tt:Source>
                                    <tt:SimpleItemDescription Name="VideoSourceConfigurationToken" Type="tt:ReferenceToken"/>
                                    <tt:SimpleItemDescription Name="VideoAnalyticsConfigurationToken" Type="tt:ReferenceToken"/>
                                    <tt:SimpleItemDescription Name="Rule" Type="xs:string"/>
                                </tt:Source>
                                <tt:Data>
                                    <tt:SimpleItemDescription Name="IsTamper" Type="xs:boolean"/>
                                </tt:Data>
                            </tt:MessageDescription>
                        </Tamper>
                    </TamperDetector>
                </tns1:RuleEngine>
            </wstop:TopicSet>
            <tev:MessageContentFilterDialect>http://www.onvif.org/ver10/tev/messageContentFilter/ItemFilter</tev:MessageContentFilterDialect>
            <tev:MessageContentSchemaLocation>http://www.onvif.org/onvif/ver10/schema/onvif.xsd</tev:MessageContentSchemaLocation>
        </tev:GetEventPropertiesResponse>
`)

var eventRenewResponse = []byte(`
        <tev:RenewResponse>
            <wsnt:TerminationTime>2023-11-25T18:08:15Z</wsnt:TerminationTime>
            <wsnt:CurrentTime>2023-11-24T14:50:25Z</wsnt:CurrentTime>
        </tev:RenewResponse>
`)

func TestEventPropertiesUnmarshalXML(t *testing.T) {
	res := GetEventPropertiesResponse{}
	err := xml.Unmarshal(eventPropertiesData, &res)
	require.NoError(t, err)
	assert.Equal(t, FixedTopicSet(true), *res.FixedTopicSet)
	assert.Equal(t, xsd.AnyURI("http://www.onvif.org/ver10/tev/messageContentFilter/ItemFilter"), *res.MessageContentFilterDialect)
	assert.Equal(t, xsd.AnyURI("http://www.onvif.org/onvif/ver10/schema/onvif.xsd"), *res.MessageContentSchemaLocation)
	userAlarm, exists := map[string]any(*res.TopicSet)["tns1:UserAlarm"]
	assert.True(t, exists)
	_, exists = (userAlarm).(map[string]any)["tnshik:IllegalAccess"]
	assert.True(t, exists)
	ruleEngine, exists := map[string]any(*res.TopicSet)["tns1:RuleEngine"]
	assert.True(t, exists)
	tamperDetector, exists := (ruleEngine).(map[string]any)["TamperDetector"]
	assert.True(t, exists)
	tamper, exists := (tamperDetector).(map[string]any)["Tamper"]
	assert.True(t, exists)
	_, exists = (tamper).(map[string]any)["tt:MessageDescription"]
	assert.True(t, exists)
}

func TestRenewFunction_Response(t *testing.T) {
	res := RenewResponse{}
	err := xml.Unmarshal(eventRenewResponse, &res)
	require.NoError(t, err)
	assert.NotNil(t, res.TerminationTime)
	assert.NotNil(t, res.CurrentTime)
	assert.Equal(t, xsd.String("2023-11-25T18:08:15Z"), *res.TerminationTime)
	assert.Equal(t, xsd.String("2023-11-24T14:50:25Z"), *res.CurrentTime)
}
