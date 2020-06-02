package controllers


import (
	"github.com/baxiang/go-miro-practice/internal/pkg/transports/http"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

)

func CreateInitControllersFn(pc *RatingsController) http.InitControllers {
	return func(r *gin.Engine) {
		r.GET("/rating/:productID", pc.Get)
	}
}

var ProviderSet = wire.NewSet(NewRatingsController, CreateInitControllersFn)