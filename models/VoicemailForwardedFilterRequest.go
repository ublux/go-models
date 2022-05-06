package models

import "go.mongodb.org/mongo-driver/bson/primitive"
import . "github.com/ublux/go-models/enums"

type VoicemailForwardedFilterRequest struct {
	DateCreated_EQ                       primitive.DateTime `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE                      primitive.DateTime `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE                      primitive.DateTime `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ                       primitive.DateTime `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE                      primitive.DateTime `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE                      primitive.DateTime `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	DurationInSeconds_EQ                 int32              `bson:"durationInSeconds_EQ" json:"durationInSeconds_EQ"`
	DurationInSeconds_GTE                int32              `bson:"durationInSeconds_GTE" json:"durationInSeconds_GTE"`
	DurationInSeconds_LTE                int32              `bson:"durationInSeconds_LTE" json:"durationInSeconds_LTE"`
	Email_CON                            string             `bson:"email_CON" json:"email_CON"`
	Email_EQ                             string             `bson:"email_EQ" json:"email_EQ"`
	Email_REG                            string             `bson:"email_REG" json:"email_REG"`
	ErrorMessage_CON                     string             `bson:"errorMessage_CON" json:"errorMessage_CON"`
	ErrorMessage_EQ                      string             `bson:"errorMessage_EQ" json:"errorMessage_EQ"`
	ErrorMessage_REG                     string             `bson:"errorMessage_REG" json:"errorMessage_REG"`
	Id_CON                               string             `bson:"id_CON" json:"id_CON"`
	Id_EQ                                string             `bson:"id_EQ" json:"id_EQ"`
	Id_REG                               string             `bson:"id_REG" json:"id_REG"`
	IdExtensionForwarde_CON              string             `bson:"idExtensionForwarde_CON" json:"idExtensionForwarde_CON"`
	IdExtensionForwarde_EQ               string             `bson:"idExtensionForwarde_EQ" json:"idExtensionForwarde_EQ"`
	IdExtensionForwarde_REG              string             `bson:"idExtensionForwarde_REG" json:"idExtensionForwarde_REG"`
	IdLineThatForwardedVoicemail_CON     string             `bson:"idLineThatForwardedVoicemail_CON" json:"idLineThatForwardedVoicemail_CON"`
	IdLineThatForwardedVoicemail_EQ      string             `bson:"idLineThatForwardedVoicemail_EQ" json:"idLineThatForwardedVoicemail_EQ"`
	IdLineThatForwardedVoicemail_REG     string             `bson:"idLineThatForwardedVoicemail_REG" json:"idLineThatForwardedVoicemail_REG"`
	IdsLinesThatCanListenToVoicemail_CON string             `bson:"idsLinesThatCanListenToVoicemail_CON" json:"idsLinesThatCanListenToVoicemail_CON"`
	IdsLinesThatCanListenToVoicemail_EQ  string             `bson:"idsLinesThatCanListenToVoicemail_EQ" json:"idsLinesThatCanListenToVoicemail_EQ"`
	IdsLinesThatCanListenToVoicemail_REG string             `bson:"idsLinesThatCanListenToVoicemail_REG" json:"idsLinesThatCanListenToVoicemail_REG"`
	VoicemailMp3_FileSizeInBytes_EQ      int32              `bson:"voicemailMp3_FileSizeInBytes_EQ" json:"voicemailMp3_FileSizeInBytes_EQ"`
	VoicemailMp3_FileSizeInBytes_GTE     int32              `bson:"voicemailMp3_FileSizeInBytes_GTE" json:"voicemailMp3_FileSizeInBytes_GTE"`
	VoicemailMp3_FileSizeInBytes_LTE     int32              `bson:"voicemailMp3_FileSizeInBytes_LTE" json:"voicemailMp3_FileSizeInBytes_LTE"`
	VoicemailMp3_Id_CON                  string             `bson:"voicemailMp3_Id_CON" json:"voicemailMp3_Id_CON"`
	VoicemailMp3_Id_EQ                   string             `bson:"voicemailMp3_Id_EQ" json:"voicemailMp3_Id_EQ"`
	VoicemailMp3_Id_REG                  string             `bson:"voicemailMp3_Id_REG" json:"voicemailMp3_Id_REG"`
	VoicemailMp3_Md5Hash_CON             string             `bson:"voicemailMp3_Md5Hash_CON" json:"voicemailMp3_Md5Hash_CON"`
	VoicemailMp3_Md5Hash_EQ              string             `bson:"voicemailMp3_Md5Hash_EQ" json:"voicemailMp3_Md5Hash_EQ"`
	VoicemailMp3_Md5Hash_REG             string             `bson:"voicemailMp3_Md5Hash_REG" json:"voicemailMp3_Md5Hash_REG"`
	VoicemailMp3_Url_CON                 string             `bson:"voicemailMp3_Url_CON" json:"voicemailMp3_Url_CON"`
	VoicemailMp3_Url_EQ                  string             `bson:"voicemailMp3_Url_EQ" json:"voicemailMp3_Url_EQ"`
	VoicemailMp3_Url_REG                 string             `bson:"voicemailMp3_Url_REG" json:"voicemailMp3_Url_REG"`
	VoicemailType_EQ                     VoicemailType      `bson:"voicemailType_EQ" json:"voicemailType_EQ"`
	VoicemailWav_FileSizeInBytes_EQ      int32              `bson:"voicemailWav_FileSizeInBytes_EQ" json:"voicemailWav_FileSizeInBytes_EQ"`
	VoicemailWav_FileSizeInBytes_GTE     int32              `bson:"voicemailWav_FileSizeInBytes_GTE" json:"voicemailWav_FileSizeInBytes_GTE"`
	VoicemailWav_FileSizeInBytes_LTE     int32              `bson:"voicemailWav_FileSizeInBytes_LTE" json:"voicemailWav_FileSizeInBytes_LTE"`
	VoicemailWav_Id_CON                  string             `bson:"voicemailWav_Id_CON" json:"voicemailWav_Id_CON"`
	VoicemailWav_Id_EQ                   string             `bson:"voicemailWav_Id_EQ" json:"voicemailWav_Id_EQ"`
	VoicemailWav_Id_REG                  string             `bson:"voicemailWav_Id_REG" json:"voicemailWav_Id_REG"`
	VoicemailWav_Md5Hash_CON             string             `bson:"voicemailWav_Md5Hash_CON" json:"voicemailWav_Md5Hash_CON"`
	VoicemailWav_Md5Hash_EQ              string             `bson:"voicemailWav_Md5Hash_EQ" json:"voicemailWav_Md5Hash_EQ"`
	VoicemailWav_Md5Hash_REG             string             `bson:"voicemailWav_Md5Hash_REG" json:"voicemailWav_Md5Hash_REG"`
	VoicemailWav_Url_CON                 string             `bson:"voicemailWav_Url_CON" json:"voicemailWav_Url_CON"`
	VoicemailWav_Url_EQ                  string             `bson:"voicemailWav_Url_EQ" json:"voicemailWav_Url_EQ"`
	VoicemailWav_Url_REG                 string             `bson:"voicemailWav_Url_REG" json:"voicemailWav_Url_REG"`
}

