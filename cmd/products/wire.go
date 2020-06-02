//+build wireinject

package main

import (
	"github.com/baxiang/go-miro-practice/internal/app/products"
	"github.com/baxiang/go-miro-practice/internal/app/products/grpcclients"
	"github.com/baxiang/go-miro-practice/internal/app/products/controllers"
	"github.com/baxiang/go-miro-practice/internal/app/products/services"
	"github.com/baxiang/go-miro-practice/internal/pkg/app"
	"github.com/baxiang/go-miro-practice/internal/pkg/config"
	"github.com/baxiang/go-miro-practice/internal/pkg/consul"
	"github.com/baxiang/go-miro-practice/internal/pkg/jaeger"
	"github.com/baxiang/go-miro-practice/internal/pkg/log"
	"github.com/baxiang/go-miro-practice/internal/pkg/transports/grpc"
	"github.com/baxiang/go-miro-practice/internal/pkg/transports/http"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	consul.ProviderSet,
	jaeger.ProviderSet,
	http.ProviderSet,
	grpc.ProviderSet,
	grpcclients.ProviderSet,
	controllers.ProviderSet,
	services.ProviderSet,
	products.ProviderSet,
)


func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}