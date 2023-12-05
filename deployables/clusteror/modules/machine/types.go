package machine

import (
	"github.com/shaharby7/Cloudy/deployables/clusteror/types"
)

type NewOptions struct {
	ProviderCode types.ProviderCode `json:"provider_code"`
	MachineSpecs types.MachineSpecs `json:"machine_specs"`
}

type CreateOptions struct{}

type CreateResult struct{}

type TerminateOptions struct{}

type TerminateResult struct{}
