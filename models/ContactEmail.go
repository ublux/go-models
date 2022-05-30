package models

import . "github.com/ublux/go-models/enums"

type ContactEmail struct {
	Email       string         `bson:"email" json:"email"`
	Label       LabelEmailType `bson:"label" json:"label"`
	SearchIndex string         `bson:"searchIndex" json:"searchIndex"`
}

// BUILDER from bson map:
func BuildContactEmail(m map[string]interface{}, x *ContactEmail) {
	if val, ok := m["email"]; ok && val != nil {
		x.Email = val.(string)
	}
	if val, ok := m["label"]; ok && val != nil {
		x.Label = LabelEmailType("Label_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["searchIndex"]; ok && val != nil {
		x.SearchIndex = val.(string)
	}
}
