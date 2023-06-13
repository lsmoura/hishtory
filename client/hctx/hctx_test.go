package hctx

import (
	"context"
	"testing"
)

func TestCtxConfig(t *testing.T) {
	t.Parallel()

	config := ClientConfig{
		UserSecret: "shhhh",
		DeviceId:   "some-id",
	}

	ctx := WithConf(context.Background(), config)

	ctxConfig := GetConf(ctx)

	if ctxConfig.UserSecret != config.UserSecret {
		t.Errorf("expected %s, got %s", config.UserSecret, ctxConfig.UserSecret)
	}
	if ctxConfig.DeviceId != config.DeviceId {
		t.Errorf("expected %s, got %s", config.DeviceId, ctxConfig.DeviceId)
	}
}
