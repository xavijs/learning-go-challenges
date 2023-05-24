package ad

import (
	"fmt"
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
		return nil, ErrorDescriptionTooLongException(description)
	}

	return &Ad{Id: id, Title: title, Description: description, Price: price, PublishedAt: publishedAt}, nil
}

func ErrorDescriptionTooLongException(description string) error {
	return fmt.Errorf("ad creation error description too long. Description: %v", description)
}

type Id struct {
	Value string
}
