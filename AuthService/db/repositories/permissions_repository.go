package db

import (
	"AuthInGo/models"
	"database/sql"

)

type PermissionsRepository interface {
	GetPermissionById(id int64) (*models.Permission, error)
	GetPermissionByName(name string) (*models.Permission, error)
	GetAllPermissions() ([]*models.Permission, error)
	CreatePermission(name string, description string, resource string, actions string) (*models.Permission, error)
	DeletePermissionById(id int64) error
	UpdatePermission(id int64, name string, description string,resource string,actions string) (*models.Permission, error)
}

type PermissionsRepositoryImpl struct {
	db *sql.DB
}

func NewPermissionsRepository(_db *sql.DB) PermissionsRepository {
	return &PermissionsRepositoryImpl{
		db: _db,
	}
}
