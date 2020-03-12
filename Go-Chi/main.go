package main

import (
	"Go-Chi/router"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	router.Post_router()
	http.ListenAndServe(":8005", router.Logger())
}