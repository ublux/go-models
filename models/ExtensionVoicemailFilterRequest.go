package models

import "go.mongodb.org/mongo-driver/bson/primitive"
import . "github.com/ublux/go-models/enums"

type ExtensionVoicemailFilterRequest struct {
	DateCreated_EQ                       primitive.DateTime `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE                      primitive.DateTime `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE                      primitive.DateTime `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ                       primitive.DateTime `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE                      primitive.DateTime `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE                      primitive.DateTime `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	Email_CON                            string             `bson:"email_CON" json:"email_CON"`
	Email_EQ                             string             `bson:"email_EQ" json:"email_EQ"`
	Email_REG                            string             `bson:"email_REG" json:"email_REG"`
	ExtensionType_EQ                     ExtensionType      `bson:"extensionType_EQ" json:"extensionType_EQ"`
	Id_CON                               string             `bson:"id_CON" json:"id_CON"`
	Id_EQ                                string             `bson:"id_EQ" json:"id_EQ"`
	Id_REG                               string             `bson:"id_REG" json:"id_REG"`
	IdAudio_CON                          string             `bson:"idAudio_CON" json:"idAudio_CON"`
	IdAudio_EQ                           string             `bson:"idAudio_EQ" json:"idAudio_EQ"`
	IdAudio_REG                          string             `bson:"idAudio_REG" json:"idAudio_REG"`
	IdMusicOnHoldGroup_CON               string             `bson:"idMusicOnHoldGroup_CON" json:"idMusicOnHoldGroup_CON"`
	IdMusicOnHoldGroup_EQ                string             `bson:"idMusicOnHoldGroup_EQ" json:"idMusicOnHoldGroup_EQ"`
	IdMusicOnHoldGroup_REG               string             `bson:"idMusicOnHoldGroup_REG" json:"idMusicOnHoldGroup_REG"`
	IdsLinesThatCanListenToVoicemail_CON string             `bson:"idsLinesThatCanListenToVoicemail_CON" json:"idsLinesThatCanListenToVoicemail_CON"`
	IdsLinesThatCanListenToVoicemail_EQ  string             `bson:"idsLinesThatCanListenToVoicemail_EQ" json:"idsLinesThatCanListenToVoicemail_EQ"`
	IdsLinesThatCanListenToVoicemail_REG string             `bson:"idsLinesThatCanListenToVoicemail_REG" json:"idsLinesThatCanListenToVoicemail_REG"`
	InjectExtensionNameToCallerId_EQ     bool               `bson:"injectExtensionNameToCallerId_EQ" json:"injectExtensionNameToCallerId_EQ"`
	Number_CON                           string             `bson:"number_CON" json:"number_CON"`
	Number_EQ                            string             `bson:"number_EQ" json:"number_EQ"`
	Number_REG                           string             `bson:"number_REG" json:"number_REG"`
	TextToSpeech_CON                     string             `bson:"textToSpeech_CON" json:"textToSpeech_CON"`
	TextToSpeech_EQ                      string             `bson:"textToSpeech_EQ" json:"textToSpeech_EQ"`
	TextToSpeech_REG                     string             `bson:"textToSpeech_REG" json:"textToSpeech_REG"`
	TextToSpeechVoiceId_CON              string             `bson:"textToSpeechVoiceId_CON" json:"textToSpeechVoiceId_CON"`
	TextToSpeechVoiceId_EQ               string             `bson:"textToSpeechVoiceId_EQ" json:"textToSpeechVoiceId_EQ"`
	TextToSpeechVoiceId_REG              string             `bson:"textToSpeechVoiceId_REG" json:"textToSpeechVoiceId_REG"`
}

