// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package copper

import (
	"github.com/google/wire"
	"github.com/tusharsoni/copper/v2/cconfig"
	"github.com/tusharsoni/copper/v2/clogger"
)

// Injectors from wire.go:

func InitApp() (*App, error) {
	flags := NewFlags()
	dir := flags.ConfigDir
	env := flags.Env
	config, err := cconfig.New(dir, env)
	if err != nil {
		return nil, err
	}
	logger, err := clogger.NewWithConfig(config)
	if err != nil {
		return nil, err
	}
	lifecycle := NewLifecycle(logger)
	app := New(lifecycle, config, logger)
	return app, nil
}

// wire.go:

var WireModule = wire.NewSet(wire.FieldsOf(new(*App), "Lifecycle", "Config", "Logger"))