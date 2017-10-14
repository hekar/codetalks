package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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

func createTalk() *Talk {
	return &Talk{
		Name: "Bob",
		URL: "http://google.ca",
		ThumbnailURL: "http://google.ca",
		Tags: []string{ "tag" },
	}
}

func bodyJson(i interface{}) string {
	bytes, err := json.Marshal(i)
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
	h.Bind(e.Group("/api"))
	context := e.NewContext(req, rec)
	e.Router().Find(method, route, context)

	return context
}

func TestGetUser(t *testing.T) {
	e := echo.New()
	h := createAPI()
	rec := httptest.NewRecorder()

	route := "/api/v1/user/1"
	c := doRequest(e, h, rec,
		echo.GET, route, "")

	expected := "{\"id\":1,\"name\":\"This is user\",\"emails\":[\"user1@example.com\"]}"

	if assert.NoError(t, h.getUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}

func TestPostUser(t *testing.T) {
	e := echo.New()
	h := createAPI()
	rec := httptest.NewRecorder()

	user := createUser()
	user.ID = 2
	userJSON := bodyJson(user)

	c := doRequest(e, h, rec,
		echo.POST, "/api/v1/user", userJSON)

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
	expected := bodyJson(user)

	route := "/api/v1/user/" + strconv.Itoa(user.ID)

	c := doRequest(e, h, rec,
		echo.PUT, route, expected)

	if assert.NoError(t, h.putUser(c)) {
		actual := rec.Body.String()
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, actual)
	}
}

func TestGetTalk(t *testing.T) {
	e := echo.New()
	h := createAPI()
	rec := httptest.NewRecorder()

	route := "/api/v1/talk/1"

	c := doRequest(e, h, rec,
		echo.GET, route, "")

	expected := "{\"id\":1,\"name\":\"Bjarne Stroustrup - The Essence of C++\",\"url\":\"https://www.youtube.com/watch?v=86xWVb4XIyE\",\"thumbnailUrl\":\"https://i.ytimg.com/an_webp/86xWVb4XIyE/mqdefault_6s.webp?du=3000\\u0026sqp=CKb708sF\\u0026rs=AOn4CLDgRM5ZQwHj8tre1P0MLtd84ZGw4w\",\"tags\":[\"cpp\"]}"
	if assert.NoError(t, h.getTalk(c)) {
		actual := rec.Body.String()
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, actual)
	}
}

func TestListTalk(t *testing.T) {
	e := echo.New()
	h := createAPI()
	rec := httptest.NewRecorder()

	route := "/api/v1/talk/"

	c := doRequest(e, h, rec,
		echo.GET, route, "")

	expected := "{\"limit\":10,\"offset\":0,\"talks\":[{\"id\":1,\"name\":\"Bjarne Stroustrup - The Essence of C++\",\"url\":\"https://www.youtube.com/watch?v=86xWVb4XIyE\",\"thumbnailUrl\":\"https://i.ytimg.com/an_webp/86xWVb4XIyE/mqdefault_6s.webp?du=3000\\u0026sqp=CKb708sF\\u0026rs=AOn4CLDgRM5ZQwHj8tre1P0MLtd84ZGw4w\",\"tags\":[\"cpp\"]},{\"id\":4,\"name\":\"JavaScript does NOT offer zero-cost abstractions\",\"url\":\"https://www.youtube.com/watch?v=yLv3hafmSas\\u0026list=PL00z3DSeZW7zF104m6m055Wgm-7i9e4rV\\u0026index=4\",\"thumbnailUrl\":\"https://i.ytimg.com/vi/yLv3hafmSas/hqdefault.jpg?sqp=-oaymwEWCKgBEF5IWvKriqkDCQgBFQAAiEIYAQ==\\u0026rs=AOn4CLDd15vgesB1PaB2S0S2w7gBalEh0Q\",\"tags\":[\"javascript\",\"performance\"]},{\"id\":3,\"name\":\"Progressive, Performant, Polymer: Pick Three - Google I/O 2016\",\"url\":\"https://www.youtube.com/watch?v=J4i0xJnQUzU\\u0026index=2\\u0026list=PL00z3DSeZW7wDVgFVboA-5rBwkI-8WT_R\",\"thumbnailUrl\":\"https://i.ytimg.com/vi/J4i0xJnQUzU/hqdefault.jpg?sqp=-oaymwEWCKgBEF5IWvKriqkDCQgBFQAAiEIYAQ==\\u0026rs=AOn4CLCg9ScmRIx1VdTWzpEIumVVL3SYQw\",\"tags\":[\"google\",\"polymer\",\"web\"]},{\"id\":6,\"name\":\"RxJS In-Depth – Ben Lesh\",\"url\":\"https://www.youtube.com/watch?v=KOOT7BArVHQ\\u0026list=PL00z3DSeZW7zF104m6m055Wgm-7i9e4rV\\u0026index=11\",\"thumbnailUrl\":\"https://i.ytimg.com/vi/KOOT7BArVHQ/hqdefault.jpg?sqp=-oaymwEWCKgBEF5IWvKriqkDCQgBFQAAiEIYAQ==\\u0026rs=AOn4CLAre3UZs8p6VH7EJtdNUXHiwac5ZA\",\"tags\":[\"rxjs\",\"javascript\"]},{\"id\":2,\"name\":\"Tech Talk: Linus Torvalds on git\",\"url\":\"https://www.youtube.com/watch?v=4XpnKHJAok8\",\"thumbnailUrl\":\"https://i.ytimg.com/vi/4XpnKHJAok8/hqdefault.jpg?sqp=-oaymwEXCPYBEIoBSFryq4qpAwkIARUAAIhCGAE=\\u0026rs=AOn4CLAlzaJGQnDKZxr4ufeSuaLDOamRjg\",\"tags\":[\"linus\",\"git\"]},{\"id\":5,\"name\":\"The Vulkan Graphics API - what it means for Linux\",\"url\":\"https://www.youtube.com/watch?v=ynyO3O3zd3E\\u0026list=PL00z3DSeZW7zF104m6m055Wgm-7i9e4rV\\u0026index=9\",\"thumbnailUrl\":\"https://i.ytimg.com/vi/ynyO3O3zd3E/hqdefault.jpg?sqp=-oaymwEWCKgBEF5IWvKriqkDCQgBFQAAiEIYAQ==\\u0026rs=AOn4CLAFrPoNSa8eoipULf2x95u_ZknvFQ\",\"tags\":[\"linux\",\"vulkan\",\"graphics\"]},{\"id\":7,\"name\":\"This is a talk\",\"url\":\"https://youtube.com\",\"thumbnailUrl\":\"http://image.example.com\",\"tags\":[\"youtube\"]},{\"id\":8,\"name\":\"This is a talk\",\"url\":\"https://youtube.com\",\"thumbnailUrl\":\"http://image.example.com\",\"tags\":[\"youtube\"]}]}"

	if assert.NoError(t, h.searchTalk(c)) {
		actual := rec.Body.String()
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, actual)
	}
}

