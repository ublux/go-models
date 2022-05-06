package models

import . "github.com/ublux/go-models/enums"
import "time"

type AccountFilterRequest struct {
	AccountSecrets_PinPhone_CON                                       string         `bson:"accountSecrets_PinPhone_CON" json:"accountSecrets_PinPhone_CON"`
	AccountSecrets_PinPhone_EQ                                        string         `bson:"accountSecrets_PinPhone_EQ" json:"accountSecrets_PinPhone_EQ"`
	AccountSecrets_PinPhone_REG                                       string         `bson:"accountSecrets_PinPhone_REG" json:"accountSecrets_PinPhone_REG"`
	AccountSecrets_PinSpy_CON                                         string         `bson:"accountSecrets_PinSpy_CON" json:"accountSecrets_PinSpy_CON"`
	AccountSecrets_PinSpy_EQ                                          string         `bson:"accountSecrets_PinSpy_EQ" json:"accountSecrets_PinSpy_EQ"`
	AccountSecrets_PinSpy_REG                                         string         `bson:"accountSecrets_PinSpy_REG" json:"accountSecrets_PinSpy_REG"`
	AccountSettings_ContactCallerIdTemplate_CON                       string         `bson:"accountSettings_ContactCallerIdTemplate_CON" json:"accountSettings_ContactCallerIdTemplate_CON"`
	AccountSettings_ContactCallerIdTemplate_EQ                        string         `bson:"accountSettings_ContactCallerIdTemplate_EQ" json:"accountSettings_ContactCallerIdTemplate_EQ"`
	AccountSettings_ContactCallerIdTemplate_REG                       string         `bson:"accountSettings_ContactCallerIdTemplate_REG" json:"accountSettings_ContactCallerIdTemplate_REG"`
	AccountSettings_TurnOnRecordingOfExternalCallsWhenCreatingLine_EQ bool           `bson:"accountSettings_TurnOnRecordingOfExternalCallsWhenCreatingLine_EQ" json:"accountSettings_TurnOnRecordingOfExternalCallsWhenCreatingLine_EQ"`
	AccountSettings_TurnOnRecordingOfInternalCallsWhenCreatingLine_EQ bool           `bson:"accountSettings_TurnOnRecordingOfInternalCallsWhenCreatingLine_EQ" json:"accountSettings_TurnOnRecordingOfInternalCallsWhenCreatingLine_EQ"`
	CompanyName_CON                                                   string         `bson:"companyName_CON" json:"companyName_CON"`
	CompanyName_EQ                                                    string         `bson:"companyName_EQ" json:"companyName_EQ"`
	CompanyName_REG                                                   string         `bson:"companyName_REG" json:"companyName_REG"`
	CountriesThatCanCallInternationally_CON                           CountryIsoCode `bson:"countriesThatCanCallInternationally_CON" json:"countriesThatCanCallInternationally_CON"`
	CountriesThatCanCallLocally_CON                                   CountryIsoCode `bson:"countriesThatCanCallLocally_CON" json:"countriesThatCanCallLocally_CON"`
	DateCreated_EQ                                                    time.Time      `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE                                                   time.Time      `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE                                                   time.Time      `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ                                                    time.Time      `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE                                                   time.Time      `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE                                                   time.Time      `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	HasGrantedSupportAccess_EQ                                        bool           `bson:"hasGrantedSupportAccess_EQ" json:"hasGrantedSupportAccess_EQ"`
	Id_CON                                                            string         `bson:"id_CON" json:"id_CON"`
	Id_EQ                                                             string         `bson:"id_EQ" json:"id_EQ"`
	Id_REG                                                            string         `bson:"id_REG" json:"id_REG"`
	IdsCloudServicePbxs_CON                                           string         `bson:"idsCloudServicePbxs_CON" json:"idsCloudServicePbxs_CON"`
	IdsCloudServicePbxs_EQ                                            string         `bson:"idsCloudServicePbxs_EQ" json:"idsCloudServicePbxs_EQ"`
	IdsCloudServicePbxs_REG                                           string         `bson:"idsCloudServicePbxs_REG" json:"idsCloudServicePbxs_REG"`
	MailingAddress_ApparmentOrSuiteNumber_CON                         string         `bson:"mailingAddress_ApparmentOrSuiteNumber_CON" json:"mailingAddress_ApparmentOrSuiteNumber_CON"`
	MailingAddress_ApparmentOrSuiteNumber_EQ                          string         `bson:"mailingAddress_ApparmentOrSuiteNumber_EQ" json:"mailingAddress_ApparmentOrSuiteNumber_EQ"`
	MailingAddress_ApparmentOrSuiteNumber_REG                         string         `bson:"mailingAddress_ApparmentOrSuiteNumber_REG" json:"mailingAddress_ApparmentOrSuiteNumber_REG"`
	MailingAddress_BusinessName_CON                                   string         `bson:"mailingAddress_BusinessName_CON" json:"mailingAddress_BusinessName_CON"`
	MailingAddress_BusinessName_EQ                                    string         `bson:"mailingAddress_BusinessName_EQ" json:"mailingAddress_BusinessName_EQ"`
	MailingAddress_BusinessName_REG                                   string         `bson:"mailingAddress_BusinessName_REG" json:"mailingAddress_BusinessName_REG"`
	MailingAddress_City_CON                                           string         `bson:"mailingAddress_City_CON" json:"mailingAddress_City_CON"`
	MailingAddress_City_EQ                                            string         `bson:"mailingAddress_City_EQ" json:"mailingAddress_City_EQ"`
	MailingAddress_City_REG                                           string         `bson:"mailingAddress_City_REG" json:"mailingAddress_City_REG"`
	MailingAddress_Country_EQ                                         CountryIsoCode `bson:"mailingAddress_Country_EQ" json:"mailingAddress_Country_EQ"`
	MailingAddress_RecipientName_CON                                  string         `bson:"mailingAddress_RecipientName_CON" json:"mailingAddress_RecipientName_CON"`
	MailingAddress_RecipientName_EQ                                   string         `bson:"mailingAddress_RecipientName_EQ" json:"mailingAddress_RecipientName_EQ"`
	MailingAddress_RecipientName_REG                                  string         `bson:"mailingAddress_RecipientName_REG" json:"mailingAddress_RecipientName_REG"`
	MailingAddress_State_CON                                          string         `bson:"mailingAddress_State_CON" json:"mailingAddress_State_CON"`
	MailingAddress_State_EQ                                           string         `bson:"mailingAddress_State_EQ" json:"mailingAddress_State_EQ"`
	MailingAddress_State_REG                                          string         `bson:"mailingAddress_State_REG" json:"mailingAddress_State_REG"`
	MailingAddress_StreetAddress_CON                                  string         `bson:"mailingAddress_StreetAddress_CON" json:"mailingAddress_StreetAddress_CON"`
	MailingAddress_StreetAddress_EQ                                   string         `bson:"mailingAddress_StreetAddress_EQ" json:"mailingAddress_StreetAddress_EQ"`
	MailingAddress_StreetAddress_REG                                  string         `bson:"mailingAddress_StreetAddress_REG" json:"mailingAddress_StreetAddress_REG"`
	MailingAddress_ZipCode_CON                                        string         `bson:"mailingAddress_ZipCode_CON" json:"mailingAddress_ZipCode_CON"`
	MailingAddress_ZipCode_EQ                                         string         `bson:"mailingAddress_ZipCode_EQ" json:"mailingAddress_ZipCode_EQ"`
	MailingAddress_ZipCode_REG                                        string         `bson:"mailingAddress_ZipCode_REG" json:"mailingAddress_ZipCode_REG"`
}

