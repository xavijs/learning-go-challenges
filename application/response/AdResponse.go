package response

import "learning-go-challenges/domain/ad"

type AdResponse struct {
	Id          string
	Title       string
	Description string
	Price       uint
	PublishedAt string
}

func FromDomain(domainAd ad.Ad) AdResponse {
	return AdResponse{
		Id:          domainAd.Id.Value,
		Title:       domainAd.Title,
		Description: domainAd.Description,
		Price:       domainAd.Price,
		PublishedAt: domainAd.PublishedAt.String(),
	}
}
