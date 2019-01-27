package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikhailbadin/csp-aggregator/controllers"
)

// NewRouter create new service router
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	{
		router.NoRoute(pageNotFoundHandler)
		router.POST("/csp_report", controllers.WriteReportHandler)
		router.POST("/csp_report_only", controllers.WriteReportOnlyHandler)
	}
	return router
}

func pageNotFoundHandler(c *gin.Context) {
	c.String(http.StatusNotFound, "")
}
