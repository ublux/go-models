package models

type TrunkTerminationGroupUpdateRequest struct {
	Id string `bson:"id" json:"id"`
}

// Implementing interface IUbluxDocumentId
func (x TrunkTerminationGroupUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildTrunkTerminationGroupUpdateRequest(m map[string]interface{}, x *TrunkTerminationGroupUpdateRequest) {
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
}
