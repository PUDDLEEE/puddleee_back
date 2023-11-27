package app

import (
	"github.com/PUDDLEEE/puddleee_back/pkg/config"
	"github.com/gin-gonic/gin"
)

type Application struct {
	Router *gin.Engine
	Config *config.Config
}
