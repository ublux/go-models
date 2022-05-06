package models

type TrunkOriginationForwardUpdateRequest struct {
	Id string `bson:"id" json:"id"`
}

// Implementing interface IUbluxDocumentId
func (x TrunkOriginationForwardUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildTrunkOriginationForwardUpdateRequest(m map[string]interface{}, x *TrunkOriginationForwardUpdateRequest) {
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
}
