package models

import "time"
import . "github.com/ublux/go-models/enums"

type WebHookFilterRequest struct {
	DateCreated_EQ  time.Time    `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE time.Time    `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE time.Time    `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ  time.Time    `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE time.Time    `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE time.Time    `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	Headers_CON     string       `bson:"headers_CON" json:"headers_CON"`
	Headers_EQ      string       `bson:"headers_EQ" json:"headers_EQ"`
	Headers_REG     string       `bson:"headers_REG" json:"headers_REG"`
	Id_CON          string       `bson:"id_CON" json:"id_CON"`
	Id_EQ           string       `bson:"id_EQ" json:"id_EQ"`
	Id_REG          string       `bson:"id_REG" json:"id_REG"`
	Url_CON         string       `bson:"url_CON" json:"url_CON"`
	Url_EQ          string       `bson:"url_EQ" json:"url_EQ"`
	Url_REG         string       `bson:"url_REG" json:"url_REG"`
	WebHookEvent_EQ WebHookEvent `bson:"webHookEvent_EQ" json:"webHookEvent_EQ"`
}

// BUILDER from bson map:
func BuildWebHookFilterRequest(m map[string]interface{}, x *WebHookFilterRequest) {
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
	if val, ok := m["headers_CON"]; ok && val != nil {
		x.Headers_CON = val.(string)
	}
	if val, ok := m["headers_EQ"]; ok && val != nil {
		x.Headers_EQ = val.(string)
	}
	if val, ok := m["headers_REG"]; ok && val != nil {
		x.Headers_REG = val.(string)
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
	if val, ok := m["url_CON"]; ok && val != nil {
		x.Url_CON = val.(string)
	}
	if val, ok := m["url_EQ"]; ok && val != nil {
		x.Url_EQ = val.(string)
	}
	if val, ok := m["url_REG"]; ok && val != nil {
		x.Url_REG = val.(string)
	}
	if val, ok := m["webHookEvent_EQ"]; ok && val != nil {
		x.WebHookEvent_EQ = WebHookEvent("WebHookEvent_EQ_" + val.(string))
	} // is NOT readonly obtained from map
}
