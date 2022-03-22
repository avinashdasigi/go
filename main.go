package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type employee struct {
	id   int
	name string
}

func main() {
	fmt.Println("Hello World")
	db, err := sql.Open("mysql", "root:Hello@123@tcp(127.0.0.1:3306)/ems")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Successfully connected to Database")

	/*insert, err := db.Query("INSERT INTO Employee VALUES(2, 'Ram')")
	if err != nil {
		panic(err)
	}
	defer insert.Close()*/

	results, err := db.Query("SELECT *FROM Employee")
	if err != nil {
		panic(err)
	}
	fmt.Println(results)
	for results.Next() {
		var emp employee
		err = results.Scan(&emp.id, &emp.name)
		if err != nil {
			panic(err)
		}
		fmt.Println(emp.id, emp.name)
	}
	defer results.Close()

}
