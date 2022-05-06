package models

import "time"
import . "github.com/ublux/go-models/enums"

type PowerDialerGroupFilterRequest struct {
	DateCreated_EQ                      time.Time              `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE                     time.Time              `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE                     time.Time              `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ                      time.Time              `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE                     time.Time              `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE                     time.Time              `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	Description_CON                     string                 `bson:"description_CON" json:"description_CON"`
	Description_EQ                      string                 `bson:"description_EQ" json:"description_EQ"`
	Description_REG                     string                 `bson:"description_REG" json:"description_REG"`
	ErrorMessage_CON                    string                 `bson:"errorMessage_CON" json:"errorMessage_CON"`
	ErrorMessage_EQ                     string                 `bson:"errorMessage_EQ" json:"errorMessage_EQ"`
	ErrorMessage_REG                    string                 `bson:"errorMessage_REG" json:"errorMessage_REG"`
	FriendlyName_CON                    string                 `bson:"friendlyName_CON" json:"friendlyName_CON"`
	FriendlyName_EQ                     string                 `bson:"friendlyName_EQ" json:"friendlyName_EQ"`
	FriendlyName_REG                    string                 `bson:"friendlyName_REG" json:"friendlyName_REG"`
	Id_CON                              string                 `bson:"id_CON" json:"id_CON"`
	Id_EQ                               string                 `bson:"id_EQ" json:"id_EQ"`
	Id_REG                              string                 `bson:"id_REG" json:"id_REG"`
	IdCallerIdMask_CON                  string                 `bson:"idCallerIdMask_CON" json:"idCallerIdMask_CON"`
	IdCallerIdMask_EQ                   string                 `bson:"idCallerIdMask_EQ" json:"idCallerIdMask_EQ"`
	IdCallerIdMask_REG                  string                 `bson:"idCallerIdMask_REG" json:"idCallerIdMask_REG"`
	IdCallFlow_CON                      string                 `bson:"idCallFlow_CON" json:"idCallFlow_CON"`
	IdCallFlow_EQ                       string                 `bson:"idCallFlow_EQ" json:"idCallFlow_EQ"`
	IdCallFlow_REG                      string                 `bson:"idCallFlow_REG" json:"idCallFlow_REG"`
	IdExtension_CON                     string                 `bson:"idExtension_CON" json:"idExtension_CON"`
	IdExtension_EQ                      string                 `bson:"idExtension_EQ" json:"idExtension_EQ"`
	IdExtension_REG                     string                 `bson:"idExtension_REG" json:"idExtension_REG"`
	IdVoipNumberPhone_CON               string                 `bson:"idVoipNumberPhone_CON" json:"idVoipNumberPhone_CON"`
	IdVoipNumberPhone_EQ                string                 `bson:"idVoipNumberPhone_EQ" json:"idVoipNumberPhone_EQ"`
	IdVoipNumberPhone_REG               string                 `bson:"idVoipNumberPhone_REG" json:"idVoipNumberPhone_REG"`
	NumberOfConcurrentCalls_EQ          int32                  `bson:"numberOfConcurrentCalls_EQ" json:"numberOfConcurrentCalls_EQ"`
	NumberOfConcurrentCalls_GTE         int32                  `bson:"numberOfConcurrentCalls_GTE" json:"numberOfConcurrentCalls_GTE"`
	NumberOfConcurrentCalls_LTE         int32                  `bson:"numberOfConcurrentCalls_LTE" json:"numberOfConcurrentCalls_LTE"`
	PowerDialerExecutingRecordIndex_EQ  int32                  `bson:"powerDialerExecutingRecordIndex_EQ" json:"powerDialerExecutingRecordIndex_EQ"`
	PowerDialerExecutingRecordIndex_GTE int32                  `bson:"powerDialerExecutingRecordIndex_GTE" json:"powerDialerExecutingRecordIndex_GTE"`
	PowerDialerExecutingRecordIndex_LTE int32                  `bson:"powerDialerExecutingRecordIndex_LTE" json:"powerDialerExecutingRecordIndex_LTE"`
	PowerDialerGroupStatus_EQ           PowerDialerGroupStatus `bson:"powerDialerGroupStatus_EQ" json:"powerDialerGroupStatus_EQ"`
	PowerDialers_ContactName_CON        string                 `bson:"powerDialers_ContactName_CON" json:"powerDialers_ContactName_CON"`
	PowerDialers_ContactName_EQ         string                 `bson:"powerDialers_ContactName_EQ" json:"powerDialers_ContactName_EQ"`
	PowerDialers_ContactName_REG        string                 `bson:"powerDialers_ContactName_REG" json:"powerDialers_ContactName_REG"`
	PowerDialers_CountryIsoCode_EQ      CountryIsoCode         `bson:"powerDialers_CountryIsoCode_EQ" json:"powerDialers_CountryIsoCode_EQ"`
	PowerDialers_ErrorMessage_CON       string                 `bson:"powerDialers_ErrorMessage_CON" json:"powerDialers_ErrorMessage_CON"`
	PowerDialers_ErrorMessage_EQ        string                 `bson:"powerDialers_ErrorMessage_EQ" json:"powerDialers_ErrorMessage_EQ"`
	PowerDialers_ErrorMessage_REG       string                 `bson:"powerDialers_ErrorMessage_REG" json:"powerDialers_ErrorMessage_REG"`
	PowerDialers_IdContact_CON          string                 `bson:"powerDialers_IdContact_CON" json:"powerDialers_IdContact_CON"`
	PowerDialers_IdContact_EQ           string                 `bson:"powerDialers_IdContact_EQ" json:"powerDialers_IdContact_EQ"`
	PowerDialers_IdContact_REG          string                 `bson:"powerDialers_IdContact_REG" json:"powerDialers_IdContact_REG"`
	PowerDialers_NumberOfAttempts_EQ    int32                  `bson:"powerDialers_NumberOfAttempts_EQ" json:"powerDialers_NumberOfAttempts_EQ"`
	PowerDialers_NumberOfAttempts_GTE   int32                  `bson:"powerDialers_NumberOfAttempts_GTE" json:"powerDialers_NumberOfAttempts_GTE"`
	PowerDialers_NumberOfAttempts_LTE   int32                  `bson:"powerDialers_NumberOfAttempts_LTE" json:"powerDialers_NumberOfAttempts_LTE"`
	PowerDialers_PhoneNumber_CON        string                 `bson:"powerDialers_PhoneNumber_CON" json:"powerDialers_PhoneNumber_CON"`
	PowerDialers_PhoneNumber_EQ         string                 `bson:"powerDialers_PhoneNumber_EQ" json:"powerDialers_PhoneNumber_EQ"`
	PowerDialers_PhoneNumber_REG        string                 `bson:"powerDialers_PhoneNumber_REG" json:"powerDialers_PhoneNumber_REG"`
	PowerDialers_PowerDialerStatus_EQ   PowerDialerStatus      `bson:"powerDialers_PowerDialerStatus_EQ" json:"powerDialers_PowerDialerStatus_EQ"`
	PowerDialers_PowerDialerType_EQ     PowerDialerType        `bson:"powerDialers_PowerDialerType_EQ" json:"powerDialers_PowerDialerType_EQ"`
}

