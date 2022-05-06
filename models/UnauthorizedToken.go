package models

import "time"

type UnauthorizedToken struct {
	DateCreated    time.Time `bson:"dateCreated" json:"dateCreated"`
	DateDeleted    time.Time `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated    time.Time `bson:"dateUpdated" json:"dateUpdated"`
	ExpirationDate time.Time `bson:"expirationDate" json:"expirationDate"`
	Id             string    `bson:"id" json:"id"`
	IdAccount      string    `bson:"idAccount" json:"idAccount"`
	IdIdentity     string    `bson:"idIdentity" json:"idIdentity"`
	Jwt            string    `bson:"jwt" json:"jwt"`
}

// Implementing interface IUbluxDocument
func (x UnauthorizedToken) GetDateDeleted() time.Time {
	return x.DateDeleted
}
func (x UnauthorizedToken) GetDateCreated() time.Time {
	return x.DateCreated
}
func (x UnauthorizedToken) GetDateUpdated() time.Time {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x UnauthorizedToken) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x UnauthorizedToken) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildUnauthorizedToken(m map[string]interface{}, x *UnauthorizedToken) {
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(time.Time)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(time.Time)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(time.Time)
	}
	if val, ok := m["expirationDate"]; ok && val != nil {
		x.ExpirationDate = val.(time.Time)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["idAccount"]; ok && val != nil {
		x.IdAccount = val.(string)
	}
	if val, ok := m["idIdentity"]; ok && val != nil {
		x.IdIdentity = val.(string)
	}
	if val, ok := m["jwt"]; ok && val != nil {
		x.Jwt = val.(string)
	}
}
