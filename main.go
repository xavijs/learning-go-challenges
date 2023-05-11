package main

import (
	"fmt"
	"learning-go-challenges/application"
	"learning-go-challenges/infrastructure/repository"
	"math/rand"
	"strconv"
)

func main() {
	adRepository := repository.NewInMemoryAdRepository()
	postAdService := application.NewPostAdService(adRepository)
	findAdService := application.NewFindAdService(adRepository)
	listAdsService := application.NewListAdsService(adRepository)

	fmt.Println("Welcome to new Marketplace!!")
	fmt.Println("Insert your Ad")

	title := "Titulo anuncio 1"
	description := "Description anuncio 1"
	var price uint = 99

	postAdResponse := postAdService.Execute(application.PostAdRequest{
		Title:       title,
		Description: description,
		Price:       price,
	})

	fmt.Println("Posted Ad is:", postAdResponse)
	fmt.Println("Finding AdId:", postAdResponse.AdResponse.Id)

	foundAd := findAdService.Execute(application.FindAdRequest{AdId: postAdResponse.AdResponse.Id})

	fmt.Println("Found Ad: ", foundAd)

	fmt.Println("Creating several Ads")
	for i := 0; i < 10; i++ {
		postAdService.Execute(
			application.PostAdRequest{
				Title:       "Anuncio " + strconv.Itoa(i),
				Description: "Descripcion" + strconv.Itoa(i),
				Price:       uint(rand.Uint32()),
			},
		)
	}
	fmt.Printf("%v Ads in repository: %v \n", len(adRepository.FindAll()), adRepository.FindAll())

	var limitListedAds uint = 2
	listedAds := listAdsService.Execute(application.ListAdsRequest{Limit: limitListedAds})
	fmt.Printf("Listing %v ads: %v \n", len(listedAds.Ads), listedAds)
}
