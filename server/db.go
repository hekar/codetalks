package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// User model for database
type User struct {
	ID     int      `json:"id",sql:"type:serial"`
	Name   string   `json:"name"`
	Emails []string `json:"emails"`
}

// UserTalk model for database
type UserTalk struct {
	ID         int    `json:"id",sql:"type:serial"`
	UserID     int    `json:"userId"`
	TalkID     int    `json:"talkId"`
	Status     int    `json:"status"` // 0 - plan to watch, 1 completed, 2 watching, 3 dropped
	Comments   string `json:"comments"`
	DateViewed int    `json:"dateViewed"`
}

// Talk model
type Talk struct {
	ID           int      `json:"id",sql:"type:serial"`
	Name         string   `json:"name"`
	URL          string   `json:"url"`
	ThumbnailURL string   `json:"thumbnailUrl"`
	Tags         []string `json:"tags"`
}

// Talk Profile model
type TalkProfile struct {
	TalkID     int    `json:"talkId"`
	Presenter  string `json:"presenter"`
	Summary    string `json:"summary"`
	Site       string `json:"site"`
	Duration   string `json:"duration"`
	DatePosted string `json:"datePosted"`
}

// Talk Stats model
type TalkStats struct {
	TalkID     int `json:"talkId"`
	Score      int `json:"score"`
	Viewed     int `json:"viewed"`
	Rank       int `json:"rank"`
	Bookmarked int `json:"bookmarked"`
}

// TalkPopular List of popular talks
type TalkPopular struct {
	ID     int `json:"id",sql:"type:serial"`
	TalkID int `json:"talkId"`
	Rank   int `json:"rank"`
}

