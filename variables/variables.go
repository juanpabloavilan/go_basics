package variables

import (
	"fmt"
)

func RunVariables() {
	fmt.Println("Variables")
	var username string = "Juan "
	fmt.Println("Variables", username)
}

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name
}
