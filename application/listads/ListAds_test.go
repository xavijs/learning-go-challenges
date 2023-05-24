package listads

import (
	"github.com/stretchr/testify/assert"
	. "learning-go-challenges/application/response"
	"learning-go-challenges/domain/ad"
	"learning-go-challenges/fixtures"
	"learning-go-challenges/mocks"
	"testing"
)

var (
	mockedAdRepository = new(mocks.AdRepository)
	service            = NewListAdsService(mockedAdRepository)
)

func TestFindEmptyListOfAdsWithLimitUnderMaximumOf5(t *testing.T) {
	adRepositoryStub := mockedAdRepository.EXPECT().FindAll().Return(&[]ad.Ad{}, nil)
	defer adRepositoryStub.Unset()

	response := service.Execute(ListAdsRequest{Limit: 3})

	expectedResponse := ListAdsResponse{}
	assert.Equal(t, expectedResponse, response)
}

func TestFindEmptyListOfAdsWithLimitUpperMaximumOf5(t *testing.T) {
	adRepositoryStub := mockedAdRepository.EXPECT().FindAll().Return(&[]ad.Ad{}, nil)
	defer adRepositoryStub.Unset()

	response := service.Execute(ListAdsRequest{Limit: 8})

	expectedResponse := ListAdsResponse{}
	assert.Equal(t, expectedResponse, response)
}

func TestFindListOfAdsWithLimitBiggerThanExistingAdsInRepository(t *testing.T) {
	adRepositoryStub := mockedAdRepository.EXPECT().FindAll().Return(&[]ad.Ad{fixtures.FirstAd, fixtures.SecondAd}, nil)
	defer adRepositoryStub.Unset()

	response := service.Execute(ListAdsRequest{Limit: 8})

	expectedResponse := ListAdsResponse{
		Ads: []*AdResponse{
			{
				Id:          fixtures.FirstAd.Id.Value,
				Title:       fixtures.FirstAd.Title,
				Description: fixtures.FirstAd.Description,
				Price:       fixtures.FirstAd.Price,
				PublishedAt: fixtures.FirstAd.PublishedAt.String(),
			},
			{
				Id:          fixtures.SecondAd.Id.Value,
				Title:       fixtures.SecondAd.Title,
				Description: fixtures.SecondAd.Description,
				Price:       fixtures.SecondAd.Price,
				PublishedAt: fixtures.SecondAd.PublishedAt.String(),
			},
		},
	}
	assert.Equal(t, expectedResponse, response)
}

func TestFindListOfAdsWith2ElementsLimitUnderMaximumOf5(t *testing.T) {
	adRepositoryStub := mockedAdRepository.EXPECT().FindAll().Return(&[]ad.Ad{fixtures.FirstAd, fixtures.SecondAd}, nil)
	defer adRepositoryStub.Unset()

	response := service.Execute(ListAdsRequest{Limit: 4})

	expectedResponse := ListAdsResponse{
		Ads: []*AdResponse{
			{
				Id:          fixtures.FirstAd.Id.Value,
				Title:       fixtures.FirstAd.Title,
				Description: fixtures.FirstAd.Description,
				Price:       fixtures.FirstAd.Price,
				PublishedAt: fixtures.FirstAd.PublishedAt.String(),
			},
			{
				Id:          fixtures.SecondAd.Id.Value,
				Title:       fixtures.SecondAd.Title,
				Description: fixtures.SecondAd.Description,
				Price:       fixtures.SecondAd.Price,
				PublishedAt: fixtures.SecondAd.PublishedAt.String(),
			},
		},
	}
	assert.Equal(t, expectedResponse, response)
}

func TestFindListOfAdsWithMoreThan5ElementsReturnsFirst5Elements(t *testing.T) {
	adRepositoryStub := mockedAdRepository.EXPECT().FindAll().Return(
		&[]ad.Ad{
			fixtures.FirstAd,
			fixtures.SecondAd,
			fixtures.FirstAd,
			fixtures.SecondAd,
			fixtures.FirstAd,
			fixtures.SecondAd,
		},
		nil,
	)
	defer adRepositoryStub.Unset()

	response := service.Execute(ListAdsRequest{Limit: 8})

	expectedResponse := ListAdsResponse{
		Ads: []*AdResponse{
			{
				Id:          fixtures.FirstAd.Id.Value,
				Title:       fixtures.FirstAd.Title,
				Description: fixtures.FirstAd.Description,
				Price:       fixtures.FirstAd.Price,
				PublishedAt: fixtures.FirstAd.PublishedAt.String(),
			},
			{
				Id:          fixtures.SecondAd.Id.Value,
				Title:       fixtures.SecondAd.Title,
				Description: fixtures.SecondAd.Description,
				Price:       fixtures.SecondAd.Price,
				PublishedAt: fixtures.SecondAd.PublishedAt.String(),
			},
			{
				Id:          fixtures.FirstAd.Id.Value,
				Title:       fixtures.FirstAd.Title,
				Description: fixtures.FirstAd.Description,
				Price:       fixtures.FirstAd.Price,
				PublishedAt: fixtures.FirstAd.PublishedAt.String(),
			},
			{
				Id:          fixtures.SecondAd.Id.Value,
				Title:       fixtures.SecondAd.Title,
				Description: fixtures.SecondAd.Description,
				Price:       fixtures.SecondAd.Price,
				PublishedAt: fixtures.SecondAd.PublishedAt.String(),
			},
			{
				Id:          fixtures.FirstAd.Id.Value,
				Title:       fixtures.FirstAd.Title,
				Description: fixtures.FirstAd.Description,
				Price:       fixtures.FirstAd.Price,
				PublishedAt: fixtures.FirstAd.PublishedAt.String(),
			},
		},
	}
	assert.Equal(t, expectedResponse, response)
}
