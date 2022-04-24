package models

import . "github.com/ublux/go-models/enums"

type LineKey struct {
	IdExtension string      `bson:"idExtension" json:"idExtension"`
	Label       string      `bson:"label" json:"label"`
	LineIndex   int32       `bson:"lineIndex" json:"lineIndex"`
	LineKeyType LineKeyType `bson:"lineKeyType" json:"lineKeyType"`
	Value       string      `bson:"value" json:"value"`
}

// BUILDER from bson map:
func BuildLineKey(m map[string]interface{}, x *LineKey) {
	if val, ok := m["idExtension"]; ok && val != nil {
		x.IdExtension = val.(string)
	}
	if val, ok := m["label"]; ok && val != nil {
		x.Label = val.(string)
	}
	if val, ok := m["lineIndex"]; ok && val != nil {
		x.LineIndex = val.(int32)
	}
	if val, ok := m["lineKeyType"]; ok && val != nil {
		x.LineKeyType = LineKeyType("LineKeyType_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["value"]; ok && val != nil {
		x.Value = val.(string)
	}
}
