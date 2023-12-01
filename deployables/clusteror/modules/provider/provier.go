package provider

import (
	"context"
)

type ProviderCode int

const (
	ProviderCode_FAKE1 ProviderCode = iota
	ProviderCode_FAKE2
)

var providerByCode map[ProviderCode]Provider = map[ProviderCode]Provider{
	ProviderCode_FAKE1: &Fake{Code: ProviderCode_FAKE1},
	ProviderCode_FAKE2: &Fake{Code: ProviderCode_FAKE2},
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
