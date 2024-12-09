package main

import (
    "log"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    route(r)
    if err := db_connect(); err != nil {
        log.Fatal(err)
    }
    migrations_up()
    r.LoadHTMLGlob("html/*")
    r.Run(":3000")
}
