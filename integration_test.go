package faceit_test

import (
	"context"
	"fmt"
	"io"
	"net"
	"testing"

	"github.com/bool64/ctxd"
	"github.com/cucumber/godog"
	"github.com/dohernandez/faceit/internal/domain/model"
	"github.com/dohernandez/faceit/internal/platform/app"
	"github.com/dohernandez/faceit/internal/platform/config"
	"github.com/dohernandez/faceit/internal/platform/storage"
	service "github.com/dohernandez/go-grpc-service"
	sapp "github.com/dohernandez/go-grpc-service/app"
	sconfig "github.com/dohernandez/go-grpc-service/config"
	"github.com/dohernandez/go-grpc-service/must"
	"github.com/dohernandez/servers"
)

func TestIntegration(t *testing.T) {
	ctx := context.Background()

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	// load configurations
	err := sconfig.WithEnvFiles(".env.integration-test")
	must.NotFail(ctxd.WrapError(ctx, err, "failed to load env from .env.integration-test"))

	var cfg config.Config

	err = sconfig.LoadConfig(&cfg)
	must.NotFail(ctxd.WrapError(ctx, err, "failed to load configurations"))

	cfg.Environment = "test"
	cfg.Logger.Output = io.Discard

	// initialize listeners
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.AppGRPCPort))
	must.NotFail(ctxd.WrapError(ctx, err, "failed to init GRPC service listener"))

	restTListener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.AppRESTPort))
	must.NotFail(ctxd.WrapError(ctx, err, "failed to init REST service listener"))

	healthListener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.AppHealthPort))
	must.NotFail(ctxd.WrapError(ctx, err, "failed to init health service listener"))

	// initialize locator
	deps, err := app.NewServiceLocator(
		&cfg,
		sapp.WithGRPC(
			servers.WithListener(grpcListener, true),
		),
		sapp.WithGRPCRest(
			servers.WithAddrAssigned(),
			servers.WithListener(restTListener, true),
		),
		sapp.WithHealthCheck(
			servers.WithAddrAssigned(),
			servers.WithListener(healthListener, true),
		),
	)
	must.NotFail(ctxd.WrapError(ctx, err, "failed to init service locator"))

	service.RunFeatures(t, ctx, &service.FeaturesConfig{
		FeaturePath: "features",
		Locator:     deps.Locator,
		FeatureContextFunc: func(_ *testing.T, _ *godog.ScenarioContext) {
			// Add step definitions
		},
		Tables: map[string]any{
			storage.UserTable: new(model.User),
		},
	})
}
