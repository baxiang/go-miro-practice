// +build wireinject

package services

import (
	"github.com/baxiang/go-miro-practice/api/proto"
	"github.com/baxiang/go-miro-practice/internal/pkg/config"
	"github.com/baxiang/go-miro-practice/internal/pkg/log"
	"github.com/google/wire"
)

var productProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	ProviderSet,
)

func CreateProductsService(cf string,
	detailsSvc proto.DetailsClient,
	ratingsSvc proto.RatingsClient,
	reviewsSvc proto.ReviewsClient) (ProductsService, error) {
	panic(wire.Build(productProviderSet))
}