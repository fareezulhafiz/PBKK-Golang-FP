package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
    r.GET("/", func (c *gin.Context) {
        c.Redirect(http.StatusMovedPermanently, "/products")
    })
    r.GET("/products",            ProductList)
    r.GET("/products/:id",        ProductView)
    r.GET("/products/:id/edit",   ProductEdit)
    r.GET("/products/create",     ProductCreate)
    r.POST("/products",           ProductStore)
    r.POST("/products/:id",       ProductUpdate)
    r.DELETE("/products/:id",     ProductDelete)
    r.GET("/categories",          CategoryList)
    r.GET("/categories/:id",      CategoryView)
    r.GET("/categories/:id/edit", CategoryEdit)
    r.GET("/categories/create",   CategoryCreate)
    r.POST("/categories",         CategoryStore)
    r.POST("/categories/:id",     CategoryUpdate)
    r.DELETE("/categories/:id",   CategoryDelete)
}
