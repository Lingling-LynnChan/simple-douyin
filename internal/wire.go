//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package internal

import (
	"douyin-simple/internal/biz/usecase"
	"douyin-simple/internal/conf"
	"douyin-simple/internal/data"
	"douyin-simple/internal/server"
	"douyin-simple/internal/service"

	"github.com/google/wire"
)

func WireApp(*conf.Server, *conf.Data) *Bootstrapper {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, usecase.ProviderSet, service.ProviderSet, New))
}