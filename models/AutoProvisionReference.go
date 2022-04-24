package models

type AutoProvisionReference struct {
	Id                    string `bson:"_id" json:"id"`
	DateCreated           string `bson:"dateCreated" json:"dateCreated"`
	DateDeleted           string `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated           string `bson:"dateUpdated" json:"dateUpdated"`
	IdPhone               string `bson:"idPhone" json:"idPhone"`
	IdPhoneToExchangeWith string `bson:"idPhoneToExchangeWith" json:"idPhoneToExchangeWith"`
	RequestedDisconnect   bool   `bson:"requestedDisconnect" json:"requestedDisconnect"`
}

// Implementing interface IUbluxDocument
func (x AutoProvisionReference) GetDateCreated() string {
	return x.DateCreated
}
func (x AutoProvisionReference) GetDateDeleted() string {
	return x.DateDeleted
}
func (x AutoProvisionReference) GetDateUpdated() string {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x AutoProvisionReference) GetId() string {
	return x.Id
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildAutoProvisionReference(m map[string]interface{}, x *AutoProvisionReference) {
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
