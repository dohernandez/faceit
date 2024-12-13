package service

import (
	"context"
	"errors"
	"github.com/bufbuild/protovalidate-go"
	"github.com/dohernandez/faceit/internal/domain/model"
	api "github.com/dohernandez/faceit/internal/platform/service/pb"
	"github.com/dohernandez/go-grpc-service/database"
	"github.com/dohernandez/servers"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

// AddUser defines the use case to add a user.
type AddUser interface {
	AddUser(ctx context.Context, u *model.User) error
}

// AddUser add new user.
//
// Receives a request with user data. Responses whether the user was added successfully or not.
func (s *FaceitService) AddUser(ctx context.Context, req *api.User) (*emptypb.Empty, error) {
	// Validate request.
	val, err := protovalidate.New(
		protovalidate.WithMessages(
			&api.User{},
		),
	)
	if err != nil {
		return nil, servers.WrapError(codes.Internal, err, "create proto validator")
	}

	fieldMsgErrs, ok := isUserValid(req, val, true)
	if !ok {
		return nil, servers.Error(codes.InvalidArgument, "validation error", fieldMsgErrs)
	}

	// Add user.
	us := &model.User{
		ID: uuid.MustParse(req.GetId()), // Safe to ignore panic as it was validated before.
		UserState: model.UserState{
			PasswordHash: req.GetPasswordHash(),
			Email:        req.GetEmail(),
			FirstName:    req.GetFirstName(),
			LastName:     req.GetLastName(),
			Nickname:     req.GetNickname(),
			Country:      req.GetCountry(),
		},
	}

	err = s.deps.AddUser().AddUser(ctx, us)
	if err != nil {
		if errors.Is(err, database.ErrAlreadyExists) {
			return nil, servers.WrapError(codes.AlreadyExists, err, "user already exists")
		}

		return nil, servers.WrapError(codes.Internal, err, "ups, something went wrong!")
	}

	_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "204")) //nolint:errcheck

	return &emptypb.Empty{}, nil
}
