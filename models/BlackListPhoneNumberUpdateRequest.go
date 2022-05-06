package models

type BlackListPhoneNumberUpdateRequest struct {
	FriendlyName string `bson:"friendlyName" json:"friendlyName"`
	Id           string `bson:"id" json:"id"`
	PhoneNumber  string `bson:"phoneNumber" json:"phoneNumber"`
}

// Implementing interface IUbluxDocumentId
func (x BlackListPhoneNumberUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildBlackListPhoneNumberUpdateRequest(m map[string]interface{}, x *BlackListPhoneNumberUpdateRequest) {
	if val, ok := m["friendlyName"]; ok && val != nil {
		x.FriendlyName = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["phoneNumber"]; ok && val != nil {
		x.PhoneNumber = val.(string)
	}
}
