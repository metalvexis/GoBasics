package basics

import (
	"fmt"
	"testing"
)

func TestMethods(t *testing.T) {
	human := &Animal{}

	fmt.Printf("Species %s said %s\n", human.Specie, human.Speak("Hi"))
}
