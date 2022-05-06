package models

import "go.mongodb.org/mongo-driver/bson/primitive"
import . "github.com/ublux/go-models/enums"

type LineKeyGroupFilterRequest struct {
	DateCreated_EQ           primitive.DateTime `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE          primitive.DateTime `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE          primitive.DateTime `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ           primitive.DateTime `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE          primitive.DateTime `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE          primitive.DateTime `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	Description_CON          string             `bson:"description_CON" json:"description_CON"`
	Description_EQ           string             `bson:"description_EQ" json:"description_EQ"`
	Description_REG          string             `bson:"description_REG" json:"description_REG"`
	FriendlyName_CON         string             `bson:"friendlyName_CON" json:"friendlyName_CON"`
	FriendlyName_EQ          string             `bson:"friendlyName_EQ" json:"friendlyName_EQ"`
	FriendlyName_REG         string             `bson:"friendlyName_REG" json:"friendlyName_REG"`
	Id_CON                   string             `bson:"id_CON" json:"id_CON"`
	Id_EQ                    string             `bson:"id_EQ" json:"id_EQ"`
	Id_REG                   string             `bson:"id_REG" json:"id_REG"`
	LineKeys_IdExtension_CON string             `bson:"lineKeys_IdExtension_CON" json:"lineKeys_IdExtension_CON"`
	LineKeys_IdExtension_EQ  string             `bson:"lineKeys_IdExtension_EQ" json:"lineKeys_IdExtension_EQ"`
	LineKeys_IdExtension_REG string             `bson:"lineKeys_IdExtension_REG" json:"lineKeys_IdExtension_REG"`
	LineKeys_Label_CON       string             `bson:"lineKeys_Label_CON" json:"lineKeys_Label_CON"`
	LineKeys_Label_EQ        string             `bson:"lineKeys_Label_EQ" json:"lineKeys_Label_EQ"`
	LineKeys_Label_REG       string             `bson:"lineKeys_Label_REG" json:"lineKeys_Label_REG"`
	LineKeys_LineIndex_EQ    int32              `bson:"lineKeys_LineIndex_EQ" json:"lineKeys_LineIndex_EQ"`
	LineKeys_LineIndex_GTE   int32              `bson:"lineKeys_LineIndex_GTE" json:"lineKeys_LineIndex_GTE"`
	LineKeys_LineIndex_LTE   int32              `bson:"lineKeys_LineIndex_LTE" json:"lineKeys_LineIndex_LTE"`
	LineKeys_LineKeyType_EQ  LineKeyType        `bson:"lineKeys_LineKeyType_EQ" json:"lineKeys_LineKeyType_EQ"`
	LineKeys_Value_CON       string             `bson:"lineKeys_Value_CON" json:"lineKeys_Value_CON"`
	LineKeys_Value_EQ        string             `bson:"lineKeys_Value_EQ" json:"lineKeys_Value_EQ"`
	LineKeys_Value_REG       string             `bson:"lineKeys_Value_REG" json:"lineKeys_Value_REG"`
}

// BUILDER from bson map:
func BuildLineKeyGroupFilterRequest(m map[string]interface{}, x *LineKeyGroupFilterRequest) {
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
	if val, ok := m["lineKeys_IdExtension_CON"]; ok && val != nil {
		x.LineKeys_IdExtension_CON = val.(string)
	}
	if val, ok := m["lineKeys_IdExtension_EQ"]; ok && val != nil {
		x.LineKeys_IdExtension_EQ = val.(string)
	}
	if val, ok := m["lineKeys_IdExtension_REG"]; ok && val != nil {
		x.LineKeys_IdExtension_REG = val.(string)
	}
	if val, ok := m["lineKeys_Label_CON"]; ok && val != nil {
		x.LineKeys_Label_CON = val.(string)
	}
	if val, ok := m["lineKeys_Label_EQ"]; ok && val != nil {
		x.LineKeys_Label_EQ = val.(string)
	}
	if val, ok := m["lineKeys_Label_REG"]; ok && val != nil {
		x.LineKeys_Label_REG = val.(string)
	}
	if val, ok := m["lineKeys_LineIndex_EQ"]; ok && val != nil {
		x.LineKeys_LineIndex_EQ = val.(int32)
	}
	if val, ok := m["lineKeys_LineIndex_GTE"]; ok && val != nil {
		x.LineKeys_LineIndex_GTE = val.(int32)
	}
	if val, ok := m["lineKeys_LineIndex_LTE"]; ok && val != nil {
		x.LineKeys_LineIndex_LTE = val.(int32)
	}
	if val, ok := m["lineKeys_LineKeyType_EQ"]; ok && val != nil {
		x.LineKeys_LineKeyType_EQ = LineKeyType("LineKeys_LineKeyType_EQ_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["lineKeys_Value_CON"]; ok && val != nil {
		x.LineKeys_Value_CON = val.(string)
	}
	if val, ok := m["lineKeys_Value_EQ"]; ok && val != nil {
		x.LineKeys_Value_EQ = val.(string)
	}
	if val, ok := m["lineKeys_Value_REG"]; ok && val != nil {
		x.LineKeys_Value_REG = val.(string)
	}
}
