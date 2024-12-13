package service

import (
	"context"
	"fmt"
	"github.com/bool64/ctxd"
	"github.com/bufbuild/protovalidate-go"
	"github.com/dohernandez/faceit/internal/domain/model"
	api "github.com/dohernandez/faceit/internal/platform/service/pb"
	"github.com/dohernandez/servers"
	"google.golang.org/grpc/codes"
	"strconv"
)

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
			&api.UsersByCountry{},
		),
	)
	if err != nil {
		return nil, servers.WrapError(codes.Internal, err, "create proto validator", nil)
	}

	fieldMsgErrs, ok := isUserValid(req, val, false)
	if !ok {
		return nil, servers.Error(codes.InvalidArgument, "validation error", fieldMsgErrs)
	}

	// Parse page token
	var offset uint64

	if pageToken := req.GetPageToken(); pageToken != "" {
		// Convert the page token to an uint64.
		_, err = fmt.Sscanf(pageToken, "%d", &offset)
		if err != nil {
			return nil, servers.WrapError(codes.InvalidArgument, err, "parse page token")
		}
	}

	limit := req.GetPageSize()

	if limit == 0 {
		limit = defaultLimit
	}

	// List users by country.
	users, err := s.deps.ListUsersByCountry().ListUsersByCountry(ctx, req.GetCountry(), limit, offset)
	if err != nil {
		return nil, servers.WrapError(codes.Internal, err, "ups, something went wrong!")
	}

	// Prepare next page token
	nextPageToken := ""

	if len(users) == int(limit) { //nolint:gosec // It is safe to ignore as it is not a security issue
		nextPageToken = strconv.FormatUint(offset+limit, 10)
	}

	// Map users to response users.
	list := make([]*api.User, 0, len(users))

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
