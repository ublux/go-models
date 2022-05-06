package models

import "time"
import . "github.com/ublux/go-models/enums"

type CustomerInfoFilterRequest struct {
	AirNetworksCustomerInfo_IdAirNetworksProvince_CON string         `bson:"airNetworksCustomerInfo_IdAirNetworksProvince_CON" json:"airNetworksCustomerInfo_IdAirNetworksProvince_CON"`
	AirNetworksCustomerInfo_IdAirNetworksProvince_EQ  string         `bson:"airNetworksCustomerInfo_IdAirNetworksProvince_EQ" json:"airNetworksCustomerInfo_IdAirNetworksProvince_EQ"`
	AirNetworksCustomerInfo_IdAirNetworksProvince_REG string         `bson:"airNetworksCustomerInfo_IdAirNetworksProvince_REG" json:"airNetworksCustomerInfo_IdAirNetworksProvince_REG"`
	AirNetworksCustomerInfo_IdNumber_CON              string         `bson:"airNetworksCustomerInfo_IdNumber_CON" json:"airNetworksCustomerInfo_IdNumber_CON"`
	AirNetworksCustomerInfo_IdNumber_EQ               string         `bson:"airNetworksCustomerInfo_IdNumber_EQ" json:"airNetworksCustomerInfo_IdNumber_EQ"`
	AirNetworksCustomerInfo_IdNumber_REG              string         `bson:"airNetworksCustomerInfo_IdNumber_REG" json:"airNetworksCustomerInfo_IdNumber_REG"`
	AirNetworksCustomerInfo_IdType_CON                string         `bson:"airNetworksCustomerInfo_IdType_CON" json:"airNetworksCustomerInfo_IdType_CON"`
	AirNetworksCustomerInfo_IdType_EQ                 string         `bson:"airNetworksCustomerInfo_IdType_EQ" json:"airNetworksCustomerInfo_IdType_EQ"`
	AirNetworksCustomerInfo_IdType_REG                string         `bson:"airNetworksCustomerInfo_IdType_REG" json:"airNetworksCustomerInfo_IdType_REG"`
	AirNetworksCustomerInfo_Population_CON            string         `bson:"airNetworksCustomerInfo_Population_CON" json:"airNetworksCustomerInfo_Population_CON"`
	AirNetworksCustomerInfo_Population_EQ             string         `bson:"airNetworksCustomerInfo_Population_EQ" json:"airNetworksCustomerInfo_Population_EQ"`
	AirNetworksCustomerInfo_Population_REG            string         `bson:"airNetworksCustomerInfo_Population_REG" json:"airNetworksCustomerInfo_Population_REG"`
	AirNetworksCustomerInfo_Province_CON              string         `bson:"airNetworksCustomerInfo_Province_CON" json:"airNetworksCustomerInfo_Province_CON"`
	AirNetworksCustomerInfo_Province_EQ               string         `bson:"airNetworksCustomerInfo_Province_EQ" json:"airNetworksCustomerInfo_Province_EQ"`
	AirNetworksCustomerInfo_Province_REG              string         `bson:"airNetworksCustomerInfo_Province_REG" json:"airNetworksCustomerInfo_Province_REG"`
	DateCreated_EQ                                    time.Time      `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE                                   time.Time      `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE                                   time.Time      `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ                                    time.Time      `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE                                   time.Time      `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE                                   time.Time      `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	FullName_CON                                      string         `bson:"fullName_CON" json:"fullName_CON"`
	FullName_EQ                                       string         `bson:"fullName_EQ" json:"fullName_EQ"`
	FullName_REG                                      string         `bson:"fullName_REG" json:"fullName_REG"`
	Id_CON                                            string         `bson:"id_CON" json:"id_CON"`
	Id_EQ                                             string         `bson:"id_EQ" json:"id_EQ"`
	Id_REG                                            string         `bson:"id_REG" json:"id_REG"`
	MailingAddress_ApparmentOrSuiteNumber_CON         string         `bson:"mailingAddress_ApparmentOrSuiteNumber_CON" json:"mailingAddress_ApparmentOrSuiteNumber_CON"`
	MailingAddress_ApparmentOrSuiteNumber_EQ          string         `bson:"mailingAddress_ApparmentOrSuiteNumber_EQ" json:"mailingAddress_ApparmentOrSuiteNumber_EQ"`
	MailingAddress_ApparmentOrSuiteNumber_REG         string         `bson:"mailingAddress_ApparmentOrSuiteNumber_REG" json:"mailingAddress_ApparmentOrSuiteNumber_REG"`
	MailingAddress_BusinessName_CON                   string         `bson:"mailingAddress_BusinessName_CON" json:"mailingAddress_BusinessName_CON"`
	MailingAddress_BusinessName_EQ                    string         `bson:"mailingAddress_BusinessName_EQ" json:"mailingAddress_BusinessName_EQ"`
	MailingAddress_BusinessName_REG                   string         `bson:"mailingAddress_BusinessName_REG" json:"mailingAddress_BusinessName_REG"`
	MailingAddress_City_CON                           string         `bson:"mailingAddress_City_CON" json:"mailingAddress_City_CON"`
	MailingAddress_City_EQ                            string         `bson:"mailingAddress_City_EQ" json:"mailingAddress_City_EQ"`
	MailingAddress_City_REG                           string         `bson:"mailingAddress_City_REG" json:"mailingAddress_City_REG"`
	MailingAddress_Country_EQ                         CountryIsoCode `bson:"mailingAddress_Country_EQ" json:"mailingAddress_Country_EQ"`
	MailingAddress_RecipientName_CON                  string         `bson:"mailingAddress_RecipientName_CON" json:"mailingAddress_RecipientName_CON"`
	MailingAddress_RecipientName_EQ                   string         `bson:"mailingAddress_RecipientName_EQ" json:"mailingAddress_RecipientName_EQ"`
	MailingAddress_RecipientName_REG                  string         `bson:"mailingAddress_RecipientName_REG" json:"mailingAddress_RecipientName_REG"`
	MailingAddress_State_CON                          string         `bson:"mailingAddress_State_CON" json:"mailingAddress_State_CON"`
	MailingAddress_State_EQ                           string         `bson:"mailingAddress_State_EQ" json:"mailingAddress_State_EQ"`
	MailingAddress_State_REG                          string         `bson:"mailingAddress_State_REG" json:"mailingAddress_State_REG"`
	MailingAddress_StreetAddress_CON                  string         `bson:"mailingAddress_StreetAddress_CON" json:"mailingAddress_StreetAddress_CON"`
	MailingAddress_StreetAddress_EQ                   string         `bson:"mailingAddress_StreetAddress_EQ" json:"mailingAddress_StreetAddress_EQ"`
	MailingAddress_StreetAddress_REG                  string         `bson:"mailingAddress_StreetAddress_REG" json:"mailingAddress_StreetAddress_REG"`
	MailingAddress_ZipCode_CON                        string         `bson:"mailingAddress_ZipCode_CON" json:"mailingAddress_ZipCode_CON"`
	MailingAddress_ZipCode_EQ                         string         `bson:"mailingAddress_ZipCode_EQ" json:"mailingAddress_ZipCode_EQ"`
	MailingAddress_ZipCode_REG                        string         `bson:"mailingAddress_ZipCode_REG" json:"mailingAddress_ZipCode_REG"`
}

