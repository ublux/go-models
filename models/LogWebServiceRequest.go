package models

import "go.mongodb.org/mongo-driver/bson/primitive"
import . "github.com/ublux/go-models/enums"

type LogWebServiceRequest struct {
	Charge                 uint32                 `bson:"charge" json:"charge"`
	ConcurrentRequests     int32                  `bson:"concurrentRequests" json:"concurrentRequests"`
	DateCreated            primitive.DateTime     `bson:"dateCreated" json:"dateCreated"`
	DateDeleted            primitive.DateTime     `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated            primitive.DateTime     `bson:"dateUpdated" json:"dateUpdated"`
	HttpMethod             string                 `bson:"httpMethod" json:"httpMethod"`
	HttpResponseStatusCode HttpResponseStatusCode `bson:"httpResponseStatusCode" json:"httpResponseStatusCode"`
	Id                     string                 `bson:"_id" json:"id"`
	IdentityChargeSum      uint64                 `bson:"identityChargeSum" json:"identityChargeSum"`
	Ip                     string                 `bson:"ip" json:"ip"`
	Penalty                uint32                 `bson:"penalty" json:"penalty"`
	QueryString            string                 `bson:"queryString" json:"queryString"`
	RequestBody            string                 `bson:"requestBody" json:"requestBody"`
	RequestUrl             string                 `bson:"requestUrl" json:"requestUrl"`
	ResponseBody           string                 `bson:"responseBody" json:"responseBody"`
	ResponseTime           int32                  `bson:"responseTime" json:"responseTime"`
	SaveQueryString        bool                   `bson:"saveQueryString" json:"saveQueryString"`
	SaveRequestBody        bool                   `bson:"saveRequestBody" json:"saveRequestBody"`
	SaveResponseBody       bool                   `bson:"saveResponseBody" json:"saveResponseBody"`
	UbluxSession           UbluxSession           `bson:"ubluxSession" json:"ubluxSession"`
	UserAgent              string                 `bson:"userAgent" json:"userAgent"`
}

// Implementing interface IUbluxDocument
func (x LogWebServiceRequest) GetDateDeleted() primitive.DateTime {
	return x.DateDeleted
}
func (x LogWebServiceRequest) GetDateCreated() primitive.DateTime {
	return x.DateCreated
}
func (x LogWebServiceRequest) GetDateUpdated() primitive.DateTime {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x LogWebServiceRequest) GetId() string {
	return x.Id
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildLogWebServiceRequest(m map[string]interface{}, x *LogWebServiceRequest) {
	if val, ok := m["charge"]; ok && val != nil {
		x.Charge = val.(uint32)
	}
	if val, ok := m["concurrentRequests"]; ok && val != nil {
		x.ConcurrentRequests = val.(int32)
	}
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(primitive.DateTime)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(primitive.DateTime)
	}
	if val, ok := m["httpMethod"]; ok && val != nil {
		x.HttpMethod = val.(string)
	}
	if val, ok := m["httpResponseStatusCode"]; ok && val != nil {
		x.HttpResponseStatusCode = HttpResponseStatusCode("HttpResponseStatusCode_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["identityChargeSum"]; ok && val != nil {
		x.IdentityChargeSum = val.(uint64)
	}
	if val, ok := m["ip"]; ok && val != nil {
		x.Ip = val.(string)
	}
	if val, ok := m["penalty"]; ok && val != nil {
		x.Penalty = val.(uint32)
	}
	if val, ok := m["queryString"]; ok && val != nil {
		x.QueryString = val.(string)
	}
	if val, ok := m["requestBody"]; ok && val != nil {
		x.RequestBody = val.(string)
	}
	if val, ok := m["requestUrl"]; ok && val != nil {
		x.RequestUrl = val.(string)
	}
	if val, ok := m["responseBody"]; ok && val != nil {
		x.ResponseBody = val.(string)
	}
	if val, ok := m["responseTime"]; ok && val != nil {
		x.ResponseTime = val.(int32)
	}
	if val, ok := m["saveQueryString"]; ok && val != nil {
		x.SaveQueryString = val.(bool)
	}
	if val, ok := m["saveRequestBody"]; ok && val != nil {
		x.SaveRequestBody = val.(bool)
	}
	if val, ok := m["saveResponseBody"]; ok && val != nil {
		x.SaveResponseBody = val.(bool)
	}
	if val, ok := m["ubluxSession"]; ok && val != nil {
		BuildUbluxSession(val.(map[string]interface{}), &x.UbluxSession)
	}
	if val, ok := m["userAgent"]; ok && val != nil {
		x.UserAgent = val.(string)
	}
}
