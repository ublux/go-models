package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CallFlowFilterRequest struct {
	DateCreated_EQ   primitive.DateTime `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE  primitive.DateTime `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE  primitive.DateTime `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ   primitive.DateTime `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE  primitive.DateTime `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE  primitive.DateTime `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	Description_CON  string             `bson:"description_CON" json:"description_CON"`
	Description_EQ   string             `bson:"description_EQ" json:"description_EQ"`
	Description_REG  string             `bson:"description_REG" json:"description_REG"`
	FriendlyName_CON string             `bson:"friendlyName_CON" json:"friendlyName_CON"`
	FriendlyName_EQ  string             `bson:"friendlyName_EQ" json:"friendlyName_EQ"`
	FriendlyName_REG string             `bson:"friendlyName_REG" json:"friendlyName_REG"`
	Id_CON           string             `bson:"id_CON" json:"id_CON"`
	Id_EQ            string             `bson:"id_EQ" json:"id_EQ"`
	Id_REG           string             `bson:"id_REG" json:"id_REG"`
	XmlTree_CON      string             `bson:"xmlTree_CON" json:"xmlTree_CON"`
	XmlTree_EQ       string             `bson:"xmlTree_EQ" json:"xmlTree_EQ"`
	XmlTree_REG      string             `bson:"xmlTree_REG" json:"xmlTree_REG"`
}

// BUILDER from bson map:
func BuildCallFlowFilterRequest(m map[string]interface{}, x *CallFlowFilterRequest) {
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
	if val, ok := m["xmlTree_CON"]; ok && val != nil {
		x.XmlTree_CON = val.(string)
	}
	if val, ok := m["xmlTree_EQ"]; ok && val != nil {
		x.XmlTree_EQ = val.(string)
	}
	if val, ok := m["xmlTree_REG"]; ok && val != nil {
		x.XmlTree_REG = val.(string)
	}
}
