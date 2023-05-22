package uuid

import "github.com/google/uuid"

type UUIDGenerator interface {
	GenerateAsString() string
}

type RandomUUIDGenerator struct{}

func (RandomUUIDGenerator) GenerateAsString() string {
	return uuid.NewString()
}
