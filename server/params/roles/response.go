package params

type RoleListview *[]RoleSingleView

type RoleSingleView struct {
	ID   int
	Name string
}

type RoleAddResponse struct {
	Message string
}
