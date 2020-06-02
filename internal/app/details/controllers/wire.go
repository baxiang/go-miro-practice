package controllers

import (
	"github.com/baxiang/go-miro-practice/internal/app/details/repositories"
	"github.com/baxiang/go-miro-practice/internal/app/details/services"
	"github.com/baxiang/go-miro-practice/internal/pkg/config"
	"github.com/baxiang/go-miro-practice/internal/pkg/database"
	"github.com/baxiang/go-miro-practice/internal/pkg/log"
	"github.com/google/wire"
)

var detailControllerProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	services.ProviderSet,
	ProviderSet,
)


func CreateDetailsController(cf string, sto repositories.DetailsRepository) (*DetailsController, error) {
	panic(wire.Build(detailControllerProviderSet))
}