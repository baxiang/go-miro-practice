package controllers

import (
	"github.com/baxiang/go-miro-practice/internal/pkg/transports/http"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func CreateInitControllersFn(pc *DetailsController) http.InitControllers {
	return func(r *gin.Engine) {
		r.GET("/detail/:id", pc.Get)
	}
}

var ProviderSet = wire.NewSet(NewDetailsController, CreateInitControllersFn)