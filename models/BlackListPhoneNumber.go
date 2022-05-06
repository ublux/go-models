package models

import "time"

type BlackListPhoneNumber struct {
	DateCreated                  time.Time `bson:"dateCreated" json:"dateCreated"`
	DateDeleted                  time.Time `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated                  time.Time `bson:"dateUpdated" json:"dateUpdated"`
	FriendlyName                 string    `bson:"friendlyName" json:"friendlyName"`
	Id                           string    `bson:"id" json:"id"`
	IdAccount                    string    `bson:"idAccount" json:"idAccount"`
	IdAudioToPlayIfCallIsBlocked string    `bson:"idAudioToPlayIfCallIsBlocked" json:"idAudioToPlayIfCallIsBlocked"`
	PhoneNumber                  string    `bson:"phoneNumber" json:"phoneNumber"`
}

// Implementing interface IUbluxDocument
func (x BlackListPhoneNumber) GetDateDeleted() time.Time {
	return x.DateDeleted
}
func (x BlackListPhoneNumber) GetDateCreated() time.Time {
	return x.DateCreated
}
func (x BlackListPhoneNumber) GetDateUpdated() time.Time {
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
		x.DateCreated = val.(time.Time)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(time.Time)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(time.Time)
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
