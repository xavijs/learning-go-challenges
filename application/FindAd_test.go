package application

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	applicationResponse "learning-go-challenges/application/response"
	"learning-go-challenges/domain/ad"
	"learning-go-challenges/mocks"
	"testing"
	"time"
)

var (
	firstAd = ad.Ad{
		Id:          ad.Id{Value: "660ba87b-794b-417e-a537-dddb042fe82f"},
		Title:       "First Ad Title",
		Description: "First Ad description",
		Price:       4,
		PublishedAt: time.Time{},
	}
)

func TestFindAnExistingAd(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedAdRepository := mocks.NewMockAdRepository(ctrl)
	findAdService := NewFindAdService(mockedAdRepository)

	mockedAdRepository.EXPECT().FindBy(gomock.Any()).Return(firstAd)

	response := findAdService.Execute(FindAdRequest{AdId: "660ba87b-794b-417e-a537-dddb042fe82f"})

	assert.Equal(
		t,
		FindAdResponse{
			AdResponse: applicationResponse.AdResponse{
				Id:          "660ba87b-794b-417e-a537-dddb042fe82f",
				Title:       firstAd.Title,
				Description: firstAd.Description,
				Price:       firstAd.Price,
				PublishedAt: firstAd.PublishedAt.String(),
			},
		},
		response,
	)
}
