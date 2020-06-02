// +build wireinject

package services

import (
	"github.com/baxiang/go-miro-practice/internal/app/ratings/repositories"
	"github.com/baxiang/go-miro-practice/internal/pkg/config"
	"github.com/baxiang/go-miro-practice/internal/pkg/database"
	"github.com/baxiang/go-miro-practice/internal/pkg/log"
	"github.com/google/wire"
)

var ratingsServerProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	ProviderSet,
)

func CreateRatingsService(cf string, sto repositories.RatingsRepository) (RatingsService, error) {
	panic(wire.Build(ratingsServerProviderSet))
}