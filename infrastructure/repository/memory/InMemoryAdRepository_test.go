package memory

import (
	"github.com/stretchr/testify/assert"
	"learning-go-challenges/domain/ad"
	"learning-go-challenges/fixtures"
	"testing"
)

var (
	memory             []ad.Ad
	inMemoryRepository = NewInMemoryAdRepository(&memory)
)

func TestFindAllWith2Ads(t *testing.T) {
	memory = []ad.Ad{fixtures.FirstAd, fixtures.SecondAd}

	var ads, _ = inMemoryRepository.FindAll()

	assert.Contains(t, *ads, fixtures.FirstAd, fixtures.SecondAd)
	assert.Len(t, *ads, 2)
}

func TestFindAdByExistingId(t *testing.T) {
	memory = []ad.Ad{fixtures.FirstAd, fixtures.SecondAd}

	foundAd, _ := inMemoryRepository.FindBy(fixtures.FirstAd.Id)

	assert.Equal(t, fixtures.FirstAd, *foundAd, "Expected to found First Ad")
}

func TestPersistAnAd(t *testing.T) {
	memory = []ad.Ad{}

	inMemoryRepository.Persist(&fixtures.FirstAd)

	assert.Contains(t, memory, fixtures.FirstAd)
}
