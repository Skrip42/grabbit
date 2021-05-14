package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skrip42/grabbitLayer/internal/executor"
	"github.com/skrip42/grabbitLayer/internal/message"
	"github.com/skrip42/grabbitLayer/internal/request"
)


func main() {
    router := gin.Default()
    ex := executor.New()
    router.POST("/", func(c *gin.Context) {
        req := request.Request{}
        c.BindJSON(&req)
        queueName := c.DefaultQuery("queue", "grabbit_default")
        callback := c.Query("callback")
        msg := message.New(req, callback, queueName)
        go ex.Execute(msg)
        c.String(http.StatusOK, msg.Callback)
    })
    router.GET("/test", func(c *gin.Context) {
        data, err := c.GetRawData()
        if err != nil {
            panic(err)
        }
        c.String(http.StatusOK, string(data))
    })

    router.Run(":88")
}
