package models

type VoicemailForwardedUpdateRequest struct {
	Id string `bson:"id" json:"id"`
}

// Implementing interface IUbluxDocumentId
func (x VoicemailForwardedUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildVoicemailForwardedUpdateRequest(m map[string]interface{}, x *VoicemailForwardedUpdateRequest) {
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
}
