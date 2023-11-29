package main

import (
	"context"

	"github.com/shaharby7/Cloudy/deployables/clusteror"
)

func main() {
	clusteror.Initiate().Start(context.Background())
}
