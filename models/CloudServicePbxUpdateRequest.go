package models

type CloudServicePbxUpdateRequest struct {
	Id string `bson:"id" json:"id"`
}

// Implementing interface IUbluxDocumentId
func (x CloudServicePbxUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildCloudServicePbxUpdateRequest(m map[string]interface{}, x *CloudServicePbxUpdateRequest) {
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
}
