package postad

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	applicationResponse "learning-go-challenges/application/response"
	"learning-go-challenges/domain/ad"
	"learning-go-challenges/mocks"
	"testing"
	"time"
)

var (
	anUUID              = "dfb485bf-8d6e-4d11-a895-9c4a31d7dda8"
	aTimestamp          = time.Now().UTC()
	mockedAdRepository  = new(mocks.AdRepository)
	mockedUUIDGenerator = new(mocks.UUIDGenerator)
	mockedClock         = new(mocks.Clock)
	service             = PostAdService{AdRepository: mockedAdRepository, UUIDGenerator: mockedUUIDGenerator, Clock: mockedClock}
)

func TestPostAnAd(t *testing.T) {
	mockedAdRepository.EXPECT().Persist(mock.AnythingOfType("Ad")).Return()
	mockedUUIDGenerator.EXPECT().GenerateAsString().Return(anUUID)
	mockedClock.EXPECT().NowAsUTC().Return(aTimestamp)

	response := service.Execute(PostAdRequest{
		Title:       "Titulo 1",
		Description: "Descripcion 1",
		Price:       99,
	})

	expectedDomainAd := ad.Ad{
		Id:          ad.Id{Value: anUUID},
		Title:       "Titulo 1",
		Description: "Descripcion 1",
		Price:       99,
		PublishedAt: aTimestamp,
	}
	expectedAdResponse := applicationResponse.AdResponse{
		Id:          expectedDomainAd.Id.Value,
		Title:       expectedDomainAd.Title,
		Description: expectedDomainAd.Description,
		Price:       expectedDomainAd.Price,
		PublishedAt: expectedDomainAd.PublishedAt.String(),
	}
	mockedAdRepository.AssertCalled(t, "Persist", expectedDomainAd)
	assert.Equal(t, PostAdResponse{AdResponse: expectedAdResponse}, response)
}
