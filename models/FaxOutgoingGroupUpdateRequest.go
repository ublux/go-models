package models

type FaxOutgoingGroupUpdateRequest struct {
	Id string `bson:"id" json:"id"`
}

// Implementing interface IUbluxDocumentId
func (x FaxOutgoingGroupUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildFaxOutgoingGroupUpdateRequest(m map[string]interface{}, x *FaxOutgoingGroupUpdateRequest) {
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
}
