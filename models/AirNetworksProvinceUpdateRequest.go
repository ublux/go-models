package models

type AirNetworksProvinceUpdateRequest struct {
	Id string `bson:"id" json:"id"`
}

// Implementing interface IUbluxDocumentId
func (x AirNetworksProvinceUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildAirNetworksProvinceUpdateRequest(m map[string]interface{}, x *AirNetworksProvinceUpdateRequest) {
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
}
