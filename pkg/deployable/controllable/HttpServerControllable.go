package controllable

import (
	"context"
	"fmt"
	"net/http"

	"github.com/shaharby7/Cloudy/pkg/deployable/constants"
)

type TServerInput = []byte
type TServerOutput = []byte
type ServerControllableConfig struct{ Port int }

func NewServerControllableConfig(Port int) *ServerControllableConfig {
	return &ServerControllableConfig{Port: Port}
}

type HttpServerActionables = TActionables[TServerInput, TServerOutput]

type SHttpServerActionable[TActionableInput any, TActionableOutput any] struct {
	SActionable[TServerInput, TServerOutput, TActionableInput, TActionableOutput]
}

func NewHttpServerActionable[TActionableInput any, TActionableOutput any](
	RunActionable FInnerRunActionable[TActionableInput, TActionableOutput],
	MarshalActionableInput FMarshalActionableInput[TServerInput, TActionableInput],
	MarshalControllableOutput FMarshalControllableOutput[TActionableOutput, TServerOutput],
	OnError func(context.Context, error) (TServerOutput, error),
) *SHttpServerActionable[TActionableInput, TActionableOutput] {
	return &SHttpServerActionable[TActionableInput, TActionableOutput]{
		*NewActionable(
			RunActionable, MarshalActionableInput, MarshalControllableOutput, OnError,
		),
	}
}

type SHttpServerControllable struct {
	SControllable[TServerInput, TServerOutput]
	config ServerControllableConfig
	server *http.Server
}

func NewHttpServerControllable(
	ControllerName string,
	actionables HttpServerActionables,
	config ServerControllableConfig,
) *SHttpServerControllable {
	var server *http.Server
	return &SHttpServerControllable{
		*NewControllable(
			constants.HTTP_SERVER,
			ControllerName,
			func(execute func(actionableName string, input TServerInput) (TServerOutput, error)) error {
				mux := http.NewServeMux()
				mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
					execute("hi", []byte{}) // TODO
				})
				server = &http.Server{
					Addr:    fmt.Sprintf(":%d", config.Port),
					Handler: mux,
				}
				go server.ListenAndServe() // TODO: handle errors and cancellations
				return nil
			},
			func(ctx context.Context, err error) ([]byte, error) { return []byte{}, nil }, //TODO
			actionables,
		),
		config,
		server,
	}
}
