//+build wireinject

package controllers

import (
	"github.com/baxiang/go-miro-practice/internal/app/ratings/repositories"
	"github.com/baxiang/go-miro-practice/internal/app/ratings/services"
	"github.com/baxiang/go-miro-practice/internal/pkg/config"
	"github.com/baxiang/go-miro-practice/internal/pkg/database"
	"github.com/baxiang/go-miro-practice/internal/pkg/log"
	"github.com/google/wire"
)

var rateControllerProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	services.ProviderSet,
	ProviderSet,
)

func CreateRatingsController(cf string, sto repositories.RatingsRepository) (*RatingsController, error) {
	panic(wire.Build(rateControllerProviderSet))
}