package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UnauthorizedToken struct {
	DateCreated    primitive.DateTime `bson:"dateCreated" json:"dateCreated"`
	DateDeleted    primitive.DateTime `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated    primitive.DateTime `bson:"dateUpdated" json:"dateUpdated"`
	ExpirationDate primitive.DateTime `bson:"expirationDate" json:"expirationDate"`
	Id             string             `bson:"_id" json:"id"`
	IdAccount      string             `bson:"idAccount" json:"idAccount"`
	IdIdentity     string             `bson:"idIdentity" json:"idIdentity"`
	Jwt            string             `bson:"jwt" json:"jwt"`
}

// Implementing interface IUbluxDocument
func (x UnauthorizedToken) GetDateDeleted() primitive.DateTime {
	return x.DateDeleted
}
func (x UnauthorizedToken) GetDateCreated() primitive.DateTime {
	return x.DateCreated
}
func (x UnauthorizedToken) GetDateUpdated() primitive.DateTime {
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
		x.DateCreated = val.(primitive.DateTime)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(primitive.DateTime)
	}
	if val, ok := m["expirationDate"]; ok && val != nil {
		x.ExpirationDate = val.(primitive.DateTime)
	}
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
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
