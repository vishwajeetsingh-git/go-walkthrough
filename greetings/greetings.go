package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name ")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}

func Hellos(names []string) (map[string]string, error){
	//messages := make(map[string]string) // this is correct syntax
	// it initializes the map and returns a reference to it. The make function is used to create slices, maps, and channels in Go. When you use make to create a map, it allocates memory for the map and initializes it so that you can start adding key-value pairs to it immediately.
	messages := map[string]string{} // this is also correct syntax. It creates an empty map literal and assigns it to the variable messages. This syntax is more concise and achieves the same result as using make.
	//var messages map[string]string // this is wrong practice. It declares a variable messages of type map[string]string, but it does not initialize it. The variable messages will have the zero value of a map, which is nil. If you try to add key-value pairs to a nil map, you will get a runtime panic.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}