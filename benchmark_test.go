package faceit_test

import (
	"context"
	"fmt"
	"io"
	"net"
	"testing"

	"github.com/bool64/ctxd"
	"github.com/dohernandez/faceit/internal/platform/app"
	"github.com/dohernandez/faceit/internal/platform/config"
	service "github.com/dohernandez/go-grpc-service"
	sapp "github.com/dohernandez/go-grpc-service/app"
	sconfig "github.com/dohernandez/go-grpc-service/config"
	"github.com/dohernandez/go-grpc-service/must"
	"github.com/dohernandez/servers"
)

func BenchmarkIntegration(b *testing.B) {
	ctx := context.Background()

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
	)
	must.NotFail(ctxd.WrapError(ctx, err, "failed to init service locator"))

	service.RunBenchmark(b, ctx, &service.BenchmarkConfig{
		Locator: deps.Locator,
		TestCases: []service.BenchmarkCases{
			{
				Name: "test case name",
				Uri:  "/uri",
				Data: nil, // map[string]any
			},
		},
	})
}