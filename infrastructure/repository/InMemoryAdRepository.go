package repository

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
func (receiver *InMemoryAdRepository) Persist(ad Ad) {
	*receiver.memory = append(*receiver.memory, ad)
	fmt.Println("AdRepository memory:", receiver.memory)
}

func (receiver *InMemoryAdRepository) FindBy(id Id) Ad {
	for _, ad := range *receiver.memory {
		if ad.Id == id {
			return ad
		}
	}
	return Empty()
}

func (receiver *InMemoryAdRepository) FindAll() []Ad {
	return *receiver.memory
}
