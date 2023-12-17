package cluster

import (
	"context"

	"github.com/shaharby7/Cloudy/deployables/clusteror/types"
)

func Create(ctx context.Context, options *types.ClusterSpecs) (*sCluster, error) {
	cluster, err := New(ctx, options)
	if err != nil {
		return nil, err
	}
	cluster.Create(ctx, nil)
	return cluster, nil
}
