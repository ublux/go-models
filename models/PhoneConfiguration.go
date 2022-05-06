package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PhoneConfiguration struct {
	DateCreated    primitive.DateTime `bson:"dateCreated" json:"dateCreated"`
	DateDeleted    primitive.DateTime `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated    primitive.DateTime `bson:"dateUpdated" json:"dateUpdated"`
	Description    string             `bson:"description" json:"description"`
	FrienlyName    string             `bson:"frienlyName" json:"frienlyName"`
	Id             string             `bson:"id" json:"id"`
	IdAccount      string             `bson:"idAccount" json:"idAccount"`
	IdLineKeyGroup string             `bson:"idLineKeyGroup" json:"idLineKeyGroup"`
}

// Implementing interface IUbluxDocument
func (x PhoneConfiguration) GetDateDeleted() primitive.DateTime {
	return x.DateDeleted
}
func (x PhoneConfiguration) GetDateCreated() primitive.DateTime {
	return x.DateCreated
}
func (x PhoneConfiguration) GetDateUpdated() primitive.DateTime {
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
		x.DateCreated = val.(primitive.DateTime)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(primitive.DateTime)
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
