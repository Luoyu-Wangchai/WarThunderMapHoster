package main

import (
	"fmt"
	"thunder_hoster/config"
)

func main() {
	config.InitConfig()

	router := RouterSetup()
	router.Run(fmt.Sprintf(":%d", config.Cfg.Port))
}
