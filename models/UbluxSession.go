package models

import (
	. "github.com/ublux/go-models/enums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UbluxSession struct {
	ExpirationDate string       `bson:"expirationDate" json:"expirationDate"`
	IdAccount      string       `bson:"idAccount" json:"idAccount"`
	IdentityType   IdentityType `bson:"identityType" json:"identityType"`
	IdIdentity     string       `bson:"idIdentity" json:"idIdentity"`
	UbluxRoles     []UbluxRole  `bson:"ubluxRoles" json:"ubluxRoles"`
}

// Implementing interface IReferncesAccount
func (x UbluxSession) GetIdAccount() string {
	return x.IdAccount
}

// BUILDER from bson map:
func BuildUbluxSession(m map[string]interface{}, x *UbluxSession) {
	if val, ok := m["expirationDate"]; ok && val != nil {
		x.ExpirationDate = val.(string)
	}
	if val, ok := m["idAccount"]; ok && val != nil {
		x.IdAccount = val.(string)
	}
	if val, ok := m["identityType"]; ok && val != nil {
		x.IdentityType = IdentityType("IdentityType_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["idIdentity"]; ok && val != nil {
		x.IdIdentity = val.(string)
	}
	if val, ok := m["ubluxRoles"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				x.UbluxRoles = append(x.UbluxRoles, UbluxRole("UbluxRoles_"+val.(string)))
			} // is NOT readonly obtained from map
		}
	}
}
