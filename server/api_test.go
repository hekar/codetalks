package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"fmt"

	"strconv"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func createDb() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
	})

	err := CreateSchema(db)
	if err != nil {
		panic(err)
	}

	return db
}

func createAPI() *API {
	db := createDb()
	mc := memcache.New("localhost:11211")
	h := &API{
		Db: db,
		Mc: mc,
	}
	return h
}

func createUser() *User {
	return &User{
		Name:   "Bob",
		Emails: []string{"bob"},
	}
}

func userJson(user *User) string {
	bytes, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func doRequest(
	e *echo.Echo, h *API, rec *httptest.ResponseRecorder,
	method, route, body string) echo.Context {
	req := httptest.NewRequest(method, route,
		strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return e.NewContext(req, rec)
}

func TestPostUser(t *testing.T) {
	e := echo.New()
	h := createAPI()
	rec := httptest.NewRecorder()

	userJSON := userJson(createUser())

	c := doRequest(e, h, rec,
		echo.POST, "/v1/user", userJSON)

	// Assertions
	if assert.NoError(t, h.postUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}

func TestPutUser(t *testing.T) {
	e := echo.New()
	h := createAPI()
	rec := httptest.NewRecorder()

	user := createUser()
	user.ID = 1
	updateJson := userJson(user)

	route := "/api/v1/user/" + strconv.Itoa(user.ID)
	fmt.Printf("FASDFSDF %v %v\n", route, updateJson)

	c := doRequest(e, h, rec,
		echo.PUT, route, updateJson)

	// Assertions
	if assert.NoError(t, h.putUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, updateJson, rec.Body.String())
	}
}
