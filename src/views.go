package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func ProductList(c *gin.Context) {
    c.HTML(http.StatusOK, Product{}.TableName() + "_list.html", gin.H{
        "ps": Product{}.GetAll(),
    })
}

func ProductView(c *gin.Context) {
    m := Product{}.Get("id", c.Param("id")).(Product)

    if (m.Id == 0) {
        c.JSON(http.StatusNotFound, gin.H{"error": "404 page not found"})
    } else {
        c.HTML(http.StatusOK, m.TableName() + "_view.html", gin.H{
            "m": m,
        })
    }
}

func ProductEdit(c *gin.Context) {
    m := Product{}.Get("id", c.Param("id")).(Product)

    if (m.Id == 0) {
        c.JSON(http.StatusNotFound, gin.H{"error": "404 page not found"})
    } else {
        c.HTML(http.StatusOK, m.TableName() + "_edit.html", gin.H{
            "m": m,
            "cs": Category{}.GetAll().([]Category),
        })
    }
}

func ProductCreate(c *gin.Context) {
    c.HTML(http.StatusOK, Product{}.TableName() + "_create.html", gin.H{
        "cs": Category{}.GetAll().([]Category),
    })
}

func ProductStore(c *gin.Context) {
    m := Product{}.Validate(c)

    if m != nil {
        DbStore(*m.(*Product))
        c.Redirect(http.StatusFound, "/products")
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": "form error"})
    }
}

func ProductUpdate(c *gin.Context) {
    m := Product{}.Validate(c)

    if m != nil {
        DbUpdate(*m.(*Product), c.Param("id"))
        c.Redirect(http.StatusFound, "/products")
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": "form error"})
    }
}

func ProductDelete(c *gin.Context) {
    DbDelete(Product{}, "id", c.Param("id"))
    c.Status(http.StatusNoContent)
}

func CategoryList(c *gin.Context) {
    c.HTML(200, Category{}.TableName() + "_list.html", gin.H{
        "cs": Category{}.GetAll().([]Category),
    })
}

func CategoryView(c *gin.Context) {
    m := Category{}.Get("id", c.Param("id")).(Category)

    if (m.Id == 0) {
        c.JSON(http.StatusNotFound, gin.H{"error": "404 page not found"})
    } else {
        c.HTML(200, m.TableName() + "_view.html", gin.H{
            "m": m,
        })
    }
}

func CategoryEdit(c *gin.Context) {
    m := Category{}.Get("id", c.Param("id")).(Category)

    if (m.Id == 0) {
        c.JSON(http.StatusNotFound, gin.H{"error": "404 page not found"})
    } else {
        c.HTML(200, m.TableName() + "_edit.html", gin.H{
            "m": m,
        })
    }
}

func CategoryCreate(c *gin.Context) {
    c.HTML(200, Category{}.TableName() + "_create.html", nil)
}

func CategoryStore(c *gin.Context) {
    m := Category{}.Validate(c)

    if m != nil {
        DbStore(*m.(*Category))
        c.Redirect(http.StatusFound, "/categories")
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": "form error"})
    }
}

func CategoryUpdate(c *gin.Context) {
    m := Category{}.Validate(c)

    if m != nil {
        DbUpdate(*m.(*Category), c.Param("id"))
        c.Redirect(http.StatusFound, "/categories")
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": "form error"})
    }
}

func CategoryDelete(c *gin.Context) {
    DbDelete(Product{}, "category", c.Param("id"))
    DbDelete(Category{}, "id", c.Param("id"))
    c.Status(http.StatusNoContent)
}
