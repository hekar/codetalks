package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"fmt"

	"strconv"

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
	h := &API{db}
	return h
}

func createUser() (*User, string) {
	user := &User{
		Name:   "Bob",
		Emails: []string{"bob"},
	}
	bytes, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	userJSON := string(bytes)

	return user, userJSON
}

func doRequest(method, route, body string) (echo.Context, *API, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, route,
		strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := createAPI()

	return c, h, rec
}

func TestCreateUser(t *testing.T) {

	_, userJSON := createUser()

	c, h, rec := doRequest(echo.POST, "/v1/user", userJSON)

	// Assertions
	if assert.NoError(t, h.postUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}

func TestPutUser(t *testing.T) {
	_, userJSON := createUser()

	c, h, rec := doRequest(echo.POST, "/v1/user", userJSON)
	user := &User{}

	h.postUser(c)

	if err := json.Unmarshal(rec.Body.Bytes(), user); err != nil {
		panic(err)
	}

	bytes, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	updatedUserJSON := string(bytes)

	route := "/api/v1/user/" + strconv.FormatInt(user.ID, 10)
	fmt.Printf("FASDFSDF %v %v\n", route, updatedUserJSON)
	c, h, rec = doRequest(echo.PUT, route, updatedUserJSON)

	// Assertions
	if assert.NoError(t, h.putUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, updatedUserJSON, rec.Body.String())
	}
}
