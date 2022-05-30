package models

import . "github.com/ublux/go-models/enums"

type RunningApplicationInstance struct {
	CloudServiceType            CloudServiceType `bson:"cloudServiceType" json:"cloudServiceType"`
	Id                          string           `bson:"id" json:"id"`
	NumberOfOperationsExecuting int64            `bson:"numberOfOperationsExecuting" json:"numberOfOperationsExecuting"`
}

// BUILDER from bson map:
func BuildRunningApplicationInstance(m map[string]interface{}, x *RunningApplicationInstance) {
	if val, ok := m["cloudServiceType"]; ok && val != nil {
		x.CloudServiceType = CloudServiceType("CloudServiceType_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["numberOfOperationsExecuting"]; ok && val != nil {
		x.NumberOfOperationsExecuting = val.(int64)
	}
}
