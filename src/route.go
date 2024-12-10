package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func route(r *gin.Engine) {
    r.GET("/", func (c *gin.Context) {
        c.Redirect(http.StatusMovedPermanently, "/products")
    })
    r.GET("/products", product_list)
    r.GET("/products/:id", product_view)
    r.GET("/products/:id/edit", product_edit)
    r.GET("/products/create", product_create)
    r.POST("/products", product_store)
    r.POST("/products/:id", product_update)
    r.DELETE("/products/:id", product_delete)
    r.GET("/categories", category_list)
    r.GET("/categories/:id", category_view)
    r.GET("/categories/:id/edit", category_edit)
    r.GET("/categories/create", category_create)
    r.POST("/categories", category_store)
    r.POST("/categories/:id", category_update)
    r.DELETE("/categories/:id", category_delete)
}
