package machine

import (
	"context"

	"github.com/google/uuid"
	"github.com/shaharby7/Cloudy/deployables/clusteror/modules/provider"
	"github.com/shaharby7/Cloudy/deployables/clusteror/types"
)

type sMachine struct {
	ID                  string                    `json:"id"`
	ProviderCode        types.ProviderCode        `json:"provider"`
	MachineProviderCode types.MachineProviderCode `json:"machine_provider_code"`
	Status              types.MachineStatus       `json:"status"`
	Specs               *types.MachineSpecs       `json:"specs"`
}

func New(ctx context.Context, options *NewOptions) (*sMachine, error) {
	id := uuid.NewString()
	return &sMachine{ID: id, ProviderCode: options.ProviderCode, Status: types.NEW, Specs: &options.MachineSpecs}, nil
}

func (machine *sMachine) Create(ctx context.Context, options *CreateOptions) (*CreateResult, error) {
	createResult := &CreateResult{}
	prov := provider.GetProviderByCode(machine.ProviderCode)
	machine.setStatus(ctx, types.CREATING)
	result, err := prov.CreateMachine(ctx, &provider.CreateMachineOptions{
		Specs: machine.Specs,
	})
	if err != nil {
		machine.setStatus(ctx, types.ERROR)
		return createResult, err
	}
	machine.setStatus(ctx, types.RUNNING)
	machine.MachineProviderCode = *result.MachineProviderCode
	return &CreateResult{}, nil
}

func (machine *sMachine) Terminate(ctx context.Context, options *TerminateOptions) (*TerminateResult, error) {
	terminateResult := &TerminateResult{}
	prov := provider.GetProviderByCode(machine.ProviderCode)
	machine.setStatus(ctx, types.TERMINATING)
	_, err := prov.TerminateMachine(ctx, &provider.TerminateMachineOptions{
		MachineProviderCode: &machine.MachineProviderCode,
	})
	if err != nil {
		machine.setStatus(ctx, types.ERROR)
		return terminateResult, err
	}
	machine.setStatus(ctx, types.TERMINATED)
	return &TerminateResult{}, nil
}

func (machine *sMachine) setStatus(ctx context.Context, status types.MachineStatus) {
	machine.Status = status
}
