package models

import . "github.com/ublux/go-models/enums"

type MailingAddress struct {
	ApparmentOrSuiteNumber string         `bson:"apparmentOrSuiteNumber" json:"apparmentOrSuiteNumber"`
	BusinessName           string         `bson:"businessName" json:"businessName"`
	City                   string         `bson:"city" json:"city"`
	Country                CountryIsoCode `bson:"country" json:"country"`
	RecipientName          string         `bson:"recipientName" json:"recipientName"`
	State                  string         `bson:"state" json:"state"`
	StreetAddress          string         `bson:"streetAddress" json:"streetAddress"`
	ZipCode                string         `bson:"zipCode" json:"zipCode"`
}

// BUILDER from bson map:
func BuildMailingAddress(m map[string]interface{}, x *MailingAddress) {
	if val, ok := m["apparmentOrSuiteNumber"]; ok && val != nil {
		x.ApparmentOrSuiteNumber = val.(string)
	}
	if val, ok := m["businessName"]; ok && val != nil {
		x.BusinessName = val.(string)
	}
	if val, ok := m["city"]; ok && val != nil {
		x.City = val.(string)
	}
	if val, ok := m["country"]; ok && val != nil {
		x.Country = CountryIsoCode("Country_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["recipientName"]; ok && val != nil {
		x.RecipientName = val.(string)
	}
	if val, ok := m["state"]; ok && val != nil {
		x.State = val.(string)
	}
	if val, ok := m["streetAddress"]; ok && val != nil {
		x.StreetAddress = val.(string)
	}
	if val, ok := m["zipCode"]; ok && val != nil {
		x.ZipCode = val.(string)
	}
}
