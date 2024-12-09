package main

import (
    "context"
    "fmt"
    "reflect"
    "strings"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgxpool"
)

type Model interface {
    tableName() string
    migrate()
    drop()
    validate(c *gin.Context) any
}

func db_exec(sql string, values ...any) {
    fmt.Println(sql, values)
    db.Exec(context.Background(), sql, values...)
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

func db_create(m Model, args []string) {
    sql := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS "%s" (%s)`,
        m.tableName(), strings.Join(args, ","))

    db_exec(sql)
}

func db_drop(m Model) {
    sql := fmt.Sprintf(`DROP TABLE "%s"`, m.tableName())

    db_exec(sql)
}

func db_get[T Model](m T) []T {
    var res []T
    var err error
    var rows pgx.Rows
    var sql string

    sql = fmt.Sprintf(`SELECT * FROM "%s"`, m.tableName())
    rows, err = db.Query(context.Background(), sql)
    if err != nil {
        return nil
    }
    res, _ = pgx.CollectRows(rows, pgx.RowToStructByName[T])
    return res
}

func db_get_by[T Model](m T, f string, v string) T {
    sql := fmt.Sprintf(`SELECT * FROM "%s" WHERE "%s"=$1`, m.tableName(), f)
    row, err := db.Query(context.Background(), sql, v)

    if err != nil {
        return m
    }
    res, _ := pgx.CollectOneRow(row, pgx.RowToStructByName[T])
    return res
}

func db_store[T Model](m T) {
    v := reflect.ValueOf(m)
    keys := make([]string, v.NumField() - 1)
    fields := make([]string, v.NumField() - 1)
    values := make([]any, v.NumField() - 1)

    for i := 1; i < v.NumField(); i++ {
        keys[i - 1] = fmt.Sprintf("$%v", i)
        fields[i - 1] = v.Type().Field(i).Name
        values[i - 1] = fmt.Sprintf("%v", v.Field(i))
    }
    sql := fmt.Sprintf(`INSERT INTO "%s" (%s) VALUES (%s)`,
        m.tableName(), strings.Join(fields, ","), strings.Join(keys, ","))
    db_exec(sql, values...)
}

func db_update[T Model](m T, id string) {
    v := reflect.ValueOf(m)
    keys := make([]string, v.NumField() - 1)
    fields := make([]string, v.NumField() - 1)
    values := make([]any, v.NumField())

    values[0] = id
    for i := 1; i < v.NumField(); i++ {
        keys[i - 1] = fmt.Sprintf("$%v", i + 1)
        fields[i - 1] = v.Type().Field(i).Name
        values[i] = fmt.Sprintf("%v", v.Field(i))
    }
    sql := fmt.Sprintf(`UPDATE "%s" SET (%s) = (%s) WHERE id=$1`,
        m.tableName(), strings.Join(fields, ","), strings.Join(keys, ","))
    db_exec(sql, values...)
}

func db_delete(m Model, f string, v string) {
    sql := fmt.Sprintf(`DELETE FROM "%s" WHERE "%s"=$1`, m.tableName(), f)

    db_exec(sql, v)
}
