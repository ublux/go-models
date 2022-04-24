package models

type ValidationError struct {
	ErrorMessage  string      `bson:"errorMessage" json:"errorMessage"`
	PropertyName  string      `bson:"propertyName" json:"propertyName"`
	PropertyValue interface{} `bson:"propertyValue" json:"propertyValue"`
}

// BUILDER from bson map:
func BuildValidationError(m map[string]interface{}, x *ValidationError) {
	if val, ok := m["errorMessage"]; ok && val != nil {
		x.ErrorMessage = val.(string)
	}
	if val, ok := m["propertyName"]; ok && val != nil {
		x.PropertyName = val.(string)
	}
	if val, ok := m["propertyValue"]; ok && val != nil {
		x.PropertyValue = val
	}
}
