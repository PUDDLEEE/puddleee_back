package main

import (
	"fmt"

	"github.com/PUDDLEEE/puddleee_back/pkg/config"
)

func main() {
	c := config.InitConfig()
	fmt.Println(c.Server.Port)
}
