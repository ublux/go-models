package models

import . "github.com/ublux/go-models/enums"

type CloudService interface {
	IUbluxDocument
	IUbluxDocumentId

	GetCloudServiceType() CloudServiceType
	GetCountryIsoCode() CountryIsoCode
	GetLocalnet() string
	GetExternalIps() []string
	GetIsFailover() bool
	GetNat() bool
	GetIsHealthy() bool
}
