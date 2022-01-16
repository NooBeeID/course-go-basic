package server

import (
	"database/sql"
	"fmt"
	"go-web-template/server/controllers"
	"net/http"
)

func StartServer(router *http.ServeMux, port string, db *sql.DB) {
	buildRoute(router, db)

	fileServer := http.FileServer(http.Dir("static/assets"))

	router.Handle("/static/", http.StripPrefix("/static", fileServer))

	fmt.Println("Server running at", port)
	http.ListenAndServe(port, router)
}

func buildRoute(router *http.ServeMux, db *sql.DB) {
	homeRoute(router, db)
	employeeRoute(router, db)
}

func homeRoute(router *http.ServeMux, db *sql.DB) {
	homeController := controllers.NewHomeController()

	router.HandleFunc("/", homeController.Index)

}

func employeeRoute(router *http.ServeMux, db *sql.DB) {
	employeeController := controllers.NewEmployeeController(db)

	router.HandleFunc("/employees", employeeController.Index)
	router.HandleFunc("/employees/update", employeeController.UpdateByID)
	router.HandleFunc("/employees/add", employeeController.Add)
	router.HandleFunc("/employees/delete", employeeController.DeleteByID)
}
