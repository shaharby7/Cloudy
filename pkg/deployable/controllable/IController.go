package controllable

import (
	"context"
	"sync"
)

type IController interface {
	Start(ctx context.Context, waitGroup *sync.WaitGroup) error
	OnError(error) error
	Execute(ctx context.Context, actionName string, in any) (out any, err error)
}

type IAction[InType any, OutType any, ActionConfig any] interface {
	RunAction(ctx context.Context, in InType) (out OutType, err error)
	OnActionError(error) error
	GetConfig() ActionConfig
}

type IActions[InType any, OutType any, ActionConfig any] map[string]IAction[InType, OutType, ActionConfig]

type ControllerBase[InType any, OutType any, ActionConfig any] struct {
	Ctx     context.Context
	Actions IActions[InType, OutType, ActionConfig]
}

func NewControllerBase[InType any, OutType any, ActionConfig any](ctx context.Context, actions IActions[InType, OutType, ActionConfig]) *ControllerBase[InType, OutType, ActionConfig] {
	// ctx = context.WithValue(ctx, "controllerType", 1)
	return &ControllerBase[InType, OutType, ActionConfig]{
		ctx, actions,
	}
}

func (controller *ControllerBase[InType, OutType, ActionConfig]) Execute(ctx context.Context, actionName string, in InType) (out OutType, err error) {
	return controller.Actions[actionName].RunAction(ctx, in)
}
