package main

import (
    "context"
    "fmt"
    "net/http"
    "reflect"
    "strings"
    "time"

    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgxpool"
)

type Model interface {
    tableName() string
    parseForm(*http.Request) bool
}

func db_exec(sql string, values ...any) {
    fmt.Println(sql, values)
    fmt.Println(db.Exec(context.Background(), sql, values...))
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

func db_create(t Model, args []string) {
    sql := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS "%s" (%s)`,
        t.tableName(), strings.Join(args, ","))

    db_exec(sql)
}

func db_drop(t Model) {
    sql := fmt.Sprintf(`DROP TABLE "$1"`, t.tableName())

    db_exec(sql)
}

func db_get[T Model](t T) []T {
    sql := fmt.Sprintf(`SELECT * FROM "%s"`, t.tableName())
    rows, err := db.Query(context.Background(), sql)

    if err != nil {
        return nil
    }
    res, err := pgx.CollectRows(rows, pgx.RowToStructByName[T])
    return res
}

func db_get_by_id[T Model](t T, id string) T {
    sql := fmt.Sprintf(`SELECT * FROM "%s" WHERE id=$1`, t.tableName())
    row, err := db.Query(context.Background(), sql, id)

    if err != nil {
        return t
    }
    res, _ := pgx.CollectOneRow(row, pgx.RowToStructByName[T])
    return res
}

func db_store(t Model) {
    v := reflect.ValueOf(t)
    keys := make([]string, v.NumField())
    values := make([]any, v.NumField())

    for i := 0; i < v.NumField(); i++ {
        keys[i] = fmt.Sprintf("$%v", i + 1)
        values[i] = fmt.Sprintf("%v", v.Field(i))
    }
    values[0] = "DEFAULT"
    sql := fmt.Sprintf(`INSERT INTO "%s" VALUES (%s)`,
        t.tableName(), strings.Join(keys, ","))
    db_exec(sql, values...)
}

func db_update(t Model, id string) {
    //v := reflect.ValueOf(t)
    //keys := make([]string, v.NumField())
    //fields := make([]string, v.NumField())
    //values := make([]any, v.NumField() + 1)

    //values[0] = id
    //for i := 0; i < v.NumField(); i++ {
    //    keys[i] = fmt.Sprintf("$%v", i + 2)
    //    fields[i] = v.Type().Field(i).Name
    //    values[i + 1] = fmt.Sprintf("%v", reflect.ValueOf(v.Field(i).Interface()))
    //}
    //sql := fmt.Sprintf(`UPDATE "%s" SET (%s) = (%s) WHERE id=$1`,
    //    t.tableName(), strings.Join(fields, ","), strings.Join(keys, ","))
    //db_exec(sql, values...)
}

func db_delete(t Model, id string) {
    sql := fmt.Sprintf(`DELETE FROM "%s" WHERE id=$1`, t.tableName())

    db_exec(sql, id)
}
