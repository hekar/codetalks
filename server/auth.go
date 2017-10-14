package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-pg/pg"
	"github.com/labstack/echo"
)

type AuthAPI struct {
	Db *pg.DB
	Mc *memcache.Client
}

// Bind register routers and group information
func (api *AuthAPI) Bind(group *echo.Group) {
	group.POST("/login", api.login)
	group.GET("/ping", api.ping)
}

func (api *AuthAPI) login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	fmt.Printf("username: %v, password %v\n", username, password)

	if username == "testuser" && password == "test" {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "TestUser"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

func (api *AuthAPI) ping(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
