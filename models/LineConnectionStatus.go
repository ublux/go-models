package models

type LineConnectionStatus struct {
	Id            string `bson:"_id" json:"id"`
	DateConnected string `bson:"dateConnected" json:"dateConnected"`
	IpLAN         string `bson:"ipLAN" json:"ipLAN"`
	IpWAN         string `bson:"ipWAN" json:"ipWAN"`
	IsConnected   bool   `bson:"isConnected" json:"isConnected"`
	PortLAN       int32  `bson:"portLAN" json:"portLAN"`
	PortWAN       int32  `bson:"portWAN" json:"portWAN"`
	UserAgent     string `bson:"userAgent" json:"userAgent"`
}

// Implementing interface IUbluxSubDocument
// Implementing interface IUbluxDocumentId
func (x LineConnectionStatus) GetId() string {
	return x.Id
}

// Implementing interface UbluxSubDocument

// BUILDER from bson map:
func BuildLineConnectionStatus(m map[string]interface{}, x *LineConnectionStatus) {
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["dateConnected"]; ok && val != nil {
		x.DateConnected = val.(string)
	}
	if val, ok := m["ipLAN"]; ok && val != nil {
		x.IpLAN = val.(string)
	}
	if val, ok := m["ipWAN"]; ok && val != nil {
		x.IpWAN = val.(string)
	}
	if val, ok := m["isConnected"]; ok && val != nil {
		x.IsConnected = val.(bool)
	}
	if val, ok := m["portLAN"]; ok && val != nil {
		x.PortLAN = val.(int32)
	}
	if val, ok := m["portWAN"]; ok && val != nil {
		x.PortWAN = val.(int32)
	}
	if val, ok := m["userAgent"]; ok && val != nil {
		x.UserAgent = val.(string)
	}
}
