package fakeprovider

import (
	"os"
	"strconv"

	"github.com/shaharby7/Cloudy/internal/fakeprovider/app"
)

var config = &deployable.DeployableConfig{
	ProjectName:          "fakeprovider",
	RequiredEnvVariables: []string{"PORT"},
}

var port, _ = strconv.Atoi(os.Getenv("PORT"))

var fakeProviderHTTPListener = &deployable.HTTPServerController{
	Config: &deployable.HTTPServerControllerConfig{Port: port},
	Router: app.FakeProviderHTTPRouter,
}

var FakeProviderDeployable = deployable.NewDeployable(
	*config,
	[]deployable.IController{fakeProviderHTTPListener},
)
