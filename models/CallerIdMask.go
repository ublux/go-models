package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CallerIdMask struct {
	DateCreated            primitive.DateTime `bson:"dateCreated" json:"dateCreated"`
	DateDeleted            primitive.DateTime `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated            primitive.DateTime `bson:"dateUpdated" json:"dateUpdated"`
	FriendlyName           string             `bson:"friendlyName" json:"friendlyName"`
	Id                     string             `bson:"_id" json:"id"`
	IdAccount              string             `bson:"idAccount" json:"idAccount"`
	PhoneNumber            string             `bson:"phoneNumber" json:"phoneNumber"`
	RandomVerificationCode string             `bson:"randomVerificationCode" json:"randomVerificationCode"`
}

// Implementing interface IUbluxDocument
func (x CallerIdMask) GetDateDeleted() primitive.DateTime {
	return x.DateDeleted
}
func (x CallerIdMask) GetDateCreated() primitive.DateTime {
	return x.DateCreated
}
func (x CallerIdMask) GetDateUpdated() primitive.DateTime {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x CallerIdMask) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x CallerIdMask) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildCallerIdMask(m map[string]interface{}, x *CallerIdMask) {
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(primitive.DateTime)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(primitive.DateTime)
	}
	if val, ok := m["friendlyName"]; ok && val != nil {
		x.FriendlyName = val.(string)
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
	if val, ok := m["phoneNumber"]; ok && val != nil {
		x.PhoneNumber = val.(string)
	}
	if val, ok := m["randomVerificationCode"]; ok && val != nil {
		x.RandomVerificationCode = val.(string)
	}
}
