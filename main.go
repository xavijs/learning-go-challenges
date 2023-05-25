package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"learning-go-challenges/application/findAd"
	"learning-go-challenges/application/listAds"
	"learning-go-challenges/application/postAd"
	"learning-go-challenges/domain/clock"
	"learning-go-challenges/domain/uuid"
	"learning-go-challenges/infrastructure/http"
	"learning-go-challenges/infrastructure/repository/postgresrepository"
	netHttp "net/http"
)

var (
	adRepository   = postgresrepository.NewPostgresAdRepository(InitDb())
	postAdService  = postad.NewPostAdService(adRepository, uuid.RandomUUIDGenerator{}, clock.RealClock{})
	findAdService  = findad.NewFindAdService(adRepository)
	listAdsService = listads.NewListAdsService(adRepository)
	HttpController = SetupHttpRouter()
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

		postedAd, err := postAdService.Execute(postad.PostAdRequest{
			Title:       request.Title,
			Description: request.Description,
			Price:       request.Price,
		})

		if err != nil {
			c.JSON(netHttp.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(netHttp.StatusCreated, http.FromApplicationResponse(postedAd.AdResponse))
	})
	return r
}

func main() {
	HttpController.Run(":8080")
}

func InitDb() *gorm.DB {
	println("Initializing DB")
	dbParams := "host=localhost user=gochallenges password=123123 dbname=gogogo port=5431 sslmode=disable TimeZone=Europe/Madrid"
	dbConnection, _ := gorm.Open(postgres.Open(dbParams), &gorm.Config{})

	var createAdsTableSql = `CREATE TABLE ads (
								id VARCHAR PRIMARY KEY,
								title VARCHAR NOT NULL,
								description TEXT NOT NULL,
								price INTEGER NOT NULL,
								published_at TIMESTAMP
							);
`
	dbConnection.Exec("TRUNCATE ads;")
	dbConnection.Exec(createAdsTableSql)
	return dbConnection
}
