// +build wireinject

package controllers

import (
	"github.com/baxiang/go-miro-practice/internal/app/reviews/repositories"
	"github.com/baxiang/go-miro-practice/internal/app/reviews/services"
	"github.com/baxiang/go-miro-practice/internal/pkg/config"
	"github.com/baxiang/go-miro-practice/internal/pkg/log"
	"github.com/google/wire"
)

var ReviewControllerProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	services.ProviderSet,
	ProviderSet,
)

func CreateReviewsController(cf string, sto repositories.ReviewsRepository) (*ReviewsController, error) {
	panic(wire.Build(ReviewControllerProviderSet))
}