package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ExtensionVoicemailUpdateRequest struct {
	Email                            []string `bson:"email" json:"email"`
	Id                               string   `bson:"id" json:"id"`
	IdAudio                          string   `bson:"idAudio" json:"idAudio"`
	IdMusicOnHoldGroup               string   `bson:"idMusicOnHoldGroup" json:"idMusicOnHoldGroup"`
	IdsLinesThatCanListenToVoicemail []string `bson:"idsLinesThatCanListenToVoicemail" json:"idsLinesThatCanListenToVoicemail"`
	InjectExtensionNameToCallerId    bool     `bson:"injectExtensionNameToCallerId" json:"injectExtensionNameToCallerId"`
	Number                           string   `bson:"number" json:"number"`
	TextToSpeech                     string   `bson:"textToSpeech" json:"textToSpeech"`
	TextToSpeechVoiceId              string   `bson:"textToSpeechVoiceId" json:"textToSpeechVoiceId"`
}

// Implementing interface IUbluxDocumentId
func (x ExtensionVoicemailUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildExtensionVoicemailUpdateRequest(m map[string]interface{}, x *ExtensionVoicemailUpdateRequest) {
	if val, ok := m["email"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				if val != nil {
					x.Email = append(x.Email, val.(string))
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
	if val, ok := m["idAudio"]; ok && val != nil {
		x.IdAudio = val.(string)
	}
	if val, ok := m["idMusicOnHoldGroup"]; ok && val != nil {
		x.IdMusicOnHoldGroup = val.(string)
	}
	if val, ok := m["idsLinesThatCanListenToVoicemail"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				if val != nil {
					x.IdsLinesThatCanListenToVoicemail = append(x.IdsLinesThatCanListenToVoicemail, val.(string))
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
	if val, ok := m["textToSpeech"]; ok && val != nil {
		x.TextToSpeech = val.(string)
	}
	if val, ok := m["textToSpeechVoiceId"]; ok && val != nil {
		x.TextToSpeechVoiceId = val.(string)
	}
}
