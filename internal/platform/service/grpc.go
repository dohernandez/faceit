package service

import (
	"context"
	"github.com/bool64/ctxd"
	api "github.com/dohernandez/faceit/internal/platform/service/pb"
)

// FaceitServiceDeps holds the dependencies for the FaceitService.
type FaceitServiceDeps interface {
	Logger() ctxd.Logger
	GRPCAddr() string
}

// FaceitService is the gRPC service.
type FaceitService struct {
	// Uncomment this line once the grpc files were generated into the proto package.
	// UnimplementedFaceitServiceServer must be embedded to have forward compatible implementations.
	api.UnimplementedFaceitServiceServer

	deps FaceitServiceDeps
}

// NewFaceitService creates a new FaceitService.
func NewFaceitService(deps FaceitServiceDeps) *FaceitService {
	return &FaceitService{
		deps: deps,
	}
}

// AddUser add new user.
//
// Receives a request with user data. Responses whether the user was added successfully or not.
func (s *FaceitService) AddUser(ctx context.Context, req *api.AddUserRequest) (*api.AddUserResponse, error) {
	return nil, nil
}
