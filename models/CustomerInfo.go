package models

import "time"

type CustomerInfo struct {
	AirNetworksCustomerInfo AirNetworksCustomerInfo `bson:"airNetworksCustomerInfo" json:"airNetworksCustomerInfo"`
	DateCreated             time.Time               `bson:"dateCreated" json:"dateCreated"`
	DateDeleted             time.Time               `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated             time.Time               `bson:"dateUpdated" json:"dateUpdated"`
	FullName                string                  `bson:"fullName" json:"fullName"`
	Id                      string                  `bson:"id" json:"id"`
	IdAccount               string                  `bson:"idAccount" json:"idAccount"`
	MailingAddress          MailingAddress          `bson:"mailingAddress" json:"mailingAddress"`
}

// Implementing interface IUbluxDocument
func (x CustomerInfo) GetDateDeleted() time.Time {
	return x.DateDeleted
}
func (x CustomerInfo) GetDateCreated() time.Time {
	return x.DateCreated
}
func (x CustomerInfo) GetDateUpdated() time.Time {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x CustomerInfo) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x CustomerInfo) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildCustomerInfo(m map[string]interface{}, x *CustomerInfo) {
	if val, ok := m["airNetworksCustomerInfo"]; ok && val != nil {
		BuildAirNetworksCustomerInfo(val.(map[string]interface{}), &x.AirNetworksCustomerInfo)
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
	if val, ok := m["fullName"]; ok && val != nil {
		x.FullName = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["idAccount"]; ok && val != nil {
		x.IdAccount = val.(string)
	}
	if val, ok := m["mailingAddress"]; ok && val != nil {
		BuildMailingAddress(val.(map[string]interface{}), &x.MailingAddress)
	}
}
