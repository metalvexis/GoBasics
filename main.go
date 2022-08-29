package main

import (
	"fmt"
	"os"
	"strconv"

	basics "github.com/metalvexis/gobasics/basics"

	gonster "github.com/metalvexis/gobasics/gonster"
)

func main() {
	var input, _ = strconv.Atoi(os.Args[1])
	fmt.Println(basics.FizzBuzz(input))
	fmt.Println(basics.Echo(os.Args[1]))

	gonsterGenerator := gonster.ProtoGonsterGen{}

	gonsterGenerator.Generate(42)

	basics.RoundBy(6, 5)

}
