package service

import (
	"context"
	"errors"

	"github.com/bool64/ctxd"
	"github.com/dohernandez/faceit/internal/domain/model"
	api "github.com/dohernandez/faceit/internal/platform/service/pb"
	"github.com/dohernandez/go-grpc-service/database"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AddUser defines the use case to add a user.
type AddUser interface {
	AddUser(ctx context.Context, us model.UserState) (model.UserID, error)
}

// FaceitServiceDeps holds the dependencies for the FaceitService.
type FaceitServiceDeps interface {
	Logger() ctxd.Logger
	GRPCAddr() string

	AddUser() AddUser
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
	ctx = ctxd.AddFields(ctx, "service", "FaceitService")

	us := model.UserState{
		FirstName:    req.GetFirstName(),
		LastName:     req.GetLastName(),
		Nickname:     req.GetNickname(),
		PasswordHash: req.GetPasswordHash(),
		Email:        req.GetEmail(),
		Country:      req.GetCountry(),
	}

	id, err := s.deps.AddUser().AddUser(ctx, us)
	if err != nil {
		if errors.Is(err, database.ErrAlreadyExists) {
			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "409")) //nolint:errcheck

			return nil, status.Error(codes.AlreadyExists, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "201")) //nolint:errcheck

	return &api.AddUserResponse{
		Id: id.String(),
	}, nil
}
