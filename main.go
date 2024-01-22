package main

import (
	"hello_goland/common"
	"hello_goland/router"
	"log"
	"net/http"
)

func init() {
	common.LoadTemplate()
}

func main() {
	//web应用
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

}
