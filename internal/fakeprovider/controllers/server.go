package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/shaharby7/Cloudy/internal/fakeprovider/modules/vms"
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
	response, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	return &controllable.Response{Data: string(response)}, nil
}

func genericErrorHandler(ctx context.Context, err error) (controllable.TServerOutput, error) {
	return &controllable.Response{Data: fmt.Sprintf("Something went wrong: %s", err), StatusCode: 500}, nil
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

	return Server
}
