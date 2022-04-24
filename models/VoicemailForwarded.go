package models

import (
	. "github.com/ublux/go-models/enums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VoicemailForwarded struct {
	Voicemail
	IdExtensionForwarde          string `bson:"idExtensionForwarde" json:"idExtensionForwarde"`
	IdLineThatForwardedVoicemail string `bson:"idLineThatForwardedVoicemail" json:"idLineThatForwardedVoicemail"`
}

// Implementing interface IUbluxDocument
func (x VoicemailForwarded) GetDateCreated() string {
	return x.DateCreated
}
func (x VoicemailForwarded) GetDateDeleted() string {
	return x.DateDeleted
}
func (x VoicemailForwarded) GetDateUpdated() string {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x VoicemailForwarded) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildVoicemailForwarded(m map[string]interface{}, x *VoicemailForwarded) {
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
	if val, ok := m["durationInSeconds"]; ok && val != nil {
		x.DurationInSeconds = val.(int32)
	}
	if val, ok := m["email"]; ok && val != nil {
		x.Email = val.(string)
	}
	if val, ok := m["errorMessage"]; ok && val != nil {
		x.ErrorMessage = val.(string)
	}
	if val, ok := m["idExtensionForwarde"]; ok && val != nil {
		x.IdExtensionForwarde = val.(string)
	}
	if val, ok := m["idLineThatForwardedVoicemail"]; ok && val != nil {
		x.IdLineThatForwardedVoicemail = val.(string)
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
	x.VoicemailType = VoicemailType_Forwarded // readonly property
	if val, ok := m["voicemailWav"]; ok && val != nil {
		BuildStoredFile(val.(map[string]interface{}), &x.VoicemailWav)
	}
}
