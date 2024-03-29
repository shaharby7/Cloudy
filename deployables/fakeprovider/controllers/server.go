package controllers

import (
	"os"
	"strconv"

	"github.com/shaharby7/Cloudy/deployables/fakeprovider/modules/network"
	"github.com/shaharby7/Cloudy/deployables/fakeprovider/modules/vms"

	"github.com/shaharby7/Cloudy/pkg/common/serverutils"
	"github.com/shaharby7/Cloudy/pkg/deployable/controllable"
)

func GenerateServer() controllable.Controllable {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	Server := controllable.NewHttpServerControllable(
		"Fakeprovider server",
		*controllable.NewServerControllableConfig(port),
	)

	Server.RegisterActionable("/api/vms/create-vm", serverutils.MakeServerRoute(vms.CreateVM))
	Server.RegisterActionable("/api/vms/delete-vm", serverutils.MakeServerRoute(vms.DeleteVM))
	Server.RegisterActionable("/api/vms/get-vm", serverutils.MakeServerRoute(vms.GetVM))

	Server.RegisterActionable("/api/network/allocate-ip", serverutils.MakeServerRoute(network.AllocatePublicIp))

	return Server
}
