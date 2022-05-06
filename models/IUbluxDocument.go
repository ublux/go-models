package models

import "time"

type IUbluxDocument interface {
	IUbluxDocumentId

	GetDateDeleted() time.Time
	GetDateCreated() time.Time
	GetDateUpdated() time.Time
}
