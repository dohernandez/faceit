package app

import (
	"github.com/dohernandez/faceit/internal/domain/usecase"
	"github.com/dohernandez/faceit/internal/platform/config"
	"github.com/dohernandez/faceit/internal/platform/notifier"
	"github.com/dohernandez/faceit/internal/platform/service"
	"github.com/dohernandez/faceit/internal/platform/storage"
	"github.com/dohernandez/faceit/resources/swagger"
	sapp "github.com/dohernandez/go-grpc-service/app"
	"github.com/dohernandez/servers"
)

// Locator defines application resources.
type Locator struct {
	*sapp.Locator

	cfg *config.Config

	FaceitService *service.FaceitService

	notifierUser *notifier.NoopNotifier

	// storages
	storageUser *storage.User

	// use cases
	ucAddUser           *usecase.AddUser
	ucUpdateUser        *usecase.UpdateUser
	usDeleteUser        *usecase.DeleteUser
	usListUserByCountry *usecase.ListUsersByCountry
}

// NewServiceLocator creates application locator.
func NewServiceLocator(cfg *config.Config, opts ...sapp.Option) (*Locator, error) {
	// Init postgres database
	opts = append(opts, sapp.WithPostgresDBx())

	upl, err := sapp.NewServiceLocator(cfg.Config, opts...)
	if err != nil {
		return nil, err
	}

	l := &Locator{
		Locator: upl,
		cfg:     cfg,
	}

	// setting up storage dependencies
	l.setupStorage()

	// setting up notifier dependencies
	l.notifierUser = notifier.NewNoopNotifier()

	// setting up use cases dependencies
	l.setupUsecaseDependencies()

	l.FaceitService = service.NewFaceitService(l)

	// Setup services with health check
	err = l.SetupServices(l.FaceitService, swagger.SwgJSON, servers.WithHealthCheck(l.withHealthChecks()...))
	if err != nil {
		return nil, err
	}

	return l, nil
}

// setupStorage sets up storage dependencies (platform).
func (l *Locator) setupStorage() {
	l.storageUser = storage.NewUser(l.Storage)
}

// setupUsecaseDependencies sets up use case dependencies (domain).
func (l *Locator) setupUsecaseDependencies() {
	l.ucAddUser = usecase.NewAddUser(l.storageUser, l.notifierUser, l.CtxdLogger())
	l.ucUpdateUser = usecase.NewUpdateUser(l.storageUser, l.notifierUser, l.CtxdLogger())
	l.usDeleteUser = usecase.NewDeleteUser(l.storageUser, l.notifierUser, l.CtxdLogger())
	l.usListUserByCountry = usecase.NewListUsersByCountry(l.storageUser, l.CtxdLogger())
}

// AddUser returns the usecase.AddUser use case.
func (l *Locator) AddUser() service.AddUser {
	return l.ucAddUser
}

// UpdateUser returns the usecase.UpdateUser use case.
func (l *Locator) UpdateUser() service.UpdateUser {
	return l.ucUpdateUser
}

// DeleteUser returns the usecase.DeleteUser use case.
func (l *Locator) DeleteUser() service.DeleteUser {
	return l.usDeleteUser
}

// ListUsersByCountry returns the usecase.ListUsersByCountry use case.
func (l *Locator) ListUsersByCountry() service.ListUsersByCountry {
	return l.usListUserByCountry
}
