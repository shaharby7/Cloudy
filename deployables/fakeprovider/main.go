package fakeprovider

import (
	"context"
	"fmt"

	"github.com/shaharby7/Cloudy/deployables/fakeprovider/controllers"
	"github.com/shaharby7/Cloudy/deployables/fakeprovider/databases"
	"github.com/shaharby7/Cloudy/deployables/fakeprovider/services"
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
			ProjectName:          "fakeprovider",
			RequiredEnvVariables: []string{"PORT", "REDIS_DOMAIN", "REDIS_PORT", "SENSOR_ADDRESS"},
		},
		*myLoggable,
		func(ctx context.Context, err error) { fmt.Println(err) },
		"/home/shahar/Projects/Cloudy/local/fakeprovider.env", // todo - infer from deployer
	)

	initiateClientsOrExit()

	dep.RegisterControllable(controllers.GenerateServer())

	return dep
}

func initiateClientsOrExit() {
	var err error
	err = databases.InitiateRedisClient()
	if err != nil {
		err = fmt.Errorf("could not connect to redis: %s", err)
		panic(err.Error())
	}
	err = services.InitiateK8SClient()
	if err != nil {
		err = fmt.Errorf("could not initiate k8s client: %s", err)
		panic(err.Error())
	}
}
