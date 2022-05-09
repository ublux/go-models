package models

type VoicemailUpdateRequest struct {
	Id string `bson:"id" json:"id"`
}

// Implementing interface IUbluxDocumentId
func (x VoicemailUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildVoicemailUpdateRequest(m map[string]interface{}, x *VoicemailUpdateRequest) {
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
}
