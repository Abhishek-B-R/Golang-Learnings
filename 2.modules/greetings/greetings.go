package greetings

import (
	"fmt"
	"errors"
)

func Hello(name string) (string,error){
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(RandomFormat(),name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error){
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}