package vms

import (
	"context"

	"github.com/shaharby7/Cloudy/deployables/fakeprovider/services"
)

func GetVM(ctx context.Context, options *GetVmOptions) (*GetVmResults, error) {
	ip, err := services.GetMachineIp(ctx, options.Machine_id)
	if err != nil {
		return &GetVmResults{}, err
	}
	return &GetVmResults{options.Machine_id, ip}, nil
}
