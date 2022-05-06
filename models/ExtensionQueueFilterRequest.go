package models

import "go.mongodb.org/mongo-driver/bson/primitive"
import . "github.com/ublux/go-models/enums"

type ExtensionQueueFilterRequest struct {
	DateCreated_EQ                                                                                     primitive.DateTime `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE                                                                                    primitive.DateTime `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE                                                                                    primitive.DateTime `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ                                                                                     primitive.DateTime `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE                                                                                    primitive.DateTime `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE                                                                                    primitive.DateTime `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	ExtensionType_EQ                                                                                   ExtensionType      `bson:"extensionType_EQ" json:"extensionType_EQ"`
	Id_CON                                                                                             string             `bson:"id_CON" json:"id_CON"`
	Id_EQ                                                                                              string             `bson:"id_EQ" json:"id_EQ"`
	Id_REG                                                                                             string             `bson:"id_REG" json:"id_REG"`
	IdExtensionIfTimeout_CON                                                                           string             `bson:"idExtensionIfTimeout_CON" json:"idExtensionIfTimeout_CON"`
	IdExtensionIfTimeout_EQ                                                                            string             `bson:"idExtensionIfTimeout_EQ" json:"idExtensionIfTimeout_EQ"`
	IdExtensionIfTimeout_REG                                                                           string             `bson:"idExtensionIfTimeout_REG" json:"idExtensionIfTimeout_REG"`
	IdMusicOnHoldGroup_CON                                                                             string             `bson:"idMusicOnHoldGroup_CON" json:"idMusicOnHoldGroup_CON"`
	IdMusicOnHoldGroup_EQ                                                                              string             `bson:"idMusicOnHoldGroup_EQ" json:"idMusicOnHoldGroup_EQ"`
	IdMusicOnHoldGroup_REG                                                                             string             `bson:"idMusicOnHoldGroup_REG" json:"idMusicOnHoldGroup_REG"`
	IdsAudios_CON                                                                                      string             `bson:"idsAudios_CON" json:"idsAudios_CON"`
	IdsAudios_EQ                                                                                       string             `bson:"idsAudios_EQ" json:"idsAudios_EQ"`
	IdsAudios_REG                                                                                      string             `bson:"idsAudios_REG" json:"idsAudios_REG"`
	IdsLines_CON                                                                                       string             `bson:"idsLines_CON" json:"idsLines_CON"`
	IdsLines_EQ                                                                                        string             `bson:"idsLines_EQ" json:"idsLines_EQ"`
	IdsLines_REG                                                                                       string             `bson:"idsLines_REG" json:"idsLines_REG"`
	InjectExtensionNameToCallerId_EQ                                                                   bool               `bson:"injectExtensionNameToCallerId_EQ" json:"injectExtensionNameToCallerId_EQ"`
	Number_CON                                                                                         string             `bson:"number_CON" json:"number_CON"`
	Number_EQ                                                                                          string             `bson:"number_EQ" json:"number_EQ"`
	Number_REG                                                                                         string             `bson:"number_REG" json:"number_REG"`
	PlayPositionAnnouncements_EQ                                                                       bool               `bson:"playPositionAnnouncements_EQ" json:"playPositionAnnouncements_EQ"`
	QueueTimeoutInMinutes_EQ                                                                           int32              `bson:"queueTimeoutInMinutes_EQ" json:"queueTimeoutInMinutes_EQ"`
	QueueTimeoutInMinutes_GTE                                                                          int32              `bson:"queueTimeoutInMinutes_GTE" json:"queueTimeoutInMinutes_GTE"`
	QueueTimeoutInMinutes_LTE                                                                          int32              `bson:"queueTimeoutInMinutes_LTE" json:"queueTimeoutInMinutes_LTE"`
	RetryFrequency_EQ                                                                                  int32              `bson:"retryFrequency_EQ" json:"retryFrequency_EQ"`
	RetryFrequency_GTE                                                                                 int32              `bson:"retryFrequency_GTE" json:"retryFrequency_GTE"`
	RetryFrequency_LTE                                                                                 int32              `bson:"retryFrequency_LTE" json:"retryFrequency_LTE"`
	RingInUse_EQ                                                                                       bool               `bson:"ringInUse_EQ" json:"ringInUse_EQ"`
	RingTimeInSeconds_EQ                                                                               int32              `bson:"ringTimeInSeconds_EQ" json:"ringTimeInSeconds_EQ"`
	RingTimeInSeconds_GTE                                                                              int32              `bson:"ringTimeInSeconds_GTE" json:"ringTimeInSeconds_GTE"`
	RingTimeInSeconds_LTE                                                                              int32              `bson:"ringTimeInSeconds_LTE" json:"ringTimeInSeconds_LTE"`
	SendEmailNotificationIfCallIsNotAnswered_Email_CON                                                 string             `bson:"sendEmailNotificationIfCallIsNotAnswered_Email_CON" json:"sendEmailNotificationIfCallIsNotAnswered_Email_CON"`
	SendEmailNotificationIfCallIsNotAnswered_Email_EQ                                                  string             `bson:"sendEmailNotificationIfCallIsNotAnswered_Email_EQ" json:"sendEmailNotificationIfCallIsNotAnswered_Email_EQ"`
	SendEmailNotificationIfCallIsNotAnswered_Email_REG                                                 string             `bson:"sendEmailNotificationIfCallIsNotAnswered_Email_REG" json:"sendEmailNotificationIfCallIsNotAnswered_Email_REG"`
	SendEmailNotificationIfCallIsNotAnswered_PreventSendingNotificationIfCallLastsLessThanNSeconds_EQ  int32              `bson:"sendEmailNotificationIfCallIsNotAnswered_PreventSendingNotificationIfCallLastsLessThanNSeconds_EQ" json:"sendEmailNotificationIfCallIsNotAnswered_PreventSendingNotificationIfCallLastsLessThanNSeconds_EQ"`
	SendEmailNotificationIfCallIsNotAnswered_PreventSendingNotificationIfCallLastsLessThanNSeconds_GTE int32              `bson:"sendEmailNotificationIfCallIsNotAnswered_PreventSendingNotificationIfCallLastsLessThanNSeconds_GTE" json:"sendEmailNotificationIfCallIsNotAnswered_PreventSendingNotificationIfCallLastsLessThanNSeconds_GTE"`
	SendEmailNotificationIfCallIsNotAnswered_PreventSendingNotificationIfCallLastsLessThanNSeconds_LTE int32              `bson:"sendEmailNotificationIfCallIsNotAnswered_PreventSendingNotificationIfCallLastsLessThanNSeconds_LTE" json:"sendEmailNotificationIfCallIsNotAnswered_PreventSendingNotificationIfCallLastsLessThanNSeconds_LTE"`
	SendEmailNotificationIfCallTakesToLongToBeAnswered_Email_CON                                       string             `bson:"sendEmailNotificationIfCallTakesToLongToBeAnswered_Email_CON" json:"sendEmailNotificationIfCallTakesToLongToBeAnswered_Email_CON"`
	SendEmailNotificationIfCallTakesToLongToBeAnswered_Email_EQ                                        string             `bson:"sendEmailNotificationIfCallTakesToLongToBeAnswered_Email_EQ" json:"sendEmailNotificationIfCallTakesToLongToBeAnswered_Email_EQ"`
	SendEmailNotificationIfCallTakesToLongToBeAnswered_Email_REG                                       string             `bson:"sendEmailNotificationIfCallTakesToLongToBeAnswered_Email_REG" json:"sendEmailNotificationIfCallTakesToLongToBeAnswered_Email_REG"`
	SendEmailNotificationIfCallTakesToLongToBeAnswered_TimeInSeconds_EQ                                int32              `bson:"sendEmailNotificationIfCallTakesToLongToBeAnswered_TimeInSeconds_EQ" json:"sendEmailNotificationIfCallTakesToLongToBeAnswered_TimeInSeconds_EQ"`
	SendEmailNotificationIfCallTakesToLongToBeAnswered_TimeInSeconds_GTE                               int32              `bson:"sendEmailNotificationIfCallTakesToLongToBeAnswered_TimeInSeconds_GTE" json:"sendEmailNotificationIfCallTakesToLongToBeAnswered_TimeInSeconds_GTE"`
	SendEmailNotificationIfCallTakesToLongToBeAnswered_TimeInSeconds_LTE                               int32              `bson:"sendEmailNotificationIfCallTakesToLongToBeAnswered_TimeInSeconds_LTE" json:"sendEmailNotificationIfCallTakesToLongToBeAnswered_TimeInSeconds_LTE"`
}

