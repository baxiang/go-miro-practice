package controllers

import (
	"github.com/baxiang/go-miro-practice/internal/pkg/transports/http"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func CreateInitControllersFn(pc *ProductsController) http.InitControllers {
	return func(r *gin.Engine) {
		r.GET("/product/:id", pc.Get)
	}
}

var ProviderSet = wire.NewSet(NewProductsController, CreateInitControllersFn)