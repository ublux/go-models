package models

type FaxEmailInfoUpdateRequest struct {
	Id string `bson:"id" json:"id"`
}

// Implementing interface IUbluxDocumentId
func (x FaxEmailInfoUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildFaxEmailInfoUpdateRequest(m map[string]interface{}, x *FaxEmailInfoUpdateRequest) {
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
}
