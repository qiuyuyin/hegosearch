package server

import (
    "github.com/gin-gonic/gin"
)

func Router(sever *SearchSever) *gin.Engine {
    var Router = gin.Default()
    Router.GET("/", func(context *gin.Context) {
        context.JSON(200, "ping pong")
    })
    Router.POST("/api/search", sever.Search)
    return Router
}
