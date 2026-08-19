package onvif

import (
	"encoding/xml"
	"strconv"
	"strings"
)

/*
The ONVIF schema models a number of attributes as whitespace-separated xs:list
values, for example:

	<xs:simpleType name="StringAttrList"><xs:list itemType="xs:string"/></xs:simpleType>
	<xs:simpleType name="IntList"><xs:list itemType="xs:int"/></xs:simpleType>
	<xs:simpleType name="FloatList"><xs:list itemType="xs:float"/></xs:simpleType>

encoding/xml has no built-in codec for a slice held in an attribute: on unmarshal
it assigns the whole raw value as a single item, and on marshal it emits one
repeated attribute per item, which is not well-formed XML. The types below carry
explicit attribute codecs so a list attribute round trips as the schema defines
it.
*/

// StringAttrList is a whitespace-separated list of strings held in an attribute.
type StringAttrList []string

// MarshalXMLAttr joins the items into a single whitespace-separated attribute.
func (list StringAttrList) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if len(list) == 0 {
		return xml.Attr{}, nil
	}
	return xml.Attr{Name: name, Value: strings.Join(list, " ")}, nil
}

// UnmarshalXMLAttr splits the attribute value on any run of whitespace.
func (list *StringAttrList) UnmarshalXMLAttr(attr xml.Attr) error {
	*list = strings.Fields(attr.Value)
	return nil
}

// IntAttrList is a whitespace-separated list of ints held in an attribute.
type IntAttrList []int

// MarshalXMLAttr joins the items into a single whitespace-separated attribute.
func (list IntAttrList) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if len(list) == 0 {
		return xml.Attr{}, nil
	}
	items := make([]string, len(list))
	for i, item := range list {
		items[i] = strconv.Itoa(item)
	}
	return xml.Attr{Name: name, Value: strings.Join(items, " ")}, nil
}

// UnmarshalXMLAttr splits the attribute value on any run of whitespace.
func (list *IntAttrList) UnmarshalXMLAttr(attr xml.Attr) error {
	fields := strings.Fields(attr.Value)
	items := make(IntAttrList, 0, len(fields))
	for _, field := range fields {
		item, err := strconv.Atoi(field)
		if err != nil {
			return err
		}
		items = append(items, item)
	}
	*list = items
	return nil
}

// FloatAttrList is a whitespace-separated list of floats held in an attribute.
type FloatAttrList []float64

// MarshalXMLAttr joins the items into a single whitespace-separated attribute.
func (list FloatAttrList) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if len(list) == 0 {
		return xml.Attr{}, nil
	}
	items := make([]string, len(list))
	for i, item := range list {
		items[i] = strconv.FormatFloat(item, 'g', -1, 64)
	}
	return xml.Attr{Name: name, Value: strings.Join(items, " ")}, nil
}

// UnmarshalXMLAttr splits the attribute value on any run of whitespace.
func (list *FloatAttrList) UnmarshalXMLAttr(attr xml.Attr) error {
	fields := strings.Fields(attr.Value)
	items := make(FloatAttrList, 0, len(fields))
	for _, field := range fields {
		item, err := strconv.ParseFloat(field, 64)
		if err != nil {
			return err
		}
		items = append(items, item)
	}
	*list = items
	return nil
}
