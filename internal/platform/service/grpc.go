package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"

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
