package memory

import (
	"fmt"
	. "learning-go-challenges/domain/ad"
)

type InMemoryAdRepository struct {
	memory *[]Ad
}

func NewInMemoryAdRepository(memory *[]Ad) *InMemoryAdRepository {
	return &InMemoryAdRepository{memory: memory}
}
func (receiver *InMemoryAdRepository) Persist(ad *Ad) error {
	*receiver.memory = append(*receiver.memory, *ad)
	fmt.Println("AdRepository memory:", receiver.memory)
	return nil
}

func (receiver *InMemoryAdRepository) FindBy(id Id) (*Ad, error) {
	for _, ad := range *receiver.memory {
		if ad.Id == id {
			return &ad, nil
		}
	}
	return nil, nil
}

func (receiver *InMemoryAdRepository) FindAll() (*[]Ad, error) {
	return receiver.memory, nil
}
