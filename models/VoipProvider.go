package models

import . "github.com/ublux/go-models/enums"
import "go.mongodb.org/mongo-driver/bson/primitive"

type VoipProvider struct {
	Country                CountryIsoCode     `bson:"country" json:"country"`
	DateCreated            primitive.DateTime `bson:"dateCreated" json:"dateCreated"`
	DateDeleted            primitive.DateTime `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated            primitive.DateTime `bson:"dateUpdated" json:"dateUpdated"`
	FriendlyName           string             `bson:"friendlyName" json:"friendlyName"`
	Id                     string             `bson:"_id" json:"id"`
	ProviderAccessToken    string             `bson:"providerAccessToken" json:"providerAccessToken"`
	ProviderAccountId      string             `bson:"providerAccountId" json:"providerAccountId"`
	ProviderOwnerAccountId string             `bson:"providerOwnerAccountId" json:"providerOwnerAccountId"`
	Status                 string             `bson:"status" json:"status"`
	UbluxPartner           UbluxPartner       `bson:"ubluxPartner" json:"ubluxPartner"`
	VoipCompany            VoipCompany        `bson:"voipCompany" json:"voipCompany"`
}

// Implementing interface IUbluxDocument
func (x VoipProvider) GetDateDeleted() primitive.DateTime {
	return x.DateDeleted
}
func (x VoipProvider) GetDateCreated() primitive.DateTime {
	return x.DateCreated
}
func (x VoipProvider) GetDateUpdated() primitive.DateTime {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x VoipProvider) GetId() string {
	return x.Id
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildVoipProvider(m map[string]interface{}, x *VoipProvider) {
	if val, ok := m["country"]; ok && val != nil {
		x.Country = CountryIsoCode("Country_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(primitive.DateTime)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(primitive.DateTime)
	}
	if val, ok := m["friendlyName"]; ok && val != nil {
		x.FriendlyName = val.(string)
	}
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["providerAccessToken"]; ok && val != nil {
		x.ProviderAccessToken = val.(string)
	}
	if val, ok := m["providerAccountId"]; ok && val != nil {
		x.ProviderAccountId = val.(string)
	}
	if val, ok := m["providerOwnerAccountId"]; ok && val != nil {
		x.ProviderOwnerAccountId = val.(string)
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
