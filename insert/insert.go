package insert

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func InsertTest(conn *pgx.Conn, context context.Context) {
    fmt.Println(conn)
    comm, err := conn.Exec(context, "INSERT INTO aboba (name, weight) VALUES ('Some name', 1234)")
    if err != nil {
        panic(err)
    }
    
    fmt.Println(comm)
}
