package api

import (
	"crud_echo/configs"
	"crud_echo/domain"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"time"
)

func MainJwt(c echo.Context) error {
	user := c.Get("user")
	token := user.(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)
	log.Println("user name", claims["name"], "user id", claims["jti"])
	return c.String(http.StatusOK, "rtyrety")
}
func Login(c echo.Context) error {
	//email := c.FormValue("email")
	//pass := c.FormValue("pass")

	token, err := createJwtToken()
	if err != nil {
		log.Println("jwt error", err)
		return c.String(http.StatusInternalServerError, "JWT Error")
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "its Ok",
		"token":   token,
	})
}

func createJwtToken() (string, error) {
	claims := domain.JwtClaims{
		"hassan",
		jwt.StandardClaims{
			Id:        "main_user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := rawToken.SignedString([]byte("test-test"))
	if err != nil {
		return "", err
	}
	return token, nil
}
func CheckUser(c echo.Context) error{
	var user = domain.User{}
	email := c.FormValue("email")
	pass := c.FormValue("password")
	//first_name := c.FormValue("first_name")
	fmt.Println(email , pass)
	//result := configs.CreateCon().Where("email = ? AND password = ?", email , pass).Find(&user)
	result := configs.CreateCon().Where("email = ?", &email).Find(&user)
	fmt.Println(result)
	return c.JSON(http.StatusOK, result)
	//if result {
	//
	//}
}
