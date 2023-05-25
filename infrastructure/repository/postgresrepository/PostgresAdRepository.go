package postgresrepository

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)
import "learning-go-challenges/domain/ad"

type PostgresAdRepository struct {
	dbConnection *gorm.DB
}

func NewPostgresAdRepository(dbConnection *gorm.DB) *PostgresAdRepository {
	return &PostgresAdRepository{dbConnection: dbConnection}
}

func (receiver *PostgresAdRepository) FindBy(id ad.Id) (*ad.Ad, error) {
	var row map[string]interface{}
	receiver.dbConnection.Raw(
		"SELECT id, title, description, price, published_at FROM ads WHERE id = ?;",
		id.Value,
	).Scan(&row)

	if row == nil {
		return nil, nil
	}
	foundAd := toAd(row)

	return &foundAd, nil
}

func (receiver *PostgresAdRepository) Persist(ad *ad.Ad) error {
	sql := "INSERT INTO ads (id, title, description, price, published_at) VALUES (?, ?, ?, ?, ?)"
	values := []interface{}{ad.Id.Value, ad.Title, ad.Description, ad.Price, ad.PublishedAt}

	result := receiver.dbConnection.Exec(sql, values...)
	if result.Error != nil {
		return fmt.Errorf("error persisting Ad %v. Error: %v", ad, result.Error.Error())
	}
	return nil
}

func (receiver *PostgresAdRepository) FindAll() (*[]ad.Ad, error) {
	var rows []map[string]interface{}
	receiver.dbConnection.Raw("SELECT id, title, description, price, published_at FROM ads;").Scan(&rows)

	var foundAds []ad.Ad
	for _, row := range rows {
		foundAds = append(foundAds, toAd(row))
	}

	return &foundAds, nil
}

func toAd(row map[string]interface{}) ad.Ad {
	return ad.Ad{
		Id:          ad.Id{Value: row["id"].(string)},
		Title:       row["title"].(string),
		Description: row["description"].(string),
		Price:       uint(row["price"].(int32)),
		PublishedAt: row["published_at"].(time.Time),
	}
}
