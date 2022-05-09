package models

import . "github.com/ublux/go-models/enums"
import "go.mongodb.org/mongo-driver/bson/primitive"

type WebHookUpdateRequest struct {
	Headers      []string     `bson:"headers" json:"headers"`
	Id           string       `bson:"id" json:"id"`
	Url          string       `bson:"url" json:"url"`
	WebHookEvent WebHookEvent `bson:"webHookEvent" json:"webHookEvent"`
}

// Implementing interface IUbluxDocumentId
func (x WebHookUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildWebHookUpdateRequest(m map[string]interface{}, x *WebHookUpdateRequest) {
	if val, ok := m["headers"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				if val != nil {
					x.Headers = append(x.Headers, val.(string))
				}
			}
		}
	}
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["url"]; ok && val != nil {
		x.Url = val.(string)
	}
	if val, ok := m["webHookEvent"]; ok && val != nil {
		x.WebHookEvent = WebHookEvent("WebHookEvent_" + val.(string))
	} // is NOT readonly obtained from map
}
