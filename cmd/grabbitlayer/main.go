package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/skrip42/grabbitLayer/internal/config"
	"github.com/skrip42/grabbitLayer/internal/executor"
	"github.com/skrip42/grabbitLayer/internal/message"
	"github.com/skrip42/grabbitLayer/internal/request"
)


func main() {
    rl, err := rotatelogs.New(
        "./logs/%Y-%m-%d.log",
        rotatelogs.WithRotationTime(24 * time.Hour),
        rotatelogs.WithMaxAge(-1),
        rotatelogs.WithRotationCount(30),
    )
    config := config.GetConfig()
    if err != nil {
        panic(err)
    }
    log.SetOutput(rl)
    gin.DefaultWriter = rl
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
        log.Printf("%v\n", string(data))
    })
    router.Run(":" + strconv.Itoa(config.Port))
}
