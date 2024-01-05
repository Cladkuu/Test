package Test

import "github.com/google/uuid"

func Test() string {
	return uuid.New().String()
}

func UUID() uuid.UUID {
	return uuid.New()
}
