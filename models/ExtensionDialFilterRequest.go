package models

import "time"
import . "github.com/ublux/go-models/enums"

type ExtensionDialFilterRequest struct {
	DateCreated_EQ                                                                                     time.Time       `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE                                                                                    time.Time       `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE                                                                                    time.Time       `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ                                                                                     time.Time       `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE                                                                                    time.Time       `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE                                                                                    time.Time       `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	EventActionToExecuteIfCallIsNotAnswered_EventActionType_EQ                                         EventActionType `bson:"eventActionToExecuteIfCallIsNotAnswered_EventActionType_EQ" json:"eventActionToExecuteIfCallIsNotAnswered_EventActionType_EQ"`
	ExtensionType_EQ                                                                                   ExtensionType   `bson:"extensionType_EQ" json:"extensionType_EQ"`
	Id_CON                                                                                             string          `bson:"id_CON" json:"id_CON"`
	Id_EQ                                                                                              string          `bson:"id_EQ" json:"id_EQ"`
	Id_REG                                                                                             string          `bson:"id_REG" json:"id_REG"`
	IdMusicOnHoldGroup_CON                                                                             string          `bson:"idMusicOnHoldGroup_CON" json:"idMusicOnHoldGroup_CON"`
	IdMusicOnHoldGroup_EQ                                                                              string          `bson:"idMusicOnHoldGroup_EQ" json:"idMusicOnHoldGroup_EQ"`
	IdMusicOnHoldGroup_REG                                                                             string          `bson:"idMusicOnHoldGroup_REG" json:"idMusicOnHoldGroup_REG"`
	IdsLines_CON                                                                                       string          `bson:"idsLines_CON" json:"idsLines_CON"`
	IdsLines_EQ                                                                                        string          `bson:"idsLines_EQ" json:"idsLines_EQ"`
	IdsLines_REG                                                                                       string          `bson:"idsLines_REG" json:"idsLines_REG"`
	InjectExtensionNameToCallerId_EQ                                                                   bool            `bson:"injectExtensionNameToCallerId_EQ" json:"injectExtensionNameToCallerId_EQ"`
	Number_CON                                                                                         string          `bson:"number_CON" json:"number_CON"`
	Number_EQ                                                                                          string          `bson:"number_EQ" json:"number_EQ"`
	Number_REG                                                                                         string          `bson:"number_REG" json:"number_REG"`
	RingTimeInSeconds_EQ                                                                               int32           `bson:"ringTimeInSeconds_EQ" json:"ringTimeInSeconds_EQ"`
	RingTimeInSeconds_GTE                                                                              int32           `bson:"ringTimeInSeconds_GTE" json:"ringTimeInSeconds_GTE"`
	RingTimeInSeconds_LTE                                                                              int32           `bson:"ringTimeInSeconds_LTE" json:"ringTimeInSeconds_LTE"`
	SendEmailNotificationIfCallIsNotAnswered_Email_CON                                                 string          `bson:"sendEmailNotificationIfCallIsNotAnswered_Email_CON" json:"sendEmailNotificationIfCallIsNotAnswered_Email_CON"`
	SendEmailNotificationIfCallIsNotAnswered_Email_EQ                                                  string          `bson:"sendEmailNotificationIfCallIsNotAnswered_Email_EQ" json:"sendEmailNotificationIfCallIsNotAnswered_Email_EQ"`
	SendEmailNotificationIfCallIsNotAnswered_Email_REG                                                 string          `bson:"sendEmailNotificationIfCallIsNotAnswered_Email_REG" json:"sendEmailNotificationIfCallIsNotAnswered_Email_REG"`
	SendEmailNotificationIfCallIsNotAnswered_PreventSendingNotificationIfCallLastsLessThanNSeconds_EQ  int32           `bson:"sendEmailNotificationIfCallIsNotAnswered_PreventSendingNotificationIfCallLastsLessThanNSeconds_EQ" json:"sendEmailNotificationIfCallIsNotAnswered_PreventSendingNotificationIfCallLastsLessThanNSeconds_EQ"`
	SendEmailNotificationIfCallIsNotAnswered_PreventSendingNotificationIfCallLastsLessThanNSeconds_GTE int32           `bson:"sendEmailNotificationIfCallIsNotAnswered_PreventSendingNotificationIfCallLastsLessThanNSeconds_GTE" json:"sendEmailNotificationIfCallIsNotAnswered_PreventSendingNotificationIfCallLastsLessThanNSeconds_GTE"`
	SendEmailNotificationIfCallIsNotAnswered_PreventSendingNotificationIfCallLastsLessThanNSeconds_LTE int32           `bson:"sendEmailNotificationIfCallIsNotAnswered_PreventSendingNotificationIfCallLastsLessThanNSeconds_LTE" json:"sendEmailNotificationIfCallIsNotAnswered_PreventSendingNotificationIfCallLastsLessThanNSeconds_LTE"`
}

// BUILDER from bson map:
func BuildExtensionDialFilterRequest(m map[string]interface{}, x *ExtensionDialFilterRequest) {
	if val, ok := m["dateCreated_EQ"]; ok && val != nil {
		x.DateCreated_EQ = val.(time.Time)
	}
	if val, ok := m["dateCreated_GTE"]; ok && val != nil {
		x.DateCreated_GTE = val.(time.Time)
	}
	if val, ok := m["dateCreated_LTE"]; ok && val != nil {
		x.DateCreated_LTE = val.(time.Time)
	}
	if val, ok := m["dateUpdated_EQ"]; ok && val != nil {
		x.DateUpdated_EQ = val.(time.Time)
	}
	if val, ok := m["dateUpdated_GTE"]; ok && val != nil {
		x.DateUpdated_GTE = val.(time.Time)
	}
	if val, ok := m["dateUpdated_LTE"]; ok && val != nil {
		x.DateUpdated_LTE = val.(time.Time)
	}
	if val, ok := m["eventActionToExecuteIfCallIsNotAnswered_EventActionType_EQ"]; ok && val != nil {
		x.EventActionToExecuteIfCallIsNotAnswered_EventActionType_EQ = EventActionType("EventActionToExecuteIfCallIsNotAnswered_EventActionType_EQ_" + val.(string))
	} // is NOT readonly obtained from map
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
	if val, ok := m["idMusicOnHoldGroup_CON"]; ok && val != nil {
		x.IdMusicOnHoldGroup_CON = val.(string)
	}
	if val, ok := m["idMusicOnHoldGroup_EQ"]; ok && val != nil {
		x.IdMusicOnHoldGroup_EQ = val.(string)
	}
	if val, ok := m["idMusicOnHoldGroup_REG"]; ok && val != nil {
		x.IdMusicOnHoldGroup_REG = val.(string)
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
}
