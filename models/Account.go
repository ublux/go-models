package models

import . "github.com/ublux/go-models/enums"
import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	AccountSecrets                      AccountSecrets     `bson:"accountSecrets" json:"accountSecrets"`
	AccountSettings                     AccountSettings    `bson:"accountSettings" json:"accountSettings"`
	CompanyName                         string             `bson:"companyName" json:"companyName"`
	CountriesThatCanCallInternationally []CountryIsoCode   `bson:"countriesThatCanCallInternationally" json:"countriesThatCanCallInternationally"`
	CountriesThatCanCallLocally         []CountryIsoCode   `bson:"countriesThatCanCallLocally" json:"countriesThatCanCallLocally"`
	DateCreated                         primitive.DateTime `bson:"dateCreated" json:"dateCreated"`
	DateDeleted                         primitive.DateTime `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated                         primitive.DateTime `bson:"dateUpdated" json:"dateUpdated"`
	HasGrantedSupportAccess             bool               `bson:"hasGrantedSupportAccess" json:"hasGrantedSupportAccess"`
	Id                                  string             `bson:"_id" json:"id"`
	IdCloudServicePbxFailover           string             `bson:"idCloudServicePbxFailover" json:"idCloudServicePbxFailover"`
	IdGTrunkTerminationGroup            string             `bson:"idGTrunkTerminationGroup" json:"idGTrunkTerminationGroup"`
	IdsCloudServicePbxs                 []string           `bson:"idsCloudServicePbxs" json:"idsCloudServicePbxs"`
	MailingAddress                      MailingAddress     `bson:"mailingAddress" json:"mailingAddress"`
	ReserveAccountsOnPhone              []int32            `bson:"reserveAccountsOnPhone" json:"reserveAccountsOnPhone"`
	UbluxPartner                        UbluxPartner       `bson:"ubluxPartner" json:"ubluxPartner"`
}

// Implementing interface IUbluxDocument
func (x Account) GetDateDeleted() primitive.DateTime {
	return x.DateDeleted
}
func (x Account) GetDateCreated() primitive.DateTime {
	return x.DateCreated
}
func (x Account) GetDateUpdated() primitive.DateTime {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x Account) GetId() string {
	return x.Id
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildAccount(m map[string]interface{}, x *Account) {
	if val, ok := m["accountSecrets"]; ok && val != nil {
		BuildAccountSecrets(val.(map[string]interface{}), &x.AccountSecrets)
	}
	if val, ok := m["accountSettings"]; ok && val != nil {
		BuildAccountSettings(val.(map[string]interface{}), &x.AccountSettings)
	}
	if val, ok := m["companyName"]; ok && val != nil {
		x.CompanyName = val.(string)
	}
	if val, ok := m["countriesThatCanCallInternationally"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				x.CountriesThatCanCallInternationally = append(x.CountriesThatCanCallInternationally, CountryIsoCode("CountriesThatCanCallInternationally_"+val.(string)))
			} // is NOT readonly obtained from map
		}
	}
	if val, ok := m["countriesThatCanCallLocally"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				x.CountriesThatCanCallLocally = append(x.CountriesThatCanCallLocally, CountryIsoCode("CountriesThatCanCallLocally_"+val.(string)))
			} // is NOT readonly obtained from map
		}
	}
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(primitive.DateTime)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(primitive.DateTime)
	}
	if val, ok := m["hasGrantedSupportAccess"]; ok && val != nil {
		x.HasGrantedSupportAccess = val.(bool)
	}
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["idCloudServicePbxFailover"]; ok && val != nil {
		x.IdCloudServicePbxFailover = val.(string)
	}
	if val, ok := m["idGTrunkTerminationGroup"]; ok && val != nil {
		x.IdGTrunkTerminationGroup = val.(string)
	}
	if val, ok := m["idsCloudServicePbxs"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				if val != nil {
					x.IdsCloudServicePbxs = append(x.IdsCloudServicePbxs, val.(string))
				}
			}
		}
	}
	if val, ok := m["mailingAddress"]; ok && val != nil {
		BuildMailingAddress(val.(map[string]interface{}), &x.MailingAddress)
	}
	if val, ok := m["reserveAccountsOnPhone"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				if val != nil {
					x.ReserveAccountsOnPhone = append(x.ReserveAccountsOnPhone, val.(int32))
				}
			}
		}
	}
	if val, ok := m["ubluxPartner"]; ok && val != nil {
		x.UbluxPartner = UbluxPartner("UbluxPartner_" + val.(string))
	} // is NOT readonly obtained from map
}
