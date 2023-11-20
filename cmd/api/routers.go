package main

import (
	"net/http"

	"github.com/PUDDLEEE/puddleee_back/internal/controller"
	"github.com/gin-gonic/gin"
)

func (c *application) routers() http.Handler {
	router := gin.Default()
	v1 := router.Group("/v1")
	controller.UserGroup(v1)
	return router
}
