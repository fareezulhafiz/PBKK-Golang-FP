package main

import (
    "context"
    "fmt"
    "net/http"
    "strings"
    "time"

    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgxpool"
)

type Model interface {
    toString() string
    parseForm(*http.Request) bool
}

var db *pgxpool.Pool

func db_connect() error {
    var err error

    for i := 0; i < 60; i++ {
        db, err = pgxpool.New(context.Background(), "")
        if err == nil && db.Ping(context.Background()) == nil {
            break
        }
        fmt.Println("cannot connect to db, trying again")
        time.Sleep(time.Second)
    }
    return err
}

func db_create(table string, args []string) {
    db.Exec(context.Background(),
        "CREATE TABLE IF NOT EXISTS " + table +
        "(" + strings.Join(args, ",") + ")")
}

func db_drop(table string) {
    db.Exec(context.Background(), "DROP TABLE $1", table)
}

func db_get[T Model](table string, t T) []T {
    rows, err := db.Query(context.Background(), "SELECT * FROM " + table)

    if err != nil {
        return nil
    }
    res, err := pgx.CollectRows(rows, pgx.RowToStructByName[T])
    return res
}

func db_get_by_id[T Model](table string, t T, id int) T {
    row, err := db.Query(context.Background(),
        "SELECT * FROM " + table + "WHERE id=$1", id)

    if err != nil {
        return t
    }
    res, _ := pgx.CollectOneRow(row, pgx.RowToStructByName[T])
    return res
}

func db_store[T Model](table string, t T) {
    db.Exec(context.Background(), "INSERT INTO " + table +
        " VALUES(" + t.toString() + ")")
}

//func db_update[T Model](table string, t T) {
//    db.Exec(context.Background(), "UPDATE " + table +
//        "VALUES(" + t.toString() + ")" + "WHERE id=$1", t.id)
//}