// CreateSchema create the database schema
func CreateSchema(db *pg.DB) error {
	fmt.Println("Creating schema")
	for _, model := range []interface{}{
		&User{},
		&UserTalk{},
		&Talk{},
		&TalkStats{},
		&TalkProfile{},
		&TalkPopular{},
	} {
		err := db.DropTable(model, &orm.DropTableOptions{
			IfExists: true,
		})
		if err != nil {
			return err
		}

		err = db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}

	fmt.Println("Completed database creation")

	Must(db.Insert(&User{
		Name: "This is user",
		Emails: []string{
			"user1@example.com",
		},
	}))

	Must(db.Insert(&Talk{
		ID:           1,
		Name:         "Bjarne Stroustrup - The Essence of C++",
		URL:          "https://www.youtube.com/watch?v=86xWVb4XIyE",
		ThumbnailURL: "https://i.ytimg.com/an_webp/86xWVb4XIyE/mqdefault_6s.webp?du=3000&sqp=CKb708sF&rs=AOn4CLDgRM5ZQwHj8tre1P0MLtd84ZGw4w",
		Tags: []string{
			"cpp",
		},
	}))

	Must(db.Insert(&TalkProfile{
		TalkID:     1,
		Presenter:  "Bjarne Stroustrup",
		Summary:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse maximus faucibus elementum. Mauris ullamcorper orci ac dapibus euismod. Vivamus mollis lorem tellus, a aliquam nulla aliquet nec. Proin tristique ex erat, sit amet molestie tellus luctus vel. Curabitur consectetur ac nulla luctus placerat. Donec tincidunt elementum molestie. Vivamus pulvinar nec ligula sed varius. Maecenas ante ligula, faucibus vitae ex vitae, ornare scelerisque arcu. Praesent fermentum odio at enim facilisis vehicula. Quisque risus dui, pretium ac libero viverra, eleifend pharetra nibh. Sed sem purus, molestie vitae euismod eu, gravida eu urna. Morbi egestas venenatis urna, id scelerisque augue tincidunt ultrices. ",
		Site:       "youtube",
		Duration:   "3600",
		DatePosted: "01-01-1970",
	}))

	Must(db.Insert(&Talk{
		ID:           2,
		Name:         "Tech Talk: Linus Torvalds on git",
		URL:          "https://www.youtube.com/watch?v=4XpnKHJAok8",
		ThumbnailURL: "https://i.ytimg.com/vi/4XpnKHJAok8/hqdefault.jpg?sqp=-oaymwEXCPYBEIoBSFryq4qpAwkIARUAAIhCGAE=&rs=AOn4CLAlzaJGQnDKZxr4ufeSuaLDOamRjg",
		Tags: []string{
			"linus",
			"git",
		},
	}))

	Must(db.Insert(&Talk{
		ID:           3,
		Name:         "Progressive, Performant, Polymer: Pick Three - Google I/O 2016",
		URL:          "https://www.youtube.com/watch?v=J4i0xJnQUzU&index=2&list=PL00z3DSeZW7wDVgFVboA-5rBwkI-8WT_R",
		ThumbnailURL: "https://i.ytimg.com/vi/J4i0xJnQUzU/hqdefault.jpg?sqp=-oaymwEWCKgBEF5IWvKriqkDCQgBFQAAiEIYAQ==&rs=AOn4CLCg9ScmRIx1VdTWzpEIumVVL3SYQw",
		Tags: []string{
			"google",
			"polymer",
			"web",
		},
	}))

	Must(db.Insert(&Talk{
		ID:           4,
		Name:         "JavaScript does NOT offer zero-cost abstractions",
		URL:          "https://www.youtube.com/watch?v=yLv3hafmSas&list=PL00z3DSeZW7zF104m6m055Wgm-7i9e4rV&index=4",
		ThumbnailURL: "https://i.ytimg.com/vi/yLv3hafmSas/hqdefault.jpg?sqp=-oaymwEWCKgBEF5IWvKriqkDCQgBFQAAiEIYAQ==&rs=AOn4CLDd15vgesB1PaB2S0S2w7gBalEh0Q",
		Tags: []string{
			"javascript",
			"performance",
		},
	}))

	Must(db.Insert(&Talk{
		ID:           5,
		Name:         "The Vulkan Graphics API - what it means for Linux",
		URL:          "https://www.youtube.com/watch?v=ynyO3O3zd3E&list=PL00z3DSeZW7zF104m6m055Wgm-7i9e4rV&index=9",
		ThumbnailURL: "https://i.ytimg.com/vi/ynyO3O3zd3E/hqdefault.jpg?sqp=-oaymwEWCKgBEF5IWvKriqkDCQgBFQAAiEIYAQ==&rs=AOn4CLAFrPoNSa8eoipULf2x95u_ZknvFQ",
		Tags: []string{
			"linux",
			"vulkan",
			"graphics",
		},
	}))

	Must(db.Insert(&Talk{
		ID:           6,
		Name:         "RxJS In-Depth – Ben Lesh",
		URL:          "https://www.youtube.com/watch?v=KOOT7BArVHQ&list=PL00z3DSeZW7zF104m6m055Wgm-7i9e4rV&index=11",
		ThumbnailURL: "https://i.ytimg.com/vi/KOOT7BArVHQ/hqdefault.jpg?sqp=-oaymwEWCKgBEF5IWvKriqkDCQgBFQAAiEIYAQ==&rs=AOn4CLAre3UZs8p6VH7EJtdNUXHiwac5ZA",
		Tags: []string{
			"rxjs",
			"javascript",
		},
	}))

	Must(db.Insert(&Talk{
		ID:           7,
		Name:         "This is a talk",
		URL:          "https://youtube.com",
		ThumbnailURL: "http://image.example.com",
		Tags: []string{
			"youtube",
		},
	}))

	Must(db.Insert(&Talk{
		ID:           8,
		Name:         "This is a talk",
		URL:          "https://youtube.com",
		ThumbnailURL: "http://image.example.com",
		Tags: []string{
			"youtube",
		},
	}))

	Must(db.Insert(&TalkPopular{
		ID:     1,
		TalkID: 1,
		Rank:   3,
	}))

	Must(db.Insert(&TalkPopular{
		ID:     2,
		TalkID: 2,
		Rank:   2,
	}))

	Must(db.Insert(&TalkPopular{
		ID:     3,
		TalkID: 3,
		Rank:   1,
	}))

	db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
		query, err := event.FormattedQuery()
		if err != nil {
			panic(err)
		}

		log.Printf("%s %s", time.Since(event.StartTime), query)
	})

	return nil
}
