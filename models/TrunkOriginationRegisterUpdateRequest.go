package models

type TrunkOriginationRegisterUpdateRequest struct {
	Id string `bson:"id" json:"id"`
}

// Implementing interface IUbluxDocumentId
func (x TrunkOriginationRegisterUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildTrunkOriginationRegisterUpdateRequest(m map[string]interface{}, x *TrunkOriginationRegisterUpdateRequest) {
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
}
