package models

type CustomerInfo struct {
	Id                      string                  `bson:"_id" json:"id"`
	AirNetworksCustomerInfo AirNetworksCustomerInfo `bson:"airNetworksCustomerInfo" json:"airNetworksCustomerInfo"`
	DateCreated             string                  `bson:"dateCreated" json:"dateCreated"`
	DateDeleted             string                  `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated             string                  `bson:"dateUpdated" json:"dateUpdated"`
	FullName                string                  `bson:"fullName" json:"fullName"`
	IdAccount               string                  `bson:"idAccount" json:"idAccount"`
	MailingAddress          MailingAddress          `bson:"mailingAddress" json:"mailingAddress"`
}

// Implementing interface IUbluxDocument
func (x CustomerInfo) GetDateCreated() string {
	return x.DateCreated
}
func (x CustomerInfo) GetDateDeleted() string {
	return x.DateDeleted
}
func (x CustomerInfo) GetDateUpdated() string {
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
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["airNetworksCustomerInfo"]; ok && val != nil {
		BuildAirNetworksCustomerInfo(val.(map[string]interface{}), &x.AirNetworksCustomerInfo)
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
	if val, ok := m["fullName"]; ok && val != nil {
		x.FullName = val.(string)
	}
	if val, ok := m["idAccount"]; ok && val != nil {
		x.IdAccount = val.(string)
	}
	if val, ok := m["mailingAddress"]; ok && val != nil {
		BuildMailingAddress(val.(map[string]interface{}), &x.MailingAddress)
	}
}
