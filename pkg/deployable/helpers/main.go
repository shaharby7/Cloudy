package helpers

import (
	"context"

	"github.com/shaharby7/Cloudy/pkg/deployable/constants"
)

func VerifyDeployableContext(ctx context.Context) bool {
	_, ok := ctx.Value(constants.IS_DEPLOYABLE_CTX).(bool)
	return ok
}