// BUILDER from bson map:
func BuildVoicemailForwardedFilterRequest(m map[string]interface{}, x *VoicemailForwardedFilterRequest) {
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
	if val, ok := m["durationInSeconds_EQ"]; ok && val != nil {
		x.DurationInSeconds_EQ = val.(int32)
	}
	if val, ok := m["durationInSeconds_GTE"]; ok && val != nil {
		x.DurationInSeconds_GTE = val.(int32)
	}
	if val, ok := m["durationInSeconds_LTE"]; ok && val != nil {
		x.DurationInSeconds_LTE = val.(int32)
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
	if val, ok := m["errorMessage_CON"]; ok && val != nil {
		x.ErrorMessage_CON = val.(string)
	}
	if val, ok := m["errorMessage_EQ"]; ok && val != nil {
		x.ErrorMessage_EQ = val.(string)
	}
	if val, ok := m["errorMessage_REG"]; ok && val != nil {
		x.ErrorMessage_REG = val.(string)
	}
	if val, ok := m["id_CON"]; ok && val != nil {
		x.Id_CON = val.(string)
	}
	if val, ok := m["id_EQ"]; ok && val != nil {
		x.Id_EQ = val.(string)
	}
	if val, ok := m["id_REG"]; ok && val != nil {
		x.Id_REG = val.(string)
	}
	if val, ok := m["idExtensionForwarde_CON"]; ok && val != nil {
		x.IdExtensionForwarde_CON = val.(string)
	}
	if val, ok := m["idExtensionForwarde_EQ"]; ok && val != nil {
		x.IdExtensionForwarde_EQ = val.(string)
	}
	if val, ok := m["idExtensionForwarde_REG"]; ok && val != nil {
		x.IdExtensionForwarde_REG = val.(string)
	}
	if val, ok := m["idLineThatForwardedVoicemail_CON"]; ok && val != nil {
		x.IdLineThatForwardedVoicemail_CON = val.(string)
	}
	if val, ok := m["idLineThatForwardedVoicemail_EQ"]; ok && val != nil {
		x.IdLineThatForwardedVoicemail_EQ = val.(string)
	}
	if val, ok := m["idLineThatForwardedVoicemail_REG"]; ok && val != nil {
		x.IdLineThatForwardedVoicemail_REG = val.(string)
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
	if val, ok := m["voicemailMp3_FileSizeInBytes_EQ"]; ok && val != nil {
		x.VoicemailMp3_FileSizeInBytes_EQ = val.(int32)
	}
	if val, ok := m["voicemailMp3_FileSizeInBytes_GTE"]; ok && val != nil {
		x.VoicemailMp3_FileSizeInBytes_GTE = val.(int32)
	}
	if val, ok := m["voicemailMp3_FileSizeInBytes_LTE"]; ok && val != nil {
		x.VoicemailMp3_FileSizeInBytes_LTE = val.(int32)
	}
	if val, ok := m["voicemailMp3_Id_CON"]; ok && val != nil {
		x.VoicemailMp3_Id_CON = val.(string)
	}
	if val, ok := m["voicemailMp3_Id_EQ"]; ok && val != nil {
		x.VoicemailMp3_Id_EQ = val.(string)
	}
	if val, ok := m["voicemailMp3_Id_REG"]; ok && val != nil {
		x.VoicemailMp3_Id_REG = val.(string)
	}
	if val, ok := m["voicemailMp3_Md5Hash_CON"]; ok && val != nil {
		x.VoicemailMp3_Md5Hash_CON = val.(string)
	}
	if val, ok := m["voicemailMp3_Md5Hash_EQ"]; ok && val != nil {
		x.VoicemailMp3_Md5Hash_EQ = val.(string)
	}
	if val, ok := m["voicemailMp3_Md5Hash_REG"]; ok && val != nil {
		x.VoicemailMp3_Md5Hash_REG = val.(string)
	}
	if val, ok := m["voicemailMp3_Url_CON"]; ok && val != nil {
		x.VoicemailMp3_Url_CON = val.(string)
	}
	if val, ok := m["voicemailMp3_Url_EQ"]; ok && val != nil {
		x.VoicemailMp3_Url_EQ = val.(string)
	}
	if val, ok := m["voicemailMp3_Url_REG"]; ok && val != nil {
		x.VoicemailMp3_Url_REG = val.(string)
	}
	if val, ok := m["voicemailType_EQ"]; ok && val != nil {
		x.VoicemailType_EQ = VoicemailType("VoicemailType_EQ_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["voicemailWav_FileSizeInBytes_EQ"]; ok && val != nil {
		x.VoicemailWav_FileSizeInBytes_EQ = val.(int32)
	}
	if val, ok := m["voicemailWav_FileSizeInBytes_GTE"]; ok && val != nil {
		x.VoicemailWav_FileSizeInBytes_GTE = val.(int32)
	}
	if val, ok := m["voicemailWav_FileSizeInBytes_LTE"]; ok && val != nil {
		x.VoicemailWav_FileSizeInBytes_LTE = val.(int32)
	}
	if val, ok := m["voicemailWav_Id_CON"]; ok && val != nil {
		x.VoicemailWav_Id_CON = val.(string)
	}
	if val, ok := m["voicemailWav_Id_EQ"]; ok && val != nil {
		x.VoicemailWav_Id_EQ = val.(string)
	}
	if val, ok := m["voicemailWav_Id_REG"]; ok && val != nil {
		x.VoicemailWav_Id_REG = val.(string)
	}
	if val, ok := m["voicemailWav_Md5Hash_CON"]; ok && val != nil {
		x.VoicemailWav_Md5Hash_CON = val.(string)
	}
	if val, ok := m["voicemailWav_Md5Hash_EQ"]; ok && val != nil {
		x.VoicemailWav_Md5Hash_EQ = val.(string)
	}
	if val, ok := m["voicemailWav_Md5Hash_REG"]; ok && val != nil {
		x.VoicemailWav_Md5Hash_REG = val.(string)
	}
	if val, ok := m["voicemailWav_Url_CON"]; ok && val != nil {
		x.VoicemailWav_Url_CON = val.(string)
	}
	if val, ok := m["voicemailWav_Url_EQ"]; ok && val != nil {
		x.VoicemailWav_Url_EQ = val.(string)
	}
	if val, ok := m["voicemailWav_Url_REG"]; ok && val != nil {
		x.VoicemailWav_Url_REG = val.(string)
	}
}
