package machine

import (
	"context"

	"github.com/google/uuid"
	"github.com/shaharby7/Cloudy/deployables/clusteror/modules/provider"
)

type MachineStatus int

const (
	MachineStatus_STARTING MachineStatus = iota
	MachineStatus_RUNNING
	MachineStatus_TERMINATING
	MachineStatus_UNAVAILABLE
)

type sMachine struct {
	ID       string                `json:"id"`
	Provider provider.ProviderCode `json:"provider"`
	Status   MachineStatus         `json:"status"`
}

func New(ctx context.Context, provider provider.ProviderCode) (*sMachine, error) {
	id := uuid.NewString()
	return &sMachine{ID: id, Provider: provider}, nil
}

func (machine *sMachine) Create(ctx context.Context) error {
	provider := provider.GetProviderByCode(machine.Provider)
	return provider.CreateMachine()
}

func (machine *sMachine) Destroy(ctx context.Context) error {
	provider := provider.GetProviderByCode(machine.Provider)
	return provider.DestroyMachine()
}
