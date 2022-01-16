package services

import (
	"context"
	"database/sql"
	"go-web-template/server/helper"
	"go-web-template/server/models"
	params "go-web-template/server/params/roles"
	repositories "go-web-template/server/repositories/roles"
)

type RoleServices struct {
	RoleRepository repositories.RoleRepository
	DB             *sql.DB
}

func NewRoleServices(db *sql.DB) *RoleServices {
	repository := repositories.NewRoleRepository(db)
	return &RoleServices{
		RoleRepository: repository,
		DB:             db,
	}
}

func (r *RoleServices) CreateNewRole(ctx context.Context, request *params.RoleCreate) *params.RoleAddResponse {
	role := parseRequestToModel(request)
	err := r.RoleRepository.Save(role)

	defer helper.HandleError()
	helper.HandlePanicIfError(err)

	return &params.RoleAddResponse{
		Message: "Add success !",
	}
}

func (r *RoleServices) GetRoles(ctx context.Context) *params.RoleListview {
	defer helper.HandleError()
	_, err := r.RoleRepository.FindAll()

	helper.HandlePanicIfError(err)

	var params params.RoleListview

	return &params

}

func parseRequestToModel(request *params.RoleCreate) *models.Role {
	role := models.NewRole()
	role.Name = request.Name
	return role
}
