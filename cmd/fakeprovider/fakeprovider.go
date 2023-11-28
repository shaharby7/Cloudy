package main

import (
	"context"

	"github.com/shaharby7/Cloudy/deployables/fakeprovider"
)

func main() {
	fakeprovider.InstantiateFakeproviderDeployable().Start(context.TODO())
}
