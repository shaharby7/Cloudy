package abstractmetalprovider

import (
	"github.com/shaharby7/Cloudy/pkg/accounts"
	"github.com/shaharby7/Cloudy/pkg/metalprovidersmanager/metal"
)

type MetalPrices = map[metal.HardwareTypes]float64

type InitiateHardwarePayload struct {
	MetalType metal.HardwareTypes
	Account   accounts.Account
}

type MetalProvider interface {
	ListPrices() MetalPrices

	InitiateHardware(
		*InitiateHardwarePayload,
	) metal.Metal
}

type AbstractMetalProvider struct {
	MetalProvider
	Options string
}
