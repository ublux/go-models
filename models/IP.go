package models

import "time"

type IP struct {
	DateCreated time.Time `bson:"dateCreated" json:"dateCreated"`
	DateDeleted time.Time `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated time.Time `bson:"dateUpdated" json:"dateUpdated"`
	Id          string    `bson:"id" json:"id"`
	IdIdentity  string    `bson:"idIdentity" json:"idIdentity"`
	IpOrigin    string    `bson:"ipOrigin" json:"ipOrigin"`
	IsBlack     bool      `bson:"isBlack" json:"isBlack"`
}

// Implementing interface IUbluxDocument
func (x IP) GetDateDeleted() time.Time {
	return x.DateDeleted
}
func (x IP) GetDateCreated() time.Time {
	return x.DateCreated
}
func (x IP) GetDateUpdated() time.Time {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x IP) GetId() string {
	return x.Id
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildIP(m map[string]interface{}, x *IP) {
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
	if val, ok := m["idIdentity"]; ok && val != nil {
		x.IdIdentity = val.(string)
	}
	if val, ok := m["ipOrigin"]; ok && val != nil {
		x.IpOrigin = val.(string)
	}
	if val, ok := m["isBlack"]; ok && val != nil {
		x.IsBlack = val.(bool)
	}
}