// BUILDER from bson map:
func BuildAccountFilterRequest(m map[string]interface{}, x *AccountFilterRequest) {
	if val, ok := m["accountSecrets_PinPhone_CON"]; ok && val != nil {
		x.AccountSecrets_PinPhone_CON = val.(string)
	}
	if val, ok := m["accountSecrets_PinPhone_EQ"]; ok && val != nil {
		x.AccountSecrets_PinPhone_EQ = val.(string)
	}
	if val, ok := m["accountSecrets_PinPhone_REG"]; ok && val != nil {
		x.AccountSecrets_PinPhone_REG = val.(string)
	}
	if val, ok := m["accountSecrets_PinSpy_CON"]; ok && val != nil {
		x.AccountSecrets_PinSpy_CON = val.(string)
	}
	if val, ok := m["accountSecrets_PinSpy_EQ"]; ok && val != nil {
		x.AccountSecrets_PinSpy_EQ = val.(string)
	}
	if val, ok := m["accountSecrets_PinSpy_REG"]; ok && val != nil {
		x.AccountSecrets_PinSpy_REG = val.(string)
	}
	if val, ok := m["accountSettings_ContactCallerIdTemplate_CON"]; ok && val != nil {
		x.AccountSettings_ContactCallerIdTemplate_CON = val.(string)
	}
	if val, ok := m["accountSettings_ContactCallerIdTemplate_EQ"]; ok && val != nil {
		x.AccountSettings_ContactCallerIdTemplate_EQ = val.(string)
	}
	if val, ok := m["accountSettings_ContactCallerIdTemplate_REG"]; ok && val != nil {
		x.AccountSettings_ContactCallerIdTemplate_REG = val.(string)
	}
	if val, ok := m["accountSettings_TurnOnRecordingOfExternalCallsWhenCreatingLine_EQ"]; ok && val != nil {
		x.AccountSettings_TurnOnRecordingOfExternalCallsWhenCreatingLine_EQ = val.(bool)
	}
	if val, ok := m["accountSettings_TurnOnRecordingOfInternalCallsWhenCreatingLine_EQ"]; ok && val != nil {
		x.AccountSettings_TurnOnRecordingOfInternalCallsWhenCreatingLine_EQ = val.(bool)
	}
	if val, ok := m["companyName_CON"]; ok && val != nil {
		x.CompanyName_CON = val.(string)
	}
	if val, ok := m["companyName_EQ"]; ok && val != nil {
		x.CompanyName_EQ = val.(string)
	}
	if val, ok := m["companyName_REG"]; ok && val != nil {
		x.CompanyName_REG = val.(string)
	}
	if val, ok := m["countriesThatCanCallInternationally_CON"]; ok && val != nil {
		x.CountriesThatCanCallInternationally_CON = CountryIsoCode("CountriesThatCanCallInternationally_CON_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["countriesThatCanCallLocally_CON"]; ok && val != nil {
		x.CountriesThatCanCallLocally_CON = CountryIsoCode("CountriesThatCanCallLocally_CON_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["dateCreated_EQ"]; ok && val != nil {
		x.DateCreated_EQ = val.(time.Time)
	}
	if val, ok := m["dateCreated_GTE"]; ok && val != nil {
		x.DateCreated_GTE = val.(time.Time)
	}
	if val, ok := m["dateCreated_LTE"]; ok && val != nil {
		x.DateCreated_LTE = val.(time.Time)
	}
	if val, ok := m["dateUpdated_EQ"]; ok && val != nil {
		x.DateUpdated_EQ = val.(time.Time)
	}
	if val, ok := m["dateUpdated_GTE"]; ok && val != nil {
		x.DateUpdated_GTE = val.(time.Time)
	}
	if val, ok := m["dateUpdated_LTE"]; ok && val != nil {
		x.DateUpdated_LTE = val.(time.Time)
	}
	if val, ok := m["hasGrantedSupportAccess_EQ"]; ok && val != nil {
		x.HasGrantedSupportAccess_EQ = val.(bool)
	}
	if val, ok := m["id_CON"]; ok && val != nil {
		x.Id_CON = val.(string)
	}
	if val, ok := m["id_EQ"]; ok && val != nil {
		x.Id_EQ = val.(string)
	}
	if val, ok := m["id_REG"]; ok && val != nil {
		x.Id_REG = val.(string)
	}
	if val, ok := m["idsCloudServicePbxs_CON"]; ok && val != nil {
		x.IdsCloudServicePbxs_CON = val.(string)
	}
	if val, ok := m["idsCloudServicePbxs_EQ"]; ok && val != nil {
		x.IdsCloudServicePbxs_EQ = val.(string)
	}
	if val, ok := m["idsCloudServicePbxs_REG"]; ok && val != nil {
		x.IdsCloudServicePbxs_REG = val.(string)
	}
	if val, ok := m["mailingAddress_ApparmentOrSuiteNumber_CON"]; ok && val != nil {
		x.MailingAddress_ApparmentOrSuiteNumber_CON = val.(string)
	}
	if val, ok := m["mailingAddress_ApparmentOrSuiteNumber_EQ"]; ok && val != nil {
		x.MailingAddress_ApparmentOrSuiteNumber_EQ = val.(string)
	}
	if val, ok := m["mailingAddress_ApparmentOrSuiteNumber_REG"]; ok && val != nil {
		x.MailingAddress_ApparmentOrSuiteNumber_REG = val.(string)
	}
	if val, ok := m["mailingAddress_BusinessName_CON"]; ok && val != nil {
		x.MailingAddress_BusinessName_CON = val.(string)
	}
	if val, ok := m["mailingAddress_BusinessName_EQ"]; ok && val != nil {
		x.MailingAddress_BusinessName_EQ = val.(string)
	}
	if val, ok := m["mailingAddress_BusinessName_REG"]; ok && val != nil {
		x.MailingAddress_BusinessName_REG = val.(string)
	}
	if val, ok := m["mailingAddress_City_CON"]; ok && val != nil {
		x.MailingAddress_City_CON = val.(string)
	}
	if val, ok := m["mailingAddress_City_EQ"]; ok && val != nil {
		x.MailingAddress_City_EQ = val.(string)
	}
	if val, ok := m["mailingAddress_City_REG"]; ok && val != nil {
		x.MailingAddress_City_REG = val.(string)
	}
	if val, ok := m["mailingAddress_Country_EQ"]; ok && val != nil {
		x.MailingAddress_Country_EQ = CountryIsoCode("MailingAddress_Country_EQ_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["mailingAddress_RecipientName_CON"]; ok && val != nil {
		x.MailingAddress_RecipientName_CON = val.(string)
	}
	if val, ok := m["mailingAddress_RecipientName_EQ"]; ok && val != nil {
		x.MailingAddress_RecipientName_EQ = val.(string)
	}
	if val, ok := m["mailingAddress_RecipientName_REG"]; ok && val != nil {
		x.MailingAddress_RecipientName_REG = val.(string)
	}
	if val, ok := m["mailingAddress_State_CON"]; ok && val != nil {
		x.MailingAddress_State_CON = val.(string)
	}
	if val, ok := m["mailingAddress_State_EQ"]; ok && val != nil {
		x.MailingAddress_State_EQ = val.(string)
	}
	if val, ok := m["mailingAddress_State_REG"]; ok && val != nil {
		x.MailingAddress_State_REG = val.(string)
	}
	if val, ok := m["mailingAddress_StreetAddress_CON"]; ok && val != nil {
		x.MailingAddress_StreetAddress_CON = val.(string)
	}
	if val, ok := m["mailingAddress_StreetAddress_EQ"]; ok && val != nil {
		x.MailingAddress_StreetAddress_EQ = val.(string)
	}
	if val, ok := m["mailingAddress_StreetAddress_REG"]; ok && val != nil {
		x.MailingAddress_StreetAddress_REG = val.(string)
	}
	if val, ok := m["mailingAddress_ZipCode_CON"]; ok && val != nil {
		x.MailingAddress_ZipCode_CON = val.(string)
	}
	if val, ok := m["mailingAddress_ZipCode_EQ"]; ok && val != nil {
		x.MailingAddress_ZipCode_EQ = val.(string)
	}
	if val, ok := m["mailingAddress_ZipCode_REG"]; ok && val != nil {
		x.MailingAddress_ZipCode_REG = val.(string)
	}
}
