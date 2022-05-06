package models

import "time"
import . "github.com/ublux/go-models/enums"
import "go.mongodb.org/mongo-driver/bson/primitive"

type TrunkOriginationRegister struct {
	DateCreated               time.Time            `bson:"dateCreated" json:"dateCreated"`
	DateDeleted               time.Time            `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated               time.Time            `bson:"dateUpdated" json:"dateUpdated"`
	FriendlyName              string               `bson:"friendlyName" json:"friendlyName"`
	Id                        string               `bson:"id" json:"id"`
	IdCloudServicePbx         string               `bson:"idCloudServicePbx" json:"idCloudServicePbx"`
	IdCloudServicePbxFailover string               `bson:"idCloudServicePbxFailover" json:"idCloudServicePbxFailover"`
	IdsVoipNumbers            []string             `bson:"idsVoipNumbers" json:"idsVoipNumbers"`
	IdVoipProvider            string               `bson:"idVoipProvider" json:"idVoipProvider"`
	Reg_host                  string               `bson:"reg_host" json:"reg_host"`
	Reg_password              string               `bson:"reg_password" json:"reg_password"`
	Reg_port                  int32                `bson:"reg_port" json:"reg_port"`
	Reg_username              string               `bson:"reg_username" json:"reg_username"`
	SID                       string               `bson:"sID" json:"sID"`
	TrunkOriginationType      TrunkOriginationType `bson:"trunkOriginationType" json:"trunkOriginationType"`
}

// Implementing interface IUbluxDocument
func (x TrunkOriginationRegister) GetDateDeleted() time.Time {
	return x.DateDeleted
}
func (x TrunkOriginationRegister) GetDateCreated() time.Time {
	return x.DateCreated
}
func (x TrunkOriginationRegister) GetDateUpdated() time.Time {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x TrunkOriginationRegister) GetId() string {
	return x.Id
}

// Implementing interface TrunkOrigination
func (x TrunkOriginationRegister) GetIdVoipProvider() string {
	return x.IdVoipProvider
}
func (x TrunkOriginationRegister) GetIdCloudServicePbx() string {
	return x.IdCloudServicePbx
}
func (x TrunkOriginationRegister) GetIdCloudServicePbxFailover() string {
	return x.IdCloudServicePbxFailover
}
func (x TrunkOriginationRegister) GetIdsVoipNumbers() []string {
	return x.IdsVoipNumbers
}
func (x TrunkOriginationRegister) GetTrunkOriginationType() TrunkOriginationType {
	return x.TrunkOriginationType
}
func (x TrunkOriginationRegister) GetSID() string {
	return x.SID
}
func (x TrunkOriginationRegister) GetFriendlyName() string {
	return x.FriendlyName
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildTrunkOriginationRegister(m map[string]interface{}, x *TrunkOriginationRegister) {
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(time.Time)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(time.Time)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(time.Time)
	}
	if val, ok := m["friendlyName"]; ok && val != nil {
		x.FriendlyName = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["idCloudServicePbx"]; ok && val != nil {
		x.IdCloudServicePbx = val.(string)
	}
	if val, ok := m["idCloudServicePbxFailover"]; ok && val != nil {
		x.IdCloudServicePbxFailover = val.(string)
	}
	if val, ok := m["idsVoipNumbers"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				if val != nil {
					x.IdsVoipNumbers = append(x.IdsVoipNumbers, val.(string))
				}
			}
		}
	}
	if val, ok := m["idVoipProvider"]; ok && val != nil {
		x.IdVoipProvider = val.(string)
	}
	if val, ok := m["reg_host"]; ok && val != nil {
		x.Reg_host = val.(string)
	}
	if val, ok := m["reg_password"]; ok && val != nil {
		x.Reg_password = val.(string)
	}
	if val, ok := m["reg_port"]; ok && val != nil {
		x.Reg_port = val.(int32)
	}
	if val, ok := m["reg_username"]; ok && val != nil {
		x.Reg_username = val.(string)
	}
	if val, ok := m["sID"]; ok && val != nil {
		x.SID = val.(string)
	}
	x.TrunkOriginationType = TrunkOriginationType_Register // readonly property
}
