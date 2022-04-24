package models

import (
	. "github.com/ublux/go-models/enums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WebHook struct {
	Id           string       `bson:"_id" json:"id"`
	DateCreated  string       `bson:"dateCreated" json:"dateCreated"`
	DateDeleted  string       `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated  string       `bson:"dateUpdated" json:"dateUpdated"`
	Headers      []string     `bson:"headers" json:"headers"`
	IdAccount    string       `bson:"idAccount" json:"idAccount"`
	Url          string       `bson:"url" json:"url"`
	WebHookEvent WebHookEvent `bson:"webHookEvent" json:"webHookEvent"`
}

// Implementing interface IUbluxDocument
func (x WebHook) GetDateCreated() string {
	return x.DateCreated
}
func (x WebHook) GetDateDeleted() string {
	return x.DateDeleted
}
func (x WebHook) GetDateUpdated() string {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x WebHook) GetId() string {
	return x.Id
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildWebHook(m map[string]interface{}, x *WebHook) {
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(string)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(string)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(string)
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
