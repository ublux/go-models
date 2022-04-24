package models

type Variable struct {
	JsonValue string `bson:"jsonValue" json:"jsonValue"`
	Name      string `bson:"name" json:"name"`
}

// BUILDER from bson map:
func BuildVariable(m map[string]interface{}, x *Variable) {
	if val, ok := m["jsonValue"]; ok && val != nil {
		x.JsonValue = val.(string)
	}
	if val, ok := m["name"]; ok && val != nil {
		x.Name = val.(string)
	}
}
