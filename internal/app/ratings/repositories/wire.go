// +build wireinject

package repositories

import (
	"github.com/baxiang/go-miro-practice/internal/pkg/config"
	"github.com/baxiang/go-miro-practice/internal/pkg/database"
	"github.com/baxiang/go-miro-practice/internal/pkg/log"
	"github.com/google/wire"
)

var ratingProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	ProviderSet,
)

func CreateRatingRepository(f string) (RatingsRepository, error) {
	panic(wire.Build(ratingProviderSet))
}