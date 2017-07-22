package main

import (
	"fmt"

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

// Talk model
type Talk struct {
	TalkID       int64
	Name         string
	Url          string
	ThumbnailUrl string
	Tags         []string
}

// TalkMeta Additional meta data relating to a talk
type TalkMeta struct {
	TalkID  int64
	Length  int64
	Author  string
	Summary string
}

// CreateSchema create the database schema
func CreateSchema(db *pg.DB) error {
	for _, model := range []interface{}{&User{}, &UserTalk{}, &Talk{}, &TalkMeta{}} {
		fmt.Println("Dropping tables")
		err := db.DropTable(model, &orm.DropTableOptions{
			IfExists: true,
		})
		if err != nil {
			return err
		}

		fmt.Println("Creating Tables")
		err = db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}

	fmt.Println("Completed database creation")

	user := &User{
		Name: "This is user",
		Emails: []string{
			"user1@example.com",
		},
	}
	err := db.Insert(user)
	if err != nil {
		return err
	}

	talk := &Talk{
		Name:         "This is a talk",
		Url:          "https://youtube.com",
		ThumbnailUrl: "http://image.example.com",
		Tags: []string{
			"youtube",
		},
	}
	err = db.Insert(talk)
	if err != nil {
		return err
	}

	return nil
}
