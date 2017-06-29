package main

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// User model for database
type User struct {
	ID     int64
	Name   string
	Emails []string
}

// UserTalk model for database
type UserTalk struct {
	UserID     int64
	TalkID     int64
	Status     int64 // 0 - plan to watch, 1 completed, 2 watching, 3 dropped
	Comments   string
	DateViewed int64
}

type Talk struct {
	TalkID       int64
	Name         string
	Url          string
	ThumbnailUrl string
	Tags         []string
}

type TalkMeta struct {
	TalkID  int64
	Length  int64
	Author  string
	Summary string
}

// CreateSchema create the database schema
func CreateSchema(db *pg.DB) error {
	for _, model := range []interface{}{&User{}} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
