package models

import . "github.com/ublux/go-models/enums"
import "go.mongodb.org/mongo-driver/bson/primitive"

type VoipNumberAvailableForPurchase struct {
	City                         string             `bson:"city" json:"city"`
	CountryIsoCode               CountryIsoCode     `bson:"countryIsoCode" json:"countryIsoCode"`
	DateCreated                  primitive.DateTime `bson:"dateCreated" json:"dateCreated"`
	DateDeleted                  primitive.DateTime `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated                  primitive.DateTime `bson:"dateUpdated" json:"dateUpdated"`
	Description                  string             `bson:"description" json:"description"`
	FriendlyName                 string             `bson:"friendlyName" json:"friendlyName"`
	Id                           string             `bson:"_id" json:"id"`
	IdAccount                    string             `bson:"idAccount" json:"idAccount"`
	IdCustomerInfo               string             `bson:"idCustomerInfo" json:"idCustomerInfo"`
	IdTrunkOrigination           string             `bson:"idTrunkOrigination" json:"idTrunkOrigination"`
	IdVoipProvider               string             `bson:"idVoipProvider" json:"idVoipProvider"`
	InjectFriendlyNameToCallerId bool               `bson:"injectFriendlyNameToCallerId" json:"injectFriendlyNameToCallerId"`
	IsSmsEnabled                 bool               `bson:"isSmsEnabled" json:"isSmsEnabled"`
	IsTollFree                   bool               `bson:"isTollFree" json:"isTollFree"`
	IsVoiceEnabled               bool               `bson:"isVoiceEnabled" json:"isVoiceEnabled"`
	IsWhatsappEnabled            bool               `bson:"isWhatsappEnabled" json:"isWhatsappEnabled"`
	Language                     Language           `bson:"language" json:"language"`
	Number                       string             `bson:"number" json:"number"`
	PendingToBeConfigured        bool               `bson:"pendingToBeConfigured" json:"pendingToBeConfigured"`
	ProviderId                   string             `bson:"providerId" json:"providerId"`
	RecordIncomingCalls          bool               `bson:"recordIncomingCalls" json:"recordIncomingCalls"`
	RequiresCustomerInfo         bool               `bson:"requiresCustomerInfo" json:"requiresCustomerInfo"`
	RulesFax                     []RuleFax          `bson:"rulesFax" json:"rulesFax"`
	RulesPhone                   []RulePhone        `bson:"rulesPhone" json:"rulesPhone"`
	RulesSms                     []RuleSms          `bson:"rulesSms" json:"rulesSms"`
	State                        string             `bson:"state" json:"state"`
	TimeZone                     string             `bson:"timeZone" json:"timeZone"`
	VoipNumberType               VoipNumberType     `bson:"voipNumberType" json:"voipNumberType"`
}

// Implementing interface IUbluxDocument
func (x VoipNumberAvailableForPurchase) GetDateDeleted() primitive.DateTime {
	return x.DateDeleted
}
func (x VoipNumberAvailableForPurchase) GetDateCreated() primitive.DateTime {
	return x.DateCreated
}
func (x VoipNumberAvailableForPurchase) GetDateUpdated() primitive.DateTime {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x VoipNumberAvailableForPurchase) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x VoipNumberAvailableForPurchase) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface VoipNumber
func (x VoipNumberAvailableForPurchase) GetIdCustomerInfo() string {
	return x.IdCustomerInfo
}
func (x VoipNumberAvailableForPurchase) GetIdTrunkOrigination() string {
	return x.IdTrunkOrigination
}
func (x VoipNumberAvailableForPurchase) GetIdVoipProvider() string {
	return x.IdVoipProvider
}
func (x VoipNumberAvailableForPurchase) GetProviderId() string {
	return x.ProviderId
}
func (x VoipNumberAvailableForPurchase) GetRulesPhone() []RulePhone {
	return x.RulesPhone
}
func (x VoipNumberAvailableForPurchase) GetRulesSms() []RuleSms {
	return x.RulesSms
}
func (x VoipNumberAvailableForPurchase) GetRulesFax() []RuleFax {
	return x.RulesFax
}
func (x VoipNumberAvailableForPurchase) GetVoipNumberType() VoipNumberType {
	return x.VoipNumberType
}
func (x VoipNumberAvailableForPurchase) GetInjectFriendlyNameToCallerId() bool {
	return x.InjectFriendlyNameToCallerId
}
func (x VoipNumberAvailableForPurchase) GetRecordIncomingCalls() bool {
	return x.RecordIncomingCalls
}
func (x VoipNumberAvailableForPurchase) GetNumber() string {
	return x.Number
}
func (x VoipNumberAvailableForPurchase) GetFriendlyName() string {
	return x.FriendlyName
}
func (x VoipNumberAvailableForPurchase) GetDescription() string {
	return x.Description
}
func (x VoipNumberAvailableForPurchase) GetLanguage() Language {
	return x.Language
}
func (x VoipNumberAvailableForPurchase) GetCity() string {
	return x.City
}
func (x VoipNumberAvailableForPurchase) GetState() string {
	return x.State
}
func (x VoipNumberAvailableForPurchase) GetCountryIsoCode() CountryIsoCode {
	return x.CountryIsoCode
}
func (x VoipNumberAvailableForPurchase) GetIsSmsEnabled() bool {
	return x.IsSmsEnabled
}
func (x VoipNumberAvailableForPurchase) GetIsVoiceEnabled() bool {
	return x.IsVoiceEnabled
}
func (x VoipNumberAvailableForPurchase) GetIsTollFree() bool {
	return x.IsTollFree
}
func (x VoipNumberAvailableForPurchase) GetIsWhatsappEnabled() bool {
	return x.IsWhatsappEnabled
}
func (x VoipNumberAvailableForPurchase) GetTimeZone() string {
	return x.TimeZone
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildVoipNumberAvailableForPurchase(m map[string]interface{}, x *VoipNumberAvailableForPurchase) {
	if val, ok := m["city"]; ok && val != nil {
		x.City = val.(string)
	}
	if val, ok := m["countryIsoCode"]; ok && val != nil {
		x.CountryIsoCode = CountryIsoCode("CountryIsoCode_" + val.(string))
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
	if val, ok := m["description"]; ok && val != nil {
		x.Description = val.(string)
	}
	if val, ok := m["friendlyName"]; ok && val != nil {
		x.FriendlyName = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["idAccount"]; ok && val != nil {
		x.IdAccount = val.(string)
	}
	if val, ok := m["idCustomerInfo"]; ok && val != nil {
		x.IdCustomerInfo = val.(string)
	}
	if val, ok := m["idTrunkOrigination"]; ok && val != nil {
		x.IdTrunkOrigination = val.(string)
	}
	if val, ok := m["idVoipProvider"]; ok && val != nil {
		x.IdVoipProvider = val.(string)
	}
	if val, ok := m["injectFriendlyNameToCallerId"]; ok && val != nil {
		x.InjectFriendlyNameToCallerId = val.(bool)
	}
	if val, ok := m["isSmsEnabled"]; ok && val != nil {
		x.IsSmsEnabled = val.(bool)
	}
	if val, ok := m["isTollFree"]; ok && val != nil {
		x.IsTollFree = val.(bool)
	}
	if val, ok := m["isVoiceEnabled"]; ok && val != nil {
		x.IsVoiceEnabled = val.(bool)
	}
	if val, ok := m["isWhatsappEnabled"]; ok && val != nil {
		x.IsWhatsappEnabled = val.(bool)
	}
	if val, ok := m["language"]; ok && val != nil {
		x.Language = Language("Language_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["number"]; ok && val != nil {
		x.Number = val.(string)
	}
	if val, ok := m["pendingToBeConfigured"]; ok && val != nil {
		x.PendingToBeConfigured = val.(bool)
	}
	if val, ok := m["providerId"]; ok && val != nil {
		x.ProviderId = val.(string)
	}
	if val, ok := m["recordIncomingCalls"]; ok && val != nil {
		x.RecordIncomingCalls = val.(bool)
	}
	if val, ok := m["requiresCustomerInfo"]; ok && val != nil {
		x.RequiresCustomerInfo = val.(bool)
	}
	if val, ok := m["rulesFax"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok {
			for _, val = range array {
				if val != nil {
					tmp := new(RuleFax)
					BuildRuleFax(val.(map[string]interface{}), tmp)
					x.RulesFax = append(x.RulesFax, *tmp)
				}
			}
		}
	}
	if val, ok := m["rulesPhone"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok {
			for _, val = range array {
				if val != nil {
					tmp := new(RulePhone)
					BuildRulePhone(val.(map[string]interface{}), tmp)
					x.RulesPhone = append(x.RulesPhone, *tmp)
				}
			}
		}
	}
	if val, ok := m["rulesSms"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok {
			for _, val = range array {
				if val != nil {
					tmp := new(RuleSms)
					BuildRuleSms(val.(map[string]interface{}), tmp)
					x.RulesSms = append(x.RulesSms, *tmp)
				}
			}
		}
	}
	if val, ok := m["state"]; ok && val != nil {
		x.State = val.(string)
	}
	if val, ok := m["timeZone"]; ok && val != nil {
		x.TimeZone = val.(string)
	}
	x.VoipNumberType = VoipNumberType_AvailableForPurchase // readonly property
}
