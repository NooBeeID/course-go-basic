package repositories

import (
	"database/sql"
	"go-web-template/server/models"
)

type roleRepo struct {
	DB *sql.DB
}

func NewRoleRepository(db *sql.DB) RoleRepository {
	return &roleRepo{
		DB: db,
	}
}

func (r *roleRepo) Save(role *models.Role) error {
	query := `
		INSERT INTO roles (name)
		VALUES ($1)
	`

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(role.Name)
	if err != nil {
		return err
	}

	return nil
}

func (r *roleRepo) FindAll() (*[]models.Role, error) {
	query := `
		SELECT id, name
		FROM roles
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var roles []models.Role

	for rows.Next() {
		var role models.Role
		err = rows.Scan(
			&role.ID, &role.Name,
		)
		if err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	return &roles, nil
}

func (r *roleRepo) FindByID(id string) (*models.Role, error) {
	return nil, nil
}

func (r *roleRepo) UpdateByID(id string, role *models.Role) error {
	return nil
}

func (r *roleRepo) DeleteByID(id string) error {
	return nil
}
