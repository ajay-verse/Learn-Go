package helpers

import (
	"github.com/google/uuid"
	"time"
)

func GenerateOrderID() string {
	return uuid.New().String()
}

func GetCurrentTime() time.Time {
	return time.Now()
}
