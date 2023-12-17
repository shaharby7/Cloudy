package provider

import (
	"context"

	"github.com/shaharby7/Cloudy/deployables/clusteror/types"
)

type Provider interface {
	Identify() *ProviderCommonData
	CreateMachine(ctx context.Context, options *CreateMachineOptions) (*CreateMachineResults, error)
	TerminateMachine(ctx context.Context, options *TerminateMachineOptions) (*TerminateMachineResults, error)
	AllocatePublicIp(ctx context.Context, options *AllocatePublicIpOptions) (*AllocatePublicIpResults, error)
}
type ProviderCommonData struct {
	ProviderCode types.ProviderCode
}

type CreateMachineOptions struct {
	Specs            *types.MachineSpecs
	IpAllocationCode string
}

type CreateMachineResults struct {
	MachineProviderCode *types.MachineProviderCode
}

type TerminateMachineOptions struct {
	MachineProviderCode *types.MachineProviderCode
}
type TerminateMachineResults struct {
}

type AllocatePublicIpOptions struct {
}
type AllocatePublicIpResults struct {
	Ip   string
	Code string
}
