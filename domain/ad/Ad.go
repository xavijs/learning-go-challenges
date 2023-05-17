package ad

import (
	"github.com/google/uuid"
	"time"
)

type Ad struct {
	Id          Id
	Title       string
	Description string
	Price       uint
	PublishedAt time.Time
}

func Empty() Ad {
	return Ad{}
}

func NewAd(title string, description string, price uint) Ad {
	return Ad{Id: NewAdId(), Title: title, Description: description, Price: price, PublishedAt: time.Now()}
}

type Id struct {
	Value string
}

func NewAdId() Id {
	return Id{Value: uuid.NewString()}
}
