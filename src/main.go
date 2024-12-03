package main

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", route)

    if err := db_connect(); err != nil {
        log.Fatal(err)
    }
    migrations_up()
    if err := http.ListenAndServe(":3000", nil); err != nil {
        log.Fatal(err)
    }
}
