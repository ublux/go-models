package models

type CallIncoming interface {
	IUbluxDocument
	IUbluxDocumentId
	IReferncesAccount

	GetProviderSid() string
	GetContact() Contact
	GetIdVoipProvider() string
	GetIdVoipNumberPhone() string
}
