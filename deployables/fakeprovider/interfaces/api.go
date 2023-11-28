package interfaces

type APIResponse[TResponseData any] struct {
	Success bool          `json:"success"`
	Data    TResponseData `json:"data"`
	Error   string         `json:"error"`
}

type APIErrorResponse = APIResponse[any]
