package models

import . "github.com/ublux/go-models/enums"

type ContactNumber struct {
	Label                     LabelNumber `bson:"label" json:"label"`
	Number                    string      `bson:"number" json:"number"`
	NumberInternationalFormat string      `bson:"numberInternationalFormat" json:"numberInternationalFormat"`
}

// BUILDER from bson map:
func BuildContactNumber(m map[string]interface{}, x *ContactNumber) {
	if val, ok := m["label"]; ok && val != nil {
		x.Label = LabelNumber("Label_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["number"]; ok && val != nil {
		x.Number = val.(string)
	}
	if val, ok := m["numberInternationalFormat"]; ok && val != nil {
		x.NumberInternationalFormat = val.(string)
	}
}
