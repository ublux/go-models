package models

type CallOutgoingToPSTNUpdateRequest struct {
	Id string `bson:"id" json:"id"`
}

// Implementing interface IUbluxDocumentId
func (x CallOutgoingToPSTNUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildCallOutgoingToPSTNUpdateRequest(m map[string]interface{}, x *CallOutgoingToPSTNUpdateRequest) {
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
}
