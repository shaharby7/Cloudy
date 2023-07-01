package controllable

import (
	"context"
	"sync"

	"github.com/google/uuid"
)

type IController interface {
	Start(ctx context.Context, waitGroup *sync.WaitGroup) error
	OnError(error) error
	Execute(ctx context.Context, actionName string, in any) (out any, err error)
}

type RunAction[InType any, OutType any, ActionConfig any] func(ctx context.Context, in InType) (out OutType, err error)
type OnActionError func(error) error

type IAction[InType any, OutType any, ActionConfig any] interface {
	RunAction[InType, OutType, ActionConfig]
	OnActionError
	GetConfig() ActionConfig
}
type Action[InType any, OutType any, ActionConfig any] struct {
	RunAction     func(ctx context.Context, in InType) (out OutType, err error)
	OnActionError func(error) error
	Config        ActionConfig
}

func (action *Action[InType, OutType, ActionConfig]) GetConfig() ActionConfig {
	return action.Config
}

func NewAction[InType any, OutType any, ActionConfig any](RunAction RunAction[InType, OutType, ActionConfig], OnActionError OnActionError, Config ActionConfig) *Action[InType, OutType, ActionConfig] {
	return &Action[InType, OutType, ActionConfig]{
		RunAction:     RunAction,
		OnActionError: OnActionError,
		Config:        Config,
	}
}

type IActions[InType any, OutType any, ActionConfig any] map[string]*Action[InType, OutType, ActionConfig]

type ControllerBase[InType any, OutType any, ActionConfig any] struct {
	Ctx            context.Context
	ControllerType string
	Actions        IActions[InType, OutType, ActionConfig]
}

func NewControllerBase[InType any, OutType any, ActionConfig any](ctx context.Context, controllerType string, actions IActions[InType, OutType, ActionConfig]) *ControllerBase[InType, OutType, ActionConfig] {
	ctx = context.WithValue(ctx, "controllerType", "")
	return &ControllerBase[InType, OutType, ActionConfig]{
		ctx, controllerType, actions,
	}
}

func (controller *ControllerBase[InType, OutType, ActionConfig]) Execute(ctx context.Context, actionName string, in InType) (out OutType, err error) {
	ctx = context.WithValue(ctx, "contextId", uuid.NewString())
	return controller.Actions[actionName].RunAction(ctx, in)
}
