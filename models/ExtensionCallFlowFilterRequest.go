package models

import "time"
import . "github.com/ublux/go-models/enums"

type ExtensionCallFlowFilterRequest struct {
	CallFlowLabel_CON                string        `bson:"callFlowLabel_CON" json:"callFlowLabel_CON"`
	CallFlowLabel_EQ                 string        `bson:"callFlowLabel_EQ" json:"callFlowLabel_EQ"`
	CallFlowLabel_REG                string        `bson:"callFlowLabel_REG" json:"callFlowLabel_REG"`
	DateCreated_EQ                   time.Time     `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE                  time.Time     `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE                  time.Time     `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ                   time.Time     `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE                  time.Time     `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE                  time.Time     `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	ExtensionType_EQ                 ExtensionType `bson:"extensionType_EQ" json:"extensionType_EQ"`
	Id_CON                           string        `bson:"id_CON" json:"id_CON"`
	Id_EQ                            string        `bson:"id_EQ" json:"id_EQ"`
	Id_REG                           string        `bson:"id_REG" json:"id_REG"`
	IdCallFlow_CON                   string        `bson:"idCallFlow_CON" json:"idCallFlow_CON"`
	IdCallFlow_EQ                    string        `bson:"idCallFlow_EQ" json:"idCallFlow_EQ"`
	IdCallFlow_REG                   string        `bson:"idCallFlow_REG" json:"idCallFlow_REG"`
	IdMusicOnHoldGroup_CON           string        `bson:"idMusicOnHoldGroup_CON" json:"idMusicOnHoldGroup_CON"`
	IdMusicOnHoldGroup_EQ            string        `bson:"idMusicOnHoldGroup_EQ" json:"idMusicOnHoldGroup_EQ"`
	IdMusicOnHoldGroup_REG           string        `bson:"idMusicOnHoldGroup_REG" json:"idMusicOnHoldGroup_REG"`
	InjectExtensionNameToCallerId_EQ bool          `bson:"injectExtensionNameToCallerId_EQ" json:"injectExtensionNameToCallerId_EQ"`
	Number_CON                       string        `bson:"number_CON" json:"number_CON"`
	Number_EQ                        string        `bson:"number_EQ" json:"number_EQ"`
	Number_REG                       string        `bson:"number_REG" json:"number_REG"`
	TimeZone_CON                     string        `bson:"timeZone_CON" json:"timeZone_CON"`
	TimeZone_EQ                      string        `bson:"timeZone_EQ" json:"timeZone_EQ"`
	TimeZone_REG                     string        `bson:"timeZone_REG" json:"timeZone_REG"`
}

// BUILDER from bson map:
func BuildExtensionCallFlowFilterRequest(m map[string]interface{}, x *ExtensionCallFlowFilterRequest) {
	if val, ok := m["callFlowLabel_CON"]; ok && val != nil {
		x.CallFlowLabel_CON = val.(string)
	}
	if val, ok := m["callFlowLabel_EQ"]; ok && val != nil {
		x.CallFlowLabel_EQ = val.(string)
	}
	if val, ok := m["callFlowLabel_REG"]; ok && val != nil {
		x.CallFlowLabel_REG = val.(string)
	}
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
	if val, ok := m["idCallFlow_CON"]; ok && val != nil {
		x.IdCallFlow_CON = val.(string)
	}
	if val, ok := m["idCallFlow_EQ"]; ok && val != nil {
		x.IdCallFlow_EQ = val.(string)
	}
	if val, ok := m["idCallFlow_REG"]; ok && val != nil {
		x.IdCallFlow_REG = val.(string)
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
	if val, ok := m["timeZone_CON"]; ok && val != nil {
		x.TimeZone_CON = val.(string)
	}
	if val, ok := m["timeZone_EQ"]; ok && val != nil {
		x.TimeZone_EQ = val.(string)
	}
	if val, ok := m["timeZone_REG"]; ok && val != nil {
		x.TimeZone_REG = val.(string)
	}
}
