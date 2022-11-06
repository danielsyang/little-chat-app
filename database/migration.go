package database

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func RunMigrations() {
	db := GetConnection()
	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file:///Users/danielyang/Documents/new_stuff/little-chat-app/migrations/", "postgres", driver)

	if err != nil {
		panic(err)
	}

	m.Up()

	println("migrated")
}
