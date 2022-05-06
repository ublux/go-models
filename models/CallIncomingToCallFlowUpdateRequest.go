package models

type CallIncomingToCallFlowUpdateRequest struct {
	Id string `bson:"id" json:"id"`
}

// Implementing interface IUbluxDocumentId
func (x CallIncomingToCallFlowUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildCallIncomingToCallFlowUpdateRequest(m map[string]interface{}, x *CallIncomingToCallFlowUpdateRequest) {
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
}
