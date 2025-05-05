package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	dsn := "root:@tcp(localhost:3306)/authdb"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database ping error:", err)
	}
	fmt.Println("Connected to MySQL!")
}
