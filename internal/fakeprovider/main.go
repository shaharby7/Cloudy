package fakeprovider

import (
	"encoding/json"
	"fmt"
	"github.com/nahojer/httprouter"
	"net/http"
	"log"
)

type HardwareParams struct {
	MetalType string
	Image     string
}

func initiateHardware(hardwareParams *HardwareParams) {

}

func InitFakeProvider() {

	router := httprouter.New()

	initiateHardwareHandler := func(responseWriter http.ResponseWriter, request *http.Request) {
		var requestedHardwareParams HardwareParams
		err := json.NewDecoder(request.Body).Decode(&requestedHardwareParams)

		if err != nil {
			log.Print(err.Error())
			http.Error(responseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		fmt.Printf("Requested metal %s with image %s\n", requestedHardwareParams.MetalType, requestedHardwareParams.Image)
		initiateHardware(&requestedHardwareParams)
	}
	router.HandleFunc("POST", "/initiateHardware", initiateHardwareHandler)

	http.ListenAndServe(":3333", router)
}
