package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/parkby/parkby-service/config"
	"gitlab.com/parkby/parkby-service/log"
	"gitlab.com/parkby/parkby-service/web/handler"
)

func main() {
	log.Info.Println("main start")
	logInfo()

	router := gin.Default()

	handler.InitUserEndpoint(router)

	http.ListenAndServe(":8080", router)
}

func logInfo() {
	log.Info.Println("MySQL: ", config.Config.MySQLUser+"@localhost")
}
