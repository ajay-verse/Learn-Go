package helpers

import (
	"github.com/google/uuid"
)

func GenerateOrderID() string {
	return uuid.New().String()
}
