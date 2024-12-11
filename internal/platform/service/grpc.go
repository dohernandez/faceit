package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

// AddUser defines the use case to add a user.
type AddUser interface {
	AddUser(ctx context.Context, u *model.User) error
}

// FaceitServiceDeps holds the dependencies for the FaceitService.
type FaceitServiceDeps interface {
	Logger() ctxd.Logger
	GRPCAddr() string

	AddUser() AddUser
	UpdateUser() UpdateUser
	DeleteUser() DeleteUser

	ListUsersByCountry() ListUsersByCountry
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
func (s *FaceitService) AddUser(ctx context.Context, req *api.User) (*emptypb.Empty, error) {
	ctx = ctxd.AddFields(ctx, "service", "FaceitService")

	// Validate request.
	val, err := protovalidate.New(
		protovalidate.WithMessages(
			&api.User{},
		),
	)
	if err != nil {
		return nil, servers.Error(codes.Internal, fmt.Errorf("create proto validator: %w", err), nil)
	}

	fieldMsgErrs, ok := isUserValid(req, val, true)
	if !ok {
		return nil, servers.Error(codes.InvalidArgument, errors.New("validation error"), fieldMsgErrs)
	}

	// Extra validation since User proto message only have ID as required field.
	//

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
			return nil, servers.Error(codes.AlreadyExists, err, "user already exists")
		}

		return nil, servers.Error(codes.Internal, err, nil)
	}

	_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "204")) //nolint:errcheck

	return &emptypb.Empty{}, nil
}

func isUserValid(msg proto.Message, val *protovalidate.Validator, forAdd bool) (map[string]string, bool) {
	fields := make(map[string]string)

	if err := val.Validate(msg); err != nil {
		fields = mapValidatorError(err)
	}

	req, ok := msg.(*api.User)
	if !ok {
		return fields, len(fields) == 0
	}

	if req.PasswordHash != nil && !isValidSHA256Hash(req.GetPasswordHash()) {
		fields["password_hash"] = "invalid hash"
	}

	_, err := uuid.Parse(req.GetId())
	if err != nil {
		fields["password_hash"] = err.Error()
	}

	if !forAdd {
		return fields, len(fields) == 0
	}

	const require = "required"

	if req.Email == nil {
		fields["email"] = require
	}

	if req.FirstName == nil {
		fields["first_name"] = require
	}

	if req.LastName == nil {
		fields["last_name"] = require
	}

	if req.Country == nil {
		fields["country"] = require
	}

	return fields, len(fields) == 0
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
		return nil, servers.Error(codes.Internal, fmt.Errorf("create proto validator: %w", err), nil)
	}

	fieldMsgErrs, ok := isUserValid(req, val, false)
	if !ok {
		return nil, servers.Error(codes.InvalidArgument, errors.New("validation error"), fieldMsgErrs)
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
			return nil, servers.Error(codes.NotFound, err, "user not found")
		}

		return nil, servers.Error(codes.Internal, err, nil)
	}

	_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "204")) //nolint:errcheck

	return &emptypb.Empty{}, nil
}

// DeleteUser defines the use case to delete a user.
type DeleteUser interface {
	DeleteUser(ctx context.Context, id model.UserID) error
}

// DeleteUser delete the user.
//
// Receives a request with user data. Responses whether the user was deleted successfully or not.
func (s *FaceitService) DeleteUser(ctx context.Context, req *api.UserID) (*emptypb.Empty, error) {
	ctx = ctxd.AddFields(ctx, "service", "FaceitService")

	// Validate request.
	val, err := protovalidate.New(
		protovalidate.WithMessages(
			&api.UserID{},
		),
	)
	if err != nil {
		return nil, servers.Error(codes.Internal, fmt.Errorf("create proto validator: %w", err), nil)
	}

	fieldMsgErrs, ok := isUserValid(req, val, false)
	if !ok {
		return nil, servers.Error(codes.InvalidArgument, errors.New("validation error"), fieldMsgErrs)
	}

	id := uuid.MustParse(req.GetId()) // Safe to ignore panic as it was validated before.

	if err = s.deps.DeleteUser().DeleteUser(ctx, id); err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return nil, servers.Error(codes.NotFound, err, "user not found")
		}

		return nil, servers.Error(codes.Internal, err, nil)
	}

	_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "204")) //nolint:errcheck

	return &emptypb.Empty{}, nil
}

// defaultLimit is the default limit for the list users by country. This number corresponds default page size defined
// in the proto file.
const defaultLimit = 100

// ListUsersByCountry defines the use case to list users by country.
type ListUsersByCountry interface {
	ListUsersByCountry(ctx context.Context, country string, limit, offset uint64) ([]*model.User, error)
}

// ListUsersByCountry list users by country.
//
// Receives a request with country data. Responses with a list of users.
func (s *FaceitService) ListUsersByCountry(ctx context.Context, req *api.UsersByCountry) (*api.UserList, error) {
	ctx = ctxd.AddFields(ctx, "service", "FaceitService")

	// Validate request.
	val, err := protovalidate.New(
		protovalidate.WithMessages(
			&api.UserID{},
		),
	)
	if err != nil {
		return nil, servers.Error(codes.Internal, fmt.Errorf("create proto validator: %w", err), nil)
	}

	fieldMsgErrs, ok := isUserValid(req, val, false)
	if !ok {
		return nil, servers.Error(codes.InvalidArgument, errors.New("validation error"), fieldMsgErrs)
	}

	// Parse page token
	var offset uint64

	if pageToken := req.GetPageToken(); pageToken != "" {
		// Convert the page token to an uint64.
		_, err = fmt.Sscanf(pageToken, "%d", &offset)
		if err != nil {
			return nil, servers.Error(codes.InvalidArgument, fmt.Errorf("parse page token: %w", err), nil)
		}
	}

	limit := req.GetPageSize()

	if limit == 0 {
		limit = defaultLimit
	}

	// List users by country.
	users, err := s.deps.ListUsersByCountry().ListUsersByCountry(ctx, req.GetCountry(), req.GetPageSize(), offset)
	if err != nil {
		return nil, servers.Error(codes.Internal, err, nil)
	}

	// Prepare next page token
	nextPageToken := ""

	if len(users) == int(limit) {
		nextPageToken = strconv.FormatUint(offset+limit, 10)
	}

	// Map users to response users.
	var list []*api.User

	for _, u := range users {
		list = append(list, &api.User{
			Id:           u.ID.String(),
			PasswordHash: &u.PasswordHash,
			Email:        &u.Email,
			FirstName:    &u.FirstName,
			LastName:     &u.LastName,
			Nickname:     &u.Nickname,
			Country:      &u.Country,
		})
	}

	return &api.UserList{
		Users:         list,
		NextPageToken: nextPageToken,
	}, nil
}
