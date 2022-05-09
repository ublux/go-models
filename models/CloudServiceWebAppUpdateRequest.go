package models

type CloudServiceWebAppUpdateRequest struct {
	Id string `bson:"id" json:"id"`
}

// Implementing interface IUbluxDocumentId
func (x CloudServiceWebAppUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildCloudServiceWebAppUpdateRequest(m map[string]interface{}, x *CloudServiceWebAppUpdateRequest) {
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
}
