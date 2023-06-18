package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nahojer/httprouter"
	"github.com/terra-farm/go-virtualbox"
)

type HardwareParams struct {
	MetalType string
	Image     string
}

func initiateHardware(hardwareParams *HardwareParams) {
	virtualbox.New().Start()
}

var FakeProviderHTTPRouter = httprouter.New()

var initiateHardwareHandler = func(responseWriter http.ResponseWriter, request *http.Request) {
	var requestedHardwareParams HardwareParams
	err := json.NewDecoder(request.Body).Decode(&requestedHardwareParams)

	if err != nil {
		fmt.Println(err.Error())
		http.Error(responseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	fmt.Printf("Requested metal %s with image %s\n", requestedHardwareParams.MetalType, requestedHardwareParams.Image)
	initiateHardware(&requestedHardwareParams)
}

func init() {
	FakeProviderHTTPRouter.HandleFunc("POST", "api/hardware/initiate", initiateHardwareHandler)
}
