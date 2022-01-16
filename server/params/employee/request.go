package params

import "go-web-template/server/models"

type EmployeeCreate struct {
	NIP     string
	Name    string
	Address string
}

func (e *EmployeeCreate) ParseToModel() *models.Employee {
	employee := models.NewEmployee()
	employee.Address = e.Address
	employee.Name = e.Name
	employee.NIP = e.NIP
	return employee
}
