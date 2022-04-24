package models

import . "github.com/ublux/go-models/enums"

type EventActionLeaveVoicemail struct {
	Email           string          `bson:"email" json:"email"`
	EventActionType EventActionType `bson:"eventActionType" json:"eventActionType"`
	IdAudio         string          `bson:"idAudio" json:"idAudio"`
}

// Implementing interface EventAction
func (x EventActionLeaveVoicemail) GetEventActionType() EventActionType {
	return x.EventActionType
}

// BUILDER from bson map:
func BuildEventActionLeaveVoicemail(m map[string]interface{}, x *EventActionLeaveVoicemail) {
	if val, ok := m["email"]; ok && val != nil {
		x.Email = val.(string)
	}
	x.EventActionType = EventActionType_LeaveVoicemail // readonly property
	if val, ok := m["idAudio"]; ok && val != nil {
		x.IdAudio = val.(string)
	}
}
