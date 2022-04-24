package models

import . "github.com/ublux/go-models/enums"
import "go.mongodb.org/mongo-driver/bson/primitive"

type PowerDialerGroup struct {
	Id                              string                 `bson:"_id" json:"id"`
	DateCreated                     string                 `bson:"dateCreated" json:"dateCreated"`
	DateDeleted                     string                 `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated                     string                 `bson:"dateUpdated" json:"dateUpdated"`
	Description                     string                 `bson:"description" json:"description"`
	ErrorMessage                    string                 `bson:"errorMessage" json:"errorMessage"`
	FriendlyName                    string                 `bson:"friendlyName" json:"friendlyName"`
	IdAccount                       string                 `bson:"idAccount" json:"idAccount"`
	IdCallerIdMask                  string                 `bson:"idCallerIdMask" json:"idCallerIdMask"`
	IdCallFlow                      string                 `bson:"idCallFlow" json:"idCallFlow"`
	IdExtension                     string                 `bson:"idExtension" json:"idExtension"`
	IdVoipNumberPhone               string                 `bson:"idVoipNumberPhone" json:"idVoipNumberPhone"`
	NumberOfConcurrentCalls         int32                  `bson:"numberOfConcurrentCalls" json:"numberOfConcurrentCalls"`
	PowerDialerExecutingRecordIndex int32                  `bson:"powerDialerExecutingRecordIndex" json:"powerDialerExecutingRecordIndex"`
	PowerDialerGroupStatus          PowerDialerGroupStatus `bson:"powerDialerGroupStatus" json:"powerDialerGroupStatus"`
	PowerDialers                    []PowerDialer          `bson:"powerDialers" json:"powerDialers"`
}

// Implementing interface IUbluxDocument
func (x PowerDialerGroup) GetDateCreated() string {
	return x.DateCreated
}
func (x PowerDialerGroup) GetDateDeleted() string {
	return x.DateDeleted
}
func (x PowerDialerGroup) GetDateUpdated() string {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x PowerDialerGroup) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x PowerDialerGroup) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildPowerDialerGroup(m map[string]interface{}, x *PowerDialerGroup) {
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(string)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(string)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(string)
	}
	if val, ok := m["description"]; ok && val != nil {
		x.Description = val.(string)
	}
	if val, ok := m["errorMessage"]; ok && val != nil {
		x.ErrorMessage = val.(string)
	}
	if val, ok := m["friendlyName"]; ok && val != nil {
		x.FriendlyName = val.(string)
	}
	if val, ok := m["idAccount"]; ok && val != nil {
		x.IdAccount = val.(string)
	}
	if val, ok := m["idCallerIdMask"]; ok && val != nil {
		x.IdCallerIdMask = val.(string)
	}
	if val, ok := m["idCallFlow"]; ok && val != nil {
		x.IdCallFlow = val.(string)
	}
	if val, ok := m["idExtension"]; ok && val != nil {
		x.IdExtension = val.(string)
	}
	if val, ok := m["idVoipNumberPhone"]; ok && val != nil {
		x.IdVoipNumberPhone = val.(string)
	}
	if val, ok := m["numberOfConcurrentCalls"]; ok && val != nil {
		x.NumberOfConcurrentCalls = val.(int32)
	}
	if val, ok := m["powerDialerExecutingRecordIndex"]; ok && val != nil {
		x.PowerDialerExecutingRecordIndex = val.(int32)
	}
	if val, ok := m["powerDialerGroupStatus"]; ok && val != nil {
		x.PowerDialerGroupStatus = PowerDialerGroupStatus("PowerDialerGroupStatus_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["powerDialers"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok {
			for _, val = range array {
				if val != nil {
					// determine the type to see what type of object we instantiate
					t := val.(map[string]interface{})["_t"].(string)
					switch t {
					case "PowerDialerAdvance":
						tmp := new(PowerDialerAdvance)
						BuildPowerDialerAdvance(val.(map[string]interface{}), tmp)
						x.PowerDialers = append(x.PowerDialers, *tmp)
					case "PowerDialerSimple":
						tmp := new(PowerDialerSimple)
						BuildPowerDialerSimple(val.(map[string]interface{}), tmp)
						x.PowerDialers = append(x.PowerDialers, *tmp)
					default:
						// error
					}
				}
			}
		}
	}
}
