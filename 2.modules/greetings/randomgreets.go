package greetings

import (
	"math/rand"
)

func RandomFormat() string {
	formats := []string{
		"Hi %v, Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}