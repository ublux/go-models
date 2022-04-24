package models

import . "github.com/ublux/go-models/enums"

type ChildCallAttendantTransferToPSTN struct {
	ChildCallType ChildCallType `bson:"childCallType" json:"childCallType"`
	PhoneNumber   string        `bson:"phoneNumber" json:"phoneNumber"`
}

// Implementing interface ChildCall
func (x ChildCallAttendantTransferToPSTN) GetChildCallType() ChildCallType {
	return x.ChildCallType
}

// BUILDER from bson map:
func BuildChildCallAttendantTransferToPSTN(m map[string]interface{}, x *ChildCallAttendantTransferToPSTN) {
	x.ChildCallType = ChildCallType_AttendantTransferToPSTN // readonly property
	if val, ok := m["phoneNumber"]; ok && val != nil {
		x.PhoneNumber = val.(string)
	}
}
