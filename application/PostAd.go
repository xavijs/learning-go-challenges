package application

import (
	. "learning-go-challenges/application/response"
	. "learning-go-challenges/domain/ad"
)

type PostAdRequest struct {
	Title       string
	Description string
	Price       uint
}

type PostAdResponse struct {
	AdResponse AdResponse
}

type PostAdService struct {
	AdRepository AdRepository
}

func NewPostAdService(adRepository AdRepository) *PostAdService {
	return &PostAdService{AdRepository: adRepository}
}

func (dependencies *PostAdService) Execute(request PostAdRequest) PostAdResponse {
	newAd := NewAd(request.Title, request.Description, request.Price)
	dependencies.AdRepository.Persist(newAd)

	return PostAdResponse{AdResponse: FromDomain(newAd)}
}