// BUILDER from bson map:
func BuildPowerDialerGroupFilterRequest(m map[string]interface{}, x *PowerDialerGroupFilterRequest) {
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
	if val, ok := m["description_CON"]; ok && val != nil {
		x.Description_CON = val.(string)
	}
	if val, ok := m["description_EQ"]; ok && val != nil {
		x.Description_EQ = val.(string)
	}
	if val, ok := m["description_REG"]; ok && val != nil {
		x.Description_REG = val.(string)
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
	if val, ok := m["idCallerIdMask_CON"]; ok && val != nil {
		x.IdCallerIdMask_CON = val.(string)
	}
	if val, ok := m["idCallerIdMask_EQ"]; ok && val != nil {
		x.IdCallerIdMask_EQ = val.(string)
	}
	if val, ok := m["idCallerIdMask_REG"]; ok && val != nil {
		x.IdCallerIdMask_REG = val.(string)
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
	if val, ok := m["idExtension_CON"]; ok && val != nil {
		x.IdExtension_CON = val.(string)
	}
	if val, ok := m["idExtension_EQ"]; ok && val != nil {
		x.IdExtension_EQ = val.(string)
	}
	if val, ok := m["idExtension_REG"]; ok && val != nil {
		x.IdExtension_REG = val.(string)
	}
	if val, ok := m["idVoipNumberPhone_CON"]; ok && val != nil {
		x.IdVoipNumberPhone_CON = val.(string)
	}
	if val, ok := m["idVoipNumberPhone_EQ"]; ok && val != nil {
		x.IdVoipNumberPhone_EQ = val.(string)
	}
	if val, ok := m["idVoipNumberPhone_REG"]; ok && val != nil {
		x.IdVoipNumberPhone_REG = val.(string)
	}
	if val, ok := m["numberOfConcurrentCalls_EQ"]; ok && val != nil {
		x.NumberOfConcurrentCalls_EQ = val.(int32)
	}
	if val, ok := m["numberOfConcurrentCalls_GTE"]; ok && val != nil {
		x.NumberOfConcurrentCalls_GTE = val.(int32)
	}
	if val, ok := m["numberOfConcurrentCalls_LTE"]; ok && val != nil {
		x.NumberOfConcurrentCalls_LTE = val.(int32)
	}
	if val, ok := m["powerDialerExecutingRecordIndex_EQ"]; ok && val != nil {
		x.PowerDialerExecutingRecordIndex_EQ = val.(int32)
	}
	if val, ok := m["powerDialerExecutingRecordIndex_GTE"]; ok && val != nil {
		x.PowerDialerExecutingRecordIndex_GTE = val.(int32)
	}
	if val, ok := m["powerDialerExecutingRecordIndex_LTE"]; ok && val != nil {
		x.PowerDialerExecutingRecordIndex_LTE = val.(int32)
	}
	if val, ok := m["powerDialerGroupStatus_EQ"]; ok && val != nil {
		x.PowerDialerGroupStatus_EQ = PowerDialerGroupStatus("PowerDialerGroupStatus_EQ_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["powerDialers_ContactName_CON"]; ok && val != nil {
		x.PowerDialers_ContactName_CON = val.(string)
	}
	if val, ok := m["powerDialers_ContactName_EQ"]; ok && val != nil {
		x.PowerDialers_ContactName_EQ = val.(string)
	}
	if val, ok := m["powerDialers_ContactName_REG"]; ok && val != nil {
		x.PowerDialers_ContactName_REG = val.(string)
	}
	if val, ok := m["powerDialers_CountryIsoCode_EQ"]; ok && val != nil {
		x.PowerDialers_CountryIsoCode_EQ = CountryIsoCode("PowerDialers_CountryIsoCode_EQ_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["powerDialers_ErrorMessage_CON"]; ok && val != nil {
		x.PowerDialers_ErrorMessage_CON = val.(string)
	}
	if val, ok := m["powerDialers_ErrorMessage_EQ"]; ok && val != nil {
		x.PowerDialers_ErrorMessage_EQ = val.(string)
	}
	if val, ok := m["powerDialers_ErrorMessage_REG"]; ok && val != nil {
		x.PowerDialers_ErrorMessage_REG = val.(string)
	}
	if val, ok := m["powerDialers_IdContact_CON"]; ok && val != nil {
		x.PowerDialers_IdContact_CON = val.(string)
	}
	if val, ok := m["powerDialers_IdContact_EQ"]; ok && val != nil {
		x.PowerDialers_IdContact_EQ = val.(string)
	}
	if val, ok := m["powerDialers_IdContact_REG"]; ok && val != nil {
		x.PowerDialers_IdContact_REG = val.(string)
	}
	if val, ok := m["powerDialers_NumberOfAttempts_EQ"]; ok && val != nil {
		x.PowerDialers_NumberOfAttempts_EQ = val.(int32)
	}
	if val, ok := m["powerDialers_NumberOfAttempts_GTE"]; ok && val != nil {
		x.PowerDialers_NumberOfAttempts_GTE = val.(int32)
	}
	if val, ok := m["powerDialers_NumberOfAttempts_LTE"]; ok && val != nil {
		x.PowerDialers_NumberOfAttempts_LTE = val.(int32)
	}
	if val, ok := m["powerDialers_PhoneNumber_CON"]; ok && val != nil {
		x.PowerDialers_PhoneNumber_CON = val.(string)
	}
	if val, ok := m["powerDialers_PhoneNumber_EQ"]; ok && val != nil {
		x.PowerDialers_PhoneNumber_EQ = val.(string)
	}
	if val, ok := m["powerDialers_PhoneNumber_REG"]; ok && val != nil {
		x.PowerDialers_PhoneNumber_REG = val.(string)
	}
	if val, ok := m["powerDialers_PowerDialerStatus_EQ"]; ok && val != nil {
		x.PowerDialers_PowerDialerStatus_EQ = PowerDialerStatus("PowerDialers_PowerDialerStatus_EQ_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["powerDialers_PowerDialerType_EQ"]; ok && val != nil {
		x.PowerDialers_PowerDialerType_EQ = PowerDialerType("PowerDialers_PowerDialerType_EQ_" + val.(string))
	} // is NOT readonly obtained from map
}
