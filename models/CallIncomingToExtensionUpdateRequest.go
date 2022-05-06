package models

type CallIncomingToExtensionUpdateRequest struct {
	Id string `bson:"id" json:"id"`
}

// Implementing interface IUbluxDocumentId
func (x CallIncomingToExtensionUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildCallIncomingToExtensionUpdateRequest(m map[string]interface{}, x *CallIncomingToExtensionUpdateRequest) {
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
}
