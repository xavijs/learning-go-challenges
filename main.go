package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learning-go-challenges/application/findAd"
	"learning-go-challenges/application/listAds"
	"learning-go-challenges/application/postAd"
	"learning-go-challenges/domain/ad"
	"learning-go-challenges/domain/clock"
	"learning-go-challenges/domain/uuid"
	"learning-go-challenges/infrastructure/http"
	"learning-go-challenges/infrastructure/repository"
	"math/rand"
	netHttp "net/http"
	"strconv"
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

	fmt.Println("Welcome to new Marketplace!!")
	fmt.Println("Insert your Ad")

	title := "Titulo anuncio 1"
	description := "Description anuncio 1"
	var price uint = 99

	postAdResponse := postAdService.Execute(postad.PostAdRequest{
		Title:       title,
		Description: description,
		Price:       price,
	})

	fmt.Println("Posted Ad is:", postAdResponse)
	fmt.Println("Finding AdId:", postAdResponse.AdResponse.Id)

	foundAd := findAdService.Execute(findad.FindAdRequest{AdId: postAdResponse.AdResponse.Id})

	fmt.Println("Found Ad: ", foundAd)

	fmt.Println("Creating several Ads")
	for i := 0; i < 10; i++ {
		postAdService.Execute(
			postad.PostAdRequest{
				Title:       "Anuncio " + strconv.Itoa(i),
				Description: "Descripcion" + strconv.Itoa(i),
				Price:       uint(rand.Uint32()),
			},
		)
	}
	fmt.Printf("%v Ads in repository: %v \n", len(*adRepository.FindAll()), adRepository.FindAll())

	var limitListedAds uint = 2
	listedAds := listAdsService.Execute(listads.ListAdsRequest{Limit: limitListedAds})
	fmt.Printf("Listing %v ads: %+v \n", len(listedAds.Ads), listedAds.Ads)
}
