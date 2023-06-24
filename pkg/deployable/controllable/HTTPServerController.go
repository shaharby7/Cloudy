package controllable

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"github.com/nahojer/httprouter"
)

type HTTPServerControllerConfig struct{ Port int }
type ServerInput = []byte
type ServerOutput = []byte
type ServerActionConfig = struct{ Method string }

type HTTPServerActions = IActions[*ServerInput, *ServerOutput, ServerActionConfig]

type HTTPServerController struct {
	ControllerBase[*ServerInput, *ServerOutput, ServerActionConfig]
	Config *HTTPServerControllerConfig
	Router *httprouter.Router
}

func NewHTTPServerController(ctx context.Context, actions HTTPServerActions, config *HTTPServerControllerConfig) *HTTPServerController {
	router := httprouter.New()
	controllerBase := ControllerBase[*ServerInput, *ServerOutput, ServerActionConfig]{ctx, actions}
	return &HTTPServerController{controllerBase, config, router}
}

func (controller *HTTPServerController) Start(ctx context.Context, wg *sync.WaitGroup) error {
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

func (listenerController *HTTPServerController) OnError(err error) error {
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server closed")
	}
	fmt.Println("Could not initiate server")
	fmt.Println(err)
	return err
}
