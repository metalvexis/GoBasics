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

		for round, actionLog := range simResult.ActionLog {
			var gonster1AtkMsg string
			if actionLog[0].IsAttackSuccess {
				gonster1AtkMsg = "landed"
			} else {
				gonster1AtkMsg = "missed"
			}
			gonster1Log := fmt.Sprintf("(%d%%) %d damage %s and prevented %d damage",
				actionLog[0].AttackProbability,
				actionLog[0].TotalDamage,
				gonster1AtkMsg,
				actionLog[0].DamagePrevented)

			var gonster2AtkMsg string
			if actionLog[1].IsAttackSuccess {
				gonster2AtkMsg = "landed"
			} else {
				gonster2AtkMsg = "missed"
			}
			gonster2Log := fmt.Sprintf("(%d%%) %d damage %s and prevented %d damage",
				actionLog[1].AttackProbability,
				actionLog[1].TotalDamage,
				gonster2AtkMsg,
				actionLog[1].DamagePrevented)
			fmt.Printf("Round %d ======\n%s\n%s\n\n", round+1, gonster1Log, gonster2Log)

			fmt.Printf("Gonster1: %d HP\nGonster2: %d HP\n\n",
				actionLog[0].RemainingHealth,
				actionLog[1].RemainingHealth)
		}

	}

	fmt.Println(sim)

}
