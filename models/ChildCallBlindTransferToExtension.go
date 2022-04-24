package models

import . "github.com/ublux/go-models/enums"

type ChildCallBlindTransferToExtension struct {
	ChildCallType ChildCallType `bson:"childCallType" json:"childCallType"`
	IdCall        string        `bson:"idCall" json:"idCall"`
	IdExtension   string        `bson:"idExtension" json:"idExtension"`
}

// Implementing interface ChildCallBlindTransfer
func (x ChildCallBlindTransferToExtension) GetIdCall() string {
	return x.IdCall
}

// Implementing interface ChildCall
func (x ChildCallBlindTransferToExtension) GetChildCallType() ChildCallType {
	return x.ChildCallType
}

// BUILDER from bson map:
func BuildChildCallBlindTransferToExtension(m map[string]interface{}, x *ChildCallBlindTransferToExtension) {
	x.ChildCallType = ChildCallType_BlindTransferToExtension // readonly property
	if val, ok := m["idCall"]; ok && val != nil {
		x.IdCall = val.(string)
	}
	if val, ok := m["idExtension"]; ok && val != nil {
		x.IdExtension = val.(string)
	}
}
