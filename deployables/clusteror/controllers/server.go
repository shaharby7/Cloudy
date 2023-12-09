package controllers

import (
	"github.com/shaharby7/Cloudy/pkg/common/serverutils"
	"github.com/shaharby7/Cloudy/pkg/deployable/controllable"

	"os"
	"strconv"

	"github.com/shaharby7/Cloudy/deployables/clusteror/modules/cluster"
	"github.com/shaharby7/Cloudy/deployables/clusteror/modules/machine"
	"github.com/shaharby7/Cloudy/deployables/clusteror/modules/provider"
)

func GenerateServer() controllable.Controllable {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	Server := controllable.NewHttpServerControllable(
		"Clusteror server",
		*controllable.NewServerControllableConfig(port),
	)

	Server.RegisterActionable("/api/providers/list", serverutils.MakeServerRoute(provider.ListProviders))

	Server.RegisterActionable("/api/machines/create", serverutils.MakeServerRoute(machine.CreateAndRun))

	Server.RegisterActionable("/api/clusters/create", serverutils.MakeServerRoute(cluster.Create))


	return Server
}
