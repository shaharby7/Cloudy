package deployable

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/joho/godotenv"

	"github.com/shaharby7/Cloudy/pkg/deployable/constants"
	"github.com/shaharby7/Cloudy/pkg/deployable/controllable"
	"github.com/shaharby7/Cloudy/pkg/deployable/loggable"
)

type Environments string

const (
	LOCAL Environments = "LOCAL"
)

type Deployable struct {
	ctx           context.Context
	ENV           Environments
	ProjectName   string
	Logger        loggable.Loggable
	Controllables []controllable.Controllable
	OnError       func(ctx context.Context, err error)
}

func (deployable *Deployable) Start(parentContext context.Context) {
	var deployableWaitGroup sync.WaitGroup
	ctx := deployable.initiateDeployableContext(parentContext)
	for _, controller := range deployable.Controllables {
		deployableWaitGroup.Add(1)
		err := controller.Start(ctx, &deployableWaitGroup)
		if err != nil {
			deployable.OnError(ctx, err)
		}
	}
	fmt.Printf("Project %s is running", deployable.ProjectName)
	deployableWaitGroup.Wait()
}

func (deployable *Deployable) initiateDeployableContext(parentContext context.Context) context.Context {
	ctx := context.WithValue(parentContext, constants.LOGGER_REF, &deployable.Logger)
	ctx = context.WithValue(ctx, constants.DEPLOYABLE_REF, &deployable)
	deployable.ctx = ctx
	return ctx
}

type DeployableConfig struct {
	ProjectName          string
	RequiredEnvVariables []string
}

func verifyEnvVariables(requiredEnvVariables *[]string) error {
	for _, name := range *requiredEnvVariables {
		val := os.Getenv(name)
		if "" == val {
			return errors.New(fmt.Sprintf("required ENV variable %s is not found", name))
		}
	}
	return nil
}

func NewDeployable(
	deployableConfig DeployableConfig,
	controllers []controllable.Controllable,
	logger loggable.Loggable,
	onError func(ctx context.Context, err error),
) (*Deployable, error) {
	ENV := os.Getenv("ENV")
	if "" == ENV {
		ENV = "LOCAL"
	}
	deployable := &Deployable{
		ENV:           Environments(ENV),
		ProjectName:   deployableConfig.ProjectName,
		Controllables: controllers,
		Logger:        logger,
		OnError:       onError,
	}
	if deployable.ENV == LOCAL {
		_, b, _, _ := runtime.Caller(0)
		basePath := filepath.Dir(b)
		godotenv.Load(fmt.Sprintf("%s/local/%s.env", basePath, deployable.ProjectName))
	}
	err := verifyEnvVariables(&deployableConfig.RequiredEnvVariables)
	return deployable, err
}
