package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type LineKeyGroupUpdateRequest struct {
	Description  string    `bson:"description" json:"description"`
	FriendlyName string    `bson:"friendlyName" json:"friendlyName"`
	Id           string    `bson:"id" json:"id"`
	LineKeys     []LineKey `bson:"lineKeys" json:"lineKeys"`
}

// Implementing interface IUbluxDocumentId
func (x LineKeyGroupUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildLineKeyGroupUpdateRequest(m map[string]interface{}, x *LineKeyGroupUpdateRequest) {
	if val, ok := m["description"]; ok && val != nil {
		x.Description = val.(string)
	}
	if val, ok := m["friendlyName"]; ok && val != nil {
		x.FriendlyName = val.(string)
	}
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["lineKeys"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok {
			for _, val = range array {
				if val != nil {
					tmp := new(LineKey)
					BuildLineKey(val.(map[string]interface{}), tmp)
					x.LineKeys = append(x.LineKeys, *tmp)
				}
			}
		}
	}
}
