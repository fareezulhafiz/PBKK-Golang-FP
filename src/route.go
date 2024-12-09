package main

import (
    "context"
    "fmt"
    "net/http"
    "regexp"
)

type Route struct {
    method  string
    regex   *regexp.Regexp
    handler http.HandlerFunc
}

type URLParam struct{}

var routes = []Route{
    {"GET", regexp.MustCompile("^/$"), index},

    {"GET", regexp.MustCompile("^/categories/?$"), category_index},
    {"GET", regexp.MustCompile("^/categories/([0-9]+)/edit/?$"), category_edit},
    {"GET", regexp.MustCompile("^/categories/create/?$"), category_create},
    {"POST", regexp.MustCompile("^/categories/?$"), category_store},
    {"GET", regexp.MustCompile("^/categories/([0-9]+)/?$"), category_view},
    {"PUT", regexp.MustCompile("^/categories/([0-9]+)/?$"), category_update},
    {"DELETE", regexp.MustCompile("^/categories/([0-9]+)/?$"), category_delete},

    {"GET",    regexp.MustCompile("^/products/?$"), product_index},
    {"GET",    regexp.MustCompile("^/products/([0-9]+)/edit/?$"), product_edit},
    {"GET",    regexp.MustCompile("^/products/create/?$"), product_create},
    {"POST",   regexp.MustCompile("^/products/?$"), product_store},
    {"GET",    regexp.MustCompile("^/products/([0-9]+)/?$"), product_view},
    {"PUT",    regexp.MustCompile("^/products/([0-9]+)/?$"), product_update},
    {"DELETE", regexp.MustCompile("^/products/([0-9]+)/?$"), product_delete},
}

func route(w http.ResponseWriter, r *http.Request) { // TODO method loop
    failed_method := 0

    for _, rt := range routes {
        matches := rt.regex.FindStringSubmatch(r.URL.Path)
        if len(matches) > 0 {
            if r.Method != rt.method {
                failed_method += 1
                continue
            }
            rt.handler(w, r.WithContext(
                context.WithValue(r.Context(), URLParam{}, matches[1:])))
            return
        }
    }
    if failed_method != 0 {
        w.Header().Set("Allow", r.Method)
        http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
        fmt.Println(r.Method, r.URL.Path)
    } else {
        http.NotFound(w, r)
    }
}
