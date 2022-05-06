package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UbluxDocument interface {
	IUbluxDocument
	IUbluxDocumentId

	GetDateDeleted() primitive.DateTime
	GetId() string
	GetDateCreated() primitive.DateTime
	GetDateUpdated() primitive.DateTime
}
