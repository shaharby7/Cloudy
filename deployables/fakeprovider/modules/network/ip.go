package network

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/shaharby7/Cloudy/deployables/fakeprovider/services"
)

func AllocatePublicIp(ctx context.Context, options *AllocatePublicIpOptions) (*AllocatePublicIpResults, error) {
	id := strings.ReplaceAll(uuid.NewString(), "-", "")
	result, err := services.AllocatePublicIp(
		ctx,
		&services.AllocatePublicIpOptions{AllocationId: id},
	)
	if err != nil {
		return nil, err
	}
	converted := AllocatePublicIpResults(*result)
	return &converted, err
}

type AllocatePublicIpOptions struct{}
type AllocatePublicIpResults struct {
	Ip           string `json:"ip"`
	AllocationId string `json:"ip_allocation_id"`
}
