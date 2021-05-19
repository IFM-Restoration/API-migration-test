package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	C "ifm.com/b1/models"
	"log"
	"time"

	_ "github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

type Config struct {
	SkipDefaultTransaction                   bool
	NamingStrategy                           schema.Namer
	Logger                                   logger.Interface
	NowFunc                                  func() time.Time
	DryRun                                   bool
	PrepareStmt                              bool
	DisableNestedTransaction                 bool
	AllowGlobalUpdate                        bool
	DisableAutomaticPing                     bool
	DisableForeignKeyConstraintWhenMigrating bool
}

func main() {

	// open DB
	log.Print("Connecting to Postgres...")
	//connStr := "user=amayer password=ifmrestoration dbname=api_migration_test2 host=localhost sslmode=disable"
	connStr := "user=Tester password=Tester123! dbname=postgres host=34.68.197.254 port=5432 sslmode=disable"
	db, err := gorm.Open("postgres", connStr, &Config{NamingStrategy: schema.NamingStrategy{SingularTable: true, NoLowerCase: true}})
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.LogMode(true)

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
	// create a record for work order table
	//db.Create(&workOrder)

	//var user = &UserTest{
	//	FirstName: "Anna",
	//	LastName: "Mayer",
	//}

	db.DropTableIfExists(&C.Client{})

	// create a table
	db.CreateTable(
		&C.Client{})
	//&C.Market{},
	//&MajorTrade{},
	//&InteriorExterior{},
	//&Priority{},
	//&WorkOrderAcceptanceStatus{},
	//&ExceptionReasonCode{},
	//&JobNeedType{},
	//&JobNoteType{},
	//&JobRescheduleReason{},
	//&JobScheduleActivityStatus{},
	//&TraitType{},
	//&State{},
	//&StateExitReason{},
	//&DelayReason{},
	//&AttachmentType{},
	//&AttachmentCategory{},
	//&AttachmentExtension{},
	//&NeedsRevisionReason{},
	//&ServiceCategory{},
	//&CustomerStatus{},
	//&LaborType{},
	//&Material{},
	//&Unit{},
	//&QbAccess{},
	//&Permissions{},
	//&Roles{},
	//&Users{},
	//&WorkOrder{},
	//&Resident{},
	//&RolePermission{},
	//&UserClientRole{},
	//&ClientContact{},
	//&Trait{},
	//&ClientGeography{},
	//&WorkOrderDetail{},
	//&WorkOrderState{},
	//&WorkOrderLocation{},
	//&Estimate{},
	//&Job{},
	//&Customer{},
	//&CustomerTrait{},
	//&ServiceDetail{},
	//&WorkOrderMajorTrade{},
	//&StateWorkOrderHistory{},
	//&ExceptionTransaction{},
	//&EstimateNeedsRevisionReason{},
	//&LineItem{},
	//&Pclsm{},
	//&JobNeed{},
	//&JobState{},
	//&JobTrait{},
	//&Invoice{},
	//&ResidentContactAttempt{},
	//&JobNote{},
	//&StateJobHistory{},
	//&ResidentContactInformation{},
	//&JobScheduleActivity{}

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
