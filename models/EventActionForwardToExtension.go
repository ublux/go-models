package models

import . "github.com/ublux/go-models/enums"

type EventActionForwardToExtension struct {
	EventActionType EventActionType `bson:"eventActionType" json:"eventActionType"`
	IdExtension     string          `bson:"idExtension" json:"idExtension"`
}

// Implementing interface EventAction
func (x EventActionForwardToExtension) GetEventActionType() EventActionType {
	return x.EventActionType
}

// BUILDER from bson map:
func BuildEventActionForwardToExtension(m map[string]interface{}, x *EventActionForwardToExtension) {
	x.EventActionType = EventActionType_ForwardToExtension // readonly property
	if val, ok := m["idExtension"]; ok && val != nil {
		x.IdExtension = val.(string)
	}
}
