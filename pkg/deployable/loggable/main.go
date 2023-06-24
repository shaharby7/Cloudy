package loggable

import (
	"context"
	"errors"
	"fmt"

	"github.com/shaharby7/Cloudy/pkg/deployable/constants"
	"github.com/shaharby7/Cloudy/pkg/deployable/helpers"
)

type ITarget interface {
	Log(ctx context.Context, eventName string, data *any) error
	Get(ctx context.Context) (any, error)
}

type Loggable struct {
	Targets    map[string]ITarget
	EventTypes map[string]struct{ Targets []string }
	Events     map[string]struct{ EventTypes []string }
	OnError    func(error)
}

func Log(ctx context.Context, eventName string, data *any) {
	go func() {
		ok := helpers.VerifyDeployableContext(ctx)
		if !ok {
			panic("Cannot log context that was not produced by deployable")
		}
		l := ctx.Value(constants.LOGGER_ATTR_NAME).(*Loggable)
		for _, eventTypeName := range l.Events[eventName].EventTypes {
			for _, targetName := range l.EventTypes[eventTypeName].Targets {
				err := l.Targets[targetName].Log(ctx, eventTypeName, data)
				if nil != err {
					l.OnError(err)
				}
			}
		}
	}()
}

type ConsoleTarget struct{}

func (target *ConsoleTarget) log(eventName string, contextId string, data *any) {
	fmt.Printf("%s\t%s::::%s", contextId, eventName, *data)
}

func (target *ConsoleTarget) get(contextId string) (any, error) {
	return nil, errors.New("Cannot get from target ConsoleTarget")
}
