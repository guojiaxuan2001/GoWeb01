package main

import (
	"goweb01/config"
	"goweb01/router"
)

//gin.H的源代码：
//type H map[string]any

func main() {
	config.InitConfig()

	r := router.SetupRouter()

	port := config.AppConfig.App.Port

	if port == "" {
		port = ":8080"
	}

	r.Run(port)
}
