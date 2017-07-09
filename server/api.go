package main

import (
	"strconv"

	"github.com/go-pg/pg"
	"github.com/labstack/echo"
)

type API struct {
	Db *pg.DB
}

type SearchTalk struct {
	talk []Talk
}

// Bind setup the API
func (api *API) Bind(group *echo.Group) {
	group.GET("/v1/conf", api.conf)

	group.POST("/v1/talk", api.postTalk)
	group.GET("/v1/talk", api.searchTalk);
	group.GET("/v1/talk/:id", api.getTalk)
	group.PUT("/v1/talk", api.putTalk)

	group.GET("/v1/user/:id/talk/", api.getUserTalk)
	group.GET("/v1/user/:id/talk/:talkid", api.getUserTalk)
	group.PUT("/v1/user/:id/talk/:talkid", api.putUserTalk)
	group.DELETE("/v1/user/:id/talk/:talkid", api.deleteUserTalk)

	group.POST("/v1/user", api.postUser)
	group.GET("/v1/user/:id", api.getUser)
	group.PUT("/v1/user/:id", api.putUser)
}

func (api *API) conf(c echo.Context) error {
	app := c.Get("app").(*App)
	return c.JSON(200, app.Conf.Root)
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
	var talk []Talk
	err := api.Db.Model(&talk).
		Column("talk.*").
		Select()
	if err != nil {
		return err
	}

	searchTalk := SearchTalk{
		talk: talk,
	}

	return c.JSON(200, searchTalk)
}

func (api *API) getTalk(c echo.Context) error {
	id := c.Param("id")

	var talk Talk
	err := api.Db.Model(&talk).
		Column("talk.*", "Talk").
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

	if id != strconv.FormatInt(talk.TalkID, 10) {
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
		Column("usertalk.*", "UserTalk").
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
		Column("user.*", "User").
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

	if id != strconv.FormatInt(user.ID, 10) {
		panic("userIDs do not match")
	}

	err := api.Db.Update(user)
	if err != nil {
		return err
	}

	return c.JSON(200, user)
}
