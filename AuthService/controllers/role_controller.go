package controllers

import (
	"AuthInGo/services"
	"AuthInGo/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type RoleController struct {
	RoleService services.RoleService
}

func NewRoleController(_roleService services.RoleService) *RoleController{
	return &RoleController{
		RoleService: _roleService,
	}
}

// func (rc *RoleController) AssignRoleToUser(w http.ResponseWriter, r *http.Request){
// 	userId := chi.URLParam(r,"userId")
// }

func (rc *RoleController) GetRoleById(w http.ResponseWriter, r *http.Request) {
	roleId :=chi.URLParam(r,"id")
	if roleId == ""{
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest,"RoleId is required",fmt.Errorf("missing role id"))
		return
	}
	id,err := strconv.ParseInt(roleId,10,64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid role ID", err)
		return
	}
	role,err :=rc.RoleService.GetRoleById(id)
	if err!= nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError,"failed to fetch role",err)
		return
	}
	if role == nil {
		utils.WriteJsonErrorResponse(w, http.StatusNotFound, "Role not found", fmt.Errorf("role with ID %d not found", roleId))
		return
	}
	utils.WriteJsonSuccessResponse(w,http.StatusOK,"Role fetched successfully",role)
}

func (rc *RoleController) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	roles,err :=rc.RoleService.GetAllRoles()
	if err!= nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError,"failed to fetch role",err)
		return
	}
	utils.WriteJsonSuccessResponse(w,http.StatusOK,"Role fetched successfully",roles)

}