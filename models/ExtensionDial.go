package models

import "go.mongodb.org/mongo-driver/bson/primitive"
import . "github.com/ublux/go-models/enums"

type ExtensionDial struct {
	DateCreated                              primitive.DateTime                       `bson:"dateCreated" json:"dateCreated"`
	DateDeleted                              primitive.DateTime                       `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated                              primitive.DateTime                       `bson:"dateUpdated" json:"dateUpdated"`
	EventActionToExecuteIfCallIsNotAnswered  EventAction                              `bson:"eventActionToExecuteIfCallIsNotAnswered" json:"eventActionToExecuteIfCallIsNotAnswered"`
	ExtensionType                            ExtensionType                            `bson:"extensionType" json:"extensionType"`
	Id                                       string                                   `bson:"id" json:"id"`
	IdAccount                                string                                   `bson:"idAccount" json:"idAccount"`
	IdMusicOnHoldGroup                       string                                   `bson:"idMusicOnHoldGroup" json:"idMusicOnHoldGroup"`
	IdsLines                                 []string                                 `bson:"idsLines" json:"idsLines"`
	InjectExtensionNameToCallerId            bool                                     `bson:"injectExtensionNameToCallerId" json:"injectExtensionNameToCallerId"`
	Number                                   string                                   `bson:"number" json:"number"`
	RingTimeInSeconds                        int32                                    `bson:"ringTimeInSeconds" json:"ringTimeInSeconds"`
	SendEmailNotificationIfCallIsNotAnswered SendEmailNotificationIfCallIsNotAnswered `bson:"sendEmailNotificationIfCallIsNotAnswered" json:"sendEmailNotificationIfCallIsNotAnswered"`
}

// Implementing interface IUbluxDocument
func (x ExtensionDial) GetDateDeleted() primitive.DateTime {
	return x.DateDeleted
}
func (x ExtensionDial) GetDateCreated() primitive.DateTime {
	return x.DateCreated
}
func (x ExtensionDial) GetDateUpdated() primitive.DateTime {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x ExtensionDial) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x ExtensionDial) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface Extension
func (x ExtensionDial) GetIdMusicOnHoldGroup() string {
	return x.IdMusicOnHoldGroup
}
func (x ExtensionDial) GetExtensionType() ExtensionType {
	return x.ExtensionType
}
func (x ExtensionDial) GetNumber() string {
	return x.Number
}
func (x ExtensionDial) GetInjectExtensionNameToCallerId() bool {
	return x.InjectExtensionNameToCallerId
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildExtensionDial(m map[string]interface{}, x *ExtensionDial) {
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(primitive.DateTime)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(primitive.DateTime)
	}
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
	x.ExtensionType = ExtensionType_Dial // readonly property
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["idAccount"]; ok && val != nil {
		x.IdAccount = val.(string)
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
