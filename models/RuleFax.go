package models

type RuleFax struct {
	ForwardToEmailAddress string `bson:"forwardToEmailAddress" json:"forwardToEmailAddress"`
}

// BUILDER from bson map:
func BuildRuleFax(m map[string]interface{}, x *RuleFax) {
	if val, ok := m["forwardToEmailAddress"]; ok && val != nil {
		x.ForwardToEmailAddress = val.(string)
	}
}
