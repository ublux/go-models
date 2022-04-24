package models

type RuleSms struct {
	ForwardToEmailAddress string `bson:"forwardToEmailAddress" json:"forwardToEmailAddress"`
}

// BUILDER from bson map:
func BuildRuleSms(m map[string]interface{}, x *RuleSms) {
	if val, ok := m["forwardToEmailAddress"]; ok && val != nil {
		x.ForwardToEmailAddress = val.(string)
	}
}
