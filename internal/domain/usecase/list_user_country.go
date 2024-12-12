package usecase

import (
	"context"

	"github.com/bool64/ctxd"
	"github.com/dohernandez/faceit/internal/domain/model"
)

//go:generate mockery --name=UserByCountryFinder --outpkg=mocks --output=mocks --filename=user_by_country_finder.go --with-expecter

// UserByCountryFinder defines functionality to list users by country.
type UserByCountryFinder interface {
	ListByCountry(ctx context.Context, country string, limit, offset uint64) ([]*model.User, error)
}

// ListUsersByCountry is a use case to list users by country.
type ListUsersByCountry struct {
	finder UserByCountryFinder

	logger ctxd.Logger
}

// NewListUsersByCountry creates a new ListUsersByCountry use case.
func NewListUsersByCountry(finder UserByCountryFinder, logger ctxd.Logger) *ListUsersByCountry {
	return &ListUsersByCountry{
		finder: finder,
		logger: logger,
	}
}

// ListUsersByCountry executes the list user by country use case.
func (l *ListUsersByCountry) ListUsersByCountry(ctx context.Context, country string, limit, offset uint64) ([]*model.User, error) {
	ctx = ctxd.AddFields(ctx, "use_case", "ListUsersByCountry", "country", country)

	users, err := l.finder.ListByCountry(ctx, country, limit, offset)
	if err != nil {
		return nil, ctxd.WrapError(ctx, err, "list user by country") // error contains the context fields added
	}

	l.logger.Debug(ctx, "user list by country")

	return users, nil
}
