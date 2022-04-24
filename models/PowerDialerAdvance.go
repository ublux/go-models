package models

import . "github.com/ublux/go-models/enums"

type PowerDialerAdvance struct {
	ContactName                    string            `bson:"contactName" json:"contactName"`
	CountryIsoCode                 CountryIsoCode    `bson:"countryIsoCode" json:"countryIsoCode"`
	ErrorMessage                   string            `bson:"errorMessage" json:"errorMessage"`
	IdCallFlow                     string            `bson:"idCallFlow" json:"idCallFlow"`
	IdContact                      string            `bson:"idContact" json:"idContact"`
	NumberOfAttempts               int32             `bson:"numberOfAttempts" json:"numberOfAttempts"`
	PhoneNumberInternationalFormat string            `bson:"phoneNumberInternationalFormat" json:"phoneNumberInternationalFormat"`
	PowerDialerStatus              PowerDialerStatus `bson:"powerDialerStatus" json:"powerDialerStatus"`
	PowerDialerType                PowerDialerType   `bson:"powerDialerType" json:"powerDialerType"`
}

// Implementing interface PowerDialer
func (x PowerDialerAdvance) GetPowerDialerType() PowerDialerType {
	return x.PowerDialerType
}
func (x PowerDialerAdvance) GetPhoneNumberInternationalFormat() string {
	return x.PhoneNumberInternationalFormat
}
func (x PowerDialerAdvance) GetCountryIsoCode() CountryIsoCode {
	return x.CountryIsoCode
}
func (x PowerDialerAdvance) GetIdContact() string {
	return x.IdContact
}
func (x PowerDialerAdvance) GetPowerDialerStatus() PowerDialerStatus {
	return x.PowerDialerStatus
}
func (x PowerDialerAdvance) GetErrorMessage() string {
	return x.ErrorMessage
}
func (x PowerDialerAdvance) GetNumberOfAttempts() int32 {
	return x.NumberOfAttempts
}
func (x PowerDialerAdvance) GetContactName() string {
	return x.ContactName
}

// BUILDER from bson map:
func BuildPowerDialerAdvance(m map[string]interface{}, x *PowerDialerAdvance) {
	if val, ok := m["contactName"]; ok && val != nil {
		x.ContactName = val.(string)
	}
	if val, ok := m["countryIsoCode"]; ok && val != nil {
		x.CountryIsoCode = CountryIsoCode("CountryIsoCode_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["errorMessage"]; ok && val != nil {
		x.ErrorMessage = val.(string)
	}
	if val, ok := m["idCallFlow"]; ok && val != nil {
		x.IdCallFlow = val.(string)
	}
	if val, ok := m["idContact"]; ok && val != nil {
		x.IdContact = val.(string)
	}
	if val, ok := m["numberOfAttempts"]; ok && val != nil {
		x.NumberOfAttempts = val.(int32)
	}
	if val, ok := m["phoneNumberInternationalFormat"]; ok && val != nil {
		x.PhoneNumberInternationalFormat = val.(string)
	}
	if val, ok := m["powerDialerStatus"]; ok && val != nil {
		x.PowerDialerStatus = PowerDialerStatus("PowerDialerStatus_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["powerDialerType"]; ok && val != nil {
		x.PowerDialerType = PowerDialerType("PowerDialerType_" + val.(string))
	} // is NOT readonly obtained from map
}
