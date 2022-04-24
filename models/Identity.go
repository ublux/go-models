package models

import . "github.com/ublux/go-models/enums"
import "go.mongodb.org/mongo-driver/bson/primitive"

type Identity struct {
	Id                           string       `bson:"_id" json:"id"`
	AllowConnectingFromIpRegex   string       `bson:"allowConnectingFromIpRegex" json:"allowConnectingFromIpRegex"`
	DateAuthenticated            string       `bson:"dateAuthenticated" json:"dateAuthenticated"`
	DateCreated                  string       `bson:"dateCreated" json:"dateCreated"`
	DateDeleted                  string       `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated                  string       `bson:"dateUpdated" json:"dateUpdated"`
	IdAccount                    string       `bson:"idAccount" json:"idAccount"`
	IdentityType                 IdentityType `bson:"identityType" json:"identityType"`
	IpAddressWhereAuthenticated  string       `bson:"ipAddressWhereAuthenticated" json:"ipAddressWhereAuthenticated"`
	Password                     string       `bson:"password" json:"password"`
	PreventConnectingIfIpChanges bool         `bson:"preventConnectingIfIpChanges" json:"preventConnectingIfIpChanges"`
	UbluxRoles                   []UbluxRole  `bson:"ubluxRoles" json:"ubluxRoles"`
	Username                     string       `bson:"username" json:"username"`
}

// Implementing interface IUbluxDocument
func (x Identity) GetDateCreated() string {
	return x.DateCreated
}
func (x Identity) GetDateDeleted() string {
	return x.DateDeleted
}
func (x Identity) GetDateUpdated() string {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x Identity) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x Identity) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildIdentity(m map[string]interface{}, x *Identity) {
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["allowConnectingFromIpRegex"]; ok && val != nil {
		x.AllowConnectingFromIpRegex = val.(string)
	}
	if val, ok := m["dateAuthenticated"]; ok && val != nil {
		x.DateAuthenticated = val.(string)
	}
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(string)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(string)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(string)
	}
	if val, ok := m["idAccount"]; ok && val != nil {
		x.IdAccount = val.(string)
	}
	if val, ok := m["identityType"]; ok && val != nil {
		x.IdentityType = IdentityType("IdentityType_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["ipAddressWhereAuthenticated"]; ok && val != nil {
		x.IpAddressWhereAuthenticated = val.(string)
	}
	if val, ok := m["password"]; ok && val != nil {
		x.Password = val.(string)
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
	if val, ok := m["username"]; ok && val != nil {
		x.Username = val.(string)
	}
}
