package pkg

import "github.com/google/uuid"

func GenerateUUID() string {
	uuidGenerated := uuid.New()

	return uuidGenerated.String()
}

func IsUUID(ctx string) bool {
	_, err := uuid.Parse(ctx)

	if err != nil {
		return false
	}

	return true
}
