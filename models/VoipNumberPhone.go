package models

import . "github.com/ublux/go-models/enums"
import "time"
import "go.mongodb.org/mongo-driver/bson/primitive"

type VoipNumberPhone struct {
	City                         string         `bson:"city" json:"city"`
	CountryIsoCode               CountryIsoCode `bson:"countryIsoCode" json:"countryIsoCode"`
	DateCreated                  time.Time      `bson:"dateCreated" json:"dateCreated"`
	DateDeleted                  time.Time      `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated                  time.Time      `bson:"dateUpdated" json:"dateUpdated"`
	Description                  string         `bson:"description" json:"description"`
	FriendlyName                 string         `bson:"friendlyName" json:"friendlyName"`
	Id                           string         `bson:"id" json:"id"`
	IdAccount                    string         `bson:"idAccount" json:"idAccount"`
	IdCustomerInfo               string         `bson:"idCustomerInfo" json:"idCustomerInfo"`
	IdTrunkOrigination           string         `bson:"idTrunkOrigination" json:"idTrunkOrigination"`
	IdVoipProvider               string         `bson:"idVoipProvider" json:"idVoipProvider"`
	InjectFriendlyNameToCallerId bool           `bson:"injectFriendlyNameToCallerId" json:"injectFriendlyNameToCallerId"`
	IsSmsEnabled                 bool           `bson:"isSmsEnabled" json:"isSmsEnabled"`
	IsTollFree                   bool           `bson:"isTollFree" json:"isTollFree"`
	IsVoiceEnabled               bool           `bson:"isVoiceEnabled" json:"isVoiceEnabled"`
	IsWhatsappEnabled            bool           `bson:"isWhatsappEnabled" json:"isWhatsappEnabled"`
	Language                     Language       `bson:"language" json:"language"`
	Number                       string         `bson:"number" json:"number"`
	RecordIncomingCalls          bool           `bson:"recordIncomingCalls" json:"recordIncomingCalls"`
	RulesFax                     []RuleFax      `bson:"rulesFax" json:"rulesFax"`
	RulesPhone                   []RulePhone    `bson:"rulesPhone" json:"rulesPhone"`
	RulesSms                     []RuleSms      `bson:"rulesSms" json:"rulesSms"`
	SID                          string         `bson:"sID" json:"sID"`
	State                        string         `bson:"state" json:"state"`
	TimeZone                     string         `bson:"timeZone" json:"timeZone"`
	VoipNumberType               VoipNumberType `bson:"voipNumberType" json:"voipNumberType"`
}

// Implementing interface IUbluxDocument
func (x VoipNumberPhone) GetDateDeleted() time.Time {
	return x.DateDeleted
}
func (x VoipNumberPhone) GetDateCreated() time.Time {
	return x.DateCreated
}
func (x VoipNumberPhone) GetDateUpdated() time.Time {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x VoipNumberPhone) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x VoipNumberPhone) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface VoipNumber
func (x VoipNumberPhone) GetIdCustomerInfo() string {
	return x.IdCustomerInfo
}
func (x VoipNumberPhone) GetIdTrunkOrigination() string {
	return x.IdTrunkOrigination
}
func (x VoipNumberPhone) GetIdVoipProvider() string {
	return x.IdVoipProvider
}
func (x VoipNumberPhone) GetSID() string {
	return x.SID
}
func (x VoipNumberPhone) GetRulesPhone() []RulePhone {
	return x.RulesPhone
}
func (x VoipNumberPhone) GetRulesSms() []RuleSms {
	return x.RulesSms
}
func (x VoipNumberPhone) GetRulesFax() []RuleFax {
	return x.RulesFax
}
func (x VoipNumberPhone) GetVoipNumberType() VoipNumberType {
	return x.VoipNumberType
}
func (x VoipNumberPhone) GetInjectFriendlyNameToCallerId() bool {
	return x.InjectFriendlyNameToCallerId
}
func (x VoipNumberPhone) GetRecordIncomingCalls() bool {
	return x.RecordIncomingCalls
}
func (x VoipNumberPhone) GetNumber() string {
	return x.Number
}
func (x VoipNumberPhone) GetFriendlyName() string {
	return x.FriendlyName
}
func (x VoipNumberPhone) GetDescription() string {
	return x.Description
}
func (x VoipNumberPhone) GetLanguage() Language {
	return x.Language
}
func (x VoipNumberPhone) GetCity() string {
	return x.City
}
func (x VoipNumberPhone) GetState() string {
	return x.State
}
func (x VoipNumberPhone) GetCountryIsoCode() CountryIsoCode {
	return x.CountryIsoCode
}
func (x VoipNumberPhone) GetIsSmsEnabled() bool {
	return x.IsSmsEnabled
}
func (x VoipNumberPhone) GetIsVoiceEnabled() bool {
	return x.IsVoiceEnabled
}
func (x VoipNumberPhone) GetIsTollFree() bool {
	return x.IsTollFree
}
func (x VoipNumberPhone) GetIsWhatsappEnabled() bool {
	return x.IsWhatsappEnabled
}
func (x VoipNumberPhone) GetTimeZone() string {
	return x.TimeZone
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildVoipNumberPhone(m map[string]interface{}, x *VoipNumberPhone) {
	if val, ok := m["city"]; ok && val != nil {
		x.City = val.(string)
	}
	if val, ok := m["countryIsoCode"]; ok && val != nil {
		x.CountryIsoCode = CountryIsoCode("CountryIsoCode_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(time.Time)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(time.Time)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(time.Time)
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
	if val, ok := m["recordIncomingCalls"]; ok && val != nil {
		x.RecordIncomingCalls = val.(bool)
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
	if val, ok := m["sID"]; ok && val != nil {
		x.SID = val.(string)
	}
	if val, ok := m["state"]; ok && val != nil {
		x.State = val.(string)
	}
	if val, ok := m["timeZone"]; ok && val != nil {
		x.TimeZone = val.(string)
	}
	x.VoipNumberType = VoipNumberType_Phone // readonly property
}
