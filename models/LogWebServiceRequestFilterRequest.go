package models

import "go.mongodb.org/mongo-driver/bson/primitive"
import . "github.com/ublux/go-models/enums"

type LogWebServiceRequestFilterRequest struct {
	Charge_EQ                       uint32                 `bson:"charge_EQ" json:"charge_EQ"`
	Charge_GTE                      uint32                 `bson:"charge_GTE" json:"charge_GTE"`
	Charge_LTE                      uint32                 `bson:"charge_LTE" json:"charge_LTE"`
	ConcurrentRequests_EQ           int32                  `bson:"concurrentRequests_EQ" json:"concurrentRequests_EQ"`
	ConcurrentRequests_GTE          int32                  `bson:"concurrentRequests_GTE" json:"concurrentRequests_GTE"`
	ConcurrentRequests_LTE          int32                  `bson:"concurrentRequests_LTE" json:"concurrentRequests_LTE"`
	DateCreated_EQ                  primitive.DateTime     `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE                 primitive.DateTime     `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE                 primitive.DateTime     `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ                  primitive.DateTime     `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE                 primitive.DateTime     `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE                 primitive.DateTime     `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	HttpMethod_CON                  string                 `bson:"httpMethod_CON" json:"httpMethod_CON"`
	HttpMethod_EQ                   string                 `bson:"httpMethod_EQ" json:"httpMethod_EQ"`
	HttpMethod_REG                  string                 `bson:"httpMethod_REG" json:"httpMethod_REG"`
	HttpResponseStatusCode_EQ       HttpResponseStatusCode `bson:"httpResponseStatusCode_EQ" json:"httpResponseStatusCode_EQ"`
	Id_CON                          string                 `bson:"id_CON" json:"id_CON"`
	Id_EQ                           string                 `bson:"id_EQ" json:"id_EQ"`
	Id_REG                          string                 `bson:"id_REG" json:"id_REG"`
	IdentityChargeSum_EQ            uint64                 `bson:"identityChargeSum_EQ" json:"identityChargeSum_EQ"`
	IdentityChargeSum_GTE           uint64                 `bson:"identityChargeSum_GTE" json:"identityChargeSum_GTE"`
	IdentityChargeSum_LTE           uint64                 `bson:"identityChargeSum_LTE" json:"identityChargeSum_LTE"`
	Ip_CON                          string                 `bson:"ip_CON" json:"ip_CON"`
	Ip_EQ                           string                 `bson:"ip_EQ" json:"ip_EQ"`
	Ip_REG                          string                 `bson:"ip_REG" json:"ip_REG"`
	Penalty_EQ                      uint32                 `bson:"penalty_EQ" json:"penalty_EQ"`
	Penalty_GTE                     uint32                 `bson:"penalty_GTE" json:"penalty_GTE"`
	Penalty_LTE                     uint32                 `bson:"penalty_LTE" json:"penalty_LTE"`
	QueryString_CON                 string                 `bson:"queryString_CON" json:"queryString_CON"`
	QueryString_EQ                  string                 `bson:"queryString_EQ" json:"queryString_EQ"`
	QueryString_REG                 string                 `bson:"queryString_REG" json:"queryString_REG"`
	RequestBody_CON                 string                 `bson:"requestBody_CON" json:"requestBody_CON"`
	RequestBody_EQ                  string                 `bson:"requestBody_EQ" json:"requestBody_EQ"`
	RequestBody_REG                 string                 `bson:"requestBody_REG" json:"requestBody_REG"`
	RequestUrl_CON                  string                 `bson:"requestUrl_CON" json:"requestUrl_CON"`
	RequestUrl_EQ                   string                 `bson:"requestUrl_EQ" json:"requestUrl_EQ"`
	RequestUrl_REG                  string                 `bson:"requestUrl_REG" json:"requestUrl_REG"`
	ResponseBody_CON                string                 `bson:"responseBody_CON" json:"responseBody_CON"`
	ResponseBody_EQ                 string                 `bson:"responseBody_EQ" json:"responseBody_EQ"`
	ResponseBody_REG                string                 `bson:"responseBody_REG" json:"responseBody_REG"`
	ResponseTime_EQ                 int32                  `bson:"responseTime_EQ" json:"responseTime_EQ"`
	ResponseTime_GTE                int32                  `bson:"responseTime_GTE" json:"responseTime_GTE"`
	ResponseTime_LTE                int32                  `bson:"responseTime_LTE" json:"responseTime_LTE"`
	SaveQueryString_EQ              bool                   `bson:"saveQueryString_EQ" json:"saveQueryString_EQ"`
	UbluxSession_ExpirationDate_EQ  primitive.DateTime     `bson:"ubluxSession_ExpirationDate_EQ" json:"ubluxSession_ExpirationDate_EQ"`
	UbluxSession_ExpirationDate_GTE primitive.DateTime     `bson:"ubluxSession_ExpirationDate_GTE" json:"ubluxSession_ExpirationDate_GTE"`
	UbluxSession_ExpirationDate_LTE primitive.DateTime     `bson:"ubluxSession_ExpirationDate_LTE" json:"ubluxSession_ExpirationDate_LTE"`
	UbluxSession_IdAccount_CON      string                 `bson:"ubluxSession_IdAccount_CON" json:"ubluxSession_IdAccount_CON"`
	UbluxSession_IdAccount_EQ       string                 `bson:"ubluxSession_IdAccount_EQ" json:"ubluxSession_IdAccount_EQ"`
	UbluxSession_IdAccount_REG      string                 `bson:"ubluxSession_IdAccount_REG" json:"ubluxSession_IdAccount_REG"`
	UbluxSession_IdentityType_EQ    IdentityType           `bson:"ubluxSession_IdentityType_EQ" json:"ubluxSession_IdentityType_EQ"`
	UbluxSession_IdIdentity_CON     string                 `bson:"ubluxSession_IdIdentity_CON" json:"ubluxSession_IdIdentity_CON"`
	UbluxSession_IdIdentity_EQ      string                 `bson:"ubluxSession_IdIdentity_EQ" json:"ubluxSession_IdIdentity_EQ"`
	UbluxSession_IdIdentity_REG     string                 `bson:"ubluxSession_IdIdentity_REG" json:"ubluxSession_IdIdentity_REG"`
	UbluxSession_UbluxRoles_CON     UbluxRole              `bson:"ubluxSession_UbluxRoles_CON" json:"ubluxSession_UbluxRoles_CON"`
	UserAgent_CON                   string                 `bson:"userAgent_CON" json:"userAgent_CON"`
	UserAgent_EQ                    string                 `bson:"userAgent_EQ" json:"userAgent_EQ"`
	UserAgent_REG                   string                 `bson:"userAgent_REG" json:"userAgent_REG"`
}

