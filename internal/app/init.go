package app

import (
	"fmt"
	"net/http"

	"github.com/PUDDLEEE/puddleee_back/internal/db"
	"github.com/PUDDLEEE/puddleee_back/pkg/config"
)

func InitApp() {
	app := &Application{
		Config: config.InitConfig(),
	}
	app.Client = db.InitDB()
	defer app.Client.Close()
	app.initRoutes()
	app.run()
}

func (app *Application) run() {
	addr := fmt.Sprintf("%s:%d", app.Config.Server.Host, app.Config.Server.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: app.Routes.Router,
	}
	srv.ListenAndServe()
}
