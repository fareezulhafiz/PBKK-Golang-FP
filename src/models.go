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
func (Category) TableName() string {
    return "categories"
}
func (Category) Migrate() {
    DbCreate(Category{}, []string{
        "id serial PRIMARY KEY",
        "name varchar(255) NOT NULL",
        "description text NOT NULL",
    })
}
func (Category) Drop() {
    DbDrop(Category{})
}
func (Category) Get(f string, v any) any {
    return DbGet(Category{}, f, v)
}
func (Category) GetAll() any {
    return DbGetAll(Category{})
}
func (m Category) Validate(c *gin.Context) any {
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
func (Product) TableName() string {
    return "products"
}
func (Product) Migrate() {
    DbCreate(Product{}, []string{
        "id serial PRIMARY KEY",
        "category int NOT NULL REFERENCES " + Category{}.TableName() + "(id)",
        "name varchar(255) NOT NULL",
        "price float(32) NOT NULL",
        "description text NOT NULL",
    })
}
func (Product) Drop() {
    DbDrop(Product{})
}
func (Product) Get(f string, v any) any {
    p := DbGet(Product{}, f, v)

    p.CategoryName = Category{}.Get("id", p.Category).(Category).Name
    return p
}
func (Product) GetAll() any {
    ps := DbGetAll(Product{})
    cs := Category{}.GetAll().([]Category)

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
func (m Product) Validate(c *gin.Context) any {
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
