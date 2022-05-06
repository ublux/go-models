package models

type AgreementToTermsAndConditionsUpdateRequest struct {
	Id string `bson:"id" json:"id"`
}

// Implementing interface IUbluxDocumentId
func (x AgreementToTermsAndConditionsUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildAgreementToTermsAndConditionsUpdateRequest(m map[string]interface{}, x *AgreementToTermsAndConditionsUpdateRequest) {
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
}
