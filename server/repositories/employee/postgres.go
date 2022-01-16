package repositories

import (
	"database/sql"
	"errors"
	"go-web-template/server/models"
)

type employeeRepo struct {
	DB *sql.DB
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepo{
		DB: db,
	}
}

func (e *employeeRepo) Save(employee *models.Employee) error {
	query := `
		INSERT INTO employees (
			id, nip, name, address, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
	`

	// kita akan melakukan prepare.
	// hal ini bertujuan untuk melakukan peningkatan performa
	// dan keamanan
	stmt, err := e.DB.Prepare(query)

	if err != nil {
		return err
	}

	// jangan lupa untuk menutup koneksi statements
	defer stmt.Close()

	_, err = stmt.Exec(
		employee.ID, employee.NIP, employee.Name,
		employee.Address, employee.CreatedAt, employee.UpdatedAt,
	)

	return err
}

func (e *employeeRepo) FindAll() (*[]models.Employee, error) {
	query := `
		SELECT 
			id, nip, name, address, created_at, updated_at
		FROM
			employees
	`

	stmt, err := e.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var employees []models.Employee

	for rows.Next() {
		var employee models.Employee
		err := rows.Scan(
			&employee.ID, &employee.NIP, &employee.Name,
			&employee.Address, &employee.CreatedAt, &employee.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return &employees, nil

}

func (e *employeeRepo) FindByID(id string) (*models.Employee, error) {
	query := `
		SELECT 
			id, nip, name, address
		FROM
			employees
		WHERE
			id=$1
	`
	stmt, err := e.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	var employee models.Employee

	err = row.Scan(
		&employee.ID, &employee.NIP, &employee.Name, &employee.Address,
	)

	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func (e *employeeRepo) UpdateByID(id string, employee *models.Employee) error {
	panic(errors.New(""))
}

func (e *employeeRepo) DeleteByID(id string) error {
	query := `
		DELETE FROM employees
		WHERE id=$1
	`

	stmt, err := e.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err
}