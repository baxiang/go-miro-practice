// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package grpcservers

import (
	"github.com/baxiang/go-miro-practice/internal/app/ratings/services"
	"github.com/baxiang/go-miro-practice/internal/pkg/config"
	"github.com/baxiang/go-miro-practice/internal/pkg/database"
	"github.com/baxiang/go-miro-practice/internal/pkg/log"
	"github.com/google/wire"
)

// Injectors from wire.go:

func CreateRatingsServer(cf string, service services.RatingsService) (*RatingsServer, error) {
	viper, err := config.NewConfig(cf)
	if err != nil {
		return nil, err
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.NewLogger(options)
	if err != nil {
		return nil, err
	}
	ratingsServer, err := NewRatingsServer(logger, service)
	if err != nil {
		return nil, err
	}
	return ratingsServer, nil
}

// wire.go:

var grpcServersProviderSet = wire.NewSet(log.ProviderSet, config.ProviderSet, database.ProviderSet, ProviderSet)