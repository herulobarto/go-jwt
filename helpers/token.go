package helpers

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/herulobarto/go-jwt/models"
)

var mySigningKey = []byte("mysecretkey")

type MyCustomClaims struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateToken(user *models.User) (string, error) {

	// Create claims with multiple fields populated
	claims := MyCustomClaims{
		user.ID,
		user.Name,
		user.Email,

		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Println(ss, err)

	return ss, err

}
