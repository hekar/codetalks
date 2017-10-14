package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func createAuthAPI() *AuthAPI {
	db := createDb()
	mc := memcache.New("localhost:11211")
	h := &AuthAPI{
		Db: db,
		Mc: mc,
	}
	return h
}

func doAuthRequest(
	e *echo.Echo, h *AuthAPI, rec *httptest.ResponseRecorder,
	method, route, body string) echo.Context {
	req := httptest.NewRequest(method, route,
		strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	h.Bind(e.Group("/"))
	context := e.NewContext(req, rec)
	e.Router().Find(method, route, context)

	return context
}

func TestLogin(t *testing.T) {
	e := echo.New()
	h := createAuthAPI()
	rec := httptest.NewRecorder()

	route := "/login"
	c := doAuthRequest(e, h, rec,
		echo.POST, route, "username=testuser&password=test")

	if assert.NoError(t, h.login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "token")
	}
}

func TestPing(t *testing.T) {
	e := echo.New()
	h := createAuthAPI()
	rec := httptest.NewRecorder()

	route := "/ping"
	c := doAuthRequest(e, h, rec,
		echo.GET, route, "")

	expected := "OK"

	if assert.NoError(t, h.ping(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}
