package gonster

import (
	"encoding/json"
	"fmt"
	"testing"
)

type argsGonster struct {
	seed     int64
	expected bool
}

var seedList = []argsGonster{
	{3, true},
	{5, true},
	{9, true},
	{15, true},
	{4, true},
}

func TestGonster(t *testing.T) {
	t.Log("Testing Gonster Generator")
	newGonsterGen := ProtoGonsterGen{}
	for _, test := range seedList {
		result, jsonGonster := newGonsterGen.Generate(test.seed)
		unmarshalled := Gonster{}
		json.Unmarshal(jsonGonster, &unmarshalled)
		prettified, _ := json.MarshalIndent(unmarshalled, "", "\t")
		fmt.Println(string(prettified))

		if unmarshalled.Health == 0 {
			t.Errorf("Output %q not equal to expected %t", result, test.expected)
		}
	}
}
