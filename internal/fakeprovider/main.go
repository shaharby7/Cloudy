package fakeprovider

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/shaharby7/Cloudy/pkg/deployable"
	"github.com/shaharby7/Cloudy/pkg/deployable/controllable"
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
		deployable.DeployableConfig{ProjectName: "fakeprovider", RequiredEnvVariables: []string{"PORT"}},
		*myLoggable,
		func(ctx context.Context, err error) { fmt.Println(err) },
		"/home/shahar/Projects/Cloudy/local/fakeprovider.env", // todo - infer from deployer
	)

	port, _ := strconv.Atoi(os.Getenv("PORT"))

	MyAction := controllable.NewHttpServerActionable(
		func(ctx context.Context, s string) (string, error) { return s, nil },
		func(ctx context.Context, b []byte) (string, error) { return "hi", nil },
		func(ctx context.Context, s string) ([]byte, error) { return []byte{}, nil },
		func(ctx context.Context, err error) (controllable.TServerOutput, error) { return []byte{}, nil },
	)

	MyControllable := controllable.NewHttpServerControllable(
		"shahar",
		*controllable.NewServerControllableConfig(port),
	)

	MyControllable.RegisterActionable("hi", MyAction)
	MyDeployable.RegisterControllable(MyControllable)

	return MyDeployable
}
