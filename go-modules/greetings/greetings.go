package greetings

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//Hello returns a greeting for a named string
func Hello(name string) (string, error) {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	if name == "" {
		return "", errors.New("name cannot be empty")
	}
	return fmt.Sprintf(randomFormat(), name), nil
}

func randomFormat() string {

	//A slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}

//Hellos provides the greeting messages for the slice of given messages
func Hellos(names []string) (map[string]string, error) {

	if names == nil {
		return nil, errors.New("names cannot be empty")
	}
	msgs := make(map[string]string)

	for _, name := range names {
		msg, err := Hello(name)
		if err != nil {
			return nil, err
		}
		msgs[name] = msg

	}

	return msgs, nil
}
