package main

import (
	"fmt"
	"learning-go-challenges/application/findAd"
	"learning-go-challenges/application/listAds"
	"learning-go-challenges/application/postAd"
	"learning-go-challenges/domain/ad"
	"learning-go-challenges/domain/clock"
	"learning-go-challenges/domain/uuid"
	"learning-go-challenges/infrastructure/repository"
	"math/rand"
	"strconv"
)

func main() {
	adRepository := repository.NewInMemoryAdRepository(&[]ad.Ad{})
	postAdService := postad.NewPostAdService(adRepository, uuid.RandomUUIDGenerator{}, clock.RealClock{})
	findAdService := findad.NewFindAdService(adRepository)
	listAdsService := listads.NewListAdsService(adRepository)

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
	fmt.Printf("%v Ads in repository: %v \n", len(adRepository.FindAll()), adRepository.FindAll())

	var limitListedAds uint = 2
	listedAds := listAdsService.Execute(listads.ListAdsRequest{Limit: limitListedAds})
	fmt.Printf("Listing %v ads: %v \n", len(listedAds.Ads), listedAds)
}
