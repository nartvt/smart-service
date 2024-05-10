//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/nartvt/go-core/server"
	"github.com/nartvt/smart-service/internal/biz"
	"github.com/nartvt/smart-service/internal/conf"
	"github.com/nartvt/smart-service/internal/service"

	coreConf "github.com/nartvt/go-core/conf"
	coreService "github.com/nartvt/go-core/service"
	data "github.com/nartvt/smart-service/internal/data"
)

// wireApp init kratos application.
func initApp(*coreConf.Server, *conf.Data, log.Logger) (coreService.Service, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, initService))
}