func TestSearchTalk(t *testing.T) {
	e := echo.New()
	h := createAPI()
	rec := httptest.NewRecorder()

	route := "/api/v1/talk/?q=c++"

	c := doRequest(e, h, rec,
		echo.GET, route, "")

	expected := "{\"limit\":10,\"offset\":0,\"talks\":[{\"id\":1,\"name\":\"Bjarne Stroustrup - The Essence of C++\",\"url\":\"https://www.youtube.com/watch?v=86xWVb4XIyE\",\"thumbnailUrl\":\"https://i.ytimg.com/an_webp/86xWVb4XIyE/mqdefault_6s.webp?du=3000\\u0026sqp=CKb708sF\\u0026rs=AOn4CLDgRM5ZQwHj8tre1P0MLtd84ZGw4w\",\"tags\":[\"cpp\"]},{\"id\":4,\"name\":\"JavaScript does NOT offer zero-cost abstractions\",\"url\":\"https://www.youtube.com/watch?v=yLv3hafmSas\\u0026list=PL00z3DSeZW7zF104m6m055Wgm-7i9e4rV\\u0026index=4\",\"thumbnailUrl\":\"https://i.ytimg.com/vi/yLv3hafmSas/hqdefault.jpg?sqp=-oaymwEWCKgBEF5IWvKriqkDCQgBFQAAiEIYAQ==\\u0026rs=AOn4CLDd15vgesB1PaB2S0S2w7gBalEh0Q\",\"tags\":[\"javascript\",\"performance\"]},{\"id\":3,\"name\":\"Progressive, Performant, Polymer: Pick Three - Google I/O 2016\",\"url\":\"https://www.youtube.com/watch?v=J4i0xJnQUzU\\u0026index=2\\u0026list=PL00z3DSeZW7wDVgFVboA-5rBwkI-8WT_R\",\"thumbnailUrl\":\"https://i.ytimg.com/vi/J4i0xJnQUzU/hqdefault.jpg?sqp=-oaymwEWCKgBEF5IWvKriqkDCQgBFQAAiEIYAQ==\\u0026rs=AOn4CLCg9ScmRIx1VdTWzpEIumVVL3SYQw\",\"tags\":[\"google\",\"polymer\",\"web\"]},{\"id\":2,\"name\":\"Tech Talk: Linus Torvalds on git\",\"url\":\"https://www.youtube.com/watch?v=4XpnKHJAok8\",\"thumbnailUrl\":\"https://i.ytimg.com/vi/4XpnKHJAok8/hqdefault.jpg?sqp=-oaymwEXCPYBEIoBSFryq4qpAwkIARUAAIhCGAE=\\u0026rs=AOn4CLAlzaJGQnDKZxr4ufeSuaLDOamRjg\",\"tags\":[\"linus\",\"git\"]},{\"id\":5,\"name\":\"The Vulkan Graphics API - what it means for Linux\",\"url\":\"https://www.youtube.com/watch?v=ynyO3O3zd3E\\u0026list=PL00z3DSeZW7zF104m6m055Wgm-7i9e4rV\\u0026index=9\",\"thumbnailUrl\":\"https://i.ytimg.com/vi/ynyO3O3zd3E/hqdefault.jpg?sqp=-oaymwEWCKgBEF5IWvKriqkDCQgBFQAAiEIYAQ==\\u0026rs=AOn4CLAFrPoNSa8eoipULf2x95u_ZknvFQ\",\"tags\":[\"linux\",\"vulkan\",\"graphics\"]}]}"
	if assert.NoError(t, h.searchTalk(c)) {
		actual := rec.Body.String()
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, actual)
	}
}

