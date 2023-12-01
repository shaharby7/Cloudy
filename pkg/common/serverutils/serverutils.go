package serverutils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/shaharby7/Cloudy/pkg/deployable/controllable"
	"io"
)

type APIResponse[TResponseData any] struct {
	Success bool          `json:"success"`
	Data    TResponseData `json:"data"`
	Error   string        `json:"error"`
}

type APIErrorResponse = APIResponse[any]

func MakeServerRoute[OptionsType any, ResultType any](
	fn func(context.Context, *OptionsType) (*ResultType, error),
) *controllable.SHttpServerActionable[*OptionsType, *ResultType] {
	return controllable.NewHttpServerActionable(
		fn,
		genericInputGenerator[OptionsType],
		genericOutputGenerator[ResultType],
		genericErrorHandler,
	)
}

var defaultHeaders = map[string]string{"Content-Type": "application/json"}

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
	body := APIResponse[OutputType]{Success: true, Data: *result}
	respData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return &controllable.Response{
		Data: string(respData), Headers: defaultHeaders,
	}, nil
}

func genericErrorHandler(ctx context.Context, err error) (controllable.TServerOutput, error) {
	fmt.Println(err)
	body := APIErrorResponse{Success: false, Error: err.Error()}
	respData, parseErr := json.Marshal(body)
	if parseErr != nil {
		return &controllable.Response{Data: parseErr.Error(), Headers: defaultHeaders}, nil
	}
	return &controllable.Response{Data: string(respData), StatusCode: 500, Headers: defaultHeaders}, nil
}
