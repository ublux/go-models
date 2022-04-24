package models

import . "github.com/ublux/go-models/enums"

type ChildCallAttendantTransferToExtension struct {
	ChildCallType ChildCallType `bson:"childCallType" json:"childCallType"`
	IdExtension   string        `bson:"idExtension" json:"idExtension"`
}

// Implementing interface ChildCall
func (x ChildCallAttendantTransferToExtension) GetChildCallType() ChildCallType {
	return x.ChildCallType
}

// BUILDER from bson map:
func BuildChildCallAttendantTransferToExtension(m map[string]interface{}, x *ChildCallAttendantTransferToExtension) {
	x.ChildCallType = ChildCallType_AttendantTransferToExtension // readonly property
	if val, ok := m["idExtension"]; ok && val != nil {
		x.IdExtension = val.(string)
	}
}
