package helpers

import (
	"github.com/google/uuid"
	"time"
)

func GenerateRandomID() string {
	return uuid.New().String()
}

func GetCurrentTime() time.Time {
	return time.Now()
}

func GetCurrentTimeString() string {
	return time.Now().Format("Jan 02 2006 03:04:05 PM")
}
