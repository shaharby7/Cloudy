package deployable

import (
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/nahojer/httprouter"
)

type IController interface {
	Start(waitGroup *sync.WaitGroup) error
	OnError(error) error
}

type Controller struct {
	IController
	Config any
}

type HTTPListenerControllerConfig struct{ Port int }

type HTTPListenerController struct {
	Config *HTTPListenerControllerConfig
	Router *httprouter.Router
}

func (listenerController *HTTPListenerController) Start(waitGroup *sync.WaitGroup) error {
	port := fmt.Sprintf(":%d", listenerController.Config.Port)
	go http.ListenAndServe(port, listenerController.Router)
	fmt.Printf("Listening on port %s\n", port)
	return nil
}

func (listenerController *HTTPListenerController) OnError(err error) error {
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server closed")
	}
	fmt.Println("Could not initiate server")
	fmt.Println(err)
	return err
}
