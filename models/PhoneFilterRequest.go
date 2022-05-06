package models

import "time"

type PhoneFilterRequest struct {
	DateCreated_EQ                               time.Time `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE                              time.Time `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE                              time.Time `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ                               time.Time `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE                              time.Time `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE                              time.Time `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	FriendlyName_CON                             string    `bson:"friendlyName_CON" json:"friendlyName_CON"`
	FriendlyName_EQ                              string    `bson:"friendlyName_EQ" json:"friendlyName_EQ"`
	FriendlyName_REG                             string    `bson:"friendlyName_REG" json:"friendlyName_REG"`
	Id_CON                                       string    `bson:"id_CON" json:"id_CON"`
	Id_EQ                                        string    `bson:"id_EQ" json:"id_EQ"`
	Id_REG                                       string    `bson:"id_REG" json:"id_REG"`
	IdCloudServicePbx_CON                        string    `bson:"idCloudServicePbx_CON" json:"idCloudServicePbx_CON"`
	IdCloudServicePbx_EQ                         string    `bson:"idCloudServicePbx_EQ" json:"idCloudServicePbx_EQ"`
	IdCloudServicePbx_REG                        string    `bson:"idCloudServicePbx_REG" json:"idCloudServicePbx_REG"`
	IdIdentityWebApp_CON                         string    `bson:"idIdentityWebApp_CON" json:"idIdentityWebApp_CON"`
	IdIdentityWebApp_EQ                          string    `bson:"idIdentityWebApp_EQ" json:"idIdentityWebApp_EQ"`
	IdIdentityWebApp_REG                         string    `bson:"idIdentityWebApp_REG" json:"idIdentityWebApp_REG"`
	IdPhoneConfiguration_CON                     string    `bson:"idPhoneConfiguration_CON" json:"idPhoneConfiguration_CON"`
	IdPhoneConfiguration_EQ                      string    `bson:"idPhoneConfiguration_EQ" json:"idPhoneConfiguration_EQ"`
	IdPhoneConfiguration_REG                     string    `bson:"idPhoneConfiguration_REG" json:"idPhoneConfiguration_REG"`
	Lines_CallerIdNumber_CON                     string    `bson:"lines_CallerIdNumber_CON" json:"lines_CallerIdNumber_CON"`
	Lines_CallerIdNumber_EQ                      string    `bson:"lines_CallerIdNumber_EQ" json:"lines_CallerIdNumber_EQ"`
	Lines_CallerIdNumber_REG                     string    `bson:"lines_CallerIdNumber_REG" json:"lines_CallerIdNumber_REG"`
	Lines_FriendlyName_CON                       string    `bson:"lines_FriendlyName_CON" json:"lines_FriendlyName_CON"`
	Lines_FriendlyName_EQ                        string    `bson:"lines_FriendlyName_EQ" json:"lines_FriendlyName_EQ"`
	Lines_FriendlyName_REG                       string    `bson:"lines_FriendlyName_REG" json:"lines_FriendlyName_REG"`
	Lines_Id_CON                                 string    `bson:"lines_Id_CON" json:"lines_Id_CON"`
	Lines_Id_EQ                                  string    `bson:"lines_Id_EQ" json:"lines_Id_EQ"`
	Lines_Id_REG                                 string    `bson:"lines_Id_REG" json:"lines_Id_REG"`
	Lines_LineConnectionStatus_DateConnected_EQ  time.Time `bson:"lines_LineConnectionStatus_DateConnected_EQ" json:"lines_LineConnectionStatus_DateConnected_EQ"`
	Lines_LineConnectionStatus_DateConnected_GTE time.Time `bson:"lines_LineConnectionStatus_DateConnected_GTE" json:"lines_LineConnectionStatus_DateConnected_GTE"`
	Lines_LineConnectionStatus_DateConnected_LTE time.Time `bson:"lines_LineConnectionStatus_DateConnected_LTE" json:"lines_LineConnectionStatus_DateConnected_LTE"`
	Lines_LineConnectionStatus_Id_CON            string    `bson:"lines_LineConnectionStatus_Id_CON" json:"lines_LineConnectionStatus_Id_CON"`
	Lines_LineConnectionStatus_Id_EQ             string    `bson:"lines_LineConnectionStatus_Id_EQ" json:"lines_LineConnectionStatus_Id_EQ"`
	Lines_LineConnectionStatus_Id_REG            string    `bson:"lines_LineConnectionStatus_Id_REG" json:"lines_LineConnectionStatus_Id_REG"`
	Lines_LineConnectionStatus_IpLAN_CON         string    `bson:"lines_LineConnectionStatus_IpLAN_CON" json:"lines_LineConnectionStatus_IpLAN_CON"`
	Lines_LineConnectionStatus_IpLAN_EQ          string    `bson:"lines_LineConnectionStatus_IpLAN_EQ" json:"lines_LineConnectionStatus_IpLAN_EQ"`
	Lines_LineConnectionStatus_IpLAN_REG         string    `bson:"lines_LineConnectionStatus_IpLAN_REG" json:"lines_LineConnectionStatus_IpLAN_REG"`
	Lines_LineConnectionStatus_IpWAN_CON         string    `bson:"lines_LineConnectionStatus_IpWAN_CON" json:"lines_LineConnectionStatus_IpWAN_CON"`
	Lines_LineConnectionStatus_IpWAN_EQ          string    `bson:"lines_LineConnectionStatus_IpWAN_EQ" json:"lines_LineConnectionStatus_IpWAN_EQ"`
	Lines_LineConnectionStatus_IpWAN_REG         string    `bson:"lines_LineConnectionStatus_IpWAN_REG" json:"lines_LineConnectionStatus_IpWAN_REG"`
	Lines_LineConnectionStatus_IsConnected_EQ    bool      `bson:"lines_LineConnectionStatus_IsConnected_EQ" json:"lines_LineConnectionStatus_IsConnected_EQ"`
	Lines_LineConnectionStatus_PortLAN_EQ        int32     `bson:"lines_LineConnectionStatus_PortLAN_EQ" json:"lines_LineConnectionStatus_PortLAN_EQ"`
	Lines_LineConnectionStatus_PortLAN_GTE       int32     `bson:"lines_LineConnectionStatus_PortLAN_GTE" json:"lines_LineConnectionStatus_PortLAN_GTE"`
	Lines_LineConnectionStatus_PortLAN_LTE       int32     `bson:"lines_LineConnectionStatus_PortLAN_LTE" json:"lines_LineConnectionStatus_PortLAN_LTE"`
	Lines_LineConnectionStatus_PortWAN_EQ        int32     `bson:"lines_LineConnectionStatus_PortWAN_EQ" json:"lines_LineConnectionStatus_PortWAN_EQ"`
	Lines_LineConnectionStatus_PortWAN_GTE       int32     `bson:"lines_LineConnectionStatus_PortWAN_GTE" json:"lines_LineConnectionStatus_PortWAN_GTE"`
	Lines_LineConnectionStatus_PortWAN_LTE       int32     `bson:"lines_LineConnectionStatus_PortWAN_LTE" json:"lines_LineConnectionStatus_PortWAN_LTE"`
	Lines_LineConnectionStatus_UserAgent_CON     string    `bson:"lines_LineConnectionStatus_UserAgent_CON" json:"lines_LineConnectionStatus_UserAgent_CON"`
	Lines_LineConnectionStatus_UserAgent_EQ      string    `bson:"lines_LineConnectionStatus_UserAgent_EQ" json:"lines_LineConnectionStatus_UserAgent_EQ"`
	Lines_LineConnectionStatus_UserAgent_REG     string    `bson:"lines_LineConnectionStatus_UserAgent_REG" json:"lines_LineConnectionStatus_UserAgent_REG"`
	Lines_RecordExternalCalls_EQ                 bool      `bson:"lines_RecordExternalCalls_EQ" json:"lines_RecordExternalCalls_EQ"`
	Lines_RecordInternalCalls_EQ                 bool      `bson:"lines_RecordInternalCalls_EQ" json:"lines_RecordInternalCalls_EQ"`
}

