package common

import "github.com/google/uuid"

// GenerateUUID generates random string
func GenerateUUID() (string, error) {
	uuid, err := uuid.NewRandom()
	return uuid.String(), err
}
