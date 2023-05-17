package application

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	applicationResponse "learning-go-challenges/application/response"
	"learning-go-challenges/domain/ad"
	"testing"
	"time"
)

type MockedAdRepository struct {
	mock.Mock
}

func (m *MockedAdRepository) FindBy(id ad.Id) ad.Ad {
	return m.Called(id).Get(0).(ad.Ad)
}

func (m *MockedAdRepository) Persist(ad ad.Ad) {
	m.Called(ad).Get(0)
}

func (m *MockedAdRepository) FindAll() []ad.Ad {
	return m.Called().Get(0).([]ad.Ad)
}

var (
	mockedAdRepository = new(MockedAdRepository)
	findAdService      = NewFindAdService(mockedAdRepository)
	firstAd            = ad.Ad{
		Id:          ad.Id{Value: "660ba87b-794b-417e-a537-dddb042fe82f"},
		Title:       "First Ad Title",
		Description: "First Ad description",
		Price:       4,
		PublishedAt: time.Time{},
	}
)

func TestFindAnExistingAd(t *testing.T) {
	mockedAdRepository.On("FindBy", mock.AnythingOfType("ad.Id")).Return(firstAd)

	response := findAdService.Execute(FindAdRequest{AdId: "660ba87b-794b-417e-a537-dddb042fe82f"})

	mockedAdRepository.AssertCalled(t, "FindBy", firstAd.Id)
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

func TestFindNonExistingAd(t *testing.T) {
	mockedAdRepository.On("FindBy", mock.AnythingOfType("ad.Id")).Return(ad.Empty())

	response := findAdService.Execute(FindAdRequest{AdId: "660ba87b-794b-417e-a537-dddb042fe82f"})

	mockedAdRepository.AssertCalled(t, "FindBy", firstAd.Id)
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
