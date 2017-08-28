package main

import (
    "github.com/gin-gonic/gin"
    "os"
)

func main() {
    port := ":" + os.Getenv("PORT")

    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    r.Run(port)
}
