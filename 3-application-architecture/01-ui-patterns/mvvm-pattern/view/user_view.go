package view

import "fmt"

// Render
func Render(fullName string) {
	fmt.Println("------------------------------------")
	fmt.Printf(" [UI DISPLAY]: Hello, %s!\n", fullName)
	fmt.Println("------------------------------------")
}
