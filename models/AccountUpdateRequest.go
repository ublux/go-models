package models

type AccountUpdateRequest struct {
	AccountSecrets  AccountSecrets  `bson:"accountSecrets" json:"accountSecrets"`
	AccountSettings AccountSettings `bson:"accountSettings" json:"accountSettings"`
	CompanyName     string          `bson:"companyName" json:"companyName"`
	Id              string          `bson:"id" json:"id"`
	MailingAddress  MailingAddress  `bson:"mailingAddress" json:"mailingAddress"`
}

// Implementing interface IUbluxDocumentId
func (x AccountUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildAccountUpdateRequest(m map[string]interface{}, x *AccountUpdateRequest) {
	if val, ok := m["accountSecrets"]; ok && val != nil {
		BuildAccountSecrets(val.(map[string]interface{}), &x.AccountSecrets)
	}
	if val, ok := m["accountSettings"]; ok && val != nil {
		BuildAccountSettings(val.(map[string]interface{}), &x.AccountSettings)
	}
	if val, ok := m["companyName"]; ok && val != nil {
		x.CompanyName = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["mailingAddress"]; ok && val != nil {
		BuildMailingAddress(val.(map[string]interface{}), &x.MailingAddress)
	}
}
