package gonster

import (
	"fmt"
)

func Simulate() {
	gonsterGenerator := ProtoGonsterGen{}

	newGonster1, _ := gonsterGenerator.Generate(42)
	newGonster2, _ := gonsterGenerator.Generate(43)

	if newGonster1.Health == 0 {
		fmt.Println("Gonster 1 is DEAD")
	}

	if newGonster2.Health == 0 {
		fmt.Println("Gonster 1 is DEAD")
	}
}
