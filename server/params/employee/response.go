package params

import (
	"go-web-template/server/models"

	"github.com/google/uuid"
)

type EmployeeSingleView struct {
	ID        uuid.UUID
	NIP       string
	Name      string
	Address   string
	CreatedAt string
	UpdatedAt string
}

func (e *EmployeeSingleView) makeSingleView(model models.Employee) {

}
