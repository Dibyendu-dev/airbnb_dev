package utils

import (
	"encoding/json"
	"net/http"
	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func init(){  //this init fn only called once per package level
	Validator =  NewValidator()  
}

func NewValidator()*validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}

func WriteJsonResponse(w http.ResponseWriter , status int , data any) error {
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data) //encode the data as JSON and write it to the response
}

func ReadJsonBody(r *http.Request, result any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return  decoder.Decode(result)
}

func WriteJsonSuccessResponse(w http.ResponseWriter,status int,message string,data any) error{
	response:= map[string]any{}
	response["status"]="success"
	response["message"]=message
	response["data"]=data
	return WriteJsonResponse(w,status,response)
}

func WriteJsonErrorResponse(w http.ResponseWriter,status int,message string,err error) error{
	response:= map[string]any{}
	response["status"]="error"
	response["message"]=message
	response["error"]=err.Error()
	return WriteJsonResponse(w,status,response)
}