package main

import (
	"github.com/PUDDLEEE/puddleee_back/cmd"
	_ "github.com/PUDDLEEE/puddleee_back/docs"
)

//	@title			Puddle Chat Application API
//	@version		1.0
//	@description	Need help with live action? Puddle will help you through chat service.

//	@contact.name	MilkyMilky0116
//	@contact.url	https://milkymilky0116.github.io
//	@contact.email	sjlee990129@gmail.com

//	@host		localhost:8080
//	@BasePath	/api/v1

// @securityDefinitions.basic	BasicAuth

func main() {
	cmd.Execute()
}
