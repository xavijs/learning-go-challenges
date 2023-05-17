package repository

import (
	"github.com/stretchr/testify/assert"
	"learning-go-challenges/domain/ad"
	"testing"
	"time"
)

var (
	memory             []ad.Ad
	inMemoryRepository = NewInMemoryAdRepository(&memory)
)

var (
	firstAd = ad.Ad{
		Id:          ad.NewAdId(),
		Title:       "First Ad Title",
		Description: "First Ad description",
		Price:       4,
		PublishedAt: time.Time{},
	}
	secondAd = ad.Ad{
		Id:          ad.NewAdId(),
		Title:       "Second Ad title",
		Description: "Second Ad description",
		Price:       6,
		PublishedAt: time.Time{},
	}
)

func TestFindAllWith2Ads(t *testing.T) {
	memory = []ad.Ad{firstAd, secondAd}

	var ads = inMemoryRepository.FindAll()

	assert.Contains(t, ads, firstAd, secondAd)
	assert.Len(t, ads, 2)
}

func TestFindAdByExistingId(t *testing.T) {
	memory = []ad.Ad{firstAd, secondAd}

	foundAd := inMemoryRepository.FindBy(firstAd.Id)

	assert.Equal(t, firstAd, foundAd, "Expected to found First Ad")
}

func TestPersistAnAd(t *testing.T) {
	memory = []ad.Ad{}

	inMemoryRepository.Persist(firstAd)

	assert.Contains(t, memory, firstAd)
}
