package models

type IUbluxDocument interface {
	IUbluxDocumentId

	GetDateCreated() string
	GetDateDeleted() string
	GetDateUpdated() string
}
