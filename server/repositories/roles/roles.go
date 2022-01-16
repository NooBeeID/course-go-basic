package repositories

import "go-web-template/server/models"

type RoleRepository interface {
	Save(role *models.Role) error
	FindAll() (*[]models.Role, error)
	FindByID(id string) (*models.Role, error)
	UpdateByID(id string, role *models.Role) error
	DeleteByID(id string) error
}
