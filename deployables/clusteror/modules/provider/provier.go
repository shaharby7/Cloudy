package provider

import (
	"context"
)

type ProviderCode string

var providerByCode map[ProviderCode]Provider = map[ProviderCode]Provider{
	"FAKE1": &Fake{Code: "FAKE1"},
	"FAKE2": &Fake{Code: "FAKE2"},
}

func GetProviderByCode(code ProviderCode) Provider {
	return providerByCode[code]
}

func ListProviders(ctx context.Context, _ *any) (*[]*ProviderCommonData, error) {
	var results []*ProviderCommonData = make([]*ProviderCommonData, len(providerByCode))
	index := 0
	for _, prov := range providerByCode {
		results[index] = prov.Identify()
		index++
	}
	return &results, nil
}
