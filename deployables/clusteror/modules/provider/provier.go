package provider

import "context"

func ListProviders(ctx context.Context, _ *any) (*[]*ProviderData, error) {
	myOnlyProvider := &ProviderData{Name: "fakeprovider1"}
	result := make([]*ProviderData, 1)
	result[0] = myOnlyProvider
	return &result, nil
}
