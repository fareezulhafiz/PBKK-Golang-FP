package main

import (
    "html/template"
    "net/http"
)

func get_param(r *http.Request, i int) string {
    return r.Context().Value(URLParam{}).([]string)[i]
}

func index(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/products", http.StatusSeeOther)
}


// TODO form error + success
func generic_index(w http.ResponseWriter, r *http.Request, t Model, s any) {
    tmpl, err := template.ParseFiles("html/" + t.tableName() + "/list.html")

    if err != nil {
        http.Error(w, "500 internal server error",
            http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, s)
}

func category_index(w http.ResponseWriter, r *http.Request) {
    generic_index(w, r, Category{}, category_get())
}

func product_index(w http.ResponseWriter, r *http.Request) {
    generic_index(w, r, Product{}, struct {
        Products []Product
        Categories []Category
    } {product_get(), append([]Category{{}}, category_get()...)})
}


func generic_edit(w http.ResponseWriter, r *http.Request, t Model) {
    old := db_get_by_id(t, get_param(r, 0))

    tmpl, err := template.ParseFiles("html/" + t.tableName() + "/edit.html")
    if err != nil {
        http.Error(w, "500 internal server error",
            http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, old)
}

func category_edit(w http.ResponseWriter, r *http.Request) {
    generic_edit(w, r, Category{})
}

func product_edit(w http.ResponseWriter, r *http.Request) {
    generic_edit(w, r, Product{})
}


func generic_create(w http.ResponseWriter, r *http.Request, t Model) {
    old := r.PostForm

    tmpl, err := template.ParseFiles("html/" + t.tableName() + "/create.html")
    if err != nil {
        http.Error(w, "500 internal server error",
            http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, old)
}

func category_create(w http.ResponseWriter, r *http.Request) {
    generic_create(w, r, Category{})
}

func product_create(w http.ResponseWriter, r *http.Request) {
    generic_create(w, r, Product{})
}


func generic_store(w http.ResponseWriter, r *http.Request, t Model) {
    r.ParseForm()
    if t.parseForm(r) {
        db_store(t)
        http.Redirect(w, r, "/" + t.tableName(), http.StatusSeeOther)
    }
}

func category_store(w http.ResponseWriter, r *http.Request) {
    generic_store(w, r, Category{})
}

func product_store(w http.ResponseWriter, r *http.Request) {
    generic_store(w, r, Product{})
}


func generic_update(w http.ResponseWriter, r *http.Request, t Model) {
    r.ParseForm()
    if t.parseForm(r) {
        db_update(t, get_param(r, 0))
        http.Redirect(w, r, "/" + t.tableName(), http.StatusSeeOther)
    }
}

func category_update(w http.ResponseWriter, r *http.Request) {
    generic_update(w, r, Category{})
}

func product_update(w http.ResponseWriter, r *http.Request) {
    generic_update(w, r, Product{})
}

func generic_delete(w http.ResponseWriter, r *http.Request, t Model) {
    db_delete(t, get_param(r, 0))
    http.Redirect(w, r, "/" + t.tableName(), http.StatusSeeOther)
}

func category_delete(w http.ResponseWriter, r *http.Request) {
    generic_delete(w, r, Category{})
}

func product_delete(w http.ResponseWriter, r *http.Request) {
    generic_delete(w, r, Product{})
}


func generic_view[T Model](w http.ResponseWriter, r *http.Request, t T) {
    t = db_get_by_id(t, get_param(r, 0))
    tmpl, err := template.ParseFiles("html/" + t.tableName() + "/view.html")
    if err != nil {
        http.Error(w, "500 internal server error",
            http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, t)
}

func category_view(w http.ResponseWriter, r *http.Request) {
    generic_view(w, r, Category{})
}

func product_view(w http.ResponseWriter, r *http.Request) {
    generic_view(w, r, Product{})
}