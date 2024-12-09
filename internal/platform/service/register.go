package service

import (
	"context"

	api "github.com/dohernandez/faceit/internal/platform/service/pb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// RegisterService registers the service implementation to grpc service.
func (s *FaceitService) RegisterService(r grpc.ServiceRegistrar) {
	// register grpc service
	api.RegisterFaceitServiceServer(r, s)
}

// RegisterServiceHandler registers the service implementation to mux.
func (s *FaceitService) RegisterServiceHandler(mux *runtime.ServeMux) error {
	// register rest service
	return api.RegisterFaceitServiceHandlerFromEndpoint(context.Background(), mux, s.deps.GRPCAddr(), []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
}
