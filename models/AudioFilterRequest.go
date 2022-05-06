package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AudioFilterRequest struct {
	AudioMp3_FileSizeInBytes_EQ  int32              `bson:"audioMp3_FileSizeInBytes_EQ" json:"audioMp3_FileSizeInBytes_EQ"`
	AudioMp3_FileSizeInBytes_GTE int32              `bson:"audioMp3_FileSizeInBytes_GTE" json:"audioMp3_FileSizeInBytes_GTE"`
	AudioMp3_FileSizeInBytes_LTE int32              `bson:"audioMp3_FileSizeInBytes_LTE" json:"audioMp3_FileSizeInBytes_LTE"`
	AudioMp3_Id_CON              string             `bson:"audioMp3_Id_CON" json:"audioMp3_Id_CON"`
	AudioMp3_Id_EQ               string             `bson:"audioMp3_Id_EQ" json:"audioMp3_Id_EQ"`
	AudioMp3_Id_REG              string             `bson:"audioMp3_Id_REG" json:"audioMp3_Id_REG"`
	AudioMp3_Md5Hash_CON         string             `bson:"audioMp3_Md5Hash_CON" json:"audioMp3_Md5Hash_CON"`
	AudioMp3_Md5Hash_EQ          string             `bson:"audioMp3_Md5Hash_EQ" json:"audioMp3_Md5Hash_EQ"`
	AudioMp3_Md5Hash_REG         string             `bson:"audioMp3_Md5Hash_REG" json:"audioMp3_Md5Hash_REG"`
	AudioMp3_Url_CON             string             `bson:"audioMp3_Url_CON" json:"audioMp3_Url_CON"`
	AudioMp3_Url_EQ              string             `bson:"audioMp3_Url_EQ" json:"audioMp3_Url_EQ"`
	AudioMp3_Url_REG             string             `bson:"audioMp3_Url_REG" json:"audioMp3_Url_REG"`
	AudioWav_FileSizeInBytes_EQ  int32              `bson:"audioWav_FileSizeInBytes_EQ" json:"audioWav_FileSizeInBytes_EQ"`
	AudioWav_FileSizeInBytes_GTE int32              `bson:"audioWav_FileSizeInBytes_GTE" json:"audioWav_FileSizeInBytes_GTE"`
	AudioWav_FileSizeInBytes_LTE int32              `bson:"audioWav_FileSizeInBytes_LTE" json:"audioWav_FileSizeInBytes_LTE"`
	AudioWav_Id_CON              string             `bson:"audioWav_Id_CON" json:"audioWav_Id_CON"`
	AudioWav_Id_EQ               string             `bson:"audioWav_Id_EQ" json:"audioWav_Id_EQ"`
	AudioWav_Id_REG              string             `bson:"audioWav_Id_REG" json:"audioWav_Id_REG"`
	AudioWav_Md5Hash_CON         string             `bson:"audioWav_Md5Hash_CON" json:"audioWav_Md5Hash_CON"`
	AudioWav_Md5Hash_EQ          string             `bson:"audioWav_Md5Hash_EQ" json:"audioWav_Md5Hash_EQ"`
	AudioWav_Md5Hash_REG         string             `bson:"audioWav_Md5Hash_REG" json:"audioWav_Md5Hash_REG"`
	AudioWav_Url_CON             string             `bson:"audioWav_Url_CON" json:"audioWav_Url_CON"`
	AudioWav_Url_EQ              string             `bson:"audioWav_Url_EQ" json:"audioWav_Url_EQ"`
	AudioWav_Url_REG             string             `bson:"audioWav_Url_REG" json:"audioWav_Url_REG"`
	DateCreated_EQ               primitive.DateTime `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE              primitive.DateTime `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE              primitive.DateTime `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ               primitive.DateTime `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE              primitive.DateTime `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE              primitive.DateTime `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	Description_CON              string             `bson:"description_CON" json:"description_CON"`
	Description_EQ               string             `bson:"description_EQ" json:"description_EQ"`
	Description_REG              string             `bson:"description_REG" json:"description_REG"`
	DurationInSeconds_EQ         int32              `bson:"durationInSeconds_EQ" json:"durationInSeconds_EQ"`
	DurationInSeconds_GTE        int32              `bson:"durationInSeconds_GTE" json:"durationInSeconds_GTE"`
	DurationInSeconds_LTE        int32              `bson:"durationInSeconds_LTE" json:"durationInSeconds_LTE"`
	FriendlyName_CON             string             `bson:"friendlyName_CON" json:"friendlyName_CON"`
	FriendlyName_EQ              string             `bson:"friendlyName_EQ" json:"friendlyName_EQ"`
	FriendlyName_REG             string             `bson:"friendlyName_REG" json:"friendlyName_REG"`
	Id_CON                       string             `bson:"id_CON" json:"id_CON"`
	Id_EQ                        string             `bson:"id_EQ" json:"id_EQ"`
	Id_REG                       string             `bson:"id_REG" json:"id_REG"`
}

