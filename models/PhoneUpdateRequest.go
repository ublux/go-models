package models

type PhoneUpdateRequest struct {
	FriendlyName         string `bson:"friendlyName" json:"friendlyName"`
	Id                   string `bson:"id" json:"id"`
	IdCloudServicePbx    string `bson:"idCloudServicePbx" json:"idCloudServicePbx"`
	IdIdentityWebApp     string `bson:"idIdentityWebApp" json:"idIdentityWebApp"`
	IdPhoneConfiguration string `bson:"idPhoneConfiguration" json:"idPhoneConfiguration"`
}

// Implementing interface IUbluxDocumentId
func (x PhoneUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildPhoneUpdateRequest(m map[string]interface{}, x *PhoneUpdateRequest) {
	if val, ok := m["friendlyName"]; ok && val != nil {
		x.FriendlyName = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["idCloudServicePbx"]; ok && val != nil {
		x.IdCloudServicePbx = val.(string)
	}
	if val, ok := m["idIdentityWebApp"]; ok && val != nil {
		x.IdIdentityWebApp = val.(string)
	}
	if val, ok := m["idPhoneConfiguration"]; ok && val != nil {
		x.IdPhoneConfiguration = val.(string)
	}
}
