package fakeprovider

import (
	"context"
	"fmt"

	"github.com/shaharby7/Cloudy/internal/fakeprovider/controllers"
	"github.com/shaharby7/Cloudy/internal/fakeprovider/databases"
	"github.com/shaharby7/Cloudy/internal/fakeprovider/services"
	"github.com/shaharby7/Cloudy/pkg/deployable"
	"github.com/shaharby7/Cloudy/pkg/deployable/loggable"
)

func InstantiateFakeproviderDeployable() *deployable.Deployable {
	myLoggable := &loggable.Loggable{
		Targets:    map[string]loggable.ITarget{"console": loggable.NewConsoleTarget()},
		EventTypes: map[string]struct{ Targets []string }{"info": {Targets: []string{"console"}}},
		Events:     map[string]struct{ EventTypes []string }{"my-log": {EventTypes: []string{"info"}}},
		OnError:    func(err error) { fmt.Println(err) },
	}

	MyDeployable, _ := deployable.NewDeployable(
		deployable.DeployableConfig{
			ProjectName:          "fakeprovider",
			RequiredEnvVariables: []string{"PORT", "REDIS_DOMAIN", "REDIS_PORT", "SENSOR_ADDRESS"},
		},
		*myLoggable,
		func(ctx context.Context, err error) { fmt.Println(err) },
		"/home/shahar/Projects/Cloudy/local/fakeprovider.env", // todo - infer from deployer
	)

	initiateClientsOrExit()

	MyDeployable.RegisterControllable(controllers.GenerateServer())

	return MyDeployable
}

func initiateClientsOrExit() {
	var err error
	err = databases.InitiateRedisClient()
	if err != nil {
		err = fmt.Errorf("Could not connect to redis: %s", err)
		panic(err.Error())
	}
	err = services.InitiateK8SClient()
	if err != nil {
		err = fmt.Errorf("Could not initiate k8s client: %s", err)
		panic(err.Error())
	}
}
