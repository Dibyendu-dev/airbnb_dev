package middleware

import (
	"AuthInGo/dto"
	"AuthInGo/utils"
	"context"
	"fmt"
	"net/http"
)

func UserLoginRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.LoginUserRequestDTO

		// read the json body into payload
		if err := utils.ReadJsonBody(r, &payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "invalid request body", err)
			return
		}

		// validate the payload using validator
		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "validation failed", err)
			return
		}

		ctx := context.WithValue(r.Context(),"payload",payload) // create a new context with the payload

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UserCreateRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.CreateUserRequestDTO

		// read the json body into payload
		// req body and req handler is consumed or read in stream manner
		if err := utils.ReadJsonBody(r, &payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "invalid request body", err)
			return
		}

		// validate the payload using validator
		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "validation failed", err)
			return
		}
		fmt.Println("payload recived for login:",payload)

		ctx := context.WithValue(r.Context(),"payload",payload) // create a new context with the payload

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func CreateRoleRequestValidator(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.CreateRoleRequestDTO

		// read and decode json body into the payload
		if err :=utils.ReadJsonBody(r, &payload); err!= nil{
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest,"invalid request body",err)
			return 
		}

		// validate the payload using validator instance
		if err :=utils.Validator.Struct(payload); err!= nil{
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest,"validation failed",err)
			return 
		}

		ctx := context.WithValue(r.Context(),"payload",payload)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}


func UpdateRoleRequestValidator(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.UpdateRoleRequestDTO

		// read and decode json body into the payload
		if err :=utils.ReadJsonBody(r, &payload); err!= nil{
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest,"invalid request body",err)
			return 
		}

		// validate the payload using validator instance
		if err :=utils.Validator.Struct(payload); err!= nil{
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest,"validation failed",err)
			return 
		}

		ctx := context.WithValue(r.Context(),"payload",payload)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}


func AssignPermissionRequestValidator(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.AssignPermissionRequestDTO

		// read and decode json body into the payload
		if err :=utils.ReadJsonBody(r, &payload); err!= nil{
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest,"invalid request body",err)
			return 
		}

		// validate the payload using validator instance
		if err :=utils.Validator.Struct(payload); err!= nil{
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest,"validation failed",err)
			return 
		}

		ctx := context.WithValue(r.Context(),"payload",payload)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}


func RemovePermissionRequestValidator(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.RemovePermissionRequestDTO

		// read and decode json body into the payload
		if err :=utils.ReadJsonBody(r, &payload); err!= nil{
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest,"invalid request body",err)
			return 
		}

		// validate the payload using validator instance
		if err :=utils.Validator.Struct(payload); err!= nil{
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest,"validation failed",err)
			return 
		}

		ctx := context.WithValue(r.Context(),"payload",payload)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
