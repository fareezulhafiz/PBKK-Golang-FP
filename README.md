# PBKK-Golang-FP

#Group Members
1) Ange, Simon, Axel, Duhayon (5999241027)
2) Muhammad Fareezul Hafiz bin Mohd Nasir (5999241099)

## Quick Start
```bash
go build .
./app
```

## Plan

The goal of this application is to make a small database of products in a
shop.<br>

The main page is a `products` page where we can `see` and `add` new products as
well as `edit` existing ones. Every product must have a `category`, a `name`
(at least 3 characters), a `price` and optionally a `description`.<br>

If no category exist, it's possible to create one either from the `categories`
page or directly when creating a new product, with the `create` button besides
`category`.<br>

The category must have a unique `name` (at least 3 characters) and
optionally a `description`. Every product in a category is linked to the parent
category, so if this one is modified, every product is also updated. If the
category is deleted, every subproduct is deleted as well.<br>

You can see a demo of the application on https://youtu.be/qUHR13qM2gM
