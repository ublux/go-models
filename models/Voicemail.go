package models

import "time"
import . "github.com/ublux/go-models/enums"
import "go.mongodb.org/mongo-driver/bson/primitive"

type Voicemail struct {
	DateCreated                      time.Time     `bson:"dateCreated" json:"dateCreated"`
	DateDeleted                      time.Time     `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated                      time.Time     `bson:"dateUpdated" json:"dateUpdated"`
	DurationInSeconds                int32         `bson:"durationInSeconds" json:"durationInSeconds"`
	Email                            string        `bson:"email" json:"email"`
	ErrorMessage                     string        `bson:"errorMessage" json:"errorMessage"`
	Id                               string        `bson:"id" json:"id"`
	IdsLinesThatCanListenToVoicemail []string      `bson:"idsLinesThatCanListenToVoicemail" json:"idsLinesThatCanListenToVoicemail"`
	VoicemailMp3                     StoredFile    `bson:"voicemailMp3" json:"voicemailMp3"`
	VoicemailType                    VoicemailType `bson:"voicemailType" json:"voicemailType"`
	VoicemailWav                     StoredFile    `bson:"voicemailWav" json:"voicemailWav"`
}

// Implementing interface IUbluxDocument
func (x Voicemail) GetDateDeleted() time.Time {
	return x.DateDeleted
}
func (x Voicemail) GetDateCreated() time.Time {
	return x.DateCreated
}
func (x Voicemail) GetDateUpdated() time.Time {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x Voicemail) GetId() string {
	return x.Id
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildVoicemail(m map[string]interface{}, x *Voicemail) {
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(time.Time)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(time.Time)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(time.Time)
	}
	if val, ok := m["durationInSeconds"]; ok && val != nil {
		x.DurationInSeconds = val.(int32)
	}
	if val, ok := m["email"]; ok && val != nil {
		x.Email = val.(string)
	}
	if val, ok := m["errorMessage"]; ok && val != nil {
		x.ErrorMessage = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
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
	if val, ok := m["voicemailMp3"]; ok && val != nil {
		BuildStoredFile(val.(map[string]interface{}), &x.VoicemailMp3)
	}
	x.VoicemailType = VoicemailType_Regular // readonly property
	if val, ok := m["voicemailWav"]; ok && val != nil {
		BuildStoredFile(val.(map[string]interface{}), &x.VoicemailWav)
	}
}
