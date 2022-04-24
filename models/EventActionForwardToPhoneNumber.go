package models

import . "github.com/ublux/go-models/enums"

type EventActionForwardToPhoneNumber struct {
	EventActionType EventActionType `bson:"eventActionType" json:"eventActionType"`
	PhoneNumber     string          `bson:"phoneNumber" json:"phoneNumber"`
}

// Implementing interface EventAction
func (x EventActionForwardToPhoneNumber) GetEventActionType() EventActionType {
	return x.EventActionType
}

// BUILDER from bson map:
func BuildEventActionForwardToPhoneNumber(m map[string]interface{}, x *EventActionForwardToPhoneNumber) {
	x.EventActionType = EventActionType_ForwardToPhoneNumber // readonly property
	if val, ok := m["phoneNumber"]; ok && val != nil {
		x.PhoneNumber = val.(string)
	}
}
