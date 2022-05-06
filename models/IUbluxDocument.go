package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type IUbluxDocument interface {
	IUbluxDocumentId

	GetDateDeleted() primitive.DateTime
	GetDateCreated() primitive.DateTime
	GetDateUpdated() primitive.DateTime
}
