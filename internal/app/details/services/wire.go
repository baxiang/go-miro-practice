// +build wireinject

package services

import (
	"github.com/baxiang/go-miro-practice/internal/app/details/repositories"
	"github.com/baxiang/go-miro-practice/internal/pkg/config"
	"github.com/baxiang/go-miro-practice/internal/pkg/database"
	"github.com/baxiang/go-miro-practice/internal/pkg/log"
	"github.com/google/wire"
)

var detailServiceProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	ProviderSet,
)

func CreateDetailsService(cf string, sto repositories.DetailsRepository) (DetailsService, error) {
	panic(wire.Build(detailServiceProviderSet))
}