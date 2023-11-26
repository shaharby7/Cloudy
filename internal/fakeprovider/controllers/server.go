package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/shaharby7/Cloudy/internal/fakeprovider/modules/vms"

	"github.com/shaharby7/Cloudy/internal/fakeprovider/interfaces"
	"github.com/shaharby7/Cloudy/pkg/deployable/controllable"
)

func genericInputGenerator[InputType any](ctx context.Context, input controllable.TServerInput) (*InputType, error) {
	defer input.Body.Close()
	body, err := io.ReadAll(input.Body)
	if err != nil {
		return nil, err
	}
	var options InputType
	if err := json.Unmarshal([]byte(body), &options); err != nil {
		return nil, err
	}
	return &options, nil
}

func genericOutputGenerator[OutputType any](ctx context.Context, result *OutputType) (controllable.TServerOutput, error) {
	body := &interfaces.APIResponse[OutputType]{Success: true, Data: *result}
	respData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return &controllable.Response{
		Data: string(respData),
	}, nil
}

func genericErrorHandler(ctx context.Context, err error) (controllable.TServerOutput, error) {
	fmt.Println(err)
	body := &interfaces.APIErrorResponse{Success: false, Error: err.Error()}
	respData, parseErr := json.Marshal(body)
	if parseErr != nil {
		return &controllable.Response{Data: parseErr.Error(), StatusCode: 500}, nil
	}
	return &controllable.Response{Data: string(respData), StatusCode: 500}, nil
}

func GenerateServer() controllable.Controllable {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	Server := controllable.NewHttpServerControllable(
		"Fakeprovider server",
		*controllable.NewServerControllableConfig(port),
	)

	CreateVmAction := controllable.NewHttpServerActionable(
		vms.CreateVM,
		genericInputGenerator[vms.CreateVmOptions],
		genericOutputGenerator[vms.Result],
		genericErrorHandler,
	)
	Server.RegisterActionable("/api/create-vm", CreateVmAction)

	DeleteVmAction := controllable.NewHttpServerActionable(
		vms.DeleteVM,
		genericInputGenerator[vms.DeleteVmOptions],
		genericOutputGenerator[vms.Result],
		genericErrorHandler,
	)
	Server.RegisterActionable("/api/delete-vm", DeleteVmAction)

	GetVmAction := controllable.NewHttpServerActionable(
		vms.GetVM,
		genericInputGenerator[vms.GetVmOptions],
		genericOutputGenerator[vms.GetVmResults],
		genericErrorHandler,
	)
	Server.RegisterActionable("/api/get-vm", GetVmAction)

	return Server
}
