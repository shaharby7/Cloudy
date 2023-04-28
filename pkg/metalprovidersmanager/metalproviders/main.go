package metalproviders

import (
	"github.com/shaharby7/Cloudy/pkg/metalprovidersmanager/abstractmetalprovider"
	"github.com/shaharby7/Cloudy/pkg/metalprovidersmanager/metalproviders/fakeprovider"
)

type MetalProviderIdentifier int64

const (
	Fake1 MetalProviderIdentifier = iota
	Fake2
)

// func initiateProvider[ProviderImplementation abstractmetalprovider.MetalProvider]() *ProviderImplementation {
// 	abstract:= &abstractmetalprovider.AbstractMetalProvider{};
// 	impl:= &ProviderImplementation{abstract};
// }

func initiateFakeProvider(Options string) *fakeprovider.FakeProvider {
	abstract := &abstractmetalprovider.AbstractMetalProvider{}
	impl := &fakeprovider.FakeProvider{AbstractMetalProvider: abstract, Options: Options}
	abstract.MetalProvider = impl
	return impl
}

var MetalProviders = map[MetalProviderIdentifier]abstractmetalprovider.MetalProvider{
	Fake1: initiateFakeProvider("hi"),
	Fake2: initiateFakeProvider("bye"),
}
