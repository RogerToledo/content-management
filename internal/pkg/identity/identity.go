package identity

import "github.com/google/uuid"

func GenerateID() (uuid.UUID, error) {
	UUID, err := uuid.NewV7()
	if err != nil {
		return uuid.Nil, err
	}

	return UUID, nil
}

func ParseID(id string) (uuid.UUID, error) {
	UUID, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, err
	}

	return UUID, nil
}

func ValidateID(idRequest string) (string, error) {
	_, err := ParseID(idRequest)
	if err != nil {
		return "", err
	}

	return idRequest, nil
}
