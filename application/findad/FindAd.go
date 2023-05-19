package findad

import (
	"learning-go-challenges/application/response"
	"learning-go-challenges/domain/ad"
)

type FindAdRequest struct {
	AdId string
}

type FindAdResponse struct {
	AdResponse response.AdResponse
}

type FindAdService struct {
	AdRepository ad.AdRepository
}

func NewFindAdService(adRepository ad.AdRepository) *FindAdService {
	return &FindAdService{AdRepository: adRepository}
}

func (dependencies FindAdService) Execute(request FindAdRequest) FindAdResponse {
	adId := ad.Id{Value: request.AdId}
	foundAd := dependencies.AdRepository.FindBy(adId)

	return FindAdResponse{AdResponse: response.FromDomain(foundAd)}
}
