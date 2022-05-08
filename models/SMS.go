package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SMS struct {
	Body         string             `bson:"body" json:"body"`
	Contact      Contact            `bson:"contact" json:"contact"`
	DateCreated  primitive.DateTime `bson:"dateCreated" json:"dateCreated"`
	DateDeleted  primitive.DateTime `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated  primitive.DateTime `bson:"dateUpdated" json:"dateUpdated"`
	From         string             `bson:"from" json:"from"`
	Id           string             `bson:"_id" json:"id"`
	IdAccount    string             `bson:"idAccount" json:"idAccount"`
	IdVoipNumber string             `bson:"idVoipNumber" json:"idVoipNumber"`
	IsIncoming   bool               `bson:"isIncoming" json:"isIncoming"`
	NumSegments  int32              `bson:"numSegments" json:"numSegments"`
	Status       string             `bson:"status" json:"status"`
	To           string             `bson:"to" json:"to"`
}

// Implementing interface IUbluxDocument
func (x SMS) GetDateDeleted() primitive.DateTime {
	return x.DateDeleted
}
func (x SMS) GetDateCreated() primitive.DateTime {
	return x.DateCreated
}
func (x SMS) GetDateUpdated() primitive.DateTime {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x SMS) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x SMS) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildSMS(m map[string]interface{}, x *SMS) {
	if val, ok := m["body"]; ok && val != nil {
		x.Body = val.(string)
	}
	if val, ok := m["contact"]; ok && val != nil {
		BuildContact(val.(map[string]interface{}), &x.Contact)
	}
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(primitive.DateTime)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(primitive.DateTime)
	}
	if val, ok := m["from"]; ok && val != nil {
		x.From = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["idAccount"]; ok && val != nil {
		x.IdAccount = val.(string)
	}
	if val, ok := m["idVoipNumber"]; ok && val != nil {
		x.IdVoipNumber = val.(string)
	}
	if val, ok := m["isIncoming"]; ok && val != nil {
		x.IsIncoming = val.(bool)
	}
	if val, ok := m["numSegments"]; ok && val != nil {
		x.NumSegments = val.(int32)
	}
	if val, ok := m["status"]; ok && val != nil {
		x.Status = val.(string)
	}
	if val, ok := m["to"]; ok && val != nil {
		x.To = val.(string)
	}
}
