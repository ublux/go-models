package models

import . "github.com/ublux/go-models/enums"
import "go.mongodb.org/mongo-driver/bson/primitive"

type ExtensionVoicemail struct {
	Id                               string        `bson:"_id" json:"id"`
	DateCreated                      string        `bson:"dateCreated" json:"dateCreated"`
	DateDeleted                      string        `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated                      string        `bson:"dateUpdated" json:"dateUpdated"`
	Email                            []string      `bson:"email" json:"email"`
	ExtensionType                    ExtensionType `bson:"extensionType" json:"extensionType"`
	IdAccount                        string        `bson:"idAccount" json:"idAccount"`
	IdAudio                          string        `bson:"idAudio" json:"idAudio"`
	IdMusicOnHoldGroup               string        `bson:"idMusicOnHoldGroup" json:"idMusicOnHoldGroup"`
	IdsLinesThatCanListenToVoicemail []string      `bson:"idsLinesThatCanListenToVoicemail" json:"idsLinesThatCanListenToVoicemail"`
	InjectExtensionNameToCallerId    bool          `bson:"injectExtensionNameToCallerId" json:"injectExtensionNameToCallerId"`
	Number                           string        `bson:"number" json:"number"`
	TextToSpeech                     string        `bson:"textToSpeech" json:"textToSpeech"`
	TextToSpeechVoiceId              string        `bson:"textToSpeechVoiceId" json:"textToSpeechVoiceId"`
}

// Implementing interface IUbluxDocument
func (x ExtensionVoicemail) GetDateCreated() string {
	return x.DateCreated
}
func (x ExtensionVoicemail) GetDateDeleted() string {
	return x.DateDeleted
}
func (x ExtensionVoicemail) GetDateUpdated() string {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x ExtensionVoicemail) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x ExtensionVoicemail) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface Extension
func (x ExtensionVoicemail) GetIdMusicOnHoldGroup() string {
	return x.IdMusicOnHoldGroup
}
func (x ExtensionVoicemail) GetExtensionType() ExtensionType {
	return x.ExtensionType
}
func (x ExtensionVoicemail) GetNumber() string {
	return x.Number
}
func (x ExtensionVoicemail) GetInjectExtensionNameToCallerId() bool {
	return x.InjectExtensionNameToCallerId
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildExtensionVoicemail(m map[string]interface{}, x *ExtensionVoicemail) {
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
	if val, ok := m["email"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				if val != nil {
					x.Email = append(x.Email, val.(string))
				}
			}
		}
	}
	x.ExtensionType = ExtensionType_Voicemail // readonly property
	if val, ok := m["idAccount"]; ok && val != nil {
		x.IdAccount = val.(string)
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
