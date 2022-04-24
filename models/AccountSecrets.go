package models

type AccountSecrets struct {
	PinPhone string `bson:"pinPhone" json:"pinPhone"`
	PinSpy   string `bson:"pinSpy" json:"pinSpy"`
}

// BUILDER from bson map:
func BuildAccountSecrets(m map[string]interface{}, x *AccountSecrets) {
	if val, ok := m["pinPhone"]; ok && val != nil {
		x.PinPhone = val.(string)
	}
	if val, ok := m["pinSpy"]; ok && val != nil {
		x.PinSpy = val.(string)
	}
}
