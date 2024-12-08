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
	ucAddUser *usecase.AddUser
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

	err = l.setupServices()
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
}

func (l *Locator) setupServices() error {
	l.InitGRPCService(
		servers.WithRegisterService(l.FaceitService),
	)

	err := l.InitGRPCRestService(
		servers.WithRegisterServiceHandler(l.FaceitService),
		servers.WithDocEndpoint(l.cfg.ServiceName,
			"/docs/",
			"/docs/service.swagger.json",
			swagger.SwgJSON),
		servers.WithVersionEndpoint(),
	)
	if err != nil {
		return err
	}

	l.InitMetricsService(servers.WithGRPCServer(l.GRPCService))

	return nil
}

// AddUser returns the service.AddUser use case.
func (l *Locator) AddUser() service.AddUser {
	return l.ucAddUser
}
