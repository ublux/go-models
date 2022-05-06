package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UnauthorizedTokenFilterRequest struct {
	DateCreated_EQ     primitive.DateTime `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE    primitive.DateTime `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE    primitive.DateTime `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ     primitive.DateTime `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE    primitive.DateTime `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE    primitive.DateTime `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	ExpirationDate_EQ  primitive.DateTime `bson:"expirationDate_EQ" json:"expirationDate_EQ"`
	ExpirationDate_GTE primitive.DateTime `bson:"expirationDate_GTE" json:"expirationDate_GTE"`
	ExpirationDate_LTE primitive.DateTime `bson:"expirationDate_LTE" json:"expirationDate_LTE"`
	Id_CON             string             `bson:"id_CON" json:"id_CON"`
	Id_EQ              string             `bson:"id_EQ" json:"id_EQ"`
	Id_REG             string             `bson:"id_REG" json:"id_REG"`
	IdIdentity_CON     string             `bson:"idIdentity_CON" json:"idIdentity_CON"`
	IdIdentity_EQ      string             `bson:"idIdentity_EQ" json:"idIdentity_EQ"`
	IdIdentity_REG     string             `bson:"idIdentity_REG" json:"idIdentity_REG"`
	Jwt_CON            string             `bson:"jwt_CON" json:"jwt_CON"`
	Jwt_EQ             string             `bson:"jwt_EQ" json:"jwt_EQ"`
	Jwt_REG            string             `bson:"jwt_REG" json:"jwt_REG"`
}

// BUILDER from bson map:
func BuildUnauthorizedTokenFilterRequest(m map[string]interface{}, x *UnauthorizedTokenFilterRequest) {
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
	if val, ok := m["expirationDate_EQ"]; ok && val != nil {
		x.ExpirationDate_EQ = val.(primitive.DateTime)
	}
	if val, ok := m["expirationDate_GTE"]; ok && val != nil {
		x.ExpirationDate_GTE = val.(primitive.DateTime)
	}
	if val, ok := m["expirationDate_LTE"]; ok && val != nil {
		x.ExpirationDate_LTE = val.(primitive.DateTime)
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
	if val, ok := m["idIdentity_CON"]; ok && val != nil {
		x.IdIdentity_CON = val.(string)
	}
	if val, ok := m["idIdentity_EQ"]; ok && val != nil {
		x.IdIdentity_EQ = val.(string)
	}
	if val, ok := m["idIdentity_REG"]; ok && val != nil {
		x.IdIdentity_REG = val.(string)
	}
	if val, ok := m["jwt_CON"]; ok && val != nil {
		x.Jwt_CON = val.(string)
	}
	if val, ok := m["jwt_EQ"]; ok && val != nil {
		x.Jwt_EQ = val.(string)
	}
	if val, ok := m["jwt_REG"]; ok && val != nil {
		x.Jwt_REG = val.(string)
	}
}
