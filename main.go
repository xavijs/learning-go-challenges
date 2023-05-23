package main

import (
	"github.com/gin-gonic/gin"
	"learning-go-challenges/application/findAd"
	"learning-go-challenges/application/listAds"
	"learning-go-challenges/application/postAd"
	"learning-go-challenges/domain/ad"
	"learning-go-challenges/domain/clock"
	"learning-go-challenges/domain/uuid"
	"learning-go-challenges/infrastructure/http"
	"learning-go-challenges/infrastructure/repository"
	netHttp "net/http"
)

var (
	RepositoryMemory = &[]ad.Ad{}
	adRepository     = repository.NewInMemoryAdRepository(RepositoryMemory)
	postAdService    = postad.NewPostAdService(adRepository, uuid.RandomUUIDGenerator{}, clock.RealClock{})
	findAdService    = findad.NewFindAdService(adRepository)
	listAdsService   = listads.NewListAdsService(adRepository)
	HttpController   = SetupHttpRouter()
)

func SetupHttpRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ads/:id", func(c *gin.Context) {
		response := findAdService.Execute(findad.FindAdRequest{
			AdId: c.Param("id"),
		})
		switch response.AdResponse {
		case nil:
			c.JSON(netHttp.StatusNotFound, nil)
		default:
			c.JSON(netHttp.StatusOK, http.FromApplicationResponse(response.AdResponse))
		}
	})
	r.GET("/ads", func(c *gin.Context) {
		response := listAdsService.Execute(listads.ListAdsRequest{Limit: 5})

		httpResponse := make([]http.AdHttpResponse, 0)
		for _, adResponse := range response.Ads {
			httpResponse = append(httpResponse, http.FromApplicationResponse(adResponse))
		}
		c.JSON(netHttp.StatusOK, httpResponse)
	})
	r.POST("/ads", func(c *gin.Context) {
		var request http.PostAdHttpRequest

		if err := c.BindJSON(&request); err != nil {
			c.JSON(netHttp.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		postedAd := postAdService.Execute(postad.PostAdRequest{
			Title:       request.Title,
			Description: request.Description,
			Price:       request.Price,
		})
		c.JSON(netHttp.StatusCreated, http.FromApplicationResponse(postedAd.AdResponse))
	})
	return r
}

func main() {
	HttpController.Run(":8080")
}
