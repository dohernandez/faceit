package faceit_test

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"

	"github.com/bool64/ctxd"
	"github.com/dohernandez/faceit/internal/domain/model"
	"github.com/dohernandez/faceit/internal/platform/app"
	"github.com/dohernandez/faceit/internal/platform/config"
	"github.com/dohernandez/faceit/internal/platform/storage"
	service "github.com/dohernandez/go-grpc-service"
	sapp "github.com/dohernandez/go-grpc-service/app"
	sconfig "github.com/dohernandez/go-grpc-service/config"
	"github.com/dohernandez/go-grpc-service/must"
	"github.com/dohernandez/servers"
	"github.com/google/uuid"
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
				Name:         "get list by country",
				Uri:          "/v1/users?country=UK",
				ResponseCode: http.StatusOK,
				Data: map[string]any{
					storage.UserTable: model.User{
						ID: uuid.MustParse("26ef0140-c436-4838-a271-32652c72f6f2"),
						UserState: model.UserState{
							FirstName:    "Alice",
							LastName:     "Bob",
							PasswordHash: "f6b7e19e0d867de6c0391879050e8297165728d89d7c4e9e8839972b356c4d9d",
							Email:        "alice@bob.com",
							Country:      "UK",
						},
					},
				},
			},
		},
	})
}
