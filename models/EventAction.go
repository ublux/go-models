package models

import . "github.com/ublux/go-models/enums"

type EventAction interface {
	GetEventActionType() EventActionType
}
