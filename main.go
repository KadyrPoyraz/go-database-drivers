package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

func main() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	databaseUrl := "postgres://username:password@localhost:6969/db"
	conn, err := pgx.Connect(context.Background(), databaseUrl)

	// Send the query to the server. The returned rows MUST be closed
	// before conn can be used again.
	rows, err := conn.Query(context.Background(), "select name, weight from aboba")
	if err != nil {
		panic(err)
	}

	fmt.Println(rows)

	type Item struct{
		name string
		weight int
	}
	var result []Item

	defer rows.Close()

	for rows.Next() {
		var itm Item

		err = rows.Scan(&itm.name, &itm.weight)
		if err != nil {
			panic(err)
		}

		result = append(result, itm)
	}

	fmt.Println(result)

	// Any errors encountered by rows.Next or rows.Scan will be returned here
	if rows.Err() != nil {
		panic(err)
	}

	// No errors found - do something with sum
}
