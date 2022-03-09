package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func Init() {
	fmt.Println("Go MySQL Tutorial")
	username := os.Getenv("mysql_users_username")
	password := os.Getenv("mysql_users_password")
	host := os.Getenv("mysql_users_host")
	schema := os.Getenv("mysql_users_schema")
	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, schema)
	Client, err := sql.Open("mysql", dataSource)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	err = Client.Ping()
	if err != nil {
		return
	}

	fmt.Println("Database successfully configured")

	// defer the close till after the main function has finished
	// executing
	defer Client.Close()

}
