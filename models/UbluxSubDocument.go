package models

type UbluxSubDocument interface {
	IUbluxSubDocument
	IUbluxDocumentId

	GetId() string
}
