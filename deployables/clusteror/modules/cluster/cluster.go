package cluster

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/shaharby7/Cloudy/deployables/clusteror/modules/machine"
	"github.com/shaharby7/Cloudy/deployables/clusteror/types"
)

type sCluster struct {
	Id           types.ClusterId     `json:"id"`
	Ip           string              `json:"ip"`
	JoinToken    string              `json:"token"`
	ClusterSpecs *types.ClusterSpecs `json:"cluster_specs"`
	Masters      []types.MachineId   `json:"masters"`
}

func New(ctx context.Context, options *types.ClusterSpecs) (*sCluster, error) {
	id := types.ClusterId(uuid.NewString())
	joinToken := uuid.NewString()
	cluster := &sCluster{
		Id:           id,
		ClusterSpecs: options,
		Masters:      make([]types.MachineId, 1),
		JoinToken:    joinToken,
	}
	userdata, err := compileUserData(ctx, nil)
	if err != nil {
		return cluster, fmt.Errorf("failed to generate master userdata: %s", err)
	}
	master_specs := &types.MachineSpecs{
		CPU:           1,
		RAM:           1,
		User_data_b64: userdata,
	}
	master, err := machine.New(ctx, &machine.NewOptions{
		ProviderCode: types.ProviderCode_FAKE1,
		MachineSpecs: *master_specs,
	})
	cluster.Masters[0] = master.ID
	if err != nil {
		return cluster, fmt.Errorf("failed to initiate master: %s", err)
	}
	_, err = master.Create(ctx, nil)
	if err != nil {
		return cluster, fmt.Errorf("failed to create master: %s", err)
	}
	return cluster, nil
}
