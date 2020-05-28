package controllers

import (
	"github.com/baxiang/go-miro-practice/internal/pkg/transports/http"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

)

func CreateInitControllersFn(pc *ReviewsController) http.InitControllers {
	return func(r *gin.Engine) {
		r.GET("/reviews", pc.Query)
	}
}

var ProviderSet = wire.NewSet(NewReviewsController, CreateInitControllersFn)