func TestListPopularTalks(t *testing.T) {
	e := echo.New()
	h := createAPI()
	rec := httptest.NewRecorder()

	route := "/api/v1/popular/"

	c := doRequest(e, h, rec,
		echo.GET, route, "")

	expected := "{\"limit\":0,\"offset\":0,\"talks\":[{\"id\":3,\"name\":\"Progressive, Performant, Polymer: Pick Three - Google I/O 2016\",\"url\":\"https://www.youtube.com/watch?v=J4i0xJnQUzU\\u0026index=2\\u0026list=PL00z3DSeZW7wDVgFVboA-5rBwkI-8WT_R\",\"thumbnailUrl\":\"https://i.ytimg.com/vi/J4i0xJnQUzU/hqdefault.jpg?sqp=-oaymwEWCKgBEF5IWvKriqkDCQgBFQAAiEIYAQ==\\u0026rs=AOn4CLCg9ScmRIx1VdTWzpEIumVVL3SYQw\",\"tags\":[\"google\",\"polymer\",\"web\"]},{\"id\":2,\"name\":\"Tech Talk: Linus Torvalds on git\",\"url\":\"https://www.youtube.com/watch?v=4XpnKHJAok8\",\"thumbnailUrl\":\"https://i.ytimg.com/vi/4XpnKHJAok8/hqdefault.jpg?sqp=-oaymwEXCPYBEIoBSFryq4qpAwkIARUAAIhCGAE=\\u0026rs=AOn4CLAlzaJGQnDKZxr4ufeSuaLDOamRjg\",\"tags\":[\"linus\",\"git\"]},{\"id\":1,\"name\":\"Bjarne Stroustrup - The Essence of C++\",\"url\":\"https://www.youtube.com/watch?v=86xWVb4XIyE\",\"thumbnailUrl\":\"https://i.ytimg.com/an_webp/86xWVb4XIyE/mqdefault_6s.webp?du=3000\\u0026sqp=CKb708sF\\u0026rs=AOn4CLDgRM5ZQwHj8tre1P0MLtd84ZGw4w\",\"tags\":[\"cpp\"]}]}"
	if assert.NoError(t, h.listPopularTalks(c)) {
		actual := rec.Body.String()
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, actual)
	}
}

func TestPostTalk(t *testing.T) {
	e := echo.New()
	h := createAPI()
	rec := httptest.NewRecorder()

	route := "/api/v1/talk"

	talk := createTalk()
	talk.ID = 50
	body := bodyJson(talk)

	c := doRequest(e, h, rec,
		echo.POST, route, body)

	expected := "{\"id\":50,\"name\":\"Bob\",\"url\":\"http://google.ca\",\"thumbnailUrl\":\"http://google.ca\",\"tags\":[\"tag\"]}"
	if assert.NoError(t, h.postTalk(c)) {
		actual := rec.Body.String()
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, expected, actual)
	}
}

func TestPutTalk(t *testing.T) {
	e := echo.New()
	h := createAPI()
	rec := httptest.NewRecorder()

	route := "/api/v1/talk/1"

	talk := createTalk()
	talk.ID = 1
	body := bodyJson(talk)

	c := doRequest(e, h, rec,
		echo.POST, route, body)

	expected := "{\"id\":1,\"name\":\"Bob\",\"url\":\"http://google.ca\",\"thumbnailUrl\":\"http://google.ca\",\"tags\":[\"tag\"]}"
	if assert.NoError(t, h.putTalk(c)) {
		actual := rec.Body.String()
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, actual)
	}
}
