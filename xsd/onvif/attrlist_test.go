package onvif

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The ONVIF schema models several attributes as whitespace-separated xs:list
// values, e.g.
//
//	<xs:simpleType name="StringAttrList"><xs:list itemType="xs:string"/></xs:simpleType>
//
// A bare Go slice has no attribute codec, so encoding/xml neither splits the
// value on unmarshal nor joins it on marshal -- it emits one repeated attribute
// per item, which is not even well-formed XML. These types therefore need
// explicit MarshalXMLAttr/UnmarshalXMLAttr implementations.

type stringAttrListHolder struct {
	XMLName xml.Name       `xml:"Holder"`
	List    StringAttrList `xml:"FormatsSupported,attr,omitempty"`
}

func TestUnmarshalStringAttrListAttr(t *testing.T) {
	holder := &stringAttrListHolder{}
	err := xml.Unmarshal([]byte(`<Holder FormatsSupported="MP4 MKV"/>`), holder)
	require.NoError(t, err)

	require.Len(t, holder.List, 2)
	assert.Equal(t, "MP4", holder.List[0])
	assert.Equal(t, "MKV", holder.List[1])
}

func TestMarshalStringAttrListAttr(t *testing.T) {
	data, err := xml.Marshal(&stringAttrListHolder{List: StringAttrList{"MP4", "MKV"}})
	require.NoError(t, err)

	assert.Equal(t, `<Holder FormatsSupported="MP4 MKV"></Holder>`, string(data))
}

func TestMarshalStringAttrListAttrOmitsEmpty(t *testing.T) {
	data, err := xml.Marshal(&stringAttrListHolder{})
	require.NoError(t, err)

	assert.Equal(t, `<Holder></Holder>`, string(data))
}

type intAttrListHolder struct {
	XMLName xml.Name    `xml:"Holder"`
	List    IntAttrList `xml:"GovLengthRange,attr,omitempty"`
}

func TestUnmarshalIntAttrListAttr(t *testing.T) {
	holder := &intAttrListHolder{}
	err := xml.Unmarshal([]byte(`<Holder GovLengthRange="1 300"/>`), holder)
	require.NoError(t, err)

	require.Len(t, holder.List, 2)
	assert.Equal(t, 1, holder.List[0])
	assert.Equal(t, 300, holder.List[1])
}

func TestMarshalIntAttrListAttr(t *testing.T) {
	data, err := xml.Marshal(&intAttrListHolder{List: IntAttrList{1, 300}})
	require.NoError(t, err)

	assert.Equal(t, `<Holder GovLengthRange="1 300"></Holder>`, string(data))
}

type floatAttrListHolder struct {
	XMLName xml.Name      `xml:"Holder"`
	List    FloatAttrList `xml:"FrameRatesSupported,attr,omitempty"`
}

func TestUnmarshalFloatAttrListAttr(t *testing.T) {
	holder := &floatAttrListHolder{}
	err := xml.Unmarshal([]byte(`<Holder FrameRatesSupported="30 25 12.5"/>`), holder)
	require.NoError(t, err)

	require.Len(t, holder.List, 3)
	assert.Equal(t, float64(30), holder.List[0])
	assert.Equal(t, float64(25), holder.List[1])
	assert.Equal(t, 12.5, holder.List[2])
}

func TestMarshalFloatAttrListAttr(t *testing.T) {
	data, err := xml.Marshal(&floatAttrListHolder{List: FloatAttrList{30, 12.5}})
	require.NoError(t, err)

	assert.Equal(t, `<Holder FrameRatesSupported="30 12.5"></Holder>`, string(data))
}

// Whitespace-separated lists may be padded or use any whitespace run.
func TestUnmarshalAttrListCollapsesWhitespace(t *testing.T) {
	holder := &stringAttrListHolder{}
	err := xml.Unmarshal([]byte("<Holder FormatsSupported=\"  MP4\t\tMKV  \"/>"), holder)
	require.NoError(t, err)

	assert.Equal(t, StringAttrList{"MP4", "MKV"}, holder.List)
}
