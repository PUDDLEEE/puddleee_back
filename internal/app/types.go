package app

import (
	"github.com/PUDDLEEE/puddleee_back/ent"
	"github.com/PUDDLEEE/puddleee_back/pkg/config"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	Router *gin.Engine
}

type Application struct {
	Routes Routes
	Config *config.Config
	Client *ent.Client
}
