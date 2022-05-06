package models

import . "github.com/ublux/go-models/enums"

type VoipNumberAvailableForPurchaseUpdateRequest struct {
	Description                  string   `bson:"description" json:"description"`
	FriendlyName                 string   `bson:"friendlyName" json:"friendlyName"`
	Id                           string   `bson:"id" json:"id"`
	InjectFriendlyNameToCallerId bool     `bson:"injectFriendlyNameToCallerId" json:"injectFriendlyNameToCallerId"`
	Language                     Language `bson:"language" json:"language"`
	RecordIncomingCalls          bool     `bson:"recordIncomingCalls" json:"recordIncomingCalls"`
	TimeZone                     string   `bson:"timeZone" json:"timeZone"`
}

// Implementing interface IUbluxDocumentId
func (x VoipNumberAvailableForPurchaseUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildVoipNumberAvailableForPurchaseUpdateRequest(m map[string]interface{}, x *VoipNumberAvailableForPurchaseUpdateRequest) {
	if val, ok := m["description"]; ok && val != nil {
		x.Description = val.(string)
	}
	if val, ok := m["friendlyName"]; ok && val != nil {
		x.FriendlyName = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["injectFriendlyNameToCallerId"]; ok && val != nil {
		x.InjectFriendlyNameToCallerId = val.(bool)
	}
	if val, ok := m["language"]; ok && val != nil {
		x.Language = Language("Language_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["recordIncomingCalls"]; ok && val != nil {
		x.RecordIncomingCalls = val.(bool)
	}
	if val, ok := m["timeZone"]; ok && val != nil {
		x.TimeZone = val.(string)
	}
}
