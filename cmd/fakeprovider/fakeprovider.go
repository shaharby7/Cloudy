package main

import (
	"context"

	"github.com/shaharby7/Cloudy/internal/fakeprovider"
)

func main() {
	fakeprovider.MyDeployable.Start(context.TODO())
}
