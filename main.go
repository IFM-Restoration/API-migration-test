package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"time"

	_ "github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

type WORK_ORDERS struct {
	ID                            uint
	ClientId                      uint
	ResidentId                    uint
	WorkOrderAcceptanceStatusId   uint
	Client_contact_approved_by_id uint
	Name                          string
	Uuid                          string
	Due_date                      time.Time
	Is_canceled                   bool
	Created_datetime              time.Time
	Created_by                    string
	Last_modified_datetime        time.Time
	Last_modified_by              string
	Is_deleted                    bool
}

type UserTest struct {
	gorm.Model
	FirstName string
	LastName  string
	Skills    []Skill
}

type Skill struct {
	gorm.Model
	SkillType string
	Level     string
	// pointer to UserTest struct, one-to-one, one-to-many relationships
	UserTestID uint
}

func main() {

	// Migration with Migrate package

	//   -mysql.dsn "amayer:ifmrestoration@tcp(localhost:5432)/api_migration_test?sslmode=disable"
	//m, err := migrate.New(
	//	"file://C://Users/amayer/GolandProjects/API-migration/db/migration",
	//	"postgres://amayer:ifmrestoration@localhost:5432/api_migration_test?sslmode=disable")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if err := m.Down(); err != nil {
	//	log.Fatal(err)
	//}

	// open DB
	log.Print("Connecting to Postgres...")
	connStr := "user=amayer password=ifmrestoration dbname=api_migration_test host=localhost sslmode=disable"
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	database := db.DB()
	err = database.Ping()
	if err != nil {
		panic(err.Error())
	}
	log.Println("DB successfully connected")

	//var workOrder = &WORK_ORDERS{
	//	01,
	//	1234,
	//	021,
	//	022,
	//	0,
	//	"IFM-CA12345",
	//	"01576148-eb2c-4cd2-b529-ecaa71f39893",
	//	time.Now(),
	//	false,
	//	time.Now(),
	//	"amayer",
	//	time.Now(),
	//	"",
	//	false}
	// create a record fpr work order table
	//db.Create(&workOrder)

	//var user = &UserTest{
	//	FirstName: "Anna",
	//	LastName: "Mayer",
	//}

	db.DropTableIfExists(&UserTest{}, &Skill{})
	// create a table
	db.CreateTable(&UserTest{}, &Skill{})

	//add data to table
	//db.Create(&user)

	// update table with logs
	//user.FirstName = "Constantine"
	//db.Debug().Save(&user)
	//var u UserTest
	//db.Debug().First(&u)
	//log.Println("First record with Constantine: ",u)
	//
	//// soft delete with logs (will update instead drop, only if we have gorm.Model)
	//db.Debug().Delete(&user)
	//
	//u = UserTest{}
	//// show first record
	//db.Debug().First(&u)
	//log.Println("First record: ",u)

	// create a table using struct
	//db.CreateTable(&WORK_ORDERS{})
	//log.Printf("Table created")

	//delete a table
	//db.DropTableIfExists(&WORK_ORDERS{})

}
