package main

import (
	"log"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

func main() {

	println("Hello")

	//   -mysql.dsn "amayer:ifmrestoration@tcp(localhost:5432)/api_migration?sslmode=disable"

	m, err := migrate.New(
		"file://C://Users/amayer/GolandProjects/API-migration/db/migration",
		"postgres://amayer:ifmrestoration@localhost:5432/api_migration_test?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Down(); err != nil {
		log.Fatal(err)
	}

}
