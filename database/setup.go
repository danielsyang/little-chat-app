package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	databaseUrl := "host=localhost user=postgres password=1234567890 dbname=little_chat_app port=5555 sslmode=disable"
	conn, err := sql.Open("postgres", databaseUrl)

	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	if err != nil {
		panic(err)
	}

	return conn
}
