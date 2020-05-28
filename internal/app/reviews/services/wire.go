// +build wireinject

package services

import (
	"github.com/baxiang/go-miro-practice/internal/app/reviews/repositories"
	"github.com/baxiang/go-miro-practice/internal/pkg/config"
	"github.com/baxiang/go-miro-practice/internal/pkg/log"
	"github.com/google/wire"
)

var ReviewsServiceProviderSet = wire.NewSet(
	config.ProviderSet,
	log.ProviderSet,
	//database.ProviderSet,
	ProviderSet,
)

func CreateReviewsService(cf string, sto repositories.ReviewsRepository) (ReviewsService, error) {
	panic(wire.Build(ReviewsServiceProviderSet))
}