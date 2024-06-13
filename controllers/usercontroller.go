package controllers

import (
	"net/http"

	"github.com/herulobarto/go-jwt/helpers"
	"github.com/herulobarto/go-jwt/models"
)

func Me(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("userinfo").(*helpers.MyCustomClaims)

	userResponse := &models.MyProfile{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	helpers.Response(w, 200, "My Profile", userResponse)
}
