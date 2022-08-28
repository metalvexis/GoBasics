package main

import (
	"fmt"
	"os"
	"strconv"

	echo "github.com/metalvexis/gobasics/lib/echo"
	basics "github.com/metalvexis/gobasics/lib/fizzbuzz"

	gonster "github.com/metalvexis/gobasics/gonster"
)

func main() {
	var input, _ = strconv.Atoi(os.Args[1])
	fmt.Println(basics.FizzBuzz(input))
	fmt.Println(echo.Echo(os.Args[1]))

	gonsterGenerator := gonster.ProtoGonsterGen{}

	gonsterGenerator.Generate(42)

}
