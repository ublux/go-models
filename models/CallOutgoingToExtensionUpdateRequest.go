package models

type CallOutgoingToExtensionUpdateRequest struct {
	Id string `bson:"id" json:"id"`
}

// Implementing interface IUbluxDocumentId
func (x CallOutgoingToExtensionUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildCallOutgoingToExtensionUpdateRequest(m map[string]interface{}, x *CallOutgoingToExtensionUpdateRequest) {
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
}
