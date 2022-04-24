package models

import . "github.com/ublux/go-models/enums"

type PowerDialer interface {
	GetPowerDialerType() PowerDialerType
	GetPhoneNumberInternationalFormat() string
	GetCountryIsoCode() CountryIsoCode
	GetIdContact() string
	GetPowerDialerStatus() PowerDialerStatus
	GetErrorMessage() string
	GetNumberOfAttempts() int32
	GetContactName() string
}