// BUILDER from bson map:
func BuildAudioFilterRequest(m map[string]interface{}, x *AudioFilterRequest) {
	if val, ok := m["audioMp3_FileSizeInBytes_EQ"]; ok && val != nil {
		x.AudioMp3_FileSizeInBytes_EQ = val.(int32)
	}
	if val, ok := m["audioMp3_FileSizeInBytes_GTE"]; ok && val != nil {
		x.AudioMp3_FileSizeInBytes_GTE = val.(int32)
	}
	if val, ok := m["audioMp3_FileSizeInBytes_LTE"]; ok && val != nil {
		x.AudioMp3_FileSizeInBytes_LTE = val.(int32)
	}
	if val, ok := m["audioMp3_Id_CON"]; ok && val != nil {
		x.AudioMp3_Id_CON = val.(string)
	}
	if val, ok := m["audioMp3_Id_EQ"]; ok && val != nil {
		x.AudioMp3_Id_EQ = val.(string)
	}
	if val, ok := m["audioMp3_Id_REG"]; ok && val != nil {
		x.AudioMp3_Id_REG = val.(string)
	}
	if val, ok := m["audioMp3_Md5Hash_CON"]; ok && val != nil {
		x.AudioMp3_Md5Hash_CON = val.(string)
	}
	if val, ok := m["audioMp3_Md5Hash_EQ"]; ok && val != nil {
		x.AudioMp3_Md5Hash_EQ = val.(string)
	}
	if val, ok := m["audioMp3_Md5Hash_REG"]; ok && val != nil {
		x.AudioMp3_Md5Hash_REG = val.(string)
	}
	if val, ok := m["audioMp3_Url_CON"]; ok && val != nil {
		x.AudioMp3_Url_CON = val.(string)
	}
	if val, ok := m["audioMp3_Url_EQ"]; ok && val != nil {
		x.AudioMp3_Url_EQ = val.(string)
	}
	if val, ok := m["audioMp3_Url_REG"]; ok && val != nil {
		x.AudioMp3_Url_REG = val.(string)
	}
	if val, ok := m["audioWav_FileSizeInBytes_EQ"]; ok && val != nil {
		x.AudioWav_FileSizeInBytes_EQ = val.(int32)
	}
	if val, ok := m["audioWav_FileSizeInBytes_GTE"]; ok && val != nil {
		x.AudioWav_FileSizeInBytes_GTE = val.(int32)
	}
	if val, ok := m["audioWav_FileSizeInBytes_LTE"]; ok && val != nil {
		x.AudioWav_FileSizeInBytes_LTE = val.(int32)
	}
	if val, ok := m["audioWav_Id_CON"]; ok && val != nil {
		x.AudioWav_Id_CON = val.(string)
	}
	if val, ok := m["audioWav_Id_EQ"]; ok && val != nil {
		x.AudioWav_Id_EQ = val.(string)
	}
	if val, ok := m["audioWav_Id_REG"]; ok && val != nil {
		x.AudioWav_Id_REG = val.(string)
	}
	if val, ok := m["audioWav_Md5Hash_CON"]; ok && val != nil {
		x.AudioWav_Md5Hash_CON = val.(string)
	}
	if val, ok := m["audioWav_Md5Hash_EQ"]; ok && val != nil {
		x.AudioWav_Md5Hash_EQ = val.(string)
	}
	if val, ok := m["audioWav_Md5Hash_REG"]; ok && val != nil {
		x.AudioWav_Md5Hash_REG = val.(string)
	}
	if val, ok := m["audioWav_Url_CON"]; ok && val != nil {
		x.AudioWav_Url_CON = val.(string)
	}
	if val, ok := m["audioWav_Url_EQ"]; ok && val != nil {
		x.AudioWav_Url_EQ = val.(string)
	}
	if val, ok := m["audioWav_Url_REG"]; ok && val != nil {
		x.AudioWav_Url_REG = val.(string)
	}
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
	if val, ok := m["description_CON"]; ok && val != nil {
		x.Description_CON = val.(string)
	}
	if val, ok := m["description_EQ"]; ok && val != nil {
		x.Description_EQ = val.(string)
	}
	if val, ok := m["description_REG"]; ok && val != nil {
		x.Description_REG = val.(string)
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
	if val, ok := m["friendlyName_CON"]; ok && val != nil {
		x.FriendlyName_CON = val.(string)
	}
	if val, ok := m["friendlyName_EQ"]; ok && val != nil {
		x.FriendlyName_EQ = val.(string)
	}
	if val, ok := m["friendlyName_REG"]; ok && val != nil {
		x.FriendlyName_REG = val.(string)
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
}
