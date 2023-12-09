package provider

import (
	"context"

	"github.com/shaharby7/Cloudy/deployables/clusteror/types"
)

var providerByCode map[types.ProviderCode]Provider = map[types.ProviderCode]Provider{
	types.ProviderCode_FAKE1: &Fake{ProviderCode: types.ProviderCode_FAKE1},
	types.ProviderCode_FAKE2: &Fake{ProviderCode: types.ProviderCode_FAKE2},
}

func GetProviderByCode(code types.ProviderCode) Provider {
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
