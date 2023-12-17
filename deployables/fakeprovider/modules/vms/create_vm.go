package vms

import (
	"context"
	"fmt"
	"strconv"

	"github.com/shaharby7/Cloudy/deployables/fakeprovider/databases"
	"github.com/shaharby7/Cloudy/deployables/fakeprovider/services"
)

func CreateVM(ctx context.Context, options *CreateVmOptions) (*Result, error) {
	machine_id_int, err := databases.RedisClient.Incr(LAST_MACHINE_ID_REDIS_KEY).Result()
	if err != nil {
		return nil, err
	}
	Machine_id := fmt.Sprintf("fakeprovider-vm-%s", strconv.Itoa(int(machine_id_int)))
	IpAllocationId := "X"
	if options.IpAllocationId != "" {
		IpAllocationId = options.IpAllocationId
	}
	err = services.SendCreateVM(ctx, Machine_id, options.User_data_b64, IpAllocationId)
	if err != nil {
		return &Result{}, err
	}
	return &Result{Machine_id}, nil
}
