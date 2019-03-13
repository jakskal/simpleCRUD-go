package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func init() {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	var err error
	db := string(os.Getenv("MYSQL_user") + ":" + os.Getenv("MYSQL_pass") + "@tcp(" + os.Getenv("MYSQL_dbhost") + ")/" + os.Getenv("MYSQL_dbName") + "?" + "charset=utf8")
	fmt.Println("database (", db, ")")
	DB, err = sql.Open("mysql", db)
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}
