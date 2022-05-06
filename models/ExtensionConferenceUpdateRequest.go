package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ExtensionConferenceUpdateRequest struct {
	AnnounceParticipants          bool     `bson:"announceParticipants" json:"announceParticipants"`
	Id                            string   `bson:"id" json:"id"`
	IdMusicOnHoldGroup            string   `bson:"idMusicOnHoldGroup" json:"idMusicOnHoldGroup"`
	IdsAudiosWhenOneParticipant   []string `bson:"idsAudiosWhenOneParticipant" json:"idsAudiosWhenOneParticipant"`
	InjectExtensionNameToCallerId bool     `bson:"injectExtensionNameToCallerId" json:"injectExtensionNameToCallerId"`
	Number                        string   `bson:"number" json:"number"`
	Pin                           string   `bson:"pin" json:"pin"`
}

// Implementing interface IUbluxDocumentId
func (x ExtensionConferenceUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildExtensionConferenceUpdateRequest(m map[string]interface{}, x *ExtensionConferenceUpdateRequest) {
	if val, ok := m["announceParticipants"]; ok && val != nil {
		x.AnnounceParticipants = val.(bool)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["idMusicOnHoldGroup"]; ok && val != nil {
		x.IdMusicOnHoldGroup = val.(string)
	}
	if val, ok := m["idsAudiosWhenOneParticipant"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				if val != nil {
					x.IdsAudiosWhenOneParticipant = append(x.IdsAudiosWhenOneParticipant, val.(string))
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
	if val, ok := m["pin"]; ok && val != nil {
		x.Pin = val.(string)
	}
}
