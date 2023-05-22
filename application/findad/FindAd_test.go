package findad

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	applicationResponse "learning-go-challenges/application/response"
	"learning-go-challenges/fixtures"
	"learning-go-challenges/mocks"
	"testing"
)

var (
	mockedAdRepository = new(mocks.AdRepository)
	findAdService      = NewFindAdService(mockedAdRepository)
)

func TestFindAnExistingAd(t *testing.T) {
	mockedAdRepository.EXPECT().FindBy(mock.Anything).Return(&fixtures.FirstAd)

	response := findAdService.Execute(FindAdRequest{AdId: fixtures.FirstAd.Id.Value})

	mockedAdRepository.AssertCalled(t, "FindBy", fixtures.FirstAd.Id)
	assert.Equal(
		t,
		FindAdResponse{
			AdResponse: &applicationResponse.AdResponse{
				Id:          fixtures.FirstAd.Id.Value,
				Title:       fixtures.FirstAd.Title,
				Description: fixtures.FirstAd.Description,
				Price:       fixtures.FirstAd.Price,
				PublishedAt: fixtures.FirstAd.PublishedAt.String(),
			},
		},
		response,
	)
}