// BUILDER from bson map:
func BuildExtensionQueueFilterRequest(m map[string]interface{}, x *ExtensionQueueFilterRequest) {
	if val, ok := m["dateCreated_EQ"]; ok && val != nil {
		x.DateCreated_EQ = val.(primitive.DateTime)
	}
	if val, ok := m["dateCreated_GTE"]; ok && val != nil {
		x.DateCreated_GTE = val.(primitive.DateTime)
	}
	if val, ok := m["dateCreated_LTE"]; ok && val != nil {
		x.DateCreated_LTE = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated_EQ"]; ok && val != nil {
		x.DateUpdated_EQ = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated_GTE"]; ok && val != nil {
		x.DateUpdated_GTE = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated_LTE"]; ok && val != nil {
		x.DateUpdated_LTE = val.(primitive.DateTime)
	}
	if val, ok := m["extensionType_EQ"]; ok && val != nil {
		x.ExtensionType_EQ = ExtensionType("ExtensionType_EQ_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["id_CON"]; ok && val != nil {
		x.Id_CON = val.(string)
	}
	if val, ok := m["id_EQ"]; ok && val != nil {
		x.Id_EQ = val.(string)
	}
	if val, ok := m["id_REG"]; ok && val != nil {
		x.Id_REG = val.(string)
	}
	if val, ok := m["idExtensionIfTimeout_CON"]; ok && val != nil {
		x.IdExtensionIfTimeout_CON = val.(string)
	}
	if val, ok := m["idExtensionIfTimeout_EQ"]; ok && val != nil {
		x.IdExtensionIfTimeout_EQ = val.(string)
	}
	if val, ok := m["idExtensionIfTimeout_REG"]; ok && val != nil {
		x.IdExtensionIfTimeout_REG = val.(string)
	}
	if val, ok := m["idMusicOnHoldGroup_CON"]; ok && val != nil {
		x.IdMusicOnHoldGroup_CON = val.(string)
	}
	if val, ok := m["idMusicOnHoldGroup_EQ"]; ok && val != nil {
		x.IdMusicOnHoldGroup_EQ = val.(string)
	}
	if val, ok := m["idMusicOnHoldGroup_REG"]; ok && val != nil {
		x.IdMusicOnHoldGroup_REG = val.(string)
	}
	if val, ok := m["idsAudios_CON"]; ok && val != nil {
		x.IdsAudios_CON = val.(string)
	}
	if val, ok := m["idsAudios_EQ"]; ok && val != nil {
		x.IdsAudios_EQ = val.(string)
	}
	if val, ok := m["idsAudios_REG"]; ok && val != nil {
		x.IdsAudios_REG = val.(string)
	}
	if val, ok := m["idsLines_CON"]; ok && val != nil {
		x.IdsLines_CON = val.(string)
	}
	if val, ok := m["idsLines_EQ"]; ok && val != nil {
		x.IdsLines_EQ = val.(string)
	}
	if val, ok := m["idsLines_REG"]; ok && val != nil {
		x.IdsLines_REG = val.(string)
	}
	if val, ok := m["injectExtensionNameToCallerId_EQ"]; ok && val != nil {
		x.InjectExtensionNameToCallerId_EQ = val.(bool)
	}
	if val, ok := m["number_CON"]; ok && val != nil {
		x.Number_CON = val.(string)
	}
	if val, ok := m["number_EQ"]; ok && val != nil {
		x.Number_EQ = val.(string)
	}
	if val, ok := m["number_REG"]; ok && val != nil {
		x.Number_REG = val.(string)
	}
	if val, ok := m["playPositionAnnouncements_EQ"]; ok && val != nil {
		x.PlayPositionAnnouncements_EQ = val.(bool)
	}
	if val, ok := m["queueTimeoutInMinutes_EQ"]; ok && val != nil {
		x.QueueTimeoutInMinutes_EQ = val.(int32)
	}
	if val, ok := m["queueTimeoutInMinutes_GTE"]; ok && val != nil {
		x.QueueTimeoutInMinutes_GTE = val.(int32)
	}
	if val, ok := m["queueTimeoutInMinutes_LTE"]; ok && val != nil {
		x.QueueTimeoutInMinutes_LTE = val.(int32)
	}
	if val, ok := m["retryFrequency_EQ"]; ok && val != nil {
		x.RetryFrequency_EQ = val.(int32)
	}
	if val, ok := m["retryFrequency_GTE"]; ok && val != nil {
		x.RetryFrequency_GTE = val.(int32)
	}
	if val, ok := m["retryFrequency_LTE"]; ok && val != nil {
		x.RetryFrequency_LTE = val.(int32)
	}
	if val, ok := m["ringInUse_EQ"]; ok && val != nil {
		x.RingInUse_EQ = val.(bool)
	}
	if val, ok := m["ringTimeInSeconds_EQ"]; ok && val != nil {
		x.RingTimeInSeconds_EQ = val.(int32)
	}
	if val, ok := m["ringTimeInSeconds_GTE"]; ok && val != nil {
		x.RingTimeInSeconds_GTE = val.(int32)
	}
	if val, ok := m["ringTimeInSeconds_LTE"]; ok && val != nil {
		x.RingTimeInSeconds_LTE = val.(int32)
	}
	if val, ok := m["sendEmailNotificationIfCallIsNotAnswered_Email_CON"]; ok && val != nil {
		x.SendEmailNotificationIfCallIsNotAnswered_Email_CON = val.(string)
	}
	if val, ok := m["sendEmailNotificationIfCallIsNotAnswered_Email_EQ"]; ok && val != nil {
		x.SendEmailNotificationIfCallIsNotAnswered_Email_EQ = val.(string)
	}
	if val, ok := m["sendEmailNotificationIfCallIsNotAnswered_Email_REG"]; ok && val != nil {
		x.SendEmailNotificationIfCallIsNotAnswered_Email_REG = val.(string)
	}
	if val, ok := m["sendEmailNotificationIfCallIsNotAnswered_PreventSendingNotificationIfCallLastsLessThanNSeconds_EQ"]; ok && val != nil {
		x.SendEmailNotificationIfCallIsNotAnswered_PreventSendingNotificationIfCallLastsLessThanNSeconds_EQ = val.(int32)
	}
	if val, ok := m["sendEmailNotificationIfCallIsNotAnswered_PreventSendingNotificationIfCallLastsLessThanNSeconds_GTE"]; ok && val != nil {
		x.SendEmailNotificationIfCallIsNotAnswered_PreventSendingNotificationIfCallLastsLessThanNSeconds_GTE = val.(int32)
	}
	if val, ok := m["sendEmailNotificationIfCallIsNotAnswered_PreventSendingNotificationIfCallLastsLessThanNSeconds_LTE"]; ok && val != nil {
		x.SendEmailNotificationIfCallIsNotAnswered_PreventSendingNotificationIfCallLastsLessThanNSeconds_LTE = val.(int32)
	}
	if val, ok := m["sendEmailNotificationIfCallTakesToLongToBeAnswered_Email_CON"]; ok && val != nil {
		x.SendEmailNotificationIfCallTakesToLongToBeAnswered_Email_CON = val.(string)
	}
	if val, ok := m["sendEmailNotificationIfCallTakesToLongToBeAnswered_Email_EQ"]; ok && val != nil {
		x.SendEmailNotificationIfCallTakesToLongToBeAnswered_Email_EQ = val.(string)
	}
	if val, ok := m["sendEmailNotificationIfCallTakesToLongToBeAnswered_Email_REG"]; ok && val != nil {
		x.SendEmailNotificationIfCallTakesToLongToBeAnswered_Email_REG = val.(string)
	}
	if val, ok := m["sendEmailNotificationIfCallTakesToLongToBeAnswered_TimeInSeconds_EQ"]; ok && val != nil {
		x.SendEmailNotificationIfCallTakesToLongToBeAnswered_TimeInSeconds_EQ = val.(int32)
	}
	if val, ok := m["sendEmailNotificationIfCallTakesToLongToBeAnswered_TimeInSeconds_GTE"]; ok && val != nil {
		x.SendEmailNotificationIfCallTakesToLongToBeAnswered_TimeInSeconds_GTE = val.(int32)
	}
	if val, ok := m["sendEmailNotificationIfCallTakesToLongToBeAnswered_TimeInSeconds_LTE"]; ok && val != nil {
		x.SendEmailNotificationIfCallTakesToLongToBeAnswered_TimeInSeconds_LTE = val.(int32)
	}
}
