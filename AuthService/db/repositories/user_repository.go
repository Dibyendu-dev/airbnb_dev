package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UserRepository interface { // facilitates dependency injection for repository
	GetById() (*models.User,error)
	Create() (error)
	GetAll() ([]*models.User,error)
	DeleteById(id int64) (error)
}

type UserRepositoryImpl struct {
	 db *sql.DB // can be implemented using ORM
}

func NewUserRepository(_db *sql.DB) UserRepository{
	return  &UserRepositoryImpl{
		 db:_db,
	}
}

func (u *UserRepositoryImpl) GetAll() ([]*models.User,error){

	// TODO:


	return nil,nil
}

func (u *UserRepositoryImpl) DeleteById(id int64) error{

	// TODO:


	return nil
}

func (u *UserRepositoryImpl) Create() (error){
	fmt.Println("creating user in user repository")

	query := "INSERT INTO users (username , email , password) VALUES (? , ? , ?)"

	result , err :=u.db.Exec(query, "test1234", "test@test.com", "test1234")
	
	if err != nil{
		fmt.Println("error inserting user",err)
		return err
	}

	rowAffected , rowErr :=result.RowsAffected()
	
	if rowErr!= nil {
		fmt.Println("error getting affecte rows",rowErr)
		return  rowErr
	}

	if rowAffected == 0{
		fmt.Println("no rows affected , users not created")
		return  nil
	}

	fmt.Println("Users created succesfully, rows affected",rowAffected)
	return  nil
}

func (u *UserRepositoryImpl) GetById() (*models.User,error) {
	fmt.Println("fetching user in user repository")

	// step 1: prepare the query
	query := "SELECT id, username, email, created_at, updated_at FROM users WHERE id = ? "

	// step 2: execute the query
	row :=u.db.QueryRow(query,1)

	// step 3: process the result
	user:= &models.User{}

	err:=row.Scan(&user.Id, &user.Username, &user.Email,  &user.CreatedAt, &user.UpdatedAt)

	if err!= nil{
		if err == sql.ErrNoRows{
			fmt.Println("No user found with given id")
			return nil,err
		}else {
			fmt.Println("error scanning user",err)
			return nil, err
		}
	}
	fmt.Println("user fetched successfully",user)

	return user ,nil
}
