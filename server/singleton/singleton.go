package singleton

import (
	"fmt"
	"log"

	"github.com/OsakiTsukiko/frogpond/server/config"
	"github.com/OsakiTsukiko/frogpond/server/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// global scope variables
var CFG config.Config
var DATABASE *gorm.DB

func initDatabase() {
	var err error
	// postgreSQL connection string
	dsn := "user=" + CFG.DataBase.Username +
		" password=" + CFG.DataBase.Password +
		" dbname=" + CFG.DataBase.Database +
		" host=" + CFG.DataBase.Host +
		" port=" + CFG.DataBase.Port +
		" sslmode=disable" // TODO: ADD OPTION FOR THIS
		// at the moment not required as i eitehr run the
		// database on the same machine or bridge the
		// network

	// connect to PostgreSQL database
	DATABASE, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("🚩 Failed to connect to database: %v", err)
	}

	// migrate the user model (creates the table if it doesn't exist)
	err = DATABASE.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatalf("🚩 Failed to migrate database: %v", err)
	}

	// migrate the token model (creates the table if it doesn't exist)
	err = DATABASE.AutoMigrate(&domain.Token{})
	if err != nil {
		log.Fatalf("🚩 Failed to migrate tokens table: %v", err)
	}
}

func listTablesAndEntries() {
	// Query to get the list of all tables
	var tables []struct {
		TableName string `gorm:"column:table_name"`
	}
	if err := DATABASE.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'").Scan(&tables).Error; err != nil {
		log.Fatalf("Error fetching table list: %v", err)
	}

	// Loop through each table and retrieve all rows
	for _, table := range tables {
		fmt.Println("Table:", table.TableName)
		var results []map[string]interface{}
		if err := DATABASE.Raw("SELECT * FROM " + table.TableName).Scan(&results).Error; err != nil {
			log.Printf("Error fetching data from table %s: %v", table.TableName, err)
			continue
		}

		// Print all rows from the table
		for _, row := range results {
			fmt.Println(row)
		}
	}
}

// init is run when the program starts
func init() {
	log.Println("🐸 Initializing FrogPond Singleton")
	CFG = config.LoadConfig() // load config from environment variables
	initDatabase()
	listTablesAndEntries() // TODO: MAKE DEBUG ONLY (ADD FLAG)
}
