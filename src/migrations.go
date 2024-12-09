package main

func migrations_up() {
    Category{}.migrate()
    Product{}.migrate()
}

func migrations_down() {
    Product{}.drop()
    Category{}.drop()
}
