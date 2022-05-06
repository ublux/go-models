package models

import . "github.com/ublux/go-models/enums"

type PowerDialerSimple struct {
	ContactName       string            `bson:"contactName" json:"contactName"`
	CountryIsoCode    CountryIsoCode    `bson:"countryIsoCode" json:"countryIsoCode"`
	ErrorMessage      string            `bson:"errorMessage" json:"errorMessage"`
	IdContact         string            `bson:"idContact" json:"idContact"`
	IdLine            string            `bson:"idLine" json:"idLine"`
	NumberOfAttempts  int32             `bson:"numberOfAttempts" json:"numberOfAttempts"`
	PhoneNumber       string            `bson:"phoneNumber" json:"phoneNumber"`
	PowerDialerStatus PowerDialerStatus `bson:"powerDialerStatus" json:"powerDialerStatus"`
	PowerDialerType   PowerDialerType   `bson:"powerDialerType" json:"powerDialerType"`
}

// Implementing interface PowerDialer
func (x PowerDialerSimple) GetPowerDialerType() PowerDialerType {
	return x.PowerDialerType
}
func (x PowerDialerSimple) GetPhoneNumber() string {
	return x.PhoneNumber
}
func (x PowerDialerSimple) GetCountryIsoCode() CountryIsoCode {
	return x.CountryIsoCode
}
func (x PowerDialerSimple) GetIdContact() string {
	return x.IdContact
}
func (x PowerDialerSimple) GetPowerDialerStatus() PowerDialerStatus {
	return x.PowerDialerStatus
}
func (x PowerDialerSimple) GetErrorMessage() string {
	return x.ErrorMessage
}
func (x PowerDialerSimple) GetNumberOfAttempts() int32 {
	return x.NumberOfAttempts
}
func (x PowerDialerSimple) GetContactName() string {
	return x.ContactName
}

// BUILDER from bson map:
func BuildPowerDialerSimple(m map[string]interface{}, x *PowerDialerSimple) {
	if val, ok := m["contactName"]; ok && val != nil {
		x.ContactName = val.(string)
	}
	if val, ok := m["countryIsoCode"]; ok && val != nil {
		x.CountryIsoCode = CountryIsoCode("CountryIsoCode_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["errorMessage"]; ok && val != nil {
		x.ErrorMessage = val.(string)
	}
	if val, ok := m["idContact"]; ok && val != nil {
		x.IdContact = val.(string)
	}
	if val, ok := m["idLine"]; ok && val != nil {
		x.IdLine = val.(string)
	}
	if val, ok := m["numberOfAttempts"]; ok && val != nil {
		x.NumberOfAttempts = val.(int32)
	}
	if val, ok := m["phoneNumber"]; ok && val != nil {
		x.PhoneNumber = val.(string)
	}
	if val, ok := m["powerDialerStatus"]; ok && val != nil {
		x.PowerDialerStatus = PowerDialerStatus("PowerDialerStatus_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["powerDialerType"]; ok && val != nil {
		x.PowerDialerType = PowerDialerType("PowerDialerType_" + val.(string))
	} // is NOT readonly obtained from map
}
