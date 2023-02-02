package main

import (
	"fmt"
	"github.com/exceed19-cpsk/backend-bubblebungbung/config"
	"github.com/gin-gonic/gin"
)

var appConfig config.Config

func init() {
	appConfig = config.Load()
}

func main() {
	server := gin.Default()
	server.Run(fmt.Sprint(":", appConfig.LISTENING_PORT))
}
