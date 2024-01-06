package machine

import (
	"github.com/shaharby7/Cloudy/deployables/clusteror/types"
)

type NewOptions struct {
	ProviderCode types.ProviderCode `json:"provider_code"`
	MachineSpecs types.MachineSpecs `json:"machine_specs"`
}

type CreateOptions struct {
	IpAllocationsId types.IpAllocationId
}

type CreateResult struct{}

type TerminateOptions struct{}

type TerminateResult struct{}

type NetStack struct {
	IpAllocationIds []types.IpAllocationId `json:"ip_allocation_ids"`
}
