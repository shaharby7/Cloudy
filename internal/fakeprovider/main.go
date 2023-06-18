package fakeprovider

import (
	"os"
	"strconv"

	"github.com/shaharby7/Cloudy/internal/fakeprovider/app"
	"github.com/shaharby7/Cloudy/pkg/deployable"
)

var config = &deployable.DeployableConfig{
	ProjectName:          "fakeprovider",
	RequiredEnvVariables: []string{"PORT"},
}

var port, _ = strconv.Atoi(os.Getenv("PORT"))

var fakeProviderHTTPListener = &deployable.HTTPListenerController{
	Config: &deployable.HTTPListenerControllerConfig{Port: port},
	Router: app.FakeProviderHTTPRouter,
}

var FakeProviderDeployable = deployable.NewDeployable(
	*config,
	[]deployable.IController{fakeProviderHTTPListener},
)
