package fixtures

import (
	"learning-go-challenges/domain/ad"
	"time"
)

var (
	FirstAd = ad.Ad{
		Id:          ad.Id{Value: "660ba87b-794b-417e-a537-dddb042fe82f"},
		Title:       "First Ad Title",
		Description: "First Ad description",
		Price:       4,
		PublishedAt: time.Time{},
	}
	SecondAd = ad.Ad{
		Id:          ad.Id{Value: "19236f3d-c951-4096-81c3-90e923f4d856"},
		Title:       "Second Ad title",
		Description: "Second Ad description",
		Price:       6,
		PublishedAt: time.Time{},
	}
)
