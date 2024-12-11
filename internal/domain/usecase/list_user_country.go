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

// ListUserByCountry is a use case to list users by country.
type ListUserByCountry struct {
	finder UserByCountryFinder

	logger ctxd.Logger
}

// NewListUserByCountry creates a new ListUserByCountry use case.
func NewListUserByCountry(finder UserByCountryFinder, logger ctxd.Logger) *ListUserByCountry {
	return &ListUserByCountry{
		finder: finder,
		logger: logger,
	}
}

// ListByCountry executes the list user by country use case.
func (l *ListUserByCountry) ListByCountry(ctx context.Context, country string, limit, offset uint64) ([]*model.User, error) {
	ctx = ctxd.AddFields(ctx, "use_case", "ListUserByCountry", "country", country)

	users, err := l.finder.ListByCountry(ctx, country, limit, offset)
	if err != nil {
		return nil, ctxd.WrapError(ctx, err, "list user by country") // error contains the context fields added
	}

	l.logger.Debug(ctx, "user list by country")

	return users, nil
}
