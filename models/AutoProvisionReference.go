package models

import "time"

type AutoProvisionReference struct {
	DateCreated           time.Time `bson:"dateCreated" json:"dateCreated"`
	DateDeleted           time.Time `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated           time.Time `bson:"dateUpdated" json:"dateUpdated"`
	Id                    string    `bson:"id" json:"id"`
	IdPhone               string    `bson:"idPhone" json:"idPhone"`
	IdPhoneToExchangeWith string    `bson:"idPhoneToExchangeWith" json:"idPhoneToExchangeWith"`
	RequestedDisconnect   bool      `bson:"requestedDisconnect" json:"requestedDisconnect"`
}

// Implementing interface IUbluxDocument
func (x AutoProvisionReference) GetDateDeleted() time.Time {
	return x.DateDeleted
}
func (x AutoProvisionReference) GetDateCreated() time.Time {
	return x.DateCreated
}
func (x AutoProvisionReference) GetDateUpdated() time.Time {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x AutoProvisionReference) GetId() string {
	return x.Id
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildAutoProvisionReference(m map[string]interface{}, x *AutoProvisionReference) {
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(time.Time)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(time.Time)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(time.Time)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["idPhone"]; ok && val != nil {
		x.IdPhone = val.(string)
	}
	if val, ok := m["idPhoneToExchangeWith"]; ok && val != nil {
		x.IdPhoneToExchangeWith = val.(string)
	}
	if val, ok := m["requestedDisconnect"]; ok && val != nil {
		x.RequestedDisconnect = val.(bool)
	}
}
