package cluster

import (
	"context"

	"github.com/shaharby7/Cloudy/deployables/clusteror/types"
)

func Create(ctx context.Context, options *types.ClusterSpecs) (*[]types.MachineId, error) {
	cluster, err := New(ctx, options)
	if err != nil {
		return nil, err
	}
	return &cluster.Masters, nil
}
