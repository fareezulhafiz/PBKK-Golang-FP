package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func product_index(c *gin.Context) {
    c.HTML(http.StatusOK, Product{}.tableName() + "_list.html", gin.H{
        "ps": db_get(Product{}),
        "cs": append([]Category{{}}, db_get(Category{})...),
    })
}

func product_view(c *gin.Context) {
    m := db_get_by(Product{}, "id", c.Param("id"))

    if (m.Id == 0) {
        c.JSON(http.StatusNotFound, gin.H{"error": "404 page not found"})
    } else {
        c.HTML(http.StatusOK, m.tableName() + "_view.html", gin.H{
            "m": m,
            "cs": append([]Category{{}}, db_get(Category{})...),
        })
    }
}

func product_edit(c *gin.Context) {
    m := db_get_by(Product{}, "id", c.Param("id"))

    if (m.Id == 0) {
        c.JSON(http.StatusNotFound, gin.H{"error": "404 page not found"})
    } else {
        c.HTML(http.StatusOK, m.tableName() + "_edit.html", gin.H{
            "m": m,
            "cs": db_get(Category{}),
        })
    }
}

func product_create(c *gin.Context) {
    c.HTML(http.StatusOK, Product{}.tableName() + "_create.html", gin.H{
        "m": Product{},
        "cs": db_get(Category{}),
    })
}

func product_store(c *gin.Context) {
    m := Product{}.validate(c)

    if m != nil {
        db_store(*m.(*Product))
        c.Redirect(http.StatusFound, "/products")
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": "form error"})
    }
}

func product_update(c *gin.Context) {
    m := Product{}.validate(c)

    if m != nil {
        db_update(*m.(*Product), c.Param("id"))
        c.Redirect(http.StatusFound, "/products")
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": "form error"})
    }
}

func product_delete(c *gin.Context) {
    db_delete(Product{}, "id", c.Param("id"))
    c.Status(http.StatusNoContent)
}

func category_index(c *gin.Context) {
    c.HTML(200, Category{}.tableName() + "_list.html", gin.H{
        "cs": db_get(Category{}),
    })
}

func category_view(c *gin.Context) {
    m := db_get_by(Category{}, "id", c.Param("id"))

    if (m.Id == 0) {
        c.JSON(http.StatusNotFound, gin.H{"error": "404 page not found"})
    } else {
        c.HTML(200, m.tableName() + "_view.html", gin.H{
            "m": m,
        })
    }
}

func category_edit(c *gin.Context) {
    m := db_get_by(Category{}, "id", c.Param("id"))

    if (m.Id == 0) {
        c.JSON(http.StatusNotFound, gin.H{"error": "404 page not found"})
    } else {
        c.HTML(200, m.tableName() + "_edit.html", gin.H{
            "m": m,
        })
    }
}

func category_create(c *gin.Context) {
    c.HTML(200, Category{}.tableName() + "_create.html", gin.H{
        "m": Category{},
    })
}

func category_store(c *gin.Context) {
    m := Category{}.validate(c)

    if m != nil {
        db_store(*m.(*Category))
        c.Redirect(http.StatusFound, "/categories")
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": "form error"})
    }
}

func category_update(c *gin.Context) {
    m := Category{}.validate(c)

    if m != nil {
        db_update(*m.(*Category), c.Param("id"))
        c.Redirect(http.StatusFound, "/categories")
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": "form error"})
    }
}

func category_delete(c *gin.Context) {
    db_delete(Product{}, "category", c.Param("id"))
    db_delete(Category{}, "id", c.Param("id"))
    c.Status(http.StatusNoContent)
}
