package models

import "go.mongodb.org/mongo-driver/bson/primitive"
import . "github.com/ublux/go-models/enums"

type TrunkOriginationForward struct {
	DateCreated               primitive.DateTime   `bson:"dateCreated" json:"dateCreated"`
	DateDeleted               primitive.DateTime   `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated               primitive.DateTime   `bson:"dateUpdated" json:"dateUpdated"`
	FriendlyName              string               `bson:"friendlyName" json:"friendlyName"`
	Id                        string               `bson:"id" json:"id"`
	IdCloudServicePbx         string               `bson:"idCloudServicePbx" json:"idCloudServicePbx"`
	IdCloudServicePbxFailover string               `bson:"idCloudServicePbxFailover" json:"idCloudServicePbxFailover"`
	IdsVoipNumbers            []string             `bson:"idsVoipNumbers" json:"idsVoipNumbers"`
	IdVoipProvider            string               `bson:"idVoipProvider" json:"idVoipProvider"`
	ProviderId                string               `bson:"providerId" json:"providerId"`
	SipUri                    string               `bson:"sipUri" json:"sipUri"`
	SipUriFailover            string               `bson:"sipUriFailover" json:"sipUriFailover"`
	TrunkOriginationType      TrunkOriginationType `bson:"trunkOriginationType" json:"trunkOriginationType"`
}

// Implementing interface IUbluxDocument
func (x TrunkOriginationForward) GetDateDeleted() primitive.DateTime {
	return x.DateDeleted
}
func (x TrunkOriginationForward) GetDateCreated() primitive.DateTime {
	return x.DateCreated
}
func (x TrunkOriginationForward) GetDateUpdated() primitive.DateTime {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x TrunkOriginationForward) GetId() string {
	return x.Id
}

// Implementing interface TrunkOrigination
func (x TrunkOriginationForward) GetIdVoipProvider() string {
	return x.IdVoipProvider
}
func (x TrunkOriginationForward) GetIdCloudServicePbx() string {
	return x.IdCloudServicePbx
}
func (x TrunkOriginationForward) GetIdCloudServicePbxFailover() string {
	return x.IdCloudServicePbxFailover
}
func (x TrunkOriginationForward) GetIdsVoipNumbers() []string {
	return x.IdsVoipNumbers
}
func (x TrunkOriginationForward) GetTrunkOriginationType() TrunkOriginationType {
	return x.TrunkOriginationType
}
func (x TrunkOriginationForward) GetProviderId() string {
	return x.ProviderId
}
func (x TrunkOriginationForward) GetFriendlyName() string {
	return x.FriendlyName
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildTrunkOriginationForward(m map[string]interface{}, x *TrunkOriginationForward) {
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
	if val, ok := m["providerId"]; ok && val != nil {
		x.ProviderId = val.(string)
	}
	if val, ok := m["sipUri"]; ok && val != nil {
		x.SipUri = val.(string)
	}
	if val, ok := m["sipUriFailover"]; ok && val != nil {
		x.SipUriFailover = val.(string)
	}
	x.TrunkOriginationType = TrunkOriginationType_Forward // readonly property
}
