package models

import . "github.com/ublux/go-models/enums"

type AgreementToTermsAndConditions struct {
	Id                         string                     `bson:"_id" json:"id"`
	DateCreated                string                     `bson:"dateCreated" json:"dateCreated"`
	DateDeleted                string                     `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated                string                     `bson:"dateUpdated" json:"dateUpdated"`
	Description                string                     `bson:"description" json:"description"`
	HttpRequestHeaders         string                     `bson:"httpRequestHeaders" json:"httpRequestHeaders"`
	Ip                         string                     `bson:"ip" json:"ip"`
	TermsAndConditionsCategory TermsAndConditionsCategory `bson:"termsAndConditionsCategory" json:"termsAndConditionsCategory"`
	UbluxSession               UbluxSession               `bson:"ubluxSession" json:"ubluxSession"`
}

// Implementing interface IUbluxDocument
func (x AgreementToTermsAndConditions) GetDateCreated() string {
	return x.DateCreated
}
func (x AgreementToTermsAndConditions) GetDateDeleted() string {
	return x.DateDeleted
}
func (x AgreementToTermsAndConditions) GetDateUpdated() string {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x AgreementToTermsAndConditions) GetId() string {
	return x.Id
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildAgreementToTermsAndConditions(m map[string]interface{}, x *AgreementToTermsAndConditions) {
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
	if val, ok := m["httpRequestHeaders"]; ok && val != nil {
		x.HttpRequestHeaders = val.(string)
	}
	if val, ok := m["ip"]; ok && val != nil {
		x.Ip = val.(string)
	}
	if val, ok := m["termsAndConditionsCategory"]; ok && val != nil {
		x.TermsAndConditionsCategory = TermsAndConditionsCategory("TermsAndConditionsCategory_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["ubluxSession"]; ok && val != nil {
		BuildUbluxSession(val.(map[string]interface{}), &x.UbluxSession)
	}
}
