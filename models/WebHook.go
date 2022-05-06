package models

import "time"
import . "github.com/ublux/go-models/enums"
import "go.mongodb.org/mongo-driver/bson/primitive"

type WebHook struct {
	DateCreated  time.Time    `bson:"dateCreated" json:"dateCreated"`
	DateDeleted  time.Time    `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated  time.Time    `bson:"dateUpdated" json:"dateUpdated"`
	Headers      []string     `bson:"headers" json:"headers"`
	Id           string       `bson:"id" json:"id"`
	IdAccount    string       `bson:"idAccount" json:"idAccount"`
	Url          string       `bson:"url" json:"url"`
	WebHookEvent WebHookEvent `bson:"webHookEvent" json:"webHookEvent"`
}

// Implementing interface IUbluxDocument
func (x WebHook) GetDateDeleted() time.Time {
	return x.DateDeleted
}
func (x WebHook) GetDateCreated() time.Time {
	return x.DateCreated
}
func (x WebHook) GetDateUpdated() time.Time {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x WebHook) GetId() string {
	return x.Id
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildWebHook(m map[string]interface{}, x *WebHook) {
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(time.Time)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(time.Time)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(time.Time)
	}
	if val, ok := m["headers"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				if val != nil {
					x.Headers = append(x.Headers, val.(string))
				}
			}
		}
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["idAccount"]; ok && val != nil {
		x.IdAccount = val.(string)
	}
	if val, ok := m["url"]; ok && val != nil {
		x.Url = val.(string)
	}
	if val, ok := m["webHookEvent"]; ok && val != nil {
		x.WebHookEvent = WebHookEvent("WebHookEvent_" + val.(string))
	} // is NOT readonly obtained from map
}
