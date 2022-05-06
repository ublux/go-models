package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type IP struct {
	DateCreated primitive.DateTime `bson:"dateCreated" json:"dateCreated"`
	DateDeleted primitive.DateTime `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated primitive.DateTime `bson:"dateUpdated" json:"dateUpdated"`
	Id          string             `bson:"id" json:"id"`
	IdIdentity  string             `bson:"idIdentity" json:"idIdentity"`
	IpOrigin    string             `bson:"ipOrigin" json:"ipOrigin"`
	IsBlack     bool               `bson:"isBlack" json:"isBlack"`
}

// Implementing interface IUbluxDocument
func (x IP) GetDateDeleted() primitive.DateTime {
	return x.DateDeleted
}
func (x IP) GetDateCreated() primitive.DateTime {
	return x.DateCreated
}
func (x IP) GetDateUpdated() primitive.DateTime {
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
		x.DateCreated = val.(primitive.DateTime)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(primitive.DateTime)
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
