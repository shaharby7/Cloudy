package machine

import (
	"context"

	"github.com/google/uuid"
	"github.com/shaharby7/Cloudy/deployables/clusteror/modules/network"
	"github.com/shaharby7/Cloudy/deployables/clusteror/modules/provider"
	"github.com/shaharby7/Cloudy/deployables/clusteror/types"
)

type sMachine struct {
	ID                  types.MachineId           `json:"id"`
	ProviderCode        types.ProviderCode        `json:"provider"`
	MachineProviderCode types.MachineProviderCode `json:"machine_provider_code"`
	Status              types.MachineStatus       `json:"status"`
	Specs               *types.MachineSpecs       `json:"specs"`
	NetStack            NetStack
}

func New(ctx context.Context, options *NewOptions) (*sMachine, error) {
	var id types.MachineId = types.MachineId(uuid.NewString())
	return &sMachine{
		ID:           id,
		ProviderCode: options.ProviderCode,
		Status:       types.MachineStatus_NEW,
		Specs:        &options.MachineSpecs,
	}, nil
}

func (machine *sMachine) Create(ctx context.Context, options *CreateOptions) (*CreateResult, error) {
	createResult := &CreateResult{}
	prov := provider.GetProviderByCode(machine.ProviderCode)
	machine.setStatus(ctx, types.MachineStatus_CREATING)
	ipAllocationCode := machine.handleIpAllocation(ctx, options)
	result, err := prov.CreateMachine(ctx, &provider.CreateMachineOptions{
		Specs:            machine.Specs,
		IpAllocationCode: ipAllocationCode,
	})
	if err != nil {
		machine.setStatus(ctx, types.MachineStatus_ERROR)
		return createResult, err
	}
	machine.setStatus(ctx, types.MachineStatus_RUNNING)
	machine.MachineProviderCode = *result.MachineProviderCode
	return &CreateResult{}, nil
}

func (machine *sMachine) Terminate(ctx context.Context, options *TerminateOptions) (*TerminateResult, error) {
	terminateResult := &TerminateResult{}
	prov := provider.GetProviderByCode(machine.ProviderCode)
	machine.setStatus(ctx, types.MachineStatus_TERMINATING)
	_, err := prov.TerminateMachine(ctx, &provider.TerminateMachineOptions{
		MachineProviderCode: &machine.MachineProviderCode,
	})
	if err != nil {
		machine.setStatus(ctx, types.MachineStatus_ERROR)
		return terminateResult, err
	}
	machine.setStatus(ctx, types.MachineStatus_TERMINATED)
	return &TerminateResult{}, nil
}

func (machine *sMachine) setStatus(ctx context.Context, status types.MachineStatus) {
	machine.Status = status
}

func (machine *sMachine) handleIpAllocation(ctx context.Context, options *CreateOptions) string {
	if options.IpAllocationsId == "" {
		return ""
	}
	machine.NetStack.IpAllocationIds = append(machine.NetStack.IpAllocationIds, options.IpAllocationsId)
	ipAllocation, _ := network.GetAllocationById(ctx, options.IpAllocationsId)
	ipAllocationCode := ipAllocation.Code
	return ipAllocationCode
}
