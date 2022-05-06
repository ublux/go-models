package models

import "time"
import "go.mongodb.org/mongo-driver/bson/primitive"

type FaxOutgoingGroup struct {
	ContainsError            bool          `bson:"containsError" json:"containsError"`
	DateCreated              time.Time     `bson:"dateCreated" json:"dateCreated"`
	DateDeleted              time.Time     `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated              time.Time     `bson:"dateUpdated" json:"dateUpdated"`
	FaxEmail                 FaxEmail      `bson:"faxEmail" json:"faxEmail"`
	FaxesOutgoing            []FaxOutgoing `bson:"faxesOutgoing" json:"faxesOutgoing"`
	From                     string        `bson:"from" json:"from"`
	Id                       string        `bson:"id" json:"id"`
	IdAccount                string        `bson:"idAccount" json:"idAccount"`
	IdVoipNumberFax          string        `bson:"idVoipNumberFax" json:"idVoipNumberFax"`
	SendConfirmationToEmails []string      `bson:"sendConfirmationToEmails" json:"sendConfirmationToEmails"`
}

// Implementing interface IUbluxDocument
func (x FaxOutgoingGroup) GetDateDeleted() time.Time {
	return x.DateDeleted
}
func (x FaxOutgoingGroup) GetDateCreated() time.Time {
	return x.DateCreated
}
func (x FaxOutgoingGroup) GetDateUpdated() time.Time {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x FaxOutgoingGroup) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x FaxOutgoingGroup) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildFaxOutgoingGroup(m map[string]interface{}, x *FaxOutgoingGroup) {
	if val, ok := m["containsError"]; ok && val != nil {
		x.ContainsError = val.(bool)
	}
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(time.Time)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(time.Time)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(time.Time)
	}
	if val, ok := m["faxEmail"]; ok && val != nil {
		BuildFaxEmail(val.(map[string]interface{}), &x.FaxEmail)
	}
	if val, ok := m["faxesOutgoing"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok {
			for _, val = range array {
				if val != nil {
					tmp := new(FaxOutgoing)
					BuildFaxOutgoing(val.(map[string]interface{}), tmp)
					x.FaxesOutgoing = append(x.FaxesOutgoing, *tmp)
				}
			}
		}
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
	if val, ok := m["idVoipNumberFax"]; ok && val != nil {
		x.IdVoipNumberFax = val.(string)
	}
	if val, ok := m["sendConfirmationToEmails"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				if val != nil {
					x.SendConfirmationToEmails = append(x.SendConfirmationToEmails, val.(string))
				}
			}
		}
	}
}
