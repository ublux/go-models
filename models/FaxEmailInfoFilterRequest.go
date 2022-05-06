package models

import "time"
import . "github.com/ublux/go-models/enums"

type FaxEmailInfoFilterRequest struct {
	DateCreated_EQ               time.Time            `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE              time.Time            `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE              time.Time            `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ               time.Time            `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE              time.Time            `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE              time.Time            `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	Id_CON                       string               `bson:"id_CON" json:"id_CON"`
	Id_EQ                        string               `bson:"id_EQ" json:"id_EQ"`
	Id_REG                       string               `bson:"id_REG" json:"id_REG"`
	IdLineThatValidatedEmail_CON string               `bson:"idLineThatValidatedEmail_CON" json:"idLineThatValidatedEmail_CON"`
	IdLineThatValidatedEmail_EQ  string               `bson:"idLineThatValidatedEmail_EQ" json:"idLineThatValidatedEmail_EQ"`
	IdLineThatValidatedEmail_REG string               `bson:"idLineThatValidatedEmail_REG" json:"idLineThatValidatedEmail_REG"`
	NumberOfEmailsReceived_EQ    int32                `bson:"numberOfEmailsReceived_EQ" json:"numberOfEmailsReceived_EQ"`
	NumberOfEmailsReceived_GTE   int32                `bson:"numberOfEmailsReceived_GTE" json:"numberOfEmailsReceived_GTE"`
	NumberOfEmailsReceived_LTE   int32                `bson:"numberOfEmailsReceived_LTE" json:"numberOfEmailsReceived_LTE"`
	ReplyStatus_EQ               LinkFaxToEmailStatus `bson:"replyStatus_EQ" json:"replyStatus_EQ"`
}

// BUILDER from bson map:
func BuildFaxEmailInfoFilterRequest(m map[string]interface{}, x *FaxEmailInfoFilterRequest) {
	if val, ok := m["dateCreated_EQ"]; ok && val != nil {
		x.DateCreated_EQ = val.(time.Time)
	}
	if val, ok := m["dateCreated_GTE"]; ok && val != nil {
		x.DateCreated_GTE = val.(time.Time)
	}
	if val, ok := m["dateCreated_LTE"]; ok && val != nil {
		x.DateCreated_LTE = val.(time.Time)
	}
	if val, ok := m["dateUpdated_EQ"]; ok && val != nil {
		x.DateUpdated_EQ = val.(time.Time)
	}
	if val, ok := m["dateUpdated_GTE"]; ok && val != nil {
		x.DateUpdated_GTE = val.(time.Time)
	}
	if val, ok := m["dateUpdated_LTE"]; ok && val != nil {
		x.DateUpdated_LTE = val.(time.Time)
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
	if val, ok := m["idLineThatValidatedEmail_CON"]; ok && val != nil {
		x.IdLineThatValidatedEmail_CON = val.(string)
	}
	if val, ok := m["idLineThatValidatedEmail_EQ"]; ok && val != nil {
		x.IdLineThatValidatedEmail_EQ = val.(string)
	}
	if val, ok := m["idLineThatValidatedEmail_REG"]; ok && val != nil {
		x.IdLineThatValidatedEmail_REG = val.(string)
	}
	if val, ok := m["numberOfEmailsReceived_EQ"]; ok && val != nil {
		x.NumberOfEmailsReceived_EQ = val.(int32)
	}
	if val, ok := m["numberOfEmailsReceived_GTE"]; ok && val != nil {
		x.NumberOfEmailsReceived_GTE = val.(int32)
	}
	if val, ok := m["numberOfEmailsReceived_LTE"]; ok && val != nil {
		x.NumberOfEmailsReceived_LTE = val.(int32)
	}
	if val, ok := m["replyStatus_EQ"]; ok && val != nil {
		x.ReplyStatus_EQ = LinkFaxToEmailStatus("ReplyStatus_EQ_" + val.(string))
	} // is NOT readonly obtained from map
}
