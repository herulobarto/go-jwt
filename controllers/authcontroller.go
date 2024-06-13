package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/herulobarto/go-jwt/configs"
	"github.com/herulobarto/go-jwt/helpers"
	"github.com/herulobarto/go-jwt/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var register models.Register

	if err := json.NewDecoder(r.Body).Decode(&register); err != nil {

		helpers.Response(w, 500, err.Error(), nil)

		return

	}

	defer r.Body.Close()

	if register.Password != register.PasswordConfirm {
		helpers.Response(w, 400, "Password not match", nil)
		return
	}

	passwordHash, err := helpers.HashPassword(register.Password)
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	user := models.User{
		Name:     register.Name,
		Email:    register.Email,
		Password: passwordHash,
	}

	if err := configs.DB.Create(&user).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 201, "Register Succesfully", nil)
}
