package models

import . "github.com/ublux/go-models/enums"

type ChildCallBlindTransferToPSTN struct {
	ChildCallType      ChildCallType `bson:"childCallType" json:"childCallType"`
	IdCall             string        `bson:"idCall" json:"idCall"`
	IdTrunkTermination string        `bson:"idTrunkTermination" json:"idTrunkTermination"`
	PhoneNumber        string        `bson:"phoneNumber" json:"phoneNumber"`
}

// Implementing interface ChildCallBlindTransfer
func (x ChildCallBlindTransferToPSTN) GetIdCall() string {
	return x.IdCall
}

// Implementing interface ChildCall
func (x ChildCallBlindTransferToPSTN) GetChildCallType() ChildCallType {
	return x.ChildCallType
}

// BUILDER from bson map:
func BuildChildCallBlindTransferToPSTN(m map[string]interface{}, x *ChildCallBlindTransferToPSTN) {
	x.ChildCallType = ChildCallType_BlindTransferToPSTN // readonly property
	if val, ok := m["idCall"]; ok && val != nil {
		x.IdCall = val.(string)
	}
	if val, ok := m["idTrunkTermination"]; ok && val != nil {
		x.IdTrunkTermination = val.(string)
	}
	if val, ok := m["phoneNumber"]; ok && val != nil {
		x.PhoneNumber = val.(string)
	}
}
