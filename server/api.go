package main

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-pg/pg"
	"github.com/labstack/echo"
)

// API contains common state for API controllers
type API struct {
	Db *pg.DB
	Mc *memcache.Client
}

// SearchTalk result of searching talks
type SearchTalk struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Talks  []Talk `json:"talks"`
}

// Bind setup the API
func (api *API) Bind(group *echo.Group) {
	group.GET("/v1/conf", api.conf)

	group.GET("/v1/talk/popular", api.listPopularTalks)
	group.POST("/v1/talk", api.postTalk)
	group.GET("/v1/talk", api.searchTalk)
	group.GET("/v1/talk/:id", api.getTalk)
	group.PUT("/v1/talk", api.putTalk)

	group.GET("/v1/user/:id/talk/", api.getUserTalk)
	group.GET("/v1/user/:id/talk/:talkid", api.getUserTalk)
	group.PUT("/v1/user/:id/talk/:talkid", api.putUserTalk)
	group.DELETE("/v1/user/:id/talk/:talkid", api.deleteUserTalk)

	group.POST("/v1/user", api.postUser)
	group.GET("/v1/user/:id", api.getUser)
	group.PUT("/v1/user/:id", api.putUser)

	group.GET("/v1/yt/json", api.getYoutubeJSON)
}

func (api *API) conf(c echo.Context) error {
	app := c.Get("app").(*App)
	return c.JSON(200, app.Conf.Root)
}

func (api *API) listPopularTalks(c echo.Context) error {
	var talks []Talk
	err := api.Db.Model(&talks).
		Column("talk.*").
		Join("inner join talk_populars on talk_populars.talk_id = talk.id").
		Order("talk_populars.rank ASC").
		Select()
	if err != nil {
		return err
	}

	searchTalk := SearchTalk{
		Talks: talks,
	}

	return c.JSON(200, searchTalk)
}

func (api *API) postTalk(c echo.Context) error {
	talk := new(Talk)
	if err := c.Bind(talk); err != nil {
		return err
	}

	err := api.Db.Insert(talk)
	if err != nil {
		return err
	}

	return c.JSON(200, talk)
}

func (api *API) searchTalk(c echo.Context) error {
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		offset = 0
	}

	limit := 10

	// Filter title, author, rating, site, date posted
	// Sort by rating, date posted
	var talks []Talk
	query := api.Db.Model(&talks).
		Column("talk.*")

	if c.QueryParam("q") != "" {
		query = query.Where("lower(talk.name) like lower(concat('%', ?, '%'))", c.QueryParam("q"))
	}

	err = query.
		Order("talk.name ASC").
		Limit(limit).
		Offset(offset).
		Select()
	if err != nil {
		return err
	}

	searchTalk := SearchTalk{
		Limit:  limit,
		Offset: offset,
		Talks:  talks,
	}

	return c.JSON(200, searchTalk)
}

func (api *API) getTalk(c echo.Context) error {
	id := c.Param("id")

	var talk Talk
	err := api.Db.Model(&talk).
		Column("talk.*").
		Where("talk.id = ?", id).
		Select()
	if err != nil {
		return err
	}

	return c.JSON(200, talk)
}

func (api *API) putTalk(c echo.Context) error {
	id := c.Param("id")

	talk := new(Talk)
	if err := c.Bind(talk); err != nil {
		return err
	}

	if id != strconv.Itoa(talk.ID) {
		panic("talkIDs do not match")
	}

	err := api.Db.Update(talk)
	if err != nil {
		return err
	}

	return c.JSON(200, talk)
}

func (api *API) getUserTalk(c echo.Context) error {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	talkID, err := strconv.ParseInt(c.Param("talkid"), 10, 64)
	if err != nil {
		return err
	}

	userTalk := &UserTalk{
		UserID: userID,
		TalkID: talkID,
	}
	err = api.Db.Model(&userTalk).
		Column("usertalk.*").
		Where("usertalk.id = ?", userID).
		Select()
	if err != nil {
		return err
	}

	return c.JSON(200, userTalk)
}

func (api *API) putUserTalk(c echo.Context) error {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	talkID, err := strconv.ParseInt(c.Param("talkid"), 10, 64)
	if err != nil {
		return err
	}

	userTalk := &UserTalk{
		UserID: userID,
		TalkID: talkID,
	}
	err = api.Db.Insert(&userTalk)
	if err != nil {
		return err
	}

	return c.JSON(200, userTalk)
}

func (api *API) deleteUserTalk(c echo.Context) error {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	talkID, err := strconv.ParseInt(c.Param("talkid"), 10, 64)
	if err != nil {
		return err
	}

	userTalk := &UserTalk{
		UserID: userID,
		TalkID: talkID,
	}
	err = api.Db.Delete(&userTalk)
	if err != nil {
		return err
	}

	return c.JSON(200, userTalk)
}

func (api *API) postUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}

	err := api.Db.Insert(user)
	if err != nil {
		return err
	}

	return c.JSON(200, user)
}

func (api *API) getUser(c echo.Context) error {
	id := c.Param("id")

	var user User
	err := api.Db.Model(&user).
		Column("user.*").
		Where("user.id = ?", id).
		Select()
	if err != nil {
		return err
	}

	return c.JSON(200, user)
}

func (api *API) putUser(c echo.Context) error {
	id := c.Param("id")

	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}

	if id != strconv.Itoa(user.ID) {
		panic("userIDs do not match")
	}

	err := api.Db.Update(user)
	if err != nil {
		return err
	}

	return c.JSON(200, user)
}

func (api *API) getYoutubeJSON(c echo.Context) error {
	yturl := c.Param("yturl")

	fullurl := "http://www.youtube.com/oembed?url=" + yturl + "&format=json"

	key := "yturl-" + yturl
	fromMc, err := api.Mc.Get(key)
	if err != nil {
		return err
	}

	if fromMc == nil {
		resp, err := http.Get(fullurl)
		if err != nil {
			return err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		api.Mc.Set(&memcache.Item{
			Key:   key,
			Value: []byte(body),
		})

		return c.JSON(200, body)
	}

	return c.JSON(200, fromMc)
}
