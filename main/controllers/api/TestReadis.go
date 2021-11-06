package api

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"net/http"
)

func Test(c echo.Context) error {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := client.Set("name", c.FormValue("name"), 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	t, err := client.Get("name").Result()
	if err != nil {
		fmt.Println(err)
	}
	return c.JSON(http.StatusOK, t)
}