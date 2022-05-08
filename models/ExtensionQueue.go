package models

import "go.mongodb.org/mongo-driver/bson/primitive"
import . "github.com/ublux/go-models/enums"

type ExtensionQueue struct {
	DateCreated                                        primitive.DateTime                                 `bson:"dateCreated" json:"dateCreated"`
	DateDeleted                                        primitive.DateTime                                 `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated                                        primitive.DateTime                                 `bson:"dateUpdated" json:"dateUpdated"`
	ExtensionType                                      ExtensionType                                      `bson:"extensionType" json:"extensionType"`
	Id                                                 string                                             `bson:"_id" json:"id"`
	IdAccount                                          string                                             `bson:"idAccount" json:"idAccount"`
	IdExtensionIfTimeout                               string                                             `bson:"idExtensionIfTimeout" json:"idExtensionIfTimeout"`
	IdMusicOnHoldGroup                                 string                                             `bson:"idMusicOnHoldGroup" json:"idMusicOnHoldGroup"`
	IdsAudios                                          []string                                           `bson:"idsAudios" json:"idsAudios"`
	IdsLines                                           []string                                           `bson:"idsLines" json:"idsLines"`
	InjectExtensionNameToCallerId                      bool                                               `bson:"injectExtensionNameToCallerId" json:"injectExtensionNameToCallerId"`
	Number                                             string                                             `bson:"number" json:"number"`
	PlayPositionAnnouncements                          bool                                               `bson:"playPositionAnnouncements" json:"playPositionAnnouncements"`
	QueueTimeoutInMinutes                              int32                                              `bson:"queueTimeoutInMinutes" json:"queueTimeoutInMinutes"`
	RetryFrequency                                     int32                                              `bson:"retryFrequency" json:"retryFrequency"`
	RingInUse                                          bool                                               `bson:"ringInUse" json:"ringInUse"`
	RingTimeInSeconds                                  int32                                              `bson:"ringTimeInSeconds" json:"ringTimeInSeconds"`
	SendEmailNotificationIfCallIsNotAnswered           SendEmailNotificationIfCallIsNotAnswered           `bson:"sendEmailNotificationIfCallIsNotAnswered" json:"sendEmailNotificationIfCallIsNotAnswered"`
	SendEmailNotificationIfCallTakesToLongToBeAnswered SendEmailNotificationIfCallTakesToLongToBeAnswered `bson:"sendEmailNotificationIfCallTakesToLongToBeAnswered" json:"sendEmailNotificationIfCallTakesToLongToBeAnswered"`
}

// Implementing interface IUbluxDocument
func (x ExtensionQueue) GetDateDeleted() primitive.DateTime {
	return x.DateDeleted
}
func (x ExtensionQueue) GetDateCreated() primitive.DateTime {
	return x.DateCreated
}
func (x ExtensionQueue) GetDateUpdated() primitive.DateTime {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x ExtensionQueue) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x ExtensionQueue) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface Extension
func (x ExtensionQueue) GetIdMusicOnHoldGroup() string {
	return x.IdMusicOnHoldGroup
}
func (x ExtensionQueue) GetExtensionType() ExtensionType {
	return x.ExtensionType
}
func (x ExtensionQueue) GetNumber() string {
	return x.Number
}
func (x ExtensionQueue) GetInjectExtensionNameToCallerId() bool {
	return x.InjectExtensionNameToCallerId
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildExtensionQueue(m map[string]interface{}, x *ExtensionQueue) {
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(primitive.DateTime)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(primitive.DateTime)
	}
	x.ExtensionType = ExtensionType_Queue // readonly property
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["idAccount"]; ok && val != nil {
		x.IdAccount = val.(string)
	}
	if val, ok := m["idExtensionIfTimeout"]; ok && val != nil {
		x.IdExtensionIfTimeout = val.(string)
	}
	if val, ok := m["idMusicOnHoldGroup"]; ok && val != nil {
		x.IdMusicOnHoldGroup = val.(string)
	}
	if val, ok := m["idsAudios"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				if val != nil {
					x.IdsAudios = append(x.IdsAudios, val.(string))
				}
			}
		}
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
	if val, ok := m["playPositionAnnouncements"]; ok && val != nil {
		x.PlayPositionAnnouncements = val.(bool)
	}
	if val, ok := m["queueTimeoutInMinutes"]; ok && val != nil {
		x.QueueTimeoutInMinutes = val.(int32)
	}
	if val, ok := m["retryFrequency"]; ok && val != nil {
		x.RetryFrequency = val.(int32)
	}
	if val, ok := m["ringInUse"]; ok && val != nil {
		x.RingInUse = val.(bool)
	}
	if val, ok := m["ringTimeInSeconds"]; ok && val != nil {
		x.RingTimeInSeconds = val.(int32)
	}
	if val, ok := m["sendEmailNotificationIfCallIsNotAnswered"]; ok && val != nil {
		BuildSendEmailNotificationIfCallIsNotAnswered(val.(map[string]interface{}), &x.SendEmailNotificationIfCallIsNotAnswered)
	}
	if val, ok := m["sendEmailNotificationIfCallTakesToLongToBeAnswered"]; ok && val != nil {
		BuildSendEmailNotificationIfCallTakesToLongToBeAnswered(val.(map[string]interface{}), &x.SendEmailNotificationIfCallTakesToLongToBeAnswered)
	}
}
