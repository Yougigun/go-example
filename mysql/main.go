package main

import (
	"database/sql"
	"fmt"
	// "sync"
	"time"

	// "fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/test")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Prepare statement for inserting data
	stmtIns, err := db.Prepare(
		"INSERT INTO userinfo_000(uid,token,bnc_uuid,aid,device_id,platform,os_version,sdk_version,app_version)"+
		" VALUES( ?, ?, ?, ?, ?, ?, ?, ?, ? )") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	// // Prepare statement for reading data
	// stmtOut, err := db.Prepare("SELECT squareNumber FROM squarenum WHERE number = ?")
	// if err != nil {
	// 	panic(err.Error()) // proper error handling instead of panic in your app
	// }
	// defer stmtOut.Close()

	// Insert square numbers for 0-24 in the database
	t := time.Now()
	n := 10000
	// var wg sync.WaitGroup
	// wg.Add(n)
	for i := 0; i < n; i++ {
		_, err = stmtIns.Exec(1,i,1,1,1,1,1,1,1,) // Insert tuples (i, i^2)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// go func(i int){
		// 	defer wg.Done()
		// 	_, err = stmtIns.Exec(1,i,1,1,1,1,1,1,1,) // Insert tuples (i, i^2)
		// 	if err != nil {
		// 		panic(err.Error()) // proper error handling instead of panic in your app
		// 	}
		// }(i)
	}
	// wg.Wait()
	fmt.Println(time.Since(t).Milliseconds())

	// var squareNum int // we "scan" the result in here

	// // Query the square-number of 13
	// err = stmtOut.QueryRow(13).Scan(&squareNum) // WHERE number = 13
	// if err != nil {
	// 	panic(err.Error()) // proper error handling instead of panic in your app
	// }
	// fmt.Printf("The square number of 13 is: %d", squareNum)

	// // Query another number.. 1 maybe?
	// err = stmtOut.QueryRow(1).Scan(&squareNum) // WHERE number = 1
	// if err != nil {
	// 	panic(err.Error()) // proper error handling instead of panic in your app
	// }
	// fmt.Printf("The square number of 1 is: %d", squareNum)
}
