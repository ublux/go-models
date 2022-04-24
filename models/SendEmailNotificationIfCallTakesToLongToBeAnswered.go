package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SendEmailNotificationIfCallTakesToLongToBeAnswered struct {
	Email         []string `bson:"email" json:"email"`
	TimeInSeconds int32    `bson:"timeInSeconds" json:"timeInSeconds"`
}

// BUILDER from bson map:
func BuildSendEmailNotificationIfCallTakesToLongToBeAnswered(m map[string]interface{}, x *SendEmailNotificationIfCallTakesToLongToBeAnswered) {
	if val, ok := m["email"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				if val != nil {
					x.Email = append(x.Email, val.(string))
				}
			}
		}
	}
	if val, ok := m["timeInSeconds"]; ok && val != nil {
		x.TimeInSeconds = val.(int32)
	}
}
