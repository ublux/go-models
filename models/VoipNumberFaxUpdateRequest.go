package models

import . "github.com/ublux/go-models/enums"
import "go.mongodb.org/mongo-driver/bson/primitive"

type VoipNumberFaxUpdateRequest struct {
	Description                  string      `bson:"description" json:"description"`
	FriendlyName                 string      `bson:"friendlyName" json:"friendlyName"`
	Id                           string      `bson:"id" json:"id"`
	InjectFriendlyNameToCallerId bool        `bson:"injectFriendlyNameToCallerId" json:"injectFriendlyNameToCallerId"`
	Language                     Language    `bson:"language" json:"language"`
	RecordIncomingCalls          bool        `bson:"recordIncomingCalls" json:"recordIncomingCalls"`
	RulesFax                     []RuleFax   `bson:"rulesFax" json:"rulesFax"`
	RulesPhone                   []RulePhone `bson:"rulesPhone" json:"rulesPhone"`
	RulesSms                     []RuleSms   `bson:"rulesSms" json:"rulesSms"`
	TimeZone                     string      `bson:"timeZone" json:"timeZone"`
}

// Implementing interface IUbluxDocumentId
func (x VoipNumberFaxUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildVoipNumberFaxUpdateRequest(m map[string]interface{}, x *VoipNumberFaxUpdateRequest) {
	if val, ok := m["description"]; ok && val != nil {
		x.Description = val.(string)
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
	if val, ok := m["injectFriendlyNameToCallerId"]; ok && val != nil {
		x.InjectFriendlyNameToCallerId = val.(bool)
	}
	if val, ok := m["language"]; ok && val != nil {
		x.Language = Language("Language_" + val.(string))
	} // is NOT readonly obtained from map
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
	if val, ok := m["timeZone"]; ok && val != nil {
		x.TimeZone = val.(string)
	}
}
