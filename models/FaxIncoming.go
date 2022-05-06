package models

import "time"

type FaxIncoming struct {
	DateCreated     time.Time  `bson:"dateCreated" json:"dateCreated"`
	DateDeleted     time.Time  `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated     time.Time  `bson:"dateUpdated" json:"dateUpdated"`
	ErrorMessage    string     `bson:"errorMessage" json:"errorMessage"`
	FaxStatus       string     `bson:"faxStatus" json:"faxStatus"`
	From            string     `bson:"from" json:"from"`
	Id              string     `bson:"id" json:"id"`
	IdAccount       string     `bson:"idAccount" json:"idAccount"`
	IdVoipNumberFax string     `bson:"idVoipNumberFax" json:"idVoipNumberFax"`
	NumPages        int32      `bson:"numPages" json:"numPages"`
	Pdf             StoredFile `bson:"pdf" json:"pdf"`
	To              string     `bson:"to" json:"to"`
}

// Implementing interface IUbluxDocument
func (x FaxIncoming) GetDateDeleted() time.Time {
	return x.DateDeleted
}
func (x FaxIncoming) GetDateCreated() time.Time {
	return x.DateCreated
}
func (x FaxIncoming) GetDateUpdated() time.Time {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x FaxIncoming) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x FaxIncoming) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildFaxIncoming(m map[string]interface{}, x *FaxIncoming) {
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(time.Time)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(time.Time)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(time.Time)
	}
	if val, ok := m["errorMessage"]; ok && val != nil {
		x.ErrorMessage = val.(string)
	}
	if val, ok := m["faxStatus"]; ok && val != nil {
		x.FaxStatus = val.(string)
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
	if val, ok := m["numPages"]; ok && val != nil {
		x.NumPages = val.(int32)
	}
	if val, ok := m["pdf"]; ok && val != nil {
		BuildStoredFile(val.(map[string]interface{}), &x.Pdf)
	}
	if val, ok := m["to"]; ok && val != nil {
		x.To = val.(string)
	}
}
