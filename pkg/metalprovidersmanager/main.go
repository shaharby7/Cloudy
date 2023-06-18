package metalprovidersmanager

import (
	"github.com/shaharby7/Cloudy/pkg/accounts"
	"github.com/shaharby7/Cloudy/pkg/metalprovidersmanager/abstractmetalprovider"
	"github.com/shaharby7/Cloudy/pkg/metalprovidersmanager/metal"
	"github.com/shaharby7/Cloudy/pkg/metalprovidersmanager/metalproviders"
	"reflect"
)

func genericCall[Payload interface{}, Result interface{}](
	identifier *metalproviders.MetalProviderIdentifier,
	method string,
	payload *Payload) Result {
	provider := metalproviders.MetalProviders[*identifier]
	untypedFunc := reflect.ValueOf(provider).MethodByName(method)
	var untypedResult []reflect.Value
	if payload != nil {
		inputs := make([]reflect.Value, 1)
		inputs[0] = reflect.ValueOf(payload)
		untypedResult = untypedFunc.Call(inputs)
	} else {
		untypedResult = untypedFunc.Call(nil)
	}
	return untypedResult[0].Interface().(Result)
}

func ListPrices(
	identifier *metalproviders.MetalProviderIdentifier,
) abstractmetalprovider.MetalPrices {
	return genericCall[interface{}, abstractmetalprovider.MetalPrices](
		identifier,
		"ListPrices",
		nil)
}

func InitiateHardware(identifier *metalproviders.MetalProviderIdentifier,
	metalType metal.HardwareTypes,
	account accounts.Account,
) metal.Metal {
	return genericCall[abstractmetalprovider.InitiateHardwarePayload, metal.Metal](
		identifier,
		"InitiateHardware",
		&abstractmetalprovider.InitiateHardwarePayload{MetalType: metalType, Account: account},
	)
}
