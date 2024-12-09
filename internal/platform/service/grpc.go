package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/bool64/ctxd"
	"github.com/bufbuild/protovalidate-go"
	"github.com/dohernandez/faceit/internal/domain/model"
	api "github.com/dohernandez/faceit/internal/platform/service/pb"
	"github.com/dohernandez/go-grpc-service/database"
	"github.com/dohernandez/servers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
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
	UpdateUser() UpdateUser
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

	// Validate request.
	val, err := protovalidate.New(
		protovalidate.WithMessages(
			&api.AddUserRequest{},
		),
	)
	if err != nil {
		return nil, servers.Error(codes.Internal, fmt.Errorf("create proto validator: %w", err), nil)
	}

	var fieldMsgErrs map[string]string

	if err = val.Validate(req); err != nil {
		fieldMsgErrs = mapValidatorError(err)
	}

	if !isValidSHA256Hash(req.GetPasswordHash()) {
		fieldMsgErrs["password_hash"] = "invalid hash"
	}

	if len(fieldMsgErrs) > 0 {
		return nil, servers.Error(codes.InvalidArgument, errors.New("validation error"), fieldMsgErrs)
	}

	// Add user.
	us := model.UserState{
		UserCredentials: model.UserCredentials{
			PasswordHash: req.GetPasswordHash(),
			Email:        req.GetEmail(),
		},
		UserInfo: model.UserInfo{
			FirstName: req.GetFirstName(),
			LastName:  req.GetLastName(),
			Nickname:  req.GetNickname(),
			Country:   req.GetCountry(),
		},
	}

	id, err := s.deps.AddUser().AddUser(ctx, us)
	if err != nil {
		if errors.Is(err, database.ErrAlreadyExists) {
			return nil, servers.Error(codes.AlreadyExists, err, "user already exists")
		}

		return nil, servers.Error(codes.Internal, err, nil)
	}

	_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "201")) //nolint:errcheck

	return &api.AddUserResponse{
		Id: id.String(),
	}, nil
}

func mapValidatorError(err error) map[string]string {
	valErrs, ok := err.(interface{ ToProto() *validate.Violations })
	if !ok {
		return nil
	}

	vals := valErrs.ToProto().GetViolations()
	if len(vals) == 0 {
		return nil
	}

	fieldMsg := make(map[string]string)

	for _, v := range vals {
		fieldMsg[v.GetFieldPath()] = v.GetMessage() //nolint:staticcheck // It is deprecated but still used in the proto package
	}

	return fieldMsg
}

func isValidSHA256Hash(hash string) bool {
	// Decode the string from hexadecimal
	decoded, err := hex.DecodeString(hash)
	if err != nil {
		return false // Not a valid hex string
	}

	// Check if the decoded byte slice matches the size of a SHA-256 hash
	return len(decoded) == sha256.Size
}

// UpdateUser defines the use case to update a user.
type UpdateUser interface {
	UpdateUser(ctx context.Context, id model.UserID, info model.UserInfo) error
}

// UpdateUser update the user.
//
// Receives a request with user data. Responses whether the user was updated successfully or not.
func (s *FaceitService) UpdateUser(ctx context.Context, req *api.UpdateUserRequest) (*emptypb.Empty, error) {
	// Validate request.
	val, err := protovalidate.New(
		protovalidate.WithMessages(
			&api.AddUserRequest{},
		),
	)
	if err != nil {
		return nil, servers.Error(codes.Internal, fmt.Errorf("create proto validator: %w", err), nil)
	}

	if err = val.Validate(req); err != nil {
		return nil, servers.Error(codes.Internal, fmt.Errorf("create proto validator: %w", err), nil)
	}

	// Update user.
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, servers.Error(codes.InvalidArgument, fmt.Errorf("parse user id: %w", err), nil)
	}

	info := model.UserInfo{
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Nickname:  req.GetNickname(),
		Country:   req.GetCountry(),
	}

	if err = s.deps.UpdateUser().UpdateUser(ctx, id, info); err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return nil, servers.Error(codes.NotFound, err, "user not found")
		}

		return nil, servers.Error(codes.Internal, err, nil)
	}

	_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "204")) //nolint:errcheck

	return nil, nil
}
