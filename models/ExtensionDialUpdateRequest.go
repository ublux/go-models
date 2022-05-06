package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ExtensionDialUpdateRequest struct {
	EventActionToExecuteIfCallIsNotAnswered  EventAction                              `bson:"eventActionToExecuteIfCallIsNotAnswered" json:"eventActionToExecuteIfCallIsNotAnswered"`
	Id                                       string                                   `bson:"id" json:"id"`
	IdMusicOnHoldGroup                       string                                   `bson:"idMusicOnHoldGroup" json:"idMusicOnHoldGroup"`
	IdsLines                                 []string                                 `bson:"idsLines" json:"idsLines"`
	InjectExtensionNameToCallerId            bool                                     `bson:"injectExtensionNameToCallerId" json:"injectExtensionNameToCallerId"`
	Number                                   string                                   `bson:"number" json:"number"`
	RingTimeInSeconds                        int32                                    `bson:"ringTimeInSeconds" json:"ringTimeInSeconds"`
	SendEmailNotificationIfCallIsNotAnswered SendEmailNotificationIfCallIsNotAnswered `bson:"sendEmailNotificationIfCallIsNotAnswered" json:"sendEmailNotificationIfCallIsNotAnswered"`
}

// Implementing interface IUbluxDocumentId
func (x ExtensionDialUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildExtensionDialUpdateRequest(m map[string]interface{}, x *ExtensionDialUpdateRequest) {
	if val, ok := m["eventActionToExecuteIfCallIsNotAnswered"]; ok && val != nil {
		// determine type to build
		t := val.(map[string]interface{})["_t"].(string)
		switch t {
		case "EventActionForwardToExtension":
			tmp := new(EventActionForwardToExtension)
			BuildEventActionForwardToExtension(val.(map[string]interface{}), tmp)
			x.EventActionToExecuteIfCallIsNotAnswered = tmp
		case "EventActionForwardToPhoneNumber":
			tmp := new(EventActionForwardToPhoneNumber)
			BuildEventActionForwardToPhoneNumber(val.(map[string]interface{}), tmp)
			x.EventActionToExecuteIfCallIsNotAnswered = tmp
		case "EventActionLeaveVoicemail":
			tmp := new(EventActionLeaveVoicemail)
			BuildEventActionLeaveVoicemail(val.(map[string]interface{}), tmp)
			x.EventActionToExecuteIfCallIsNotAnswered = tmp
		}
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["idMusicOnHoldGroup"]; ok && val != nil {
		x.IdMusicOnHoldGroup = val.(string)
	}
	if val, ok := m["idsLines"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				if val != nil {
					x.IdsLines = append(x.IdsLines, val.(string))
				}
			}
		}
	}
	if val, ok := m["injectExtensionNameToCallerId"]; ok && val != nil {
		x.InjectExtensionNameToCallerId = val.(bool)
	}
	if val, ok := m["number"]; ok && val != nil {
		x.Number = val.(string)
	}
	if val, ok := m["ringTimeInSeconds"]; ok && val != nil {
		x.RingTimeInSeconds = val.(int32)
	}
	if val, ok := m["sendEmailNotificationIfCallIsNotAnswered"]; ok && val != nil {
		BuildSendEmailNotificationIfCallIsNotAnswered(val.(map[string]interface{}), &x.SendEmailNotificationIfCallIsNotAnswered)
	}
}
