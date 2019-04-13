package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"listener/app/Http/Controller"
	"net/http"
)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/book/list", Controller.GetBookList)
	r.StaticFS("/public",http.Dir(fmt.Sprintf("%s", viper.Get("public"))))
	r.StaticFS("/static",http.Dir(fmt.Sprintf("%s/static", viper.Get("public"))))

	return r
}
