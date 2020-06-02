// +build wireinject

package grpcserver

import (
	"github.com/baxiang/go-miro-practice/internal/app/ratings/services"
	"github.com/baxiang/go-miro-practice/internal/pkg/config"
	"github.com/baxiang/go-miro-practice/internal/pkg/database"
	"github.com/baxiang/go-miro-practice/internal/pkg/log"
	"github.com/google/wire"
)

var grpcServersProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	ProviderSet,
)

func CreateRatingsServer(cf string, service services.RatingsService) (*RatingsServer, error) {
	panic(wire.Build(grpcServersProviderSet))
}