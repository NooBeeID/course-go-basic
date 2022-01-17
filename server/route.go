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
	menuRouteAPI(router, db)
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

func menuRouteAPI(router *http.ServeMux, db *sql.DB) {
	menuController := controllers.NewMenuController(db)

	router.HandleFunc("/api/menus", menuController.FindAll)
	router.HandleFunc("/api/menus/add", menuController.Add)
	router.HandleFunc("/api/menus/id", menuController.FindByID)
	router.HandleFunc("/api/menus/update", menuController.UpdateByID)
	router.HandleFunc("/api/menus/delete", menuController.DeleteByID)
}
