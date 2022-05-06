package models

import "go.mongodb.org/mongo-driver/bson/primitive"
import . "github.com/ublux/go-models/enums"

type ExtensionConferenceFilterRequest struct {
	AnnounceParticipants_EQ          bool               `bson:"announceParticipants_EQ" json:"announceParticipants_EQ"`
	DateCreated_EQ                   primitive.DateTime `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE                  primitive.DateTime `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE                  primitive.DateTime `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ                   primitive.DateTime `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE                  primitive.DateTime `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE                  primitive.DateTime `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	ExtensionType_EQ                 ExtensionType      `bson:"extensionType_EQ" json:"extensionType_EQ"`
	Id_CON                           string             `bson:"id_CON" json:"id_CON"`
	Id_EQ                            string             `bson:"id_EQ" json:"id_EQ"`
	Id_REG                           string             `bson:"id_REG" json:"id_REG"`
	IdMusicOnHoldGroup_CON           string             `bson:"idMusicOnHoldGroup_CON" json:"idMusicOnHoldGroup_CON"`
	IdMusicOnHoldGroup_EQ            string             `bson:"idMusicOnHoldGroup_EQ" json:"idMusicOnHoldGroup_EQ"`
	IdMusicOnHoldGroup_REG           string             `bson:"idMusicOnHoldGroup_REG" json:"idMusicOnHoldGroup_REG"`
	IdsAudiosWhenOneParticipant_CON  string             `bson:"idsAudiosWhenOneParticipant_CON" json:"idsAudiosWhenOneParticipant_CON"`
	IdsAudiosWhenOneParticipant_EQ   string             `bson:"idsAudiosWhenOneParticipant_EQ" json:"idsAudiosWhenOneParticipant_EQ"`
	IdsAudiosWhenOneParticipant_REG  string             `bson:"idsAudiosWhenOneParticipant_REG" json:"idsAudiosWhenOneParticipant_REG"`
	InjectExtensionNameToCallerId_EQ bool               `bson:"injectExtensionNameToCallerId_EQ" json:"injectExtensionNameToCallerId_EQ"`
	Number_CON                       string             `bson:"number_CON" json:"number_CON"`
	Number_EQ                        string             `bson:"number_EQ" json:"number_EQ"`
	Number_REG                       string             `bson:"number_REG" json:"number_REG"`
	Pin_CON                          string             `bson:"pin_CON" json:"pin_CON"`
	Pin_EQ                           string             `bson:"pin_EQ" json:"pin_EQ"`
	Pin_REG                          string             `bson:"pin_REG" json:"pin_REG"`
}

// BUILDER from bson map:
func BuildExtensionConferenceFilterRequest(m map[string]interface{}, x *ExtensionConferenceFilterRequest) {
	if val, ok := m["announceParticipants_EQ"]; ok && val != nil {
		x.AnnounceParticipants_EQ = val.(bool)
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
	if val, ok := m["idMusicOnHoldGroup_CON"]; ok && val != nil {
		x.IdMusicOnHoldGroup_CON = val.(string)
	}
	if val, ok := m["idMusicOnHoldGroup_EQ"]; ok && val != nil {
		x.IdMusicOnHoldGroup_EQ = val.(string)
	}
	if val, ok := m["idMusicOnHoldGroup_REG"]; ok && val != nil {
		x.IdMusicOnHoldGroup_REG = val.(string)
	}
	if val, ok := m["idsAudiosWhenOneParticipant_CON"]; ok && val != nil {
		x.IdsAudiosWhenOneParticipant_CON = val.(string)
	}
	if val, ok := m["idsAudiosWhenOneParticipant_EQ"]; ok && val != nil {
		x.IdsAudiosWhenOneParticipant_EQ = val.(string)
	}
	if val, ok := m["idsAudiosWhenOneParticipant_REG"]; ok && val != nil {
		x.IdsAudiosWhenOneParticipant_REG = val.(string)
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
	if val, ok := m["pin_CON"]; ok && val != nil {
		x.Pin_CON = val.(string)
	}
	if val, ok := m["pin_EQ"]; ok && val != nil {
		x.Pin_EQ = val.(string)
	}
	if val, ok := m["pin_REG"]; ok && val != nil {
		x.Pin_REG = val.(string)
	}
}
