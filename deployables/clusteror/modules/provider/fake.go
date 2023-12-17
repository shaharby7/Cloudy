package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/shaharby7/Cloudy/deployables/clusteror/types"
	"github.com/shaharby7/Cloudy/pkg/common/serverutils"
)

type Fake struct {
	ProviderCode types.ProviderCode
}

func (fake *Fake) CreateMachine(ctx context.Context, options *CreateMachineOptions) (*CreateMachineResults, error) {
	result := &CreateMachineResults{}
	resp, err := post[FakeproviderCreateMachineInput, FakeproviderCreateMachineOutput](
		ctx, "/api/vms/create-vm",
		&FakeproviderCreateMachineInput{
			User_data_b64:  options.Specs.User_data_b64,
			IpAllocationId: options.IpAllocationCode,
		},
	)
	if err != nil {
		return result, err
	}
	result.MachineProviderCode = (*types.MachineProviderCode)(&resp.Machine_id)
	return result, nil
}

func (fake *Fake) TerminateMachine(ctx context.Context, options *TerminateMachineOptions) (*TerminateMachineResults, error) {
	result := &TerminateMachineResults{}
	_, err := post[FakeproviderDeleteMachineInput, FakeproviderDeleteMachineOutput](
		ctx, "/api/vms/delete-vm", &FakeproviderDeleteMachineInput{Machine_id: string(*options.MachineProviderCode)},
	)
	return result, err
}

func (fake *Fake) AllocatePublicIp(ctx context.Context, options *AllocatePublicIpOptions) (*AllocatePublicIpResults, error) {
	output, err := post[FakeproviderAllocateIpInput, FakeproviderAllocateIpOutput](
		ctx, "/api/network/allocate-ip", nil,
	)
	if err != nil {
		return nil, err
	}
	result := &AllocatePublicIpResults{
		Ip:   output.Ip,
		Code: output.IpAllocationId,
	}
	return result, err
}

func (fake *Fake) Identify() *ProviderCommonData {
	return &ProviderCommonData{ProviderCode: fake.ProviderCode}
}

func post[RequestBody any, ResponseData any](ctx context.Context, path string, requestBody *RequestBody) (*ResponseData, error) {
	address := os.Getenv("FAKEPROVIDER_ADDRESS")
	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(
		fmt.Sprintf("%s%s", address, path),
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	var parsed serverutils.APIResponse[ResponseData]
	err = json.Unmarshal(respBody, &parsed)
	if err != nil {
		return nil, err
	}
	if !parsed.Success {
		return nil, errors.New(parsed.Error)
	}
	return &parsed.Data, nil
}

type FakeproviderCreateMachineInput struct {
	User_data_b64  string `json:"user_data_b64"`
	IpAllocationId string `json:"ip_allocation_id"`
}

type FakeproviderCreateMachineOutput struct {
	Machine_id string `json:"machine_id"`
}

type FakeproviderDeleteMachineInput struct {
	Machine_id string `json:"machine_id"`
}
type FakeproviderDeleteMachineOutput struct {
}

type FakeproviderAllocateIpInput struct {
}

type FakeproviderAllocateIpOutput struct {
	Ip             string `json:"ip"`
	IpAllocationId string `json:"ip_allocation_id"`
}
