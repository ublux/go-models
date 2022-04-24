package models

import . "github.com/ublux/go-models/enums"

type ChildCall interface {
	GetChildCallType() ChildCallType
}
