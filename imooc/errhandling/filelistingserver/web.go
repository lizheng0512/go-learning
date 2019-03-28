package main

import (
	"github.com/lizheng0512/go-learning/imooc/errhandling/filelistingserver/filelisting"
	"net/http"
)

func main() {
	http.HandleFunc("/list/", filelisting.HandleFileList)

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
