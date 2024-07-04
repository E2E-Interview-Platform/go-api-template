package helpers

import "github.com/google/uuid"

type IDGenerator interface {
	GenerateId() string
}

type UUIDGenerator struct{}

func (generator UUIDGenerator) GenerateId() string {
	id := uuid.New()
	return id.String()
}
