package service

import (
	"context"
	"errors"

	"github.com/bool64/ctxd"
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

// UpdateUser defines the use case to update a user.
type UpdateUser interface {
	UpdateUser(ctx context.Context, id model.UserID, info model.UserState) error
}

// UpdateUser update the user.
//
// Receives a request with user data. Responses whether the user was updated successfully or not.
func (s *FaceitService) UpdateUser(ctx context.Context, req *api.User) (*emptypb.Empty, error) {
	ctx = ctxd.AddFields(ctx, "service", "FaceitService")

	// Validate request.
	val, err := protovalidate.New(
		protovalidate.WithMessages(
			&api.User{},
		),
	)
	if err != nil {
		return nil, servers.WrapError(codes.Internal, err, "create proto validator")
	}

	fieldMsgErrs, ok := isUserValid(req, val, false)
	if !ok {
		return nil, servers.Error(codes.InvalidArgument, "validation error", fieldMsgErrs)
	}

	id := uuid.MustParse(req.GetId()) // Safe to ignore panic as it was validated before.

	us := model.UserState{
		PasswordHash: req.GetPasswordHash(),
		Email:        req.GetEmail(),
		FirstName:    req.GetFirstName(),
		LastName:     req.GetLastName(),
		Nickname:     req.GetNickname(),
		Country:      req.GetCountry(),
	}

	if us == (model.UserState{}) {
		_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "204")) //nolint:errcheck

		return &emptypb.Empty{}, nil
	}

	if err = s.deps.UpdateUser().UpdateUser(ctx, id, us); err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return nil, servers.WrapError(codes.NotFound, err, "user not found")
		}

		return nil, servers.WrapError(codes.Internal, err, "ups, something went wrong!")
	}

	_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "204")) //nolint:errcheck

	return &emptypb.Empty{}, nil
}
