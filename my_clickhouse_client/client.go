package main

import (
	"database/sql"
	"fmt"
	log "github.com/gogap/logrus"
	"github.com/gogap/logrus/hooks/file"
	"github.com/kshvakov/clickhouse"
	"time"
)

func init() {
	log.SetLevel(log.InfoLevel)
	//log.SetFormatter(&log.JSONFormatter{})
	log.AddHook(file.NewHook("logs/client1.log"))
}

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
			log.Errorf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			checkErr(err)
		}
	}

	time.Sleep(1e9)

	for i := 0; i < 5; i++ {
		go func() {
			rows, err := db.Query("select * from test.vehicle")
			time.Sleep(2e9)
			checkErr(err)
			for rows.Next() {
				var d time.Time
				var plateNumber string
				var captureTime time.Time
				if err := rows.Scan(&d, &plateNumber, &captureTime); err == nil {
					fmt.Println(d, plateNumber, captureTime)
				} else {
					log.Errorf(err.Error())
				}
			}
		}()
		//time.Sleep(1e9)
	}

	time.Sleep(20e9)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		log.Error(err)
	}
}
