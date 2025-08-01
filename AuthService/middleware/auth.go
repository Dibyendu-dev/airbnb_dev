package middleware

import (
	env "AuthInGo/config/env"
	dbConfig"AuthInGo/config/db"
	repo "AuthInGo/db/repositories"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "authorization header is required", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "authorization must be start with Bearer", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			http.Error(w, "Token is required", http.StatusUnauthorized)
			return
		}

		claims := jwt.MapClaims{}

		_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(env.GetString("JWT_SECRET", "TOKEN")), nil
		})
		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		userId, okId := claims["id"].(float64)
		email, okEmail := claims["email"].(string)

		if !okId || !okEmail {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		fmt.Println("Authenticated user ID:", int64(userId), "Email:", email)

		ctx := context.WithValue(r.Context(),"userID", strconv.FormatFloat(userId,'f',0,64))
		ctx = context.WithValue(ctx,"email",email)

		next.ServeHTTP(w,r.WithContext(ctx))

	})
}

func RequireAllRoles(roles ...string)func(http.Handler) http.Handler {

	// function that can create a middleware for checking the above set of roles

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			userIdStr := r.Context().Value("userId").(string)
			userId,err :=strconv.ParseInt(userIdStr,10,64)
			if err!=nil {
				http.Error(w, "invalid user id",http.StatusUnauthorized)
				return 
			}
			dbconn,dbErr :=dbConfig.SetupDB()
			if dbErr != nil{
				http.Error(w, "Database connection error:" +dbErr.Error(), http.StatusInternalServerError)
				return 
			}
			
			urr:= repo.NewUserRoleRepository(dbconn)
			hasAllRoles, hasAllRolesErr := urr.HasAllRoles(userId,roles)
			fmt.Println("userid",userId,"roles",roles,"hasAllRoles",hasAllRoles)
			if hasAllRolesErr != nil {
				http.Error(w, "Error checking user roles: "+hasAllRolesErr.Error(), http.StatusInternalServerError)
				return
			}
			if !hasAllRoles{
				http.Error(w, "Forbidden: You do not have the required roles", http.StatusForbidden)
				return
			}
			fmt.Println("User has all required roles:", roles)
			next.ServeHTTP(w,r)
		})
	}
}

func RequireAnyRoles(roles ...string) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			userIdStr := r.Context().Value("userId").(string)
			userId,err :=strconv.ParseInt(userIdStr,10,64)
			if err!=nil {
				http.Error(w, "invalid user id",http.StatusUnauthorized)
				return 
			}
			dbconn,dbErr :=dbConfig.SetupDB()
			if dbErr != nil{
				http.Error(w, "Database connection error:" +dbErr.Error(), http.StatusInternalServerError)
				return 
			}
			 urr:= repo.NewUserRoleRepository(dbconn)
			hasAnyRoles, hasAnyRolesErr := urr.HasAnyRole(userId,roles)
			fmt.Println("userid",userId,"roles",roles,"hasAnyRoles",hasAnyRoles)
			if hasAnyRolesErr != nil {
				http.Error(w, "Error checking user roles: "+hasAnyRolesErr.Error(), http.StatusInternalServerError)
				return
			}
			if !hasAnyRoles{
				http.Error(w, "Forbidden: You do not have the required roles", http.StatusForbidden)
				return
			}
			fmt.Println("User has all required roles:", roles)
			next.ServeHTTP(w,r)
		})
	}
}
