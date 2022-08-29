package gonster

import (
	k8Rand "k8s.io/apimachinery/pkg/util/rand"
)

type Simulator interface {
	Start(gonster1 Gonster, gonster2 Gonster, rounds int) (Simulator, error)
}

type ProtoSimulator struct {
	battleLog    [][]SimulatorRoundLog
	battleResult int
}

type SimulatorRoundLog struct {
	remainingHealth   int
	attackProbability int
	isAttackSuccess   bool
	damagePrevented   int
	totalDamage       int
}

const (
	GONSTER1_WIN int = iota
	GONSTER2_WIN
	DRAW
)

func (sim *ProtoSimulator) Start(gonster1 Gonster, gonster2 Gonster, rounds int) (ProtoSimulator, error) {
	gonster1Health := gonster1.Health
	gonster2Health := gonster2.Health

	for i := 0; i < rounds; i++ {
		gonster1AtkLog := SimulateBattle(gonster1, gonster2, 1)
		gonster2AtkLog := SimulateBattle(gonster2, gonster1, 1)

		if gonster1AtkLog.isAttackSuccess {
			gonster2Health = gonster2Health - gonster1AtkLog.totalDamage

		}

		if gonster2AtkLog.isAttackSuccess {
			gonster1Health = gonster1Health - gonster2AtkLog.totalDamage

		}

		gonster1AtkLog.remainingHealth = gonster1Health
		gonster2AtkLog.remainingHealth = gonster2Health

		sim.battleLog = append(sim.battleLog, []SimulatorRoundLog{gonster1AtkLog, gonster2AtkLog})

		isGonster1Defeated := gonster1Health <= 0
		isGonster2Defeated := gonster2Health <= 0

		if isGonster1Defeated && isGonster2Defeated {
			sim.battleResult = DRAW
			return *sim, nil
		}

		if isGonster1Defeated {
			sim.battleResult = GONSTER2_WIN
			return *sim, nil
		}

		if isGonster2Defeated {
			sim.battleResult = GONSTER1_WIN
			return *sim, nil
		}
	}

	sim.battleResult = DRAW

	return *sim, nil
}

func SimulateBattle(attacker Gonster, defender Gonster, seed int64) SimulatorRoundLog {
	k8Rand.Seed(seed)
	atkProb := 10 + (attacker.Head - defender.Leg)
	isAtkSuccess := k8Rand.IntnRange(0, 99) <= atkProb
	percentReduction := float32(defender.Torso) / 100
	damagePrevented := int(float32(attacker.Arm) * percentReduction)
	atkTotalDamage := attacker.Arm - damagePrevented
	atkLog := SimulatorRoundLog{
		attackProbability: atkProb,
		isAttackSuccess:   isAtkSuccess,
		damagePrevented:   damagePrevented,
		totalDamage:       int(atkTotalDamage),
	}

	return atkLog
}
