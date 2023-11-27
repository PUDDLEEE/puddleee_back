package cmd

import (
	"fmt"
	"net/http"

	"github.com/PUDDLEEE/puddleee_back/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run Puddlee Backend Server",
	Long: `Run Puddlee Backend Server on local.
	configuration file lies on the config directory.`,
	Run: serve,
}

func serve(cmd *cobra.Command, args []string) {
	c := config.InitConfig()
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ping",
		})
	})
	addr := fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
	r.Run(addr)
}
