package main

func migrations_up() {
    category_migrate()
    product_migrate()
}

func migrations_down() {
    category_drop()
    product_drop()
}
