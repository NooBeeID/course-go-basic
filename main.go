package main

import (
	"go-web-template/server"
	"go-web-template/server/config"
	"net/http"
)

func main() {
	run()
}

func run() {
	// ini gunanya untuk membuat sebuah router.
	// router ini akan menghandle seluruh route dari aplikasi kita
	router := http.NewServeMux()

	port := ":9999"

	db := config.CreateConnection()

	// kirim router dan port ke fungsi StartServer
	server.StartServer(router, port, db)
}
