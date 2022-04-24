package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type LineKeyGroup struct {
	Id           string    `bson:"_id" json:"id"`
	DateCreated  string    `bson:"dateCreated" json:"dateCreated"`
	DateDeleted  string    `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated  string    `bson:"dateUpdated" json:"dateUpdated"`
	Description  string    `bson:"description" json:"description"`
	FriendlyName string    `bson:"friendlyName" json:"friendlyName"`
	IdAccount    string    `bson:"idAccount" json:"idAccount"`
	LineKeys     []LineKey `bson:"lineKeys" json:"lineKeys"`
}

// Implementing interface IUbluxDocument
func (x LineKeyGroup) GetDateCreated() string {
	return x.DateCreated
}
func (x LineKeyGroup) GetDateDeleted() string {
	return x.DateDeleted
}
func (x LineKeyGroup) GetDateUpdated() string {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x LineKeyGroup) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x LineKeyGroup) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildLineKeyGroup(m map[string]interface{}, x *LineKeyGroup) {
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(string)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(string)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(string)
	}
	if val, ok := m["description"]; ok && val != nil {
		x.Description = val.(string)
	}
	if val, ok := m["friendlyName"]; ok && val != nil {
		x.FriendlyName = val.(string)
	}
	if val, ok := m["idAccount"]; ok && val != nil {
		x.IdAccount = val.(string)
	}
	if val, ok := m["lineKeys"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok {
			for _, val = range array {
				if val != nil {
					tmp := new(LineKey)
					BuildLineKey(val.(map[string]interface{}), tmp)
					x.LineKeys = append(x.LineKeys, *tmp)
				}
			}
		}
	}
}
