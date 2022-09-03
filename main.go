package main

import (
	"fmt"
	"os"
	"strconv"

	gonster "github.com/metalvexis/gobasics/gonster"
)

func main() {
	// var input, _ = strconv.Atoi(os.Args[1])
	// fmt.Println(basics.FizzBuzz(input))
	// fmt.Println(basics.Echo(os.Args[1]))

	gonsterGenerator := gonster.ProtoGonsterGen{}

	seed1, _ := strconv.Atoi(os.Args[1])
	seed2, _ := strconv.Atoi(os.Args[2])

	gonster1, _ := gonsterGenerator.Generate(seed1)
	gonster2, _ := gonsterGenerator.Generate(seed2)

	battleSimulator := gonster.ProtoSimulator{}

	battleResult, err := battleSimulator.Start(gonster1, gonster2, 100)

	if err != nil {
		fmt.Println(err)
	}

	lastActionG1 := battleResult.ActionLog[len(battleResult.ActionLog)-1][0]
	lastActionG2 := battleResult.ActionLog[len(battleResult.ActionLog)-1][1]

	fmt.Printf("Gonster1 Last Log : %+v\n\n", lastActionG1)
	fmt.Printf("Gonster2 Last Log : %+v\n\n", lastActionG2)
	fmt.Printf("Result: %s", battleResult.BattleResult)
}
