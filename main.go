package test

import "fmt"

func sayHello(name string) string {
	if len(name) == 0 {
		return "Hello Anonymous"
	}
	fmt.Printf("HELLOW")
	fmt.Printf("HELLOW")
	fmt.Printf("HELLOW")
	fmt.Printf("HELLOW")



	return fmt.Sprintf("Hello %s", name)
}
