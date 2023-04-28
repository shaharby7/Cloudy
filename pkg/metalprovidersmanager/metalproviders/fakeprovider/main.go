package fakeprovider

import (
	"fmt"

	"github.com/shaharby7/Cloudy/pkg/metalprovidersmanager/abstractmetalprovider"
	"github.com/shaharby7/Cloudy/pkg/metalprovidersmanager/metal"
)

type FakeProvider struct {
	*abstractmetalprovider.AbstractMetalProvider
	Options string
}

func (f *FakeProvider) ListPrices() abstractmetalprovider.MetalPrices {
	fmt.Println("listing prices")
	return abstractmetalprovider.MetalPrices{}
}

func (f *FakeProvider) InitiateHardware(
	*abstractmetalprovider.InitiateHardwarePayload,
) metal.Metal {
	fmt.Println("initiatingHardware")
	return metal.Metal{}
}
