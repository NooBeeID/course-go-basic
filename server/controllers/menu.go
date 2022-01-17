package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	apps "go-web-template/server/apps/api"
	"go-web-template/server/helper"
	params "go-web-template/server/params/menu"
	"go-web-template/server/services"
	"net/http"

	"github.com/google/uuid"
)

type MenuController interface {
	FindAll(w http.ResponseWriter, r *http.Request)
	FindByID(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
	DeleteByID(w http.ResponseWriter, r *http.Request)
	UpdateByID(w http.ResponseWriter, r *http.Request)
}

type menuController struct {
	DB *sql.DB
}

func NewMenuController(db *sql.DB) MenuController {
	return &menuController{
		DB: db,
	}
}

func (m *menuController) FindAll(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "GET" {
		menus := services.NewMenuService(m.DB).GetAllMenu()

		success := apps.ResponseSuccess{
			Status: http.StatusOK,
			Data:   menus,
		}

		w.WriteHeader(success.Status)
		json.NewEncoder(w).Encode(success)

	} else {
		helper.HandleNotMethodAllowed(w, method)
	}
}

func (m *menuController) FindByID(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "GET" {

		id := r.URL.Query().Get("id")

		menu := services.NewMenuService(m.DB).GetMenuByID(id)

		success := apps.ResponseSuccess{
			Status: http.StatusOK,
			Data:   menu,
		}

		w.WriteHeader(success.Status)
		json.NewEncoder(w).Encode(success)

	} else {
		helper.HandleNotMethodAllowed(w, method)
	}
}

func (m *menuController) Add(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	if method == "POST" {
		err := r.ParseForm()
		if err != nil {
			helper.HandleBadRequest(w, err)
			return
		}

		var request params.MenuCreate

		err = json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			helper.HandleBadRequest(w, err)
			return
		}

		fmt.Println(request)

		isSuccess := services.NewMenuService(m.DB).CreateNewMenu(&request)

		if isSuccess {
			var response apps.ResponseSuccess
			response.Data = "Create new menu success !"
			response.Status = http.StatusCreated

			w.WriteHeader(response.Status)
			json.NewEncoder(w).Encode(response)
			return
		} else {
			var response apps.ResponseFail
			response.Status = http.StatusBadRequest
			response.Message = "Bad Request"

			w.WriteHeader(response.Status)
			json.NewEncoder(w).Encode(response)
			return
		}

	} else if method == "OPTIONS" {
		return
	} else {
		helper.HandleNotMethodAllowed(w, method)
	}
}

func (m *menuController) DeleteByID(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "DELETE" {
		id := r.URL.Query().Get("id")

		isSuccess := services.NewMenuService(m.DB).DeleteMenuByID(id)

		if isSuccess {
			success := apps.ResponseSuccess{
				Status: http.StatusOK,
				Data:   "delete success",
			}

			w.WriteHeader(success.Status)
			json.NewEncoder(w).Encode(success)
		} else {
			helper.HandleBadRequest(w, errors.New("delete fail"))
		}
	} else {
		helper.HandleNotMethodAllowed(w, method)
		return
	}
}

func (m *menuController) UpdateByID(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "PUT" {
		id := r.URL.Query().Get("id")

		newID, err := uuid.Parse(id)
		if err != nil {
			helper.HandleBadRequest(w, err)
			return
		}

		var request params.MenuUpdate

		err = json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			helper.HandleBadRequest(w, err)
			return
		}

		request.ID = newID

		isSuccess := services.NewMenuService(m.DB).UpdateMenuByID(&request)

		if isSuccess {
			success := apps.ResponseSuccess{
				Status: http.StatusOK,
				Data:   "update success",
			}

			w.WriteHeader(success.Status)
			json.NewEncoder(w).Encode(success)
		} else {
			helper.HandleBadRequest(w, errors.New("update fail"))
		}
	} else {
		helper.HandleNotMethodAllowed(w, method)
	}
}
