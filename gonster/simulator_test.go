package gonster

import (
	"fmt"
	"testing"
)

type argsSimulator struct {
	gonster1 Gonster
	gonster2 Gonster
	expected bool
}

var simulatorList = []argsSimulator{
	{
		Gonster{Health: 50, Head: 50, Torso: 25, Arm: 15, Leg: 25},
		Gonster{Health: 50, Head: 80, Torso: 10, Arm: 10, Leg: 25},
		false,
	},
}

func TestSimulator(t *testing.T) {
	t.Log("Testing Battle Simulator")
	sim := ProtoSimulator{}

	for _, test := range simulatorList {
		simResult, _ := sim.Start(test.gonster1, test.gonster2, 10)

		for round, battleLog := range simResult.battleLog {
			var gonster1AtkMsg string
			if battleLog[0].isAttackSuccess {
				gonster1AtkMsg = "landed"
			} else {
				gonster1AtkMsg = "missed"
			}
			gonster1Log := fmt.Sprintf("(%d%%) %d damage %s and prevented %d damage",
				battleLog[0].attackProbability,
				battleLog[0].totalDamage,
				gonster1AtkMsg,
				battleLog[0].damagePrevented)

			var gonster2AtkMsg string
			if battleLog[1].isAttackSuccess {
				gonster2AtkMsg = "landed"
			} else {
				gonster2AtkMsg = "missed"
			}
			gonster2Log := fmt.Sprintf("(%d%%) %d damage %s and prevented %d damage",
				battleLog[1].attackProbability,
				battleLog[1].totalDamage,
				gonster2AtkMsg,
				battleLog[1].damagePrevented)
			fmt.Printf("Round %d ======\n%s\n%s\n\n", round+1, gonster1Log, gonster2Log)

			fmt.Printf("Gonster1: %d HP\nGonster2: %d HP\n\n",
				battleLog[0].remainingHealth,
				battleLog[1].remainingHealth)
		}

	}

	fmt.Println(sim)

}
