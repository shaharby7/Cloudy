package main

import (
	"github.com/shaharby7/Cloudy/pkg/accounts"
	"github.com/shaharby7/Cloudy/pkg/metalprovidersmanager"
	"github.com/shaharby7/Cloudy/pkg/metalprovidersmanager/metalproviders"

	"github.com/shaharby7/Cloudy/pkg/metalprovidersmanager/metal"
)

func main() {
	token := "1"
	account := accounts.Authenticate(&token)
	myProvider := metalproviders.Fake1
	metalprovidersmanager.ListPrices(&myProvider)
	metalprovidersmanager.InitiateHardware(&myProvider, metal.MyOtherHardware, account)
}
