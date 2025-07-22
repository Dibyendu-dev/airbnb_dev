package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UserRepository interface { // facilitates dependency injection for repository
	GetById() (*models.User,error)
	Create(username , email , hashedPassword string) (error)
	GetAll() ([]*models.User,error)
	DeleteById(id int64) (error)
	GetUserByEmail(email string) (*models.User,error)
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

	query := "select id,username,email,created_at,updated_at from users"

	rows,err :=u.db.Query(query)
	if err != nil{
		fmt.Println("error fetching user",err)
		return  nil,err
	}
	defer rows.Close() // ensure rows are closed after processing


	return nil,nil
}

func (u *UserRepositoryImpl) DeleteById(id int64) error{

	query := "delete from users where id = ?"
	result , err :=u.db.Exec(query, id)
	if err != nil{
		fmt.Println("error deleting user",err)
		return  err
	}
	 rowsAffected,rowErr := result.RowsAffected()
	 if rowErr != nil{
		fmt.Println("error getting rows affected",rowErr)
		return  rowErr
	 }
	 if rowsAffected == 0{
		fmt.Println("no rows affected , users not deleted")
		return  nil
	 }
	 fmt.Println("user deleted successfully,rows affected:",rowsAffected)
	return nil
}

//  goose -dir "db/migrations" mysql "root:ddas4548@tcp(127.0.0.1:3306)/auth_dev" up
//  goose -dir "db/migrations" create create_user_table sql

func (u *UserRepositoryImpl) Create(username , email , hashedPassword string) (error){
	fmt.Println("creating user in user repository")

	query := "INSERT INTO users (username , email , password) VALUES (? , ? , ?)"

	result , err :=u.db.Exec(query, username, email, hashedPassword)
	
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

	err:=row.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)

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

func (u *UserRepositoryImpl) GetUserByEmail( email string) (*models.User,error){
	fmt.Println("get user by email in repo layer")
	
	query := "select id ,email , password from users where email = ?"
	row := u.db.QueryRow(query,email)

	user := &models.User{}  // reference obj.

	err := row.Scan(&user.Id,&user.Email,&user.Password) //password is hashed
	if err!= nil{
		if err == sql.ErrNoRows{
			fmt.Println("No user found with given email")
			return nil,err
		}else {
			fmt.Println("error scanning user",err)
			return nil, err
		}
	}
	return  user,nil
}
