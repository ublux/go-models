package models

type CallerIdMask struct {
	Id                             string `bson:"_id" json:"id"`
	DateCreated                    string `bson:"dateCreated" json:"dateCreated"`
	DateDeleted                    string `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated                    string `bson:"dateUpdated" json:"dateUpdated"`
	FriendlyName                   string `bson:"friendlyName" json:"friendlyName"`
	IdAccount                      string `bson:"idAccount" json:"idAccount"`
	PhoneNumberInternationalFormat string `bson:"phoneNumberInternationalFormat" json:"phoneNumberInternationalFormat"`
	RandomVerificationCode         string `bson:"randomVerificationCode" json:"randomVerificationCode"`
}

// Implementing interface IUbluxDocument
func (x CallerIdMask) GetDateCreated() string {
	return x.DateCreated
}
func (x CallerIdMask) GetDateDeleted() string {
	return x.DateDeleted
}
func (x CallerIdMask) GetDateUpdated() string {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x CallerIdMask) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x CallerIdMask) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildCallerIdMask(m map[string]interface{}, x *CallerIdMask) {
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(string)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(string)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(string)
	}
	if val, ok := m["friendlyName"]; ok && val != nil {
		x.FriendlyName = val.(string)
	}
	if val, ok := m["idAccount"]; ok && val != nil {
		x.IdAccount = val.(string)
	}
	if val, ok := m["phoneNumberInternationalFormat"]; ok && val != nil {
		x.PhoneNumberInternationalFormat = val.(string)
	}
	if val, ok := m["randomVerificationCode"]; ok && val != nil {
		x.RandomVerificationCode = val.(string)
	}
}
