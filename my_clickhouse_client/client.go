package main

import (
	"database/sql"
	"fmt"
	"github.com/kshvakov/clickhouse"
	"time"
)

func main() {

	db, err := sql.Open("clickhouse", "tcp://192.168.125.100:9000?debug=true")
	checkErr(err)

	go func() {
		for {
			fmt.Println(db.Stats())
			time.Sleep(1e9)
		}
	}()

	if err := db.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			checkErr(err)
		}
	}

	time.Sleep(1e9)

	for i := 0; i < 5; i++ {
		go func() {
			rows, err := db.Query("select id from test.vehicle")
			checkErr(err)
			for rows.Next() {
				var id string
				if err := rows.Scan(&id); err == nil {
					fmt.Println(id)
				} else {
					fmt.Println(err)
				}
			}
		}()
	}

	time.Sleep(20e9)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
