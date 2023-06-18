package deployable

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/joho/godotenv"
)

type Environments string

const (
	LOCAL Environments = "LOCAL"
)

type Deployable struct {
	ENV         Environments
	ProjectName string
	Controllers []IController
}

func (d *Deployable) Start() {
	var deployableWaitGroup sync.WaitGroup
	for _, controller := range d.Controllers {
		deployableWaitGroup.Add(1)
		err := controller.Start(&deployableWaitGroup)
		if err != nil {
			d.OnError(err)
		}
	}
	fmt.Printf("Project %s is running", d.ProjectName)
	deployableWaitGroup.Wait()
}

func (d *Deployable) OnError(error error) {
	fmt.Printf("Project %s could not be initialized", d.ProjectName)
	fmt.Println(error)
	os.Exit(1)
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

func NewDeployable(deployableConfig DeployableConfig, controllers []IController) *Deployable {
	ENV := os.Getenv("ENV")
	if "" == ENV {
		ENV = "LOCAL"
	}
	deployable := &Deployable{
		ENV:         Environments(ENV),
		ProjectName: deployableConfig.ProjectName,
		Controllers: controllers,
	}
	if deployable.ENV == LOCAL {
		_, b, _, _ := runtime.Caller(0)
		basePath := filepath.Dir(b)
		godotenv.Load(fmt.Sprintf("%s/local/%s.env", basePath, deployable.ProjectName))
	}
	err := verifyEnvVariables(&deployableConfig.RequiredEnvVariables)
	if err != nil {
		deployable.OnError(err)
	}
	return deployable
}
