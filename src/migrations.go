package main

func MigrationsUp() {
    Category{}.Migrate()
    Product{}.Migrate()
}

func MigrationsDown() {
    Product{}.Drop()
    Category{}.Drop()
}
