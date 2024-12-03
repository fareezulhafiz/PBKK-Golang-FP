package main

import (
    "fmt"
    "net/http"
    "strconv"
)

type Category struct {
    Id          int
    Name        string
    Description string
}
func (c Category) toString() string {
    return fmt.Sprintf("DEFAULT, '%s', '%s'", c.Name, c.Description)
}
func (c Category) parseForm(r *http.Request) bool {
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
func (p Product) toString() string {
    return fmt.Sprintf("DEFAULT, %d, '%s', %d, '%s'",
        p.Category, p.Name, p.Price, p.Description)
}
func (p Product) parseForm(r *http.Request) bool {
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
    db_create("categories", []string{
        "id          serial       PRIMARY KEY",
        "name        varchar(255) NOT NULL",
        "description text         NOT NULL"})
}

func category_drop() {
    db_drop("categories")
}

func category_get() []Category {
    return db_get("categories", Category{})
}

func category_get_by_id(id int) Category {
    return db_get_by_id("categories", Category{}, id)
}

func product_migrate() {
    db_create("products", []string{
        "id          serial       PRIMARY KEY",
        "category    int          NOT NULL REFERENCES categories(id)",
        "name        varchar(255) NOT NULL",
        "price       int          NOT NULL",
        "description text         NOT NULL"})
}

func product_drop() {
    db_drop("products")
}

func product_get() []Product {
    return db_get("products", Product{})
}

func product_get_by_id(id int) Product {
    return db_get_by_id("products", Product{}, id)
}
