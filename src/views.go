package main

import (
    "fmt"
    "html/template"
    "net/http"
)

func get_param(r *http.Request, i int) string {
    return r.Context().Value(URLParam{}).([]string)[i]
}

func index(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/products", http.StatusSeeOther)
}


func generic_index(w http.ResponseWriter, r *http.Request, p string, s any) {
    tmpl, err := template.ParseFiles("html/" + p + "/list.html")

    if err != nil {
        http.Error(w, "500 internal server error",
            http.StatusInternalServerError)
        fmt.Println(err)
        return
    }
    fmt.Println(tmpl.Execute(w, s))
}

func category_index(w http.ResponseWriter, r *http.Request) {
    generic_index(w, r, "categories", category_get())
}

func product_index(w http.ResponseWriter, r *http.Request) {
    fmt.Println(append([]Category{{}}, category_get()...))
    generic_index(w, r, "products", struct {
        Products []Product
        Categories []Category
    } {product_get(), append([]Category{}, category_get()...)})
}


func generic_edit(w http.ResponseWriter, r *http.Request, p string) {
    //id := get_param(r, 0)
    //old := r.PostForm

    http.Error(w, "501 not implemented", http.StatusNotImplemented)
    //http.Redirect("/categories")
}

func category_edit(w http.ResponseWriter, r *http.Request) {
    generic_edit(w, r, "categories")
}

func product_edit(w http.ResponseWriter, r *http.Request) {
    generic_edit(w, r, "products")
}


func generic_create(w http.ResponseWriter, r *http.Request, p string) {
    old := r.PostForm

    tmpl, err := template.ParseFiles("html/" + p + "/create.html")
    if err != nil {
        http.Error(w, "500 internal server error",
            http.StatusInternalServerError)
        fmt.Println(err)
        return
    }
    tmpl.Execute(w, old)
}

func category_create(w http.ResponseWriter, r *http.Request) {
    generic_create(w, r, "categories")
}

func product_create(w http.ResponseWriter, r *http.Request) {
    generic_create(w, r, "products")
}


func generic_store(w http.ResponseWriter, r *http.Request, t Model, p string) {
    r.ParseForm()
    if t.parseForm(r) {
        db_store(p, t)
        http.Redirect(w, r, "/" + p , http.StatusSeeOther)
    }
}

func category_store(w http.ResponseWriter, r *http.Request) {
    generic_store(w, r, Category{}, "categories")
}

func product_store(w http.ResponseWriter, r *http.Request) {
    generic_store(w, r, Product{}, "products")
}


func generic_update(w http.ResponseWriter, r *http.Request, p string) {
    http.Error(w, "501 not implemented", http.StatusNotImplemented)
}

func category_update(w http.ResponseWriter, r *http.Request) {
    generic_update(w, r, "categories")
}

func product_update(w http.ResponseWriter, r *http.Request) {
    generic_update(w, r, "products")
}


func generic_delete(w http.ResponseWriter, r *http.Request, p string) {
    http.Error(w, "501 not implemented", http.StatusNotImplemented)
}

func category_delete(w http.ResponseWriter, r *http.Request) {
    generic_delete(w, r, "categories")
}

func product_delete(w http.ResponseWriter, r *http.Request) {
    generic_delete(w, r, "products")
}
