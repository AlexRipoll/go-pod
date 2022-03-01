package validate

import (
	"errors"
	"github.com/google/uuid"
)

// GenerateID generates a 128 bit (16 byte) Universal Unique IDentifier (UUID) as defined in RFC 4122.
func GenerateID() string {
	return uuid.NewString()
}

// CheckID validates if the given ID has the expected format.
func CheckID(id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return errors.New("invalid uuid provided")
	}
	return nil
}
