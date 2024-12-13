package service

import (
	"crypto/sha256"
	"encoding/hex"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/bool64/ctxd"
	"github.com/bufbuild/protovalidate-go"
	api "github.com/dohernandez/faceit/internal/platform/service/pb"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

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
