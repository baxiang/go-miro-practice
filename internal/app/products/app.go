package products

import (
	"github.com/baxiang/go-miro-practice/internal/pkg/app"
	"github.com/baxiang/go-miro-practice/internal/pkg/transports/http"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"github.com/pkg/errors"
)

type Options struct {
	Name string
}

func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal app option error")
	}

	logger.Info("load application options success")

	return o, err
}

func NewApp(o *Options, logger *zap.Logger, hs *http.Server) (*app.Application, error) {
	a, err := app.NewApp(o.Name, logger, app.HttpServerOption(hs))

	if err != nil {
		return nil, errors.Wrap(err, "new app error")
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp, NewOptions)