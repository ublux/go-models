package models

import "time"
import . "github.com/ublux/go-models/enums"
import "go.mongodb.org/mongo-driver/bson/primitive"

type Identity struct {
	AllowConnectingFromIpRegex   string       `bson:"allowConnectingFromIpRegex" json:"allowConnectingFromIpRegex"`
	DateAuthenticated            time.Time    `bson:"dateAuthenticated" json:"dateAuthenticated"`
	DateCreated                  time.Time    `bson:"dateCreated" json:"dateCreated"`
	DateDeleted                  time.Time    `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated                  time.Time    `bson:"dateUpdated" json:"dateUpdated"`
	Id                           string       `bson:"id" json:"id"`
	IdAccount                    string       `bson:"idAccount" json:"idAccount"`
	IdentityType                 IdentityType `bson:"identityType" json:"identityType"`
	IpAddressWhereAuthenticated  string       `bson:"ipAddressWhereAuthenticated" json:"ipAddressWhereAuthenticated"`
	Password                     string       `bson:"password" json:"password"`
	PreventConnectingIfIpChanges bool         `bson:"preventConnectingIfIpChanges" json:"preventConnectingIfIpChanges"`
	UbluxRoles                   []UbluxRole  `bson:"ubluxRoles" json:"ubluxRoles"`
	Username                     string       `bson:"username" json:"username"`
}

// Implementing interface IUbluxDocument
func (x Identity) GetDateDeleted() time.Time {
	return x.DateDeleted
}
func (x Identity) GetDateCreated() time.Time {
	return x.DateCreated
}
func (x Identity) GetDateUpdated() time.Time {
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
	if val, ok := m["allowConnectingFromIpRegex"]; ok && val != nil {
		x.AllowConnectingFromIpRegex = val.(string)
	}
	if val, ok := m["dateAuthenticated"]; ok && val != nil {
		x.DateAuthenticated = val.(time.Time)
	}
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(time.Time)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(time.Time)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(time.Time)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
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
