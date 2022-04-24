package models

import . "github.com/ublux/go-models/enums"

type VoipProvider struct {
	Id             string         `bson:"_id" json:"id"`
	AccessToken    string         `bson:"accessToken" json:"accessToken"`
	Country        CountryIsoCode `bson:"country" json:"country"`
	DateCreated    string         `bson:"dateCreated" json:"dateCreated"`
	DateDeleted    string         `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated    string         `bson:"dateUpdated" json:"dateUpdated"`
	FriendlyName   string         `bson:"friendlyName" json:"friendlyName"`
	OwnerAccountId string         `bson:"ownerAccountId" json:"ownerAccountId"`
	SID            string         `bson:"sID" json:"sID"`
	Status         string         `bson:"status" json:"status"`
	UbluxPartner   UbluxPartner   `bson:"ubluxPartner" json:"ubluxPartner"`
	VoipCompany    VoipCompany    `bson:"voipCompany" json:"voipCompany"`
}

// Implementing interface IUbluxDocument
func (x VoipProvider) GetDateCreated() string {
	return x.DateCreated
}
func (x VoipProvider) GetDateDeleted() string {
	return x.DateDeleted
}
func (x VoipProvider) GetDateUpdated() string {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x VoipProvider) GetId() string {
	return x.Id
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildVoipProvider(m map[string]interface{}, x *VoipProvider) {
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["accessToken"]; ok && val != nil {
		x.AccessToken = val.(string)
	}
	if val, ok := m["country"]; ok && val != nil {
		x.Country = CountryIsoCode("Country_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(string)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(string)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(string)
	}
	if val, ok := m["friendlyName"]; ok && val != nil {
		x.FriendlyName = val.(string)
	}
	if val, ok := m["ownerAccountId"]; ok && val != nil {
		x.OwnerAccountId = val.(string)
	}
	if val, ok := m["sID"]; ok && val != nil {
		x.SID = val.(string)
	}
	if val, ok := m["status"]; ok && val != nil {
		x.Status = val.(string)
	}
	if val, ok := m["ubluxPartner"]; ok && val != nil {
		x.UbluxPartner = UbluxPartner("UbluxPartner_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["voipCompany"]; ok && val != nil {
		x.VoipCompany = VoipCompany("VoipCompany_" + val.(string))
	} // is NOT readonly obtained from map
}
