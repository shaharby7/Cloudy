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

var config = &deployable.DeployableConfig{
	ProjectName:          "fakeprovider",
	RequiredEnvVariables: []string{"PORT"},
}

var port, _ = strconv.Atoi(os.Getenv("PORT"))

var myLoggable = &loggable.Loggable{
	Targets:    map[string]loggable.ITarget{"console": loggable.NewConsoleTarget()},
	EventTypes: map[string]struct{ Targets []string }{"info": {Targets: []string{"console"}}},
	Events:     map[string]struct{ EventTypes []string }{"my-log": {EventTypes: []string{"info"}}},
	OnError:    func(err error) { fmt.Println(err) },
}

var MyAction = controllable.NewHttpServerActionable(
	func(ctx context.Context, s string) (string, error) { return s, nil },
	func(ctx context.Context, b []byte) (string, error) { return "hi", nil },
	func(ctx context.Context, s string) ([]byte, error) { return []byte{}, nil },
	func(ctx context.Context, err error) (controllable.TServerOutput, error) { return []byte{}, nil },
)

var MyControllable = controllable.NewHttpServerControllable(
	"shahar",
	controllable.HttpServerActionables{"hi": MyAction},
	*controllable.NewServerControllableConfig(port),
)

var MyDeployable, _ = deployable.NewDeployable(
	deployable.DeployableConfig{ProjectName: "shahar-deployable", RequiredEnvVariables: []string{"PORT"}},
	[]controllable.Controllable{MyControllable},
	*myLoggable,
	func(ctx context.Context, err error) { fmt.Println(err) },
)

// var MyAction = controllable.NewHttpServerAction(
// 	func(ctx context.Context, in string) (out string, err error) { return fmt.Sprintf("hello %s", in), nil },
// 	func(input *[]byte) (string, error) { return "shahar", nil },
// 	func(output string) (*[]byte, error) { return &[]byte{}, nil },
// 	func(err error) error { return nil },
// 	controllable.ServerActionConfig{Method: "post"},
// )

// var MyServer = controllable.NewHttpServerController(
// 	context.Background(),
// 	controllable.IHttpServerActions{"shahar": controllable.HttpServerAction[any, any](MyAction)},
// 	&controllable.HttpServerControllerConfig{Port: port},
// )
