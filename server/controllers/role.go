package controllers

import (
	"database/sql"
	"fmt"
	apps "go-web-template/server/apps/web"
	params "go-web-template/server/params/roles"
	"go-web-template/server/services"
	"go-web-template/server/utils"
	"log"
	"net/http"
	"path"
	"text/template"
)

type RoleController interface {
	Index(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
}

type roleController struct {
	DB *sql.DB
}

func NewRoleController(db *sql.DB) RoleController {
	return &roleController{
		DB: db,
	}
}

func (role *roleController) Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("static", "pages/roles/index.html"), utils.LayoutMaster)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// roleServices := services.NewRoleServices(role.DB)

	// roles, err := roleServices.

	web := apps.RenderWeb{
		Title: "Halaman Roles",
	}
	err = tmpl.Execute(w, web)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (role *roleController) Add(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	if method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("static", "pages/roles/add.html"), utils.LayoutMaster)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		web := apps.RenderWeb{
			Title: "Tambah Role",
		}

		err = tmpl.Execute(w, web)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if method == "POST" {

		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var request params.RoleCreate = params.RoleCreate{
			Name: r.Form.Get("name"),
		}

		roleServices := services.NewRoleServices(role.DB)

		_ = roleServices.CreateNewRole(r.Context(), &request)

		w.Write([]byte(`
			<script>
				alert("Tambah data roles berhasil !");
				window.location.href="../roles"
			</script>
		`))

	} else {
		msg := fmt.Sprintf("Method %s tidak di perbolehkan", method)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
}
