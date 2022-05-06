package models

import "time"

type PhoneConfiguration struct {
	DateCreated    time.Time `bson:"dateCreated" json:"dateCreated"`
	DateDeleted    time.Time `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated    time.Time `bson:"dateUpdated" json:"dateUpdated"`
	Description    string    `bson:"description" json:"description"`
	FrienlyName    string    `bson:"frienlyName" json:"frienlyName"`
	Id             string    `bson:"id" json:"id"`
	IdAccount      string    `bson:"idAccount" json:"idAccount"`
	IdLineKeyGroup string    `bson:"idLineKeyGroup" json:"idLineKeyGroup"`
}

// Implementing interface IUbluxDocument
func (x PhoneConfiguration) GetDateDeleted() time.Time {
	return x.DateDeleted
}
func (x PhoneConfiguration) GetDateCreated() time.Time {
	return x.DateCreated
}
func (x PhoneConfiguration) GetDateUpdated() time.Time {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x PhoneConfiguration) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x PhoneConfiguration) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildPhoneConfiguration(m map[string]interface{}, x *PhoneConfiguration) {
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(time.Time)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(time.Time)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(time.Time)
	}
	if val, ok := m["description"]; ok && val != nil {
		x.Description = val.(string)
	}
	if val, ok := m["frienlyName"]; ok && val != nil {
		x.FrienlyName = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["idAccount"]; ok && val != nil {
		x.IdAccount = val.(string)
	}
	if val, ok := m["idLineKeyGroup"]; ok && val != nil {
		x.IdLineKeyGroup = val.(string)
	}
}
