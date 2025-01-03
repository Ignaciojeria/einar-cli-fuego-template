package observability

import (
	"archetype/app/shared/configuration"
	"archetype/app/shared/infrastructure/observability/strategy"
	"log/slog"

	ioc "github.com/Ignaciojeria/einar-ioc/v2"
)

func init() {
	ioc.Registry(
		newLoggerProvider,
		configuration.NewConf,
	)
}

func newLoggerProvider(conf configuration.Conf) (*slog.Logger, error) {
	// Get the observability strategy
	observabilityStrategyKey := configuration.Getenv(strategy.OBSERVABILITY_STRATEGY)
	switch observabilityStrategyKey {
	case "openobserve":
		return strategy.OpenObserveGRPCLogProvider(conf)
	case "datadog":
		return strategy.DatadogGRPCLogProvider(conf)
	default:
		return strategy.NoOpStdoutLogProvider(conf), nil
	}
}
