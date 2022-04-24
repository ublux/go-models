package models

type AirNetworksCustomerInfo struct {
	IdAirNetworksProvince string `bson:"idAirNetworksProvince" json:"idAirNetworksProvince"`
	IdNumber              string `bson:"idNumber" json:"idNumber"`
	IdType                string `bson:"idType" json:"idType"`
	Population            string `bson:"population" json:"population"`
	Province              string `bson:"province" json:"province"`
}

// BUILDER from bson map:
func BuildAirNetworksCustomerInfo(m map[string]interface{}, x *AirNetworksCustomerInfo) {
	if val, ok := m["idAirNetworksProvince"]; ok && val != nil {
		x.IdAirNetworksProvince = val.(string)
	}
	if val, ok := m["idNumber"]; ok && val != nil {
		x.IdNumber = val.(string)
	}
	if val, ok := m["idType"]; ok && val != nil {
		x.IdType = val.(string)
	}
	if val, ok := m["population"]; ok && val != nil {
		x.Population = val.(string)
	}
	if val, ok := m["province"]; ok && val != nil {
		x.Province = val.(string)
	}
}
