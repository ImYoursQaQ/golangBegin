package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) string {
	message := fmt.Sprintf("hello, %v, welcome you!", name)
	return message
}

// HelloWithError check string
func HelloWithError(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is blank")
	}

	message := fmt.Sprintf("hello, %v, welcome you!", name)
	// for unit test
	//message := fmt.Sprint(randomFormat())
	return message, nil
}

func HelloWithRandom(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is blank")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos 数组，map，循环迭代
func Hellos(names []string) (map[string]string, error) {
	// 声明一个map
	messages := make(map[string]string)
	for _, name := range names {
		message, err := HelloWithRandom(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"yoho, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
