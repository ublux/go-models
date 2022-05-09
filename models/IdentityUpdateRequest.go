package models

import . "github.com/ublux/go-models/enums"
import "go.mongodb.org/mongo-driver/bson/primitive"

type IdentityUpdateRequest struct {
	AllowConnectingFromIpRegex   string      `bson:"allowConnectingFromIpRegex" json:"allowConnectingFromIpRegex"`
	Id                           string      `bson:"id" json:"id"`
	PreventConnectingIfIpChanges bool        `bson:"preventConnectingIfIpChanges" json:"preventConnectingIfIpChanges"`
	UbluxRoles                   []UbluxRole `bson:"ubluxRoles" json:"ubluxRoles"`
}

// Implementing interface IUbluxDocumentId
func (x IdentityUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildIdentityUpdateRequest(m map[string]interface{}, x *IdentityUpdateRequest) {
	if val, ok := m["allowConnectingFromIpRegex"]; ok && val != nil {
		x.AllowConnectingFromIpRegex = val.(string)
	}
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["preventConnectingIfIpChanges"]; ok && val != nil {
		x.PreventConnectingIfIpChanges = val.(bool)
	}
	if val, ok := m["ubluxRoles"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				x.UbluxRoles = append(x.UbluxRoles, UbluxRole("UbluxRoles_"+val.(string)))
			} // is NOT readonly obtained from map
		}
	}
}
