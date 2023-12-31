package test

import (
	"fmt"
	"testing"
)

func Test_SayHello_ValidArgument(t *testing.T) {
	fmt.Printf("HELLOW")

	inputs := []struct {
		name   string
		result string
	}{
		{name: "Yemeksepeti", result: "Hello Yemeksepeti"},
		{name: "Banabi", result: "Hello Banabi"},
		{name: "Yemek", result: "Hello Yemek"},
		{name: "Yemek", result: "Hello Yemek"},
		{name: "hi", result: "Hello hi"},
	}

	for _, item := range inputs {

		result := sayHello(item.name)
		if result != item.result {
			t.Errorf("\"sayHello('%s')\" failed, expected -> %v, got -> %v", item.name, item.result, result)
		} else {
			t.Logf("\"sayHello('%s')\" succeded, expected -> %v, got -> %v", item.name, item.result, result)
		}
	}
}