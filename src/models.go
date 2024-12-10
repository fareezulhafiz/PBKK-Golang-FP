package main

import (
    "strconv"

    "github.com/gin-gonic/gin"
)

type Category struct {
    Id          int
    Name        string
    Description string
}
func (Category) tableName() string {
    return "categories"
}
func (Category) migrate() {
    db_create(Category{}, []string{
        "id serial PRIMARY KEY",
        "name varchar(255) NOT NULL",
        "description text NOT NULL"})
}
func (Category) drop() {
    db_drop(Category{})
}
func (Category) get(f string, v any) any {
    return db_get(Category{}, f, v)
}
func (Category) get_all() any {
    return db_get_all(Category{})
}
func (m Category) validate(c *gin.Context) any {
    if m.Name = c.PostForm("name"); len(m.Name) < 3 {
        return nil
    }
    m.Description = c.PostForm("description")
    m.Id = 1
    return &m
}

type Product struct {
    Id           int
    Category     int
    CategoryName string `db:"-"`
    Name         string
    Price        float64
    Description  string
}
func (Product) tableName() string {
    return "products"
}
func (Product) migrate() {
    db_create(Product{}, []string{
        "id serial PRIMARY KEY",
        "category int NOT NULL REFERENCES " + Category{}.tableName() + "(id)",
        "name varchar(255) NOT NULL",
        "price float(32) NOT NULL",
        "description text NOT NULL"})
}
func (Product) drop() {
    db_drop(Product{})
}
func (Product) get(f string, v any) any {
    p := db_get(Product{}, f, v)

    p.CategoryName = Category{}.get("id", p.Category).(Category).Name
    return p


}
func (Product) get_all() any {
    ps := db_get_all(Product{})
    cs := Category{}.get_all().([]Category)

    for i, p := range ps {
        for _, c := range cs {
            if c.Id == p.Category {
                ps[i].CategoryName = c.Name
                break
            }
        }
    }
    return ps
}
func (m Product) validate(c *gin.Context) any {
    var err error

    if m.Category, err = strconv.Atoi(c.PostForm("category")); err != nil {
        return nil
    }
    if m.Name = c.PostForm("name"); len(m.Name) < 3 || len(m.Name) > 255 {
        return nil
    }
    if m.Price, err = strconv.ParseFloat(c.PostForm("price"), 32); err != nil {
        return nil
    }
    m.Description = c.PostForm("description")
    return &m
}
