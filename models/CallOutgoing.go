package models

type CallOutgoing interface {
	IUbluxDocument
	IUbluxDocumentId
	IReferncesAccount

	GetContact() Contact
	GetIdLineThatInitiatedCall() string
}
