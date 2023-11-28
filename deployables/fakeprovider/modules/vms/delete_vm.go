package vms

import (
	"context"

	"github.com/shaharby7/Cloudy/deployables/fakeprovider/services"
)

func DeleteVM(ctx context.Context, options *DeleteVmOptions) (*Result, error) {
	err := services.SendDeleteVM(ctx, options.Machine_id)
	if err != nil {
		return &Result{}, err
	}
	return &Result{options.Machine_id}, nil
}
