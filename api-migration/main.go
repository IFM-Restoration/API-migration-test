package main

import (
	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"log"
)

func main() {

	// Migration with Migrate package

	//   -mysql.dsn "amayer:ifmrestoration@tcp(localhost:5432)/api_migration_test?sslmode=disable"

	log.Println("Hello!")
	m, err := migrate.New(

		"file://C://Users/amayer/GolandProjects/API-migration/db/migration",
		"postgres://amayer:ifmrestoration@localhost:5432/api_migration_test?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Migration completed")

}
