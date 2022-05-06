package models

import . "github.com/ublux/go-models/enums"
import "go.mongodb.org/mongo-driver/bson/primitive"

type VoipNumberPhoneFilterRequest struct {
	City_CON                           string             `bson:"city_CON" json:"city_CON"`
	City_EQ                            string             `bson:"city_EQ" json:"city_EQ"`
	City_REG                           string             `bson:"city_REG" json:"city_REG"`
	CountryIsoCode_EQ                  CountryIsoCode     `bson:"countryIsoCode_EQ" json:"countryIsoCode_EQ"`
	DateCreated_EQ                     primitive.DateTime `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE                    primitive.DateTime `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE                    primitive.DateTime `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ                     primitive.DateTime `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE                    primitive.DateTime `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE                    primitive.DateTime `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	Description_CON                    string             `bson:"description_CON" json:"description_CON"`
	Description_EQ                     string             `bson:"description_EQ" json:"description_EQ"`
	Description_REG                    string             `bson:"description_REG" json:"description_REG"`
	FriendlyName_CON                   string             `bson:"friendlyName_CON" json:"friendlyName_CON"`
	FriendlyName_EQ                    string             `bson:"friendlyName_EQ" json:"friendlyName_EQ"`
	FriendlyName_REG                   string             `bson:"friendlyName_REG" json:"friendlyName_REG"`
	Id_CON                             string             `bson:"id_CON" json:"id_CON"`
	Id_EQ                              string             `bson:"id_EQ" json:"id_EQ"`
	Id_REG                             string             `bson:"id_REG" json:"id_REG"`
	InjectFriendlyNameToCallerId_EQ    bool               `bson:"injectFriendlyNameToCallerId_EQ" json:"injectFriendlyNameToCallerId_EQ"`
	IsSmsEnabled_EQ                    bool               `bson:"isSmsEnabled_EQ" json:"isSmsEnabled_EQ"`
	IsTollFree_EQ                      bool               `bson:"isTollFree_EQ" json:"isTollFree_EQ"`
	IsVoiceEnabled_EQ                  bool               `bson:"isVoiceEnabled_EQ" json:"isVoiceEnabled_EQ"`
	IsWhatsappEnabled_EQ               bool               `bson:"isWhatsappEnabled_EQ" json:"isWhatsappEnabled_EQ"`
	Language_EQ                        Language           `bson:"language_EQ" json:"language_EQ"`
	Number_CON                         string             `bson:"number_CON" json:"number_CON"`
	Number_EQ                          string             `bson:"number_EQ" json:"number_EQ"`
	Number_REG                         string             `bson:"number_REG" json:"number_REG"`
	RecordIncomingCalls_EQ             bool               `bson:"recordIncomingCalls_EQ" json:"recordIncomingCalls_EQ"`
	RulesFax_ForwardToEmailAddress_CON string             `bson:"rulesFax_ForwardToEmailAddress_CON" json:"rulesFax_ForwardToEmailAddress_CON"`
	RulesFax_ForwardToEmailAddress_EQ  string             `bson:"rulesFax_ForwardToEmailAddress_EQ" json:"rulesFax_ForwardToEmailAddress_EQ"`
	RulesFax_ForwardToEmailAddress_REG string             `bson:"rulesFax_ForwardToEmailAddress_REG" json:"rulesFax_ForwardToEmailAddress_REG"`
	RulesPhone_DaysOfWeek_CON          DayOfWeek          `bson:"rulesPhone_DaysOfWeek_CON" json:"rulesPhone_DaysOfWeek_CON"`
	RulesPhone_IdCallFlow_CON          string             `bson:"rulesPhone_IdCallFlow_CON" json:"rulesPhone_IdCallFlow_CON"`
	RulesPhone_IdCallFlow_EQ           string             `bson:"rulesPhone_IdCallFlow_EQ" json:"rulesPhone_IdCallFlow_EQ"`
	RulesPhone_IdCallFlow_REG          string             `bson:"rulesPhone_IdCallFlow_REG" json:"rulesPhone_IdCallFlow_REG"`
	RulesPhone_IdExtension_CON         string             `bson:"rulesPhone_IdExtension_CON" json:"rulesPhone_IdExtension_CON"`
	RulesPhone_IdExtension_EQ          string             `bson:"rulesPhone_IdExtension_EQ" json:"rulesPhone_IdExtension_EQ"`
	RulesPhone_IdExtension_REG         string             `bson:"rulesPhone_IdExtension_REG" json:"rulesPhone_IdExtension_REG"`
	RulesSms_ForwardToEmailAddress_CON string             `bson:"rulesSms_ForwardToEmailAddress_CON" json:"rulesSms_ForwardToEmailAddress_CON"`
	RulesSms_ForwardToEmailAddress_EQ  string             `bson:"rulesSms_ForwardToEmailAddress_EQ" json:"rulesSms_ForwardToEmailAddress_EQ"`
	RulesSms_ForwardToEmailAddress_REG string             `bson:"rulesSms_ForwardToEmailAddress_REG" json:"rulesSms_ForwardToEmailAddress_REG"`
	State_CON                          string             `bson:"state_CON" json:"state_CON"`
	State_EQ                           string             `bson:"state_EQ" json:"state_EQ"`
	State_REG                          string             `bson:"state_REG" json:"state_REG"`
	TimeZone_CON                       string             `bson:"timeZone_CON" json:"timeZone_CON"`
	TimeZone_EQ                        string             `bson:"timeZone_EQ" json:"timeZone_EQ"`
	TimeZone_REG                       string             `bson:"timeZone_REG" json:"timeZone_REG"`
	VoipNumberType_EQ                  VoipNumberType     `bson:"voipNumberType_EQ" json:"voipNumberType_EQ"`
}

