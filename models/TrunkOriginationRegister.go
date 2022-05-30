package models

import "go.mongodb.org/mongo-driver/bson/primitive"
import . "github.com/ublux/go-models/enums"

type TrunkOriginationRegister struct {
	DateCreated               primitive.DateTime   `bson:"dateCreated" json:"dateCreated"`
	DateDeleted               primitive.DateTime   `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated               primitive.DateTime   `bson:"dateUpdated" json:"dateUpdated"`
	FriendlyName              string               `bson:"friendlyName" json:"friendlyName"`
	Id                        string               `bson:"_id" json:"id"`
	IdCloudServicePbx         string               `bson:"idCloudServicePbx" json:"idCloudServicePbx"`
	IdCloudServicePbxFailover string               `bson:"idCloudServicePbxFailover" json:"idCloudServicePbxFailover"`
	IdVoipProvider            string               `bson:"idVoipProvider" json:"idVoipProvider"`
	ProviderId                string               `bson:"providerId" json:"providerId"`
	Reg_host                  string               `bson:"reg_host" json:"reg_host"`
	Reg_password              string               `bson:"reg_password" json:"reg_password"`
	Reg_port                  int32                `bson:"reg_port" json:"reg_port"`
	Reg_username              string               `bson:"reg_username" json:"reg_username"`
	TrunkOriginationType      TrunkOriginationType `bson:"trunkOriginationType" json:"trunkOriginationType"`
}

// Implementing interface IUbluxDocument
func (x TrunkOriginationRegister) GetDateDeleted() primitive.DateTime {
	return x.DateDeleted
}
func (x TrunkOriginationRegister) GetDateCreated() primitive.DateTime {
	return x.DateCreated
}
func (x TrunkOriginationRegister) GetDateUpdated() primitive.DateTime {
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
func (x TrunkOriginationRegister) GetTrunkOriginationType() TrunkOriginationType {
	return x.TrunkOriginationType
}
func (x TrunkOriginationRegister) GetProviderId() string {
	return x.ProviderId
}
func (x TrunkOriginationRegister) GetFriendlyName() string {
	return x.FriendlyName
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildTrunkOriginationRegister(m map[string]interface{}, x *TrunkOriginationRegister) {
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(primitive.DateTime)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(primitive.DateTime)
	}
	if val, ok := m["friendlyName"]; ok && val != nil {
		x.FriendlyName = val.(string)
	}
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
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
	if val, ok := m["idVoipProvider"]; ok && val != nil {
		x.IdVoipProvider = val.(string)
	}
	if val, ok := m["providerId"]; ok && val != nil {
		x.ProviderId = val.(string)
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
	x.TrunkOriginationType = TrunkOriginationType_Register // readonly property
}
