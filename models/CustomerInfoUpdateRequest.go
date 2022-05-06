package models

type CustomerInfoUpdateRequest struct {
	FullName       string         `bson:"fullName" json:"fullName"`
	Id             string         `bson:"id" json:"id"`
	MailingAddress MailingAddress `bson:"mailingAddress" json:"mailingAddress"`
}

// Implementing interface IUbluxDocumentId
func (x CustomerInfoUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildCustomerInfoUpdateRequest(m map[string]interface{}, x *CustomerInfoUpdateRequest) {
	if val, ok := m["fullName"]; ok && val != nil {
		x.FullName = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["mailingAddress"]; ok && val != nil {
		BuildMailingAddress(val.(map[string]interface{}), &x.MailingAddress)
	}
}
