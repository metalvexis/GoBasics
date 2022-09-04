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
		simResult, _ := sim.Start(test.gonster1, test.gonster2, 50)

		for round, actionLog := range simResult.ActionLog {
			fmt.Printf("Round %d ======\n%s\n%s\n\n", round+1, createLog(&actionLog[0]), createLog(&actionLog[1]))

			fmt.Printf("Gonster1: %d HP\nGonster2: %d HP\n\n",
				actionLog[0].RemainingHealth,
				actionLog[1].RemainingHealth)
		}

	}

	fmt.Println(sim)
}

func createLog(actionLog *SimulatorActionLog) string {
	var atkMsg string

	if actionLog.IsAttackSuccess {
		atkMsg = "landed"
	} else {
		atkMsg = "missed"
	}

	return fmt.Sprintf("(%d%%) %d damage %s and prevented %d damage",
		actionLog.AttackProbability,
		actionLog.TotalDamage,
		atkMsg,
		actionLog.DamagePrevented)
}
