package uuid

import "github.com/google/uuid"

type UUIDGenerator interface {
	GenerateAsString() string
}

type RandomUUIDGenerator struct{}

func NewRandomUUIDGenerator() *RandomUUIDGenerator {
	return &RandomUUIDGenerator{}
}

func (RandomUUIDGenerator) GenerateAsString() string {
	return uuid.NewString()
}
