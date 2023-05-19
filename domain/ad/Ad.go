package ad

import (
	"time"
)

type Ad struct {
	Id          Id
	Title       string
	Description string
	Price       uint
	PublishedAt time.Time
}

func NewAd(id Id, title string, description string, price uint, publishedAt time.Time) Ad {
	return Ad{Id: id, Title: title, Description: description, Price: price, PublishedAt: publishedAt}
}

type Id struct {
	Value string
}