// BUILDER from bson map:
func BuildLogWebServiceRequestFilterRequest(m map[string]interface{}, x *LogWebServiceRequestFilterRequest) {
	if val, ok := m["charge_EQ"]; ok && val != nil {
		x.Charge_EQ = val.(uint32)
	}
	if val, ok := m["charge_GTE"]; ok && val != nil {
		x.Charge_GTE = val.(uint32)
	}
	if val, ok := m["charge_LTE"]; ok && val != nil {
		x.Charge_LTE = val.(uint32)
	}
	if val, ok := m["concurrentRequests_EQ"]; ok && val != nil {
		x.ConcurrentRequests_EQ = val.(int32)
	}
	if val, ok := m["concurrentRequests_GTE"]; ok && val != nil {
		x.ConcurrentRequests_GTE = val.(int32)
	}
	if val, ok := m["concurrentRequests_LTE"]; ok && val != nil {
		x.ConcurrentRequests_LTE = val.(int32)
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
	if val, ok := m["httpMethod_CON"]; ok && val != nil {
		x.HttpMethod_CON = val.(string)
	}
	if val, ok := m["httpMethod_EQ"]; ok && val != nil {
		x.HttpMethod_EQ = val.(string)
	}
	if val, ok := m["httpMethod_REG"]; ok && val != nil {
		x.HttpMethod_REG = val.(string)
	}
	if val, ok := m["httpResponseStatusCode_EQ"]; ok && val != nil {
		x.HttpResponseStatusCode_EQ = HttpResponseStatusCode("HttpResponseStatusCode_EQ_" + val.(string))
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
	if val, ok := m["identityChargeSum_EQ"]; ok && val != nil {
		x.IdentityChargeSum_EQ = val.(uint64)
	}
	if val, ok := m["identityChargeSum_GTE"]; ok && val != nil {
		x.IdentityChargeSum_GTE = val.(uint64)
	}
	if val, ok := m["identityChargeSum_LTE"]; ok && val != nil {
		x.IdentityChargeSum_LTE = val.(uint64)
	}
	if val, ok := m["ip_CON"]; ok && val != nil {
		x.Ip_CON = val.(string)
	}
	if val, ok := m["ip_EQ"]; ok && val != nil {
		x.Ip_EQ = val.(string)
	}
	if val, ok := m["ip_REG"]; ok && val != nil {
		x.Ip_REG = val.(string)
	}
	if val, ok := m["penalty_EQ"]; ok && val != nil {
		x.Penalty_EQ = val.(uint32)
	}
	if val, ok := m["penalty_GTE"]; ok && val != nil {
		x.Penalty_GTE = val.(uint32)
	}
	if val, ok := m["penalty_LTE"]; ok && val != nil {
		x.Penalty_LTE = val.(uint32)
	}
	if val, ok := m["queryString_CON"]; ok && val != nil {
		x.QueryString_CON = val.(string)
	}
	if val, ok := m["queryString_EQ"]; ok && val != nil {
		x.QueryString_EQ = val.(string)
	}
	if val, ok := m["queryString_REG"]; ok && val != nil {
		x.QueryString_REG = val.(string)
	}
	if val, ok := m["requestBody_CON"]; ok && val != nil {
		x.RequestBody_CON = val.(string)
	}
	if val, ok := m["requestBody_EQ"]; ok && val != nil {
		x.RequestBody_EQ = val.(string)
	}
	if val, ok := m["requestBody_REG"]; ok && val != nil {
		x.RequestBody_REG = val.(string)
	}
	if val, ok := m["requestUrl_CON"]; ok && val != nil {
		x.RequestUrl_CON = val.(string)
	}
	if val, ok := m["requestUrl_EQ"]; ok && val != nil {
		x.RequestUrl_EQ = val.(string)
	}
	if val, ok := m["requestUrl_REG"]; ok && val != nil {
		x.RequestUrl_REG = val.(string)
	}
	if val, ok := m["responseBody_CON"]; ok && val != nil {
		x.ResponseBody_CON = val.(string)
	}
	if val, ok := m["responseBody_EQ"]; ok && val != nil {
		x.ResponseBody_EQ = val.(string)
	}
	if val, ok := m["responseBody_REG"]; ok && val != nil {
		x.ResponseBody_REG = val.(string)
	}
	if val, ok := m["responseTime_EQ"]; ok && val != nil {
		x.ResponseTime_EQ = val.(int32)
	}
	if val, ok := m["responseTime_GTE"]; ok && val != nil {
		x.ResponseTime_GTE = val.(int32)
	}
	if val, ok := m["responseTime_LTE"]; ok && val != nil {
		x.ResponseTime_LTE = val.(int32)
	}
	if val, ok := m["saveQueryString_EQ"]; ok && val != nil {
		x.SaveQueryString_EQ = val.(bool)
	}
	if val, ok := m["ubluxSession_ExpirationDate_EQ"]; ok && val != nil {
		x.UbluxSession_ExpirationDate_EQ = val.(primitive.DateTime)
	}
	if val, ok := m["ubluxSession_ExpirationDate_GTE"]; ok && val != nil {
		x.UbluxSession_ExpirationDate_GTE = val.(primitive.DateTime)
	}
	if val, ok := m["ubluxSession_ExpirationDate_LTE"]; ok && val != nil {
		x.UbluxSession_ExpirationDate_LTE = val.(primitive.DateTime)
	}
	if val, ok := m["ubluxSession_IdAccount_CON"]; ok && val != nil {
		x.UbluxSession_IdAccount_CON = val.(string)
	}
	if val, ok := m["ubluxSession_IdAccount_EQ"]; ok && val != nil {
		x.UbluxSession_IdAccount_EQ = val.(string)
	}
	if val, ok := m["ubluxSession_IdAccount_REG"]; ok && val != nil {
		x.UbluxSession_IdAccount_REG = val.(string)
	}
	if val, ok := m["ubluxSession_IdentityType_EQ"]; ok && val != nil {
		x.UbluxSession_IdentityType_EQ = IdentityType("UbluxSession_IdentityType_EQ_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["ubluxSession_IdIdentity_CON"]; ok && val != nil {
		x.UbluxSession_IdIdentity_CON = val.(string)
	}
	if val, ok := m["ubluxSession_IdIdentity_EQ"]; ok && val != nil {
		x.UbluxSession_IdIdentity_EQ = val.(string)
	}
	if val, ok := m["ubluxSession_IdIdentity_REG"]; ok && val != nil {
		x.UbluxSession_IdIdentity_REG = val.(string)
	}
	if val, ok := m["ubluxSession_UbluxRoles_CON"]; ok && val != nil {
		x.UbluxSession_UbluxRoles_CON = UbluxRole("UbluxSession_UbluxRoles_CON_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["userAgent_CON"]; ok && val != nil {
		x.UserAgent_CON = val.(string)
	}
	if val, ok := m["userAgent_EQ"]; ok && val != nil {
		x.UserAgent_EQ = val.(string)
	}
	if val, ok := m["userAgent_REG"]; ok && val != nil {
		x.UserAgent_REG = val.(string)
	}
}
