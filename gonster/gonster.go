package gonster

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type Gonster struct {
	Health int
	Head   int
	Torso  int
	Arm    int
	Leg    int
}

type GonsterGenerator interface {
	Generate(seed int64) (Gonster, []byte)
}

type ProtoGonsterGen struct{}

func (pgg ProtoGonsterGen) Generate(seed int64) (Gonster, []byte) {
	rand.Seed(seed)

	health := 50
	head := rand.Int()
	torso := rand.Int()
	arm := rand.Int()
	leg := rand.Int()

	newGonster := Gonster{health, head, torso, arm, leg}

	gonsterJson, err := json.Marshal(newGonster)

	if err != nil {
		fmt.Println("Cannot serialize Gonster", err)
	}

	return newGonster, gonsterJson
}
