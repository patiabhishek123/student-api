package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/patiabhishek123/students-api/internal/types"
	"github.com/patiabhishek123/students-api/internal/utils/response"
	"github.com/go-playground/validator/v10"

)


func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info(" Creating student info")

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf(" Empty Body ")))
			return
		}

		if err !=nil {
			response.WriteJson(w,http.StatusBadRequest,response.GeneralError(err))
		}
		//request validation 

		if err :=validator.New().Struct(student);err!=nil {

			//we need to typ3ecast errror
			validateErrs :=err.(validator.ValidationErrors)
			response.WriteJson(w,http.StatusBadRequest,response.ValidationError(validateErrs))
			return;
		}
		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})
	}
}


