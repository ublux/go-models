package models

type CallerIdMaskUpdateRequest struct {
	FriendlyName string `bson:"friendlyName" json:"friendlyName"`
	Id           string `bson:"id" json:"id"`
}

// Implementing interface IUbluxDocumentId
func (x CallerIdMaskUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildCallerIdMaskUpdateRequest(m map[string]interface{}, x *CallerIdMaskUpdateRequest) {
	if val, ok := m["friendlyName"]; ok && val != nil {
		x.FriendlyName = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
}
