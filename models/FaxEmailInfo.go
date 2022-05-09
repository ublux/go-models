package models

import "go.mongodb.org/mongo-driver/bson/primitive"
import . "github.com/ublux/go-models/enums"

type FaxEmailInfo struct {
	DateCreated                                           primitive.DateTime   `bson:"dateCreated" json:"dateCreated"`
	DateDeleted                                           primitive.DateTime   `bson:"dateDeleted" json:"dateDeleted"`
	DateIdentificationTokenCreated                        primitive.DateTime   `bson:"dateIdentificationTokenCreated" json:"dateIdentificationTokenCreated"`
	DateUpdated                                           primitive.DateTime   `bson:"dateUpdated" json:"dateUpdated"`
	HaveWeSentEmailExplainingUserHasToWaitForConfirmation bool                 `bson:"haveWeSentEmailExplainingUserHasToWaitForConfirmation" json:"haveWeSentEmailExplainingUserHasToWaitForConfirmation"`
	Id                                                    string               `bson:"_id" json:"id"`
	IdAccount                                             string               `bson:"idAccount" json:"idAccount"`
	IdentificationToken                                   int32                `bson:"identificationToken" json:"identificationToken"`
	IdLineThatValidatedEmail                              string               `bson:"idLineThatValidatedEmail" json:"idLineThatValidatedEmail"`
	IsSpam                                                bool                 `bson:"isSpam" json:"isSpam"`
	NumberOfEmailsReceived                                int32                `bson:"numberOfEmailsReceived" json:"numberOfEmailsReceived"`
	ReplyStatus                                           LinkFaxToEmailStatus `bson:"replyStatus" json:"replyStatus"`
}

// Implementing interface IUbluxDocument
func (x FaxEmailInfo) GetDateDeleted() primitive.DateTime {
	return x.DateDeleted
}
func (x FaxEmailInfo) GetDateCreated() primitive.DateTime {
	return x.DateCreated
}
func (x FaxEmailInfo) GetDateUpdated() primitive.DateTime {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x FaxEmailInfo) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x FaxEmailInfo) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildFaxEmailInfo(m map[string]interface{}, x *FaxEmailInfo) {
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(primitive.DateTime)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(primitive.DateTime)
	}
	if val, ok := m["dateIdentificationTokenCreated"]; ok && val != nil {
		x.DateIdentificationTokenCreated = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(primitive.DateTime)
	}
	if val, ok := m["haveWeSentEmailExplainingUserHasToWaitForConfirmation"]; ok && val != nil {
		x.HaveWeSentEmailExplainingUserHasToWaitForConfirmation = val.(bool)
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
	if val, ok := m["identificationToken"]; ok && val != nil {
		x.IdentificationToken = val.(int32)
	}
	if val, ok := m["idLineThatValidatedEmail"]; ok && val != nil {
		x.IdLineThatValidatedEmail = val.(string)
	}
	if val, ok := m["isSpam"]; ok && val != nil {
		x.IsSpam = val.(bool)
	}
	if val, ok := m["numberOfEmailsReceived"]; ok && val != nil {
		x.NumberOfEmailsReceived = val.(int32)
	}
	if val, ok := m["replyStatus"]; ok && val != nil {
		x.ReplyStatus = LinkFaxToEmailStatus("ReplyStatus_" + val.(string))
	} // is NOT readonly obtained from map
}
