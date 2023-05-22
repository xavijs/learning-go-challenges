package ad

import (
	"learning-go-challenges/domain/ad/exception"
	"time"
)

type Ad struct {
	Id          Id
	Title       string
	Description string
	Price       uint
	PublishedAt time.Time
}

func NewAd(id Id, title string, description string, price uint, publishedAt time.Time) (*Ad, error) {
	const MaxDescriptionLength = 50

	if len(description) > MaxDescriptionLength {
		return nil, exception.AdDescriptionTooLongException{}
	}

	return &Ad{Id: id, Title: title, Description: description, Price: price, PublishedAt: publishedAt}, nil
}

type Id struct {
	Value string
}
