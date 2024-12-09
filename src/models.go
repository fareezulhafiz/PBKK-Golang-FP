package main

import (
    "net/http"
    "strconv"
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
    db_create(&Category{}, []string{
        "id serial PRIMARY KEY",
        "name varchar(255) NOT NULL",
        "description text NOT NULL"})
}
func (Category) drop() {
    db_drop(&Category{})
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
func (Product) migrate() {
    db_create(&Product{}, []string{
        "id serial PRIMARY KEY",
        "category int NOT NULL REFERENCES " + Category{}.tableName() + "(id)",
        "name varchar(255) NOT NULL",
        "price int NOT NULL",
        "description text NOT NULL"})
}
func (Product) drop() {
    db_drop(&Product{})
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
