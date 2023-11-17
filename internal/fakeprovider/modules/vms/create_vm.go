package vms

import (
	"context"
	"fmt"
	"strconv"

	"github.com/shaharby7/Cloudy/internal/fakeprovider/databases"
	"github.com/shaharby7/Cloudy/internal/fakeprovider/services"
)

func CreateVM(ctx context.Context, options *CreateVmOptions) (*Result, error) {
	machine_id_int, err := databases.RedisClient.Incr(LAST_MACHINE_ID_REDIS_KEY).Result()
	if err != nil {
		return nil, err
	}
	Machine_id := fmt.Sprintf("fakeprovider-vm-%s", strconv.Itoa(int(machine_id_int)));
	err = services.SendCreateVM(ctx, Machine_id, options.User_data_b64)
	if err != nil {
		return &Result{}, err
	}
	return &Result{Machine_id}, nil
}
