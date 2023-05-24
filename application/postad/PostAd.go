package postad

import (
	. "learning-go-challenges/application/response"
	. "learning-go-challenges/domain/ad"
	. "learning-go-challenges/domain/clock"
	. "learning-go-challenges/domain/uuid"
)

type PostAdRequest struct {
	Title       string
	Description string
	Price       uint
}

type PostAdResponse struct {
	AdResponse *AdResponse
}

type PostAdService struct {
	AdRepository  AdRepository
	UUIDGenerator UUIDGenerator
	Clock         Clock
}

func NewPostAdService(adRepository AdRepository, uuidGenerator UUIDGenerator, clock Clock) *PostAdService {
	return &PostAdService{
		AdRepository:  adRepository,
		UUIDGenerator: uuidGenerator,
		Clock:         clock,
	}
}

func (dependencies *PostAdService) Execute(request PostAdRequest) (*PostAdResponse, error) {
	newAd, err := NewAd(
		Id{Value: dependencies.UUIDGenerator.GenerateAsString()},
		request.Title,
		request.Description,
		request.Price,
		dependencies.Clock.NowAsUTC(),
	)

	if err != nil {
		return nil, err
	}

	dependencies.AdRepository.Persist(newAd)

	return &PostAdResponse{AdResponse: FromDomain(newAd)}, nil
}
