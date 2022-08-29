package gonster

import (
	"encoding/json"
	"fmt"
	"math/rand"

	k8Rand "k8s.io/apimachinery/pkg/util/rand"

	basics "github.com/metalvexis/gobasics/basics"
)

type Gonster struct {
	Health int
	Head   int
	Torso  int
	Arm    int
	Leg    int
}

type GonsterGenerator interface {
	Generate(seed int) (Gonster, []byte)
}

type ProtoGonsterGen struct{}

func (pgg ProtoGonsterGen) Generate(seed int) (Gonster, []byte) {
	rand.Seed(int64(seed))

	health := 50

	baseHead := k8Rand.IntnRange(50, 90)
	head, _ := basics.RoundBy(baseHead, 5)

	baseTorso := k8Rand.IntnRange(0, 50)
	torso, _ := basics.RoundBy(baseTorso, 5)

	arm := k8Rand.IntnRange(3, 10)

	baseLeg := k8Rand.IntnRange(0, 50)
	leg, _ := basics.RoundBy(baseLeg, 5)

	newGonster := Gonster{health, head, torso, arm, leg}

	gonsterJson, err := json.Marshal(newGonster)

	if err != nil {
		fmt.Println("Cannot serialize Gonster", err)
	}

	return newGonster, gonsterJson
}
