package db

import (
	// "database/sql"
	"fmt"
)

type UserRepository interface { // facilitates dependency injection for repository
	Create() error
}

type UserRepositoryImpl struct {
	// db *sql.DB // can be implemented using ORM
}

func NewUserRepository() UserRepository{
	return  &UserRepositoryImpl{
		// db:db,
	}
}

func (u *UserRepositoryImpl) Create() error {
	fmt.Println("creating user in user repository")
	return nil
}
