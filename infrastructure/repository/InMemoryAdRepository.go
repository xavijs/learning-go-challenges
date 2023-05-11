package repository

import (
	"fmt"
	. "learning-go-challenges/domain/ad"
)

type InMemoryAdRepository struct {
}

var memory []Ad

func NewInMemoryAdRepository() InMemoryAdRepository {
	return InMemoryAdRepository{}
}

func (repository InMemoryAdRepository) Persist(ad Ad) {
	memory = append(memory, ad)
	fmt.Println("AdRepository memory:", memory)
}

func (repository InMemoryAdRepository) FindBy(id Id) Ad {
	for index := range memory {
		if memory[index].Id == id {
			return memory[index]
		}
	}
	panic("Ad not found!")
}

func (repository InMemoryAdRepository) FindAll() []Ad {
	return memory
}
