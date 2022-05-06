package models

type PhoneConfigurationUpdateRequest struct {
	Description    string `bson:"description" json:"description"`
	FrienlyName    string `bson:"frienlyName" json:"frienlyName"`
	Id             string `bson:"id" json:"id"`
	IdLineKeyGroup string `bson:"idLineKeyGroup" json:"idLineKeyGroup"`
}

// Implementing interface IUbluxDocumentId
func (x PhoneConfigurationUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildPhoneConfigurationUpdateRequest(m map[string]interface{}, x *PhoneConfigurationUpdateRequest) {
	if val, ok := m["description"]; ok && val != nil {
		x.Description = val.(string)
	}
	if val, ok := m["frienlyName"]; ok && val != nil {
		x.FrienlyName = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["idLineKeyGroup"]; ok && val != nil {
		x.IdLineKeyGroup = val.(string)
	}
}
