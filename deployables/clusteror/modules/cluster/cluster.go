package cluster

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/shaharby7/Cloudy/deployables/clusteror/modules/machine"
	"github.com/shaharby7/Cloudy/deployables/clusteror/modules/network"
	"github.com/shaharby7/Cloudy/deployables/clusteror/modules/provider"
	"github.com/shaharby7/Cloudy/deployables/clusteror/types"
)

const MASTERS_PROVIDER_CODE = types.ProviderCode_FAKE1

var MASTERS_PROVIDER = provider.GetProviderByCode(MASTERS_PROVIDER_CODE)

type sCluster struct {
	Id           types.ClusterId     `json:"id"`
	Ip           string              `json:"ip"`
	JoinToken    string              `json:"token"`
	ClusterSpecs *types.ClusterSpecs `json:"cluster_specs"`
	Masters      []types.MachineId   `json:"masters"`
	Status       types.ClusterStatus `json:"status"`
	Errors       []error             `json:"errors"`
}

func New(ctx context.Context, options *types.ClusterSpecs) (*sCluster, error) {
	id := types.ClusterId(uuid.NewString())
	joinToken := uuid.NewString()
	cluster := &sCluster{
		Id:           id,
		ClusterSpecs: options,
		Masters:      make([]types.MachineId, 1),
		JoinToken:    joinToken,
		Errors:       make([]error, 0),
	}
	cluster.setStatus(ctx, types.ClusterStatus_NEW)
	return cluster, nil
}

func (cluster *sCluster) Create(ctx context.Context, options *CreateClusterOptions) (*CreateClusterResult, error) {
	master_ip_allocation, err := network.New(ctx, &network.NewOptions{ProviderCode: MASTERS_PROVIDER_CODE})
	if err != nil {
		return nil, cluster.handleError(ctx, err, "failed to allocate ip to master: %s")
	}
	_, err = master_ip_allocation.Create(ctx, nil)
	if err != nil {
		return nil, cluster.handleError(ctx, err, "failed to allocate ip to master: %s")
	}
	cluster.Ip = master_ip_allocation.IP
	master_userdata, err := compileUserData(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to generate master userdata: %s", err)
	}
	master_specs := &types.MachineSpecs{
		CPU:           1,
		RAM:           1,
		User_data_b64: master_userdata,
	}
	master, err := machine.New(ctx, &machine.NewOptions{
		ProviderCode: MASTERS_PROVIDER_CODE,
		MachineSpecs: *master_specs,
	})
	cluster.Masters[0] = master.ID
	if err != nil {
		return nil, cluster.handleError(ctx, err, "failed to initiate master: %s")
	}
	_, err = master.Create(ctx, &machine.CreateOptions{IpAllocationsId: master_ip_allocation.ID})
	if err != nil {
		return nil, cluster.handleError(ctx, err, "failed to create master: %s")
	}
	return nil, nil
}

func (cluster *sCluster) setStatus(ctx context.Context, status types.ClusterStatus) {
	cluster.Status = status
}

func (cluster *sCluster) handleError(ctx context.Context, err error, errTemplate string) error {
	handledErr := fmt.Errorf(errTemplate, err)
	cluster.setStatus(ctx, types.ClusterStatus_ERROR)
	cluster.Errors = append(cluster.Errors, handledErr)
	return handledErr
}
