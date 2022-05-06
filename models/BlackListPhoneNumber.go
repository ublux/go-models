package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type BlackListPhoneNumber struct {
	DateCreated                  primitive.DateTime `bson:"dateCreated" json:"dateCreated"`
	DateDeleted                  primitive.DateTime `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated                  primitive.DateTime `bson:"dateUpdated" json:"dateUpdated"`
	FriendlyName                 string             `bson:"friendlyName" json:"friendlyName"`
	Id                           string             `bson:"id" json:"id"`
	IdAccount                    string             `bson:"idAccount" json:"idAccount"`
	IdAudioToPlayIfCallIsBlocked string             `bson:"idAudioToPlayIfCallIsBlocked" json:"idAudioToPlayIfCallIsBlocked"`
	PhoneNumber                  string             `bson:"phoneNumber" json:"phoneNumber"`
}

// Implementing interface IUbluxDocument
func (x BlackListPhoneNumber) GetDateDeleted() primitive.DateTime {
	return x.DateDeleted
}
func (x BlackListPhoneNumber) GetDateCreated() primitive.DateTime {
	return x.DateCreated
}
func (x BlackListPhoneNumber) GetDateUpdated() primitive.DateTime {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x BlackListPhoneNumber) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x BlackListPhoneNumber) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildBlackListPhoneNumber(m map[string]interface{}, x *BlackListPhoneNumber) {
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
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["idAccount"]; ok && val != nil {
		x.IdAccount = val.(string)
	}
	if val, ok := m["idAudioToPlayIfCallIsBlocked"]; ok && val != nil {
		x.IdAudioToPlayIfCallIsBlocked = val.(string)
	}
	if val, ok := m["phoneNumber"]; ok && val != nil {
		x.PhoneNumber = val.(string)
	}
}
