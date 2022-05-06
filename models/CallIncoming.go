package models

type CallIncoming interface {
	IUbluxDocument
	IUbluxDocumentId
	IReferncesAccount

	GetProviderSid() string
	GetIdVoipProvider() string
	GetContact() Contact
	GetIdVoipNumberPhone() string
	GetFromInternationalFormat() string
}
