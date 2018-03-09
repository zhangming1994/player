package main

import (
	"log"
	"net/http"
	"player/router"
)

func main() {
	router.Router()
	//静态文件
	http.Handle("/static/",http.StripPrefix("/static",http.FileServer(http.Dir("static"))))

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
