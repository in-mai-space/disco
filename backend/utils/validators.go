package utils

import "github.com/google/uuid"

func ValidateID(id string) (*uuid.UUID, error) {
	idAsUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return &idAsUUID, nil
}
