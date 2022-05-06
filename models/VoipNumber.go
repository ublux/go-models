package models

import . "github.com/ublux/go-models/enums"

type VoipNumber interface {
	IUbluxDocument
	IUbluxDocumentId
	IReferncesAccount

	GetIdAccount() string
	GetIdCustomerInfo() string
	GetIdTrunkOrigination() string
	GetIdVoipProvider() string
	GetSID() string
	GetRulesPhone() []RulePhone
	GetRulesSms() []RuleSms
	GetRulesFax() []RuleFax
	GetVoipNumberType() VoipNumberType
	GetInjectFriendlyNameToCallerId() bool
	GetRecordIncomingCalls() bool
	GetNumber() string
	GetFriendlyName() string
	GetDescription() string
	GetLanguage() Language
	GetCity() string
	GetState() string
	GetCountryIsoCode() CountryIsoCode
	GetIsSmsEnabled() bool
	GetIsVoiceEnabled() bool
	GetIsTollFree() bool
	GetIsWhatsappEnabled() bool
	GetTimeZone() string
}
