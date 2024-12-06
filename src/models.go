package main

import (
    "net/http"
    "strconv"
    "fmt"
)

type Category struct {
    Id          int
    Name        string
    Description string
}
func (Category) tableName() string {
    return "categories"
}
func (c *Category) parseForm(r *http.Request) bool {
    if c.Name = r.FormValue("name"); len(c.Name) < 3 || len(c.Name) > 255 {
        return false
    }
    c.Description = r.FormValue("description")
    return true
}

type Product struct {
    Id          int
    Category    int
    Name        string
    Price       int
    Description string
}
func (Product) tableName() string {
    return "products"
}
func (p *Product) parseForm(r *http.Request) bool {
    var err error

    if p.Category, err = strconv.Atoi(r.FormValue("category")); err != nil {
        return false
    }
    if p.Name = r.FormValue("name"); len(p.Name) < 3 || len(p.Name) > 255 {
        return false
    }
    if p.Price, err = strconv.Atoi(r.FormValue("price")); err != nil {
        return false
    }
    p.Description = r.FormValue("description")
    return true
}

func category_migrate() {
    db_create(Category{}, []string{
        "id serial PRIMARY KEY",
        "name varchar(255) NOT NULL",
        "description text NOT NULL"})
}

func category_drop() {
    db_drop(Category{})
}

func category_get() []Category {
    return db_get(Category{})
}

func product_migrate() {
    db_create(Product{}, []string{
        "id serial PRIMARY KEY",
        "category int NOT NULL REFERENCES " + Category{}.tableName() + "(id)",
        "name varchar(255) NOT NULL",
        "price int NOT NULL",
        "description text NOT NULL"})
}

func product_drop() {
    db_drop(Product{})
}

func product_get() []Product {
    return db_get(Product{})
}
