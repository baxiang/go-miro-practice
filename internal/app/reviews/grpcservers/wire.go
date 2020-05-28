// +build wireinject

package grpcservers

import (
	"github.com/baxiang/go-miro-practice/internal/app/reviews/services"
	"github.com/baxiang/go-miro-practice/internal/pkg/config"
	"github.com/baxiang/go-miro-practice/internal/pkg/log"
	"github.com/google/wire"
)

var grpcReviewsProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	ProviderSet,
)

func CreateReviewsServer(cf string, service services.ReviewsService) (*ReviewsServer, error) {
	panic(wire.Build(grpcReviewsProviderSet))
}