// BUILDER from bson map:
func BuildVoipNumberPhoneFilterRequest(m map[string]interface{}, x *VoipNumberPhoneFilterRequest) {
	if val, ok := m["city_CON"]; ok && val != nil {
		x.City_CON = val.(string)
	}
	if val, ok := m["city_EQ"]; ok && val != nil {
		x.City_EQ = val.(string)
	}
	if val, ok := m["city_REG"]; ok && val != nil {
		x.City_REG = val.(string)
	}
	if val, ok := m["countryIsoCode_EQ"]; ok && val != nil {
		x.CountryIsoCode_EQ = CountryIsoCode("CountryIsoCode_EQ_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["dateCreated_EQ"]; ok && val != nil {
		x.DateCreated_EQ = val.(primitive.DateTime)
	}
	if val, ok := m["dateCreated_GTE"]; ok && val != nil {
		x.DateCreated_GTE = val.(primitive.DateTime)
	}
	if val, ok := m["dateCreated_LTE"]; ok && val != nil {
		x.DateCreated_LTE = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated_EQ"]; ok && val != nil {
		x.DateUpdated_EQ = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated_GTE"]; ok && val != nil {
		x.DateUpdated_GTE = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated_LTE"]; ok && val != nil {
		x.DateUpdated_LTE = val.(primitive.DateTime)
	}
	if val, ok := m["description_CON"]; ok && val != nil {
		x.Description_CON = val.(string)
	}
	if val, ok := m["description_EQ"]; ok && val != nil {
		x.Description_EQ = val.(string)
	}
	if val, ok := m["description_REG"]; ok && val != nil {
		x.Description_REG = val.(string)
	}
	if val, ok := m["friendlyName_CON"]; ok && val != nil {
		x.FriendlyName_CON = val.(string)
	}
	if val, ok := m["friendlyName_EQ"]; ok && val != nil {
		x.FriendlyName_EQ = val.(string)
	}
	if val, ok := m["friendlyName_REG"]; ok && val != nil {
		x.FriendlyName_REG = val.(string)
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
	if val, ok := m["injectFriendlyNameToCallerId_EQ"]; ok && val != nil {
		x.InjectFriendlyNameToCallerId_EQ = val.(bool)
	}
	if val, ok := m["isSmsEnabled_EQ"]; ok && val != nil {
		x.IsSmsEnabled_EQ = val.(bool)
	}
	if val, ok := m["isTollFree_EQ"]; ok && val != nil {
		x.IsTollFree_EQ = val.(bool)
	}
	if val, ok := m["isVoiceEnabled_EQ"]; ok && val != nil {
		x.IsVoiceEnabled_EQ = val.(bool)
	}
	if val, ok := m["isWhatsappEnabled_EQ"]; ok && val != nil {
		x.IsWhatsappEnabled_EQ = val.(bool)
	}
	if val, ok := m["language_EQ"]; ok && val != nil {
		x.Language_EQ = Language("Language_EQ_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["number_CON"]; ok && val != nil {
		x.Number_CON = val.(string)
	}
	if val, ok := m["number_EQ"]; ok && val != nil {
		x.Number_EQ = val.(string)
	}
	if val, ok := m["number_REG"]; ok && val != nil {
		x.Number_REG = val.(string)
	}
	if val, ok := m["recordIncomingCalls_EQ"]; ok && val != nil {
		x.RecordIncomingCalls_EQ = val.(bool)
	}
	if val, ok := m["rulesFax_ForwardToEmailAddress_CON"]; ok && val != nil {
		x.RulesFax_ForwardToEmailAddress_CON = val.(string)
	}
	if val, ok := m["rulesFax_ForwardToEmailAddress_EQ"]; ok && val != nil {
		x.RulesFax_ForwardToEmailAddress_EQ = val.(string)
	}
	if val, ok := m["rulesFax_ForwardToEmailAddress_REG"]; ok && val != nil {
		x.RulesFax_ForwardToEmailAddress_REG = val.(string)
	}
	if val, ok := m["rulesPhone_DaysOfWeek_CON"]; ok && val != nil {
		x.RulesPhone_DaysOfWeek_CON = DayOfWeek("RulesPhone_DaysOfWeek_CON_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["rulesPhone_IdCallFlow_CON"]; ok && val != nil {
		x.RulesPhone_IdCallFlow_CON = val.(string)
	}
	if val, ok := m["rulesPhone_IdCallFlow_EQ"]; ok && val != nil {
		x.RulesPhone_IdCallFlow_EQ = val.(string)
	}
	if val, ok := m["rulesPhone_IdCallFlow_REG"]; ok && val != nil {
		x.RulesPhone_IdCallFlow_REG = val.(string)
	}
	if val, ok := m["rulesPhone_IdExtension_CON"]; ok && val != nil {
		x.RulesPhone_IdExtension_CON = val.(string)
	}
	if val, ok := m["rulesPhone_IdExtension_EQ"]; ok && val != nil {
		x.RulesPhone_IdExtension_EQ = val.(string)
	}
	if val, ok := m["rulesPhone_IdExtension_REG"]; ok && val != nil {
		x.RulesPhone_IdExtension_REG = val.(string)
	}
	if val, ok := m["rulesSms_ForwardToEmailAddress_CON"]; ok && val != nil {
		x.RulesSms_ForwardToEmailAddress_CON = val.(string)
	}
	if val, ok := m["rulesSms_ForwardToEmailAddress_EQ"]; ok && val != nil {
		x.RulesSms_ForwardToEmailAddress_EQ = val.(string)
	}
	if val, ok := m["rulesSms_ForwardToEmailAddress_REG"]; ok && val != nil {
		x.RulesSms_ForwardToEmailAddress_REG = val.(string)
	}
	if val, ok := m["state_CON"]; ok && val != nil {
		x.State_CON = val.(string)
	}
	if val, ok := m["state_EQ"]; ok && val != nil {
		x.State_EQ = val.(string)
	}
	if val, ok := m["state_REG"]; ok && val != nil {
		x.State_REG = val.(string)
	}
	if val, ok := m["timeZone_CON"]; ok && val != nil {
		x.TimeZone_CON = val.(string)
	}
	if val, ok := m["timeZone_EQ"]; ok && val != nil {
		x.TimeZone_EQ = val.(string)
	}
	if val, ok := m["timeZone_REG"]; ok && val != nil {
		x.TimeZone_REG = val.(string)
	}
	if val, ok := m["voipNumberType_EQ"]; ok && val != nil {
		x.VoipNumberType_EQ = VoipNumberType("VoipNumberType_EQ_" + val.(string))
	} // is NOT readonly obtained from map
}