// BUILDER from bson map:
func BuildPhoneFilterRequest(m map[string]interface{}, x *PhoneFilterRequest) {
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
	if val, ok := m["idCloudServicePbx_CON"]; ok && val != nil {
		x.IdCloudServicePbx_CON = val.(string)
	}
	if val, ok := m["idCloudServicePbx_EQ"]; ok && val != nil {
		x.IdCloudServicePbx_EQ = val.(string)
	}
	if val, ok := m["idCloudServicePbx_REG"]; ok && val != nil {
		x.IdCloudServicePbx_REG = val.(string)
	}
	if val, ok := m["idIdentityWebApp_CON"]; ok && val != nil {
		x.IdIdentityWebApp_CON = val.(string)
	}
	if val, ok := m["idIdentityWebApp_EQ"]; ok && val != nil {
		x.IdIdentityWebApp_EQ = val.(string)
	}
	if val, ok := m["idIdentityWebApp_REG"]; ok && val != nil {
		x.IdIdentityWebApp_REG = val.(string)
	}
	if val, ok := m["idPhoneConfiguration_CON"]; ok && val != nil {
		x.IdPhoneConfiguration_CON = val.(string)
	}
	if val, ok := m["idPhoneConfiguration_EQ"]; ok && val != nil {
		x.IdPhoneConfiguration_EQ = val.(string)
	}
	if val, ok := m["idPhoneConfiguration_REG"]; ok && val != nil {
		x.IdPhoneConfiguration_REG = val.(string)
	}
	if val, ok := m["lines_CallerIdNumber_CON"]; ok && val != nil {
		x.Lines_CallerIdNumber_CON = val.(string)
	}
	if val, ok := m["lines_CallerIdNumber_EQ"]; ok && val != nil {
		x.Lines_CallerIdNumber_EQ = val.(string)
	}
	if val, ok := m["lines_CallerIdNumber_REG"]; ok && val != nil {
		x.Lines_CallerIdNumber_REG = val.(string)
	}
	if val, ok := m["lines_FriendlyName_CON"]; ok && val != nil {
		x.Lines_FriendlyName_CON = val.(string)
	}
	if val, ok := m["lines_FriendlyName_EQ"]; ok && val != nil {
		x.Lines_FriendlyName_EQ = val.(string)
	}
	if val, ok := m["lines_FriendlyName_REG"]; ok && val != nil {
		x.Lines_FriendlyName_REG = val.(string)
	}
	if val, ok := m["lines_Id_CON"]; ok && val != nil {
		x.Lines_Id_CON = val.(string)
	}
	if val, ok := m["lines_Id_EQ"]; ok && val != nil {
		x.Lines_Id_EQ = val.(string)
	}
	if val, ok := m["lines_Id_REG"]; ok && val != nil {
		x.Lines_Id_REG = val.(string)
	}
	if val, ok := m["lines_LineConnectionStatus_DateConnected_EQ"]; ok && val != nil {
		x.Lines_LineConnectionStatus_DateConnected_EQ = val.(time.Time)
	}
	if val, ok := m["lines_LineConnectionStatus_DateConnected_GTE"]; ok && val != nil {
		x.Lines_LineConnectionStatus_DateConnected_GTE = val.(time.Time)
	}
	if val, ok := m["lines_LineConnectionStatus_DateConnected_LTE"]; ok && val != nil {
		x.Lines_LineConnectionStatus_DateConnected_LTE = val.(time.Time)
	}
	if val, ok := m["lines_LineConnectionStatus_Id_CON"]; ok && val != nil {
		x.Lines_LineConnectionStatus_Id_CON = val.(string)
	}
	if val, ok := m["lines_LineConnectionStatus_Id_EQ"]; ok && val != nil {
		x.Lines_LineConnectionStatus_Id_EQ = val.(string)
	}
	if val, ok := m["lines_LineConnectionStatus_Id_REG"]; ok && val != nil {
		x.Lines_LineConnectionStatus_Id_REG = val.(string)
	}
	if val, ok := m["lines_LineConnectionStatus_IpLAN_CON"]; ok && val != nil {
		x.Lines_LineConnectionStatus_IpLAN_CON = val.(string)
	}
	if val, ok := m["lines_LineConnectionStatus_IpLAN_EQ"]; ok && val != nil {
		x.Lines_LineConnectionStatus_IpLAN_EQ = val.(string)
	}
	if val, ok := m["lines_LineConnectionStatus_IpLAN_REG"]; ok && val != nil {
		x.Lines_LineConnectionStatus_IpLAN_REG = val.(string)
	}
	if val, ok := m["lines_LineConnectionStatus_IpWAN_CON"]; ok && val != nil {
		x.Lines_LineConnectionStatus_IpWAN_CON = val.(string)
	}
	if val, ok := m["lines_LineConnectionStatus_IpWAN_EQ"]; ok && val != nil {
		x.Lines_LineConnectionStatus_IpWAN_EQ = val.(string)
	}
	if val, ok := m["lines_LineConnectionStatus_IpWAN_REG"]; ok && val != nil {
		x.Lines_LineConnectionStatus_IpWAN_REG = val.(string)
	}
	if val, ok := m["lines_LineConnectionStatus_IsConnected_EQ"]; ok && val != nil {
		x.Lines_LineConnectionStatus_IsConnected_EQ = val.(bool)
	}
	if val, ok := m["lines_LineConnectionStatus_PortLAN_EQ"]; ok && val != nil {
		x.Lines_LineConnectionStatus_PortLAN_EQ = val.(int32)
	}
	if val, ok := m["lines_LineConnectionStatus_PortLAN_GTE"]; ok && val != nil {
		x.Lines_LineConnectionStatus_PortLAN_GTE = val.(int32)
	}
	if val, ok := m["lines_LineConnectionStatus_PortLAN_LTE"]; ok && val != nil {
		x.Lines_LineConnectionStatus_PortLAN_LTE = val.(int32)
	}
	if val, ok := m["lines_LineConnectionStatus_PortWAN_EQ"]; ok && val != nil {
		x.Lines_LineConnectionStatus_PortWAN_EQ = val.(int32)
	}
	if val, ok := m["lines_LineConnectionStatus_PortWAN_GTE"]; ok && val != nil {
		x.Lines_LineConnectionStatus_PortWAN_GTE = val.(int32)
	}
	if val, ok := m["lines_LineConnectionStatus_PortWAN_LTE"]; ok && val != nil {
		x.Lines_LineConnectionStatus_PortWAN_LTE = val.(int32)
	}
	if val, ok := m["lines_LineConnectionStatus_UserAgent_CON"]; ok && val != nil {
		x.Lines_LineConnectionStatus_UserAgent_CON = val.(string)
	}
	if val, ok := m["lines_LineConnectionStatus_UserAgent_EQ"]; ok && val != nil {
		x.Lines_LineConnectionStatus_UserAgent_EQ = val.(string)
	}
	if val, ok := m["lines_LineConnectionStatus_UserAgent_REG"]; ok && val != nil {
		x.Lines_LineConnectionStatus_UserAgent_REG = val.(string)
	}
	if val, ok := m["lines_RecordExternalCalls_EQ"]; ok && val != nil {
		x.Lines_RecordExternalCalls_EQ = val.(bool)
	}
	if val, ok := m["lines_RecordInternalCalls_EQ"]; ok && val != nil {
		x.Lines_RecordInternalCalls_EQ = val.(bool)
	}
}
