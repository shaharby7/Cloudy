package clusteror

import (
	"context"
	"fmt"

	"github.com/shaharby7/Cloudy/deployables/clusteror/controllers"
	"github.com/shaharby7/Cloudy/pkg/deployable"
	"github.com/shaharby7/Cloudy/pkg/deployable/loggable"
)

func Initiate() *deployable.Deployable {
	myLoggable := &loggable.Loggable{
		Targets:    map[string]loggable.ITarget{"console": loggable.NewConsoleTarget()},
		EventTypes: map[string]struct{ Targets []string }{"info": {Targets: []string{"console"}}},
		Events:     map[string]struct{ EventTypes []string }{"my-log": {EventTypes: []string{"info"}}},
		OnError:    func(err error) { fmt.Println(err) },
	}

	dep, _ := deployable.NewDeployable(
		deployable.DeployableConfig{
			ProjectName:          "clusteror",
			RequiredEnvVariables: []string{"PORT", "FAKEPROVIDER_ADDRESS"},
		},
		*myLoggable,
		func(ctx context.Context, err error) { fmt.Println(err) },
		"/home/shahar/Cloudy/Cloudy/local/clusteror.env", // todo - infer from deployer
	)

	dep.RegisterControllable(controllers.GenerateServer())

	return dep
}
