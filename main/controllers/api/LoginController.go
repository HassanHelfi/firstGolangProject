package api

import (
	"crud_echo/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"time"
)

func Login(c echo.Context) error {
	//email := c.FormValue("email")
	//pass := c.FormValue("pass")

	token, err := createJwtToken();
	if err != nil {
		log.Println("jwt error", err)
		return c.String(http.StatusInternalServerError, "JWT Error")
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message" : "its Ok",
		"token" : token,
	})
}

func createJwtToken() (string, error)  {
	claims := models.JwtClaims{
		"hassan",
		jwt.StandardClaims{
			Id: "main_user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512 , claims)
	token, err := rawToken.SignedString([]byte("test-test"))
	if err != nil {
		return "" , err
	}
	return token , nil
}
