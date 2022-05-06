package models

import "time"

type UbluxDocument interface {
	IUbluxDocument
	IUbluxDocumentId

	GetDateDeleted() time.Time
	GetId() string
	GetDateCreated() time.Time
	GetDateUpdated() time.Time
}
