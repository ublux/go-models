package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SendEmailNotificationIfCallIsNotAnswered struct {
	Email                                                 []string `bson:"email" json:"email"`
	PreventSendingNotificationIfCallLastsLessThanNSeconds int32    `bson:"preventSendingNotificationIfCallLastsLessThanNSeconds" json:"preventSendingNotificationIfCallLastsLessThanNSeconds"`
}

// BUILDER from bson map:
func BuildSendEmailNotificationIfCallIsNotAnswered(m map[string]interface{}, x *SendEmailNotificationIfCallIsNotAnswered) {
	if val, ok := m["email"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				if val != nil {
					x.Email = append(x.Email, val.(string))
				}
			}
		}
	}
	if val, ok := m["preventSendingNotificationIfCallLastsLessThanNSeconds"]; ok && val != nil {
		x.PreventSendingNotificationIfCallLastsLessThanNSeconds = val.(int32)
	}
}