// BUILDER from bson map:
func BuildCustomerInfoFilterRequest(m map[string]interface{}, x *CustomerInfoFilterRequest) {
	if val, ok := m["airNetworksCustomerInfo_IdAirNetworksProvince_CON"]; ok && val != nil {
		x.AirNetworksCustomerInfo_IdAirNetworksProvince_CON = val.(string)
	}
	if val, ok := m["airNetworksCustomerInfo_IdAirNetworksProvince_EQ"]; ok && val != nil {
		x.AirNetworksCustomerInfo_IdAirNetworksProvince_EQ = val.(string)
	}
	if val, ok := m["airNetworksCustomerInfo_IdAirNetworksProvince_REG"]; ok && val != nil {
		x.AirNetworksCustomerInfo_IdAirNetworksProvince_REG = val.(string)
	}
	if val, ok := m["airNetworksCustomerInfo_IdNumber_CON"]; ok && val != nil {
		x.AirNetworksCustomerInfo_IdNumber_CON = val.(string)
	}
	if val, ok := m["airNetworksCustomerInfo_IdNumber_EQ"]; ok && val != nil {
		x.AirNetworksCustomerInfo_IdNumber_EQ = val.(string)
	}
	if val, ok := m["airNetworksCustomerInfo_IdNumber_REG"]; ok && val != nil {
		x.AirNetworksCustomerInfo_IdNumber_REG = val.(string)
	}
	if val, ok := m["airNetworksCustomerInfo_IdType_CON"]; ok && val != nil {
		x.AirNetworksCustomerInfo_IdType_CON = val.(string)
	}
	if val, ok := m["airNetworksCustomerInfo_IdType_EQ"]; ok && val != nil {
		x.AirNetworksCustomerInfo_IdType_EQ = val.(string)
	}
	if val, ok := m["airNetworksCustomerInfo_IdType_REG"]; ok && val != nil {
		x.AirNetworksCustomerInfo_IdType_REG = val.(string)
	}
	if val, ok := m["airNetworksCustomerInfo_Population_CON"]; ok && val != nil {
		x.AirNetworksCustomerInfo_Population_CON = val.(string)
	}
	if val, ok := m["airNetworksCustomerInfo_Population_EQ"]; ok && val != nil {
		x.AirNetworksCustomerInfo_Population_EQ = val.(string)
	}
	if val, ok := m["airNetworksCustomerInfo_Population_REG"]; ok && val != nil {
		x.AirNetworksCustomerInfo_Population_REG = val.(string)
	}
	if val, ok := m["airNetworksCustomerInfo_Province_CON"]; ok && val != nil {
		x.AirNetworksCustomerInfo_Province_CON = val.(string)
	}
	if val, ok := m["airNetworksCustomerInfo_Province_EQ"]; ok && val != nil {
		x.AirNetworksCustomerInfo_Province_EQ = val.(string)
	}
	if val, ok := m["airNetworksCustomerInfo_Province_REG"]; ok && val != nil {
		x.AirNetworksCustomerInfo_Province_REG = val.(string)
	}
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
	if val, ok := m["fullName_CON"]; ok && val != nil {
		x.FullName_CON = val.(string)
	}
	if val, ok := m["fullName_EQ"]; ok && val != nil {
		x.FullName_EQ = val.(string)
	}
	if val, ok := m["fullName_REG"]; ok && val != nil {
		x.FullName_REG = val.(string)
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
