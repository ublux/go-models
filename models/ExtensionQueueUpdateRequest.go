package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ExtensionQueueUpdateRequest struct {
	Id                                                 string                                             `bson:"id" json:"id"`
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

// Implementing interface IUbluxDocumentId
func (x ExtensionQueueUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildExtensionQueueUpdateRequest(m map[string]interface{}, x *ExtensionQueueUpdateRequest) {
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
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
