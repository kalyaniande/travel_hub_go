package router

import (
	"azri_hub/handlers/hotels_handler"
	"azri_hub/middleware/authentication"
	_ "azri_hub/router/routes"
	"github.com/atarantini/ginrequestid"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	router := gin.Default()
	router.Use(ginrequestid.RequestId())
	v1 := router.Group("/api/v1/hotels")
	v1.Use(authentication.Authenticate())

	v1.POST("/availability", hotels_handler.SearchandAvailability)

	return router
}
