package app

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hellofresh/health-go/v5"
)

// ErrHealthCheck occurs when health check failed.
var ErrHealthCheck = errors.New("health checks")

func (l *Locator) withHealthChecks() []health.Config {
	checks := []health.Config{
		{
			Name:      "postgres-check",
			Timeout:   time.Second * 5,
			SkipOnErr: true,
			Check: func(ctx context.Context) error {
				if err := l.Storage.DB().PingContext(ctx); err != nil {
					return fmt.Errorf("%w: ping context: %v", ErrHealthCheck, err) //nolint:errorlint
				}

				var version string
				if err := l.Storage.DB().QueryRowContext(ctx, "SELECT VERSION()").Scan(&version); err != nil {
					return fmt.Errorf("%w: query version: %s", ErrHealthCheck, err) //nolint:errorlint
				}

				return nil
			},
		},
	}

	return checks
}
