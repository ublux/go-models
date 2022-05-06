package models

type VoipProviderUpdateRequest struct {
	Id string `bson:"id" json:"id"`
}

// Implementing interface IUbluxDocumentId
func (x VoipProviderUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildVoipProviderUpdateRequest(m map[string]interface{}, x *VoipProviderUpdateRequest) {
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
}
