package models

type AudioUpdateRequest struct {
	Description  string `bson:"description" json:"description"`
	FriendlyName string `bson:"friendlyName" json:"friendlyName"`
	Id           string `bson:"id" json:"id"`
}

// Implementing interface IUbluxDocumentId
func (x AudioUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildAudioUpdateRequest(m map[string]interface{}, x *AudioUpdateRequest) {
	if val, ok := m["description"]; ok && val != nil {
		x.Description = val.(string)
	}
	if val, ok := m["friendlyName"]; ok && val != nil {
		x.FriendlyName = val.(string)
	}
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
}
