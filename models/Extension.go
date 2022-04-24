package models

import . "github.com/ublux/go-models/enums"

type Extension interface {
	IUbluxDocument
	IUbluxDocumentId
	IReferncesAccount

	GetIdAccount() string
	GetIdMusicOnHoldGroup() string
	GetExtensionType() ExtensionType
	GetNumber() string
	GetInjectExtensionNameToCallerId() bool
}
