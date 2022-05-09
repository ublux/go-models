package models

type TrunkTerminationUpdateRequest struct {
	Id string `bson:"id" json:"id"`
}

// Implementing interface IUbluxDocumentId
func (x TrunkTerminationUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildTrunkTerminationUpdateRequest(m map[string]interface{}, x *TrunkTerminationUpdateRequest) {
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
}
