package greetings_GO

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	if name == "" {
		return "", errors.New("empty name")
	}
	return message, nil
}
