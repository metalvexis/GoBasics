package gonster

import (
	"time"

	k8Rand "k8s.io/apimachinery/pkg/util/rand"
)

type Simulator interface {
	Start(gonster1 Gonster, gonster2 Gonster, rounds int) (any, error)
}

type ProtoSimulator struct {
	ActionLog    [][]SimulatorActionLog
	BattleResult string
}

type SimulatorActionLog struct {
	RemainingHealth   int
	DiceRoll          int
	AttackProbability int
	IsAttackSuccess   bool
	DamagePrevented   int
	TotalDamage       int
}

const (
	GONSTER1_WIN = "GONSTER1_WIN"
	GONSTER2_WIN = "GONSTER2_WIN"
	DRAW         = "DRAW"
)

func (sim *ProtoSimulator) Start(gonster1 Gonster, gonster2 Gonster, rounds int) (ProtoSimulator, error) {
	gonster1Health := gonster1.Health
	gonster2Health := gonster2.Health

	for i := 0; i < rounds; i++ {
		gonster1AtkLog := SimulateBattle(gonster1, gonster2, time.Now().UnixNano())
		gonster2AtkLog := SimulateBattle(gonster2, gonster1, time.Now().UnixNano())

		if gonster1AtkLog.IsAttackSuccess {
			gonster2Health = gonster2Health - gonster1AtkLog.TotalDamage
		}

		if gonster2AtkLog.IsAttackSuccess {
			gonster1Health = gonster1Health - gonster2AtkLog.TotalDamage
		}

		gonster1AtkLog.RemainingHealth = gonster1Health
		gonster2AtkLog.RemainingHealth = gonster2Health

		sim.ActionLog = append(sim.ActionLog, []SimulatorActionLog{gonster1AtkLog, gonster2AtkLog})

		isGonster1Defeated := gonster1Health <= 0
		isGonster2Defeated := gonster2Health <= 0

		if isGonster1Defeated && isGonster2Defeated {
			sim.BattleResult = GONSTER1_WIN
			return *sim, nil
		}

		if isGonster1Defeated {
			sim.BattleResult = GONSTER2_WIN
			return *sim, nil
		}

		if isGonster2Defeated {
			sim.BattleResult = GONSTER1_WIN
			return *sim, nil
		}
	}

	sim.BattleResult = DRAW

	return *sim, nil
}

func SimulateBattle(attacker Gonster, defender Gonster, seed int64) SimulatorActionLog {
	k8Rand.Seed(seed)
	atkProb := 10 + (attacker.Head - defender.Leg)
	diceRoll := k8Rand.IntnRange(0, 99)
	isAtkSuccess := diceRoll <= atkProb
	percentReduction := float32(defender.Torso) / 100
	DamagePrevented := int(float32(attacker.Arm) * percentReduction)
	atkTotalDamage := attacker.Arm - DamagePrevented
	actLog := SimulatorActionLog{
		DiceRoll:          diceRoll,
		AttackProbability: atkProb,
		IsAttackSuccess:   isAtkSuccess,
		DamagePrevented:   DamagePrevented,
		TotalDamage:       int(atkTotalDamage),
	}

	return actLog
}
