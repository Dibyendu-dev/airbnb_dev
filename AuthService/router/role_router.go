package router

import (
	"AuthInGo/controllers"
	"AuthInGo/middleware"

	"github.com/go-chi/chi/v5"
)

type RoleRouter struct {
	roleController *controllers.RoleController
}

func NewRoleRouter(_roleController *controllers.RoleController) Router {
	return &RoleRouter{
		roleController: _roleController,
	}
}

func (rr *RoleRouter) Register(r chi.Router){
	r.Get("/roles/{id}", rr.roleController.GetRoleById)
	r.Get("/roles",rr.roleController.GetAllRoles)
	r.With(middleware.CreateRoleRequestValidator).Post("/roles",rr.roleController.CreateRole)
	r.With(middleware.UpdateRoleRequestValidator).Post("/roles",rr.roleController.UpdateRole)
	r.Delete("roles/{id}",rr.roleController.DeleteRole)

	// role permission
	r.Get("/roles/{id}/permissions",rr.roleController.GetRolePermissions)
	r.Get("/role-permissions",rr.roleController.GetAllRolePermissions)
	r.With(middleware.AssignPermissionRequestValidator).Post("/roles/{id}/permissions",rr.roleController.AssignPermissionToRole)
	r.With(middleware.RemovePermissionRequestValidator).Post("/roles/{id}/permissions",rr.roleController.RemovePermissionFromRole)
	


}
