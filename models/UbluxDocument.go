package models

type UbluxDocument interface {
	IUbluxDocument
	IUbluxDocumentId

	GetId() string
	GetDateCreated() string
	GetDateDeleted() string
	GetDateUpdated() string
}
