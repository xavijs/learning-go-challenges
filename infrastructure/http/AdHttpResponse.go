package http

import "learning-go-challenges/application/response"

type AdHttpResponse struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	PublishedAt string `json:"publishedAt"`
}

func FromApplicationResponse(applicationAdResponse *response.AdResponse) AdHttpResponse {
	return AdHttpResponse{
		Id:          applicationAdResponse.Id,
		Title:       applicationAdResponse.Title,
		Description: applicationAdResponse.Description,
		Price:       applicationAdResponse.Price,
		PublishedAt: applicationAdResponse.PublishedAt,
	}
}
