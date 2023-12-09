package machine

import "context"

func CreateAndRun(ctx context.Context, options *NewOptions) (*sMachine, error) {
	machine, err := New(ctx, options)
	if err != nil {
		return machine, err
	}
	_, err = machine.Create(ctx, nil)
	return machine, err
}