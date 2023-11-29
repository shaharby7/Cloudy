package controllers

import (
	"github.com/shaharby7/Cloudy/pkg/common/serverutils"
	"github.com/shaharby7/Cloudy/pkg/deployable/controllable"

	"os"
	"strconv"

	"github.com/shaharby7/Cloudy/deployables/clusteror/modules/provider"
)

func GenerateServer() controllable.Controllable {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	Server := controllable.NewHttpServerControllable(
		"Clusteror server",
		*controllable.NewServerControllableConfig(port),
	)

	Server.RegisterActionable("/api/list-providers", serverutils.MakeServerRoute(provider.ListProviders))

	return Server
}