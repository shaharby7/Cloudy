package controllable

import (
	"context"
	"errors"
	"fmt"
	"github.com/nahojer/httprouter"
	"net/http"
	"sync"
)

type HttpServerControllerConfig struct{ Port int }
type ServerInput = []byte
type ServerOutput = []byte
type ServerActionConfig = struct{ Method string }

type IHttpServerActions = IActions[*ServerInput, *ServerOutput, ServerActionConfig]

type HttpServerControllerAction struct {
	RunAction     func(ctx context.Context, in *ServerInput) (out *ServerOutput, err error)
	OnActionError func(error) error
	Config        ServerActionConfig
}

func (httpAction *HttpServerControllerAction) GetConfig() ServerActionConfig {
	return httpAction.Config
}

func NewHttpServerAction(RunAction RunAction[*ServerInput, *ServerOutput, ServerActionConfig], OnActionError OnActionError, Config ServerActionConfig) *Action[*ServerInput, *ServerOutput, ServerActionConfig] {
	return &Action[*ServerInput, *ServerOutput, ServerActionConfig]{
		RunAction:     RunAction,
		OnActionError: OnActionError,
		Config:        Config,
	}
}

type HttpServerController struct {
	*ControllerBase[*ServerInput, *ServerOutput, ServerActionConfig]
	Config *HttpServerControllerConfig
	Router *httprouter.Router
}

func NewHttpServerController(ctx context.Context, actions IHttpServerActions, config *HttpServerControllerConfig) *HttpServerController {
	router := httprouter.New()
	controllerBase := NewControllerBase(ctx, "HttpServer", actions)
	return &HttpServerController{controllerBase, config, router}
}

func (controller *HttpServerController) Start(ctx context.Context, wg *sync.WaitGroup) error {
	for actionName := range controller.Actions {
		handler := func(responseWriter http.ResponseWriter, request *http.Request) {
			out, _ := controller.ControllerBase.Execute(ctx, actionName, &ServerInput{})
			responseWriter.Write(*out)
		}
		controller.Router.HandleFunc("post", actionName, handler)
	}
	port := fmt.Sprintf(":%d", controller.Config.Port)
	go http.ListenAndServe(port, controller.Router)
	fmt.Printf("Listening on port %s\n", port)
	return nil
}

func (listenerController *HttpServerController) OnError(err error) error {
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server closed")
	}
	fmt.Println("Could not initiate server")
	fmt.Println(err)
	return err
}
