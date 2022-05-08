package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AutoProvisionReference struct {
	DateCreated           primitive.DateTime `bson:"dateCreated" json:"dateCreated"`
	DateDeleted           primitive.DateTime `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated           primitive.DateTime `bson:"dateUpdated" json:"dateUpdated"`
	Id                    string             `bson:"_id" json:"id"`
	IdPhone               string             `bson:"idPhone" json:"idPhone"`
	IdPhoneToExchangeWith string             `bson:"idPhoneToExchangeWith" json:"idPhoneToExchangeWith"`
	RequestedDisconnect   bool               `bson:"requestedDisconnect" json:"requestedDisconnect"`
}

// Implementing interface IUbluxDocument
func (x AutoProvisionReference) GetDateDeleted() primitive.DateTime {
	return x.DateDeleted
}
func (x AutoProvisionReference) GetDateCreated() primitive.DateTime {
	return x.DateCreated
}
func (x AutoProvisionReference) GetDateUpdated() primitive.DateTime {
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
		x.DateCreated = val.(primitive.DateTime)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(primitive.DateTime)
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
