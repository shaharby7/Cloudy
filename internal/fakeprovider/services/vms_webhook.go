package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type CreateVmOptions struct {
	Machine_id    string `json:"machine_id"`
	User_data_b64 string `json:"user_data_b64"`
}

type DeleteVmOptions struct {
	Machine_id string `json:"machine_id"`
}

func post[RequestBody any](ctx context.Context, path string, requestBody *RequestBody) error {
	var address = os.Getenv("VMS_WEBHOOK_ADDRESS")
	body, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}
	resp, err := http.Post(
		fmt.Sprintf("%s%s", address, path),
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resp.Body.Close()
	var parsed string = string(respBody)
	if parsed != "success" {
		return fmt.Errorf("unsuccessful response from vms-webhook. expected \"success\" got: %s", parsed)
	}
	return nil
}

func SendCreateVM(ctx context.Context, machine_id string, user_data_b64 string) error {
	err := post(
		ctx,
		"/create-vm",
		&CreateVmOptions{
			Machine_id:    machine_id,
			User_data_b64: user_data_b64,
		},
	)
	return err
}

func SendDeleteVM(ctx context.Context, machine_id string) error {
	err := post(
		ctx,
		"/delete-vm",
		&DeleteVmOptions{
			Machine_id: machine_id,
		},
	)
	return err
}
