package service

import (
	"github.com/bool64/ctxd"
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
	// api.UnimplementedFaceitServiceServer

	deps FaceitServiceDeps
}

// NewFaceitService creates a new FaceitService.
func NewFaceitService(deps FaceitServiceDeps) *FaceitService {
	return &FaceitService{
		deps: deps,
	}
}

/*
// PostFuncName ... .
func (s *FaceitService) PostFuncName(ctx context.Context, req interface{}) (interface{}, error) {
	return nil, nil
}
*/
