package controllers

import (
	"AuthInGo/dto"
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

func (rc *RoleController) AssignRoleToUser(w http.ResponseWriter, r *http.Request){
	userId := chi.URLParam(r,"userId")
	roleId := chi.URLParam(r,"roleId")
	if userId ==""{
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "User ID is required", fmt.Errorf("missing user ID"))
		return
	}
	if roleId == ""{
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("missing role ID"))
		return
	}
	userIdInt ,err :=strconv.ParseInt(userId,10,64)
	if err!= nil{
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid user ID", err)
		return
	}
	roleIdInt ,err :=strconv.ParseInt(roleId,10,64)
	if err!= nil{
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid role ID", err)
		return
	}

	err = rc.RoleService.AssignRoleToUser(userIdInt,roleIdInt)
	if err!= nil{
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to assign role to user", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK,"Role assign to nuser successfully",nil)
}

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

func (rc *RoleController) CreateRole(w http.ResponseWriter, r *http.Request) {
	payload:= r.Context().Value("payload").(dto.CreateRoleRequestDTO)

	role,err :=rc.RoleService.CreateRole(payload.Name, payload.Description)
	if err!= nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError,"failed to create role",err)
		return
	}
	utils.WriteJsonSuccessResponse(w,http.StatusOK,"Role created successfully",role)

}

func (rc *RoleController) UpdateRole (w http.ResponseWriter, r *http.Request){
	roleId := chi.URLParam(r, "id")
	if roleId == ""{
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("missing role ID"))
		return
	}

	id,err :=strconv.ParseInt(roleId,10,64)
	if err!= nil{
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid role ID", err)
		return
	}

	payload := r.Context().Value("payload").(dto.UpdateRoleRequestDTO)
	role,err :=rc.RoleService.UpdateRole(id,payload.Name,payload.Description)
	if err!= nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError,"failed to update role",err)
		return
	}
	utils.WriteJsonSuccessResponse(w,http.StatusOK,"Role updated successfully",role)

}

func (rc *RoleController) DeleteRole (w http.ResponseWriter, r *http.Request){
	roleId := chi.URLParam(r, "id")
	if roleId == ""{
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("missing role ID"))
		return
	}

	id,err :=strconv.ParseInt(roleId,10,64)
	if err!= nil{
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid role ID", err)
		return
	}
	err = rc.RoleService.DeleteRoleById(id)
	if err!= nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError,"failed to delete role",err)
		return
	}
	utils.WriteJsonSuccessResponse(w,http.StatusOK,"Role deleted successfully",nil)
}

func (rc *RoleController) GetRolePermissions(w http.ResponseWriter, r *http.Request){
	roleId := chi.URLParam(r, "id")
	if roleId == ""{
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("missing role ID"))
		return
	}

	id,err :=strconv.ParseInt(roleId,10,64)
	if err!= nil{
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid role ID", err)
		return
	}
	rolePermissions,err :=rc.RoleService.GetRolePermissions(id)
	if err!= nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError,"failed to fetch role permission",err)
		return
	}
	utils.WriteJsonSuccessResponse(w,http.StatusOK,"Role permission fetched successfully",rolePermissions)

}

func (rc *RoleController) AssignPermissionToRole(w http.ResponseWriter, r *http.Request){
	roleId := chi.URLParam(r, "id")
	if roleId == ""{
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("missing role ID"))
		return
	}

	id,err :=strconv.ParseInt(roleId,10,64)
	if err!= nil{
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid role ID", err)
		return
	}
	payload := r.Context().Value("payload").(dto.AssignPermissionRequestDTO)
	rolePermission,err := rc.RoleService.AddPermissionToRole(id,payload.PermissionId)
	if err!= nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError,"failed to assign role permission",err)
		return
	}
	utils.WriteJsonSuccessResponse(w,http.StatusOK,"Role permission assign successfully",rolePermission)
}

func (rc *RoleController) RemovePermissionFromRole(w http.ResponseWriter, r *http.Request){
	roleId := chi.URLParam(r, "id")
	if roleId == ""{
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("missing role ID"))
		return
	}

	id,err :=strconv.ParseInt(roleId,10,64)
	if err!= nil{
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid role ID", err)
		return
	}
	payload := r.Context().Value("payload").(dto.RemovePermissionRequestDTO)
	err = rc.RoleService.RemovePermissionFromRole(id,payload.PermissionId)
	if err!= nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError,"failed to remove  permission from role",err)
		return
	}
	utils.WriteJsonSuccessResponse(w,http.StatusOK," permission removed successfully",nil)

}

func (rc *RoleController) GetAllRolePermissions(w http.ResponseWriter, r *http.Request){
	rolePermissions,err := rc.RoleService.GetAllRolePermissions()
	if err!= nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError,"failed to fetch all permission from role",err)
		return
	}
	utils.WriteJsonSuccessResponse(w,http.StatusOK," all permission fetched successfully",rolePermissions)
}