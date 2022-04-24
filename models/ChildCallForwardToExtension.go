package models

import (
	. "github.com/ublux/go-models/enums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChildCallForwardToExtension struct {
	ChildCallType          ChildCallType `bson:"childCallType" json:"childCallType"`
	ExtensionFriendlyName  string        `bson:"extensionFriendlyName" json:"extensionFriendlyName"`
	ExtensionNumber        string        `bson:"extensionNumber" json:"extensionNumber"`
	IdExtension            string        `bson:"idExtension" json:"idExtension"`
	IdLineThatAnswered     string        `bson:"idLineThatAnswered" json:"idLineThatAnswered"`
	IdsLinesThatDidNotRing []string      `bson:"idsLinesThatDidNotRing" json:"idsLinesThatDidNotRing"`
	IdsLinesThatRing       []string      `bson:"idsLinesThatRing" json:"idsLinesThatRing"`
}

// Implementing interface ChildCall
func (x ChildCallForwardToExtension) GetChildCallType() ChildCallType {
	return x.ChildCallType
}

// BUILDER from bson map:
func BuildChildCallForwardToExtension(m map[string]interface{}, x *ChildCallForwardToExtension) {
	x.ChildCallType = ChildCallType_ForwardToExtension // readonly property
	if val, ok := m["extensionFriendlyName"]; ok && val != nil {
		x.ExtensionFriendlyName = val.(string)
	}
	if val, ok := m["extensionNumber"]; ok && val != nil {
		x.ExtensionNumber = val.(string)
	}
	if val, ok := m["idExtension"]; ok && val != nil {
		x.IdExtension = val.(string)
	}
	if val, ok := m["idLineThatAnswered"]; ok && val != nil {
		x.IdLineThatAnswered = val.(string)
	}
	if val, ok := m["idsLinesThatDidNotRing"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				if val != nil {
					x.IdsLinesThatDidNotRing = append(x.IdsLinesThatDidNotRing, val.(string))
				}
			}
		}
	}
	if val, ok := m["idsLinesThatRing"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				if val != nil {
					x.IdsLinesThatRing = append(x.IdsLinesThatRing, val.(string))
				}
			}
		}
	}
}
