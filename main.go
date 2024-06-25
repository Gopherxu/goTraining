package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func getConnection() *sql.DB {
	//emp := &Empl{}
	//b, err := json.Marshal(e1)
	//fmt.Print("\nreadingfromDB Row 2", string(b))
	host := "127.0.0.1"
	port := "5432"
	user := "postgres"
	password := "postgres"
	dbname := "employee"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	connection, err := sql.Open("postgres", psqlInfo)
	fmt.Print("\nreadingfromDB Result 1", connection)
	fmt.Print("\nreadingfromDB Error 1", err)
	return connection
}
func main() {
	getConnection()
}
