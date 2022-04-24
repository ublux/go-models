package models

type AccountSettings struct {
	ContactCallerIdTemplate                        string `bson:"contactCallerIdTemplate" json:"contactCallerIdTemplate"`
	TurnOnRecordingOfExternalCallsWhenCreatingLine bool   `bson:"turnOnRecordingOfExternalCallsWhenCreatingLine" json:"turnOnRecordingOfExternalCallsWhenCreatingLine"`
	TurnOnRecordingOfInternalCallsWhenCreatingLine bool   `bson:"turnOnRecordingOfInternalCallsWhenCreatingLine" json:"turnOnRecordingOfInternalCallsWhenCreatingLine"`
}

// BUILDER from bson map:
func BuildAccountSettings(m map[string]interface{}, x *AccountSettings) {
	if val, ok := m["contactCallerIdTemplate"]; ok && val != nil {
		x.ContactCallerIdTemplate = val.(string)
	}
	if val, ok := m["turnOnRecordingOfExternalCallsWhenCreatingLine"]; ok && val != nil {
		x.TurnOnRecordingOfExternalCallsWhenCreatingLine = val.(bool)
	}
	if val, ok := m["turnOnRecordingOfInternalCallsWhenCreatingLine"]; ok && val != nil {
		x.TurnOnRecordingOfInternalCallsWhenCreatingLine = val.(bool)
	}
}