// BUILDER from bson map:
func BuildExtensionVoicemailFilterRequest(m map[string]interface{}, x *ExtensionVoicemailFilterRequest) {
	if val, ok := m["dateCreated_EQ"]; ok && val != nil {
		x.DateCreated_EQ = val.(primitive.DateTime)
	}
	if val, ok := m["dateCreated_GTE"]; ok && val != nil {
		x.DateCreated_GTE = val.(primitive.DateTime)
	}
	if val, ok := m["dateCreated_LTE"]; ok && val != nil {
		x.DateCreated_LTE = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated_EQ"]; ok && val != nil {
		x.DateUpdated_EQ = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated_GTE"]; ok && val != nil {
		x.DateUpdated_GTE = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated_LTE"]; ok && val != nil {
		x.DateUpdated_LTE = val.(primitive.DateTime)
	}
	if val, ok := m["email_CON"]; ok && val != nil {
		x.Email_CON = val.(string)
	}
	if val, ok := m["email_EQ"]; ok && val != nil {
		x.Email_EQ = val.(string)
	}
	if val, ok := m["email_REG"]; ok && val != nil {
		x.Email_REG = val.(string)
	}
	if val, ok := m["extensionType_EQ"]; ok && val != nil {
		x.ExtensionType_EQ = ExtensionType("ExtensionType_EQ_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["id_CON"]; ok && val != nil {
		x.Id_CON = val.(string)
	}
	if val, ok := m["id_EQ"]; ok && val != nil {
		x.Id_EQ = val.(string)
	}
	if val, ok := m["id_REG"]; ok && val != nil {
		x.Id_REG = val.(string)
	}
	if val, ok := m["idAudio_CON"]; ok && val != nil {
		x.IdAudio_CON = val.(string)
	}
	if val, ok := m["idAudio_EQ"]; ok && val != nil {
		x.IdAudio_EQ = val.(string)
	}
	if val, ok := m["idAudio_REG"]; ok && val != nil {
		x.IdAudio_REG = val.(string)
	}
	if val, ok := m["idMusicOnHoldGroup_CON"]; ok && val != nil {
		x.IdMusicOnHoldGroup_CON = val.(string)
	}
	if val, ok := m["idMusicOnHoldGroup_EQ"]; ok && val != nil {
		x.IdMusicOnHoldGroup_EQ = val.(string)
	}
	if val, ok := m["idMusicOnHoldGroup_REG"]; ok && val != nil {
		x.IdMusicOnHoldGroup_REG = val.(string)
	}
	if val, ok := m["idsLinesThatCanListenToVoicemail_CON"]; ok && val != nil {
		x.IdsLinesThatCanListenToVoicemail_CON = val.(string)
	}
	if val, ok := m["idsLinesThatCanListenToVoicemail_EQ"]; ok && val != nil {
		x.IdsLinesThatCanListenToVoicemail_EQ = val.(string)
	}
	if val, ok := m["idsLinesThatCanListenToVoicemail_REG"]; ok && val != nil {
		x.IdsLinesThatCanListenToVoicemail_REG = val.(string)
	}
	if val, ok := m["injectExtensionNameToCallerId_EQ"]; ok && val != nil {
		x.InjectExtensionNameToCallerId_EQ = val.(bool)
	}
	if val, ok := m["number_CON"]; ok && val != nil {
		x.Number_CON = val.(string)
	}
	if val, ok := m["number_EQ"]; ok && val != nil {
		x.Number_EQ = val.(string)
	}
	if val, ok := m["number_REG"]; ok && val != nil {
		x.Number_REG = val.(string)
	}
	if val, ok := m["textToSpeech_CON"]; ok && val != nil {
		x.TextToSpeech_CON = val.(string)
	}
	if val, ok := m["textToSpeech_EQ"]; ok && val != nil {
		x.TextToSpeech_EQ = val.(string)
	}
	if val, ok := m["textToSpeech_REG"]; ok && val != nil {
		x.TextToSpeech_REG = val.(string)
	}
	if val, ok := m["textToSpeechVoiceId_CON"]; ok && val != nil {
		x.TextToSpeechVoiceId_CON = val.(string)
	}
	if val, ok := m["textToSpeechVoiceId_EQ"]; ok && val != nil {
		x.TextToSpeechVoiceId_EQ = val.(string)
	}
	if val, ok := m["textToSpeechVoiceId_REG"]; ok && val != nil {
		x.TextToSpeechVoiceId_REG = val.(string)
	}
}
