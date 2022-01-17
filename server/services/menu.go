package services

import (
	"database/sql"
	"go-web-template/server/helper"
	"go-web-template/server/models"
	params "go-web-template/server/params/menu"
	repositories "go-web-template/server/repositories/menu"
)

type MenuServices struct {
	MenuRepository repositories.MenuRepository
	DB             *sql.DB
}

func NewMenuService(db *sql.DB) *MenuServices {
	repository := repositories.NewMenyRepository(db)
	return &MenuServices{
		MenuRepository: repository,
		DB:             db,
	}
}

func (m *MenuServices) CreateNewMenu(request *params.MenuCreate) bool {
	defer helper.HandleError()
	menu := request.ParseToModel()

	err := m.MenuRepository.Save(menu)
	helper.HandlePanicIfError(err)

	return true
}

func (m *MenuServices) GetAllMenu() *[]params.MenuSingleView {
	defer helper.HandleError()
	menus, err := m.MenuRepository.FindAll()

	helper.HandlePanicIfError(err)

	return makeMenuListView(menus)

}

func (m *MenuServices) GetMenuByID(id string) *params.MenuSingleView {
	defer helper.HandleError()
	menu, err := m.MenuRepository.FindByID(id)
	helper.HandlePanicIfError(err)

	return makeMenuSingleView(menu)
}

func (m *MenuServices) DeleteMenuByID(id string) bool {
	defer helper.HandleError()

	err := m.MenuRepository.DeleteByID(id)
	helper.HandlePanicIfError(err)

	return true
}

func (m *MenuServices) UpdateMenuByID(requrest *params.MenuUpdate) bool {
	defer helper.HandleError()
	model := requrest.ParseToModel()
	err := m.MenuRepository.UpdateByID(model)
	helper.HandlePanicIfError(err)

	return true
}

func makeMenuSingleView(model *models.Menu) *params.MenuSingleView {
	return &params.MenuSingleView{
		ID:        model.ID,
		Name:      model.Name,
		Desc:      model.Desc,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func makeMenuListView(models *[]models.Menu) *[]params.MenuSingleView {
	var menuListview []params.MenuSingleView
	for _, model := range *models {
		menuListview = append(menuListview, *makeMenuSingleView(&model))
	}

	return &menuListview
}
