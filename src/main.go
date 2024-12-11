package main

import (
    "log"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    Route(r)
    if err := DbConnect(); err != nil {
        log.Fatal(err)
    }
    r.LoadHTMLGlob("html/*")
    r.Run(":3000")
}
