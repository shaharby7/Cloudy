package network

import (
	"context"

	"github.com/google/uuid"
	"github.com/shaharby7/Cloudy/deployables/clusteror/modules/provider"
	"github.com/shaharby7/Cloudy/deployables/clusteror/types"
)

var allocations map[types.IpAllocationId]*sIpAllocation = make(map[types.IpAllocationId]*sIpAllocation, 0)

type sIpAllocation struct {
	ProviderCode types.ProviderCode   `json:"provider_code"`
	ID           types.IpAllocationId `json:"ip_allocation_id"`
	IP           string               `json:"ip"`
	Code         string               `json:"code"`
}

func New(ctx context.Context, options *NewOptions) (*sIpAllocation, error) {
	var id types.IpAllocationId = types.IpAllocationId(uuid.NewString())
	allocation := &sIpAllocation{
		ProviderCode: options.ProviderCode,
		ID:           id,
	}
	allocations[id] = allocation
	return allocation, nil
}

func (ipAllocation *sIpAllocation) Create(ctx context.Context, options *CreateOptions) (*CreateResults, error) {
	prov := provider.GetProviderByCode(ipAllocation.ProviderCode)
	providerResult, err := prov.AllocatePublicIp(ctx, nil)
	if err != nil {
		return nil, err
	}
	ipAllocation.IP = providerResult.Ip
	ipAllocation.Code = providerResult.Code
	return &CreateResults{}, nil
}

func GetAllocationById(ctx context.Context, id types.IpAllocationId) (*sIpAllocation, error) {
	return allocations[id], nil
}

type NewOptions struct {
	ProviderCode types.ProviderCode `json:"provider_code"`
}

type CreateOptions struct{}
type CreateResults struct{}
