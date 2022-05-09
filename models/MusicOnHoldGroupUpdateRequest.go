package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type MusicOnHoldGroupUpdateRequest struct {
	Description  string   `bson:"description" json:"description"`
	FriendlyName string   `bson:"friendlyName" json:"friendlyName"`
	Id           string   `bson:"id" json:"id"`
	IdsAudios    []string `bson:"idsAudios" json:"idsAudios"`
}

// Implementing interface IUbluxDocumentId
func (x MusicOnHoldGroupUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildMusicOnHoldGroupUpdateRequest(m map[string]interface{}, x *MusicOnHoldGroupUpdateRequest) {
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
	if val, ok := m["idsAudios"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				if val != nil {
					x.IdsAudios = append(x.IdsAudios, val.(string))
				}
			}
		}
	}
}
