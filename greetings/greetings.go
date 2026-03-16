package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty name!")
	}

	message := fmt.Sprintf(randFmt(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	nameMap := make(map[string]string)
	for _, name := range names {
		hello, err := Hello(name)
		if err != nil {
			return make(map[string]string), err
		}

		nameMap[name] = hello
	}

	return nameMap, nil
}

func randFmt() string {
	formats := []string{
		"Hello, %v. Nice to meet you!",
		"Ay yo %v. How's it goin?",
	}

	return formats[rand.Intn(len(formats))]
}
