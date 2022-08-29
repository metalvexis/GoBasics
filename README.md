# GoBasics

My collection of Go practice code.
 
**FizzBuzz Test**

The task is simple: Print integers 1 to N, but print “Fizz” if an integer is divisible by 3, “Buzz” if an integer is divisible by 5, and “FizzBuzz” if an integer is divisible by both 3 and 5.

**Gonster Builder**

Procedural Gonster generator for our auto battler server. A Gonster has 4 body parts: head, torso, arms, and legs. Each body part corresponds to an attribute of the Gonster. All Gonster has a health of 50.

A generated Gonster will be serialized as JSON:

```json
{
	"health": 50,
	"head": 25,
	"torso": 15,
	"arm": 4,
	"leg": 30
}
```

1. Head
    - defines the accuracy of an offensive action (50-90 rounded to nearest 5)
2. Torso
    - defines the reduction of damage received from offensive action. (0-50 rounded to nearest 5)
3. Arm
    - defines the amount of damage the offensive action of the Gonster (3 - 10)
4. Leg
    - defines the chance of dodging an offensive action from opposing Gonster. (0-50 rounded to nearest 5)

**Gonster Battle Simulator**

Given 2 Gonster entering a simulation, the simulator defines the actions a Gonster can do for every turn. The simulator is composed of 100 rounds of battle. Every round, both Gonster attacks at the same time. The simulator will run until a Gonster is left with zero health or the 100 rounds are reached. The Gonster with highest remaining positive health is the victor of the simulation. The simulator will log every action and result of each round and battle.

In every round:

- An attack will succeed depending on the “Head” attribute of the attacker and the “Leg” attribute of the defender.
- A random number will be generated from 0 to 99. The attack succeeds if the generated number is less than or equal to the probability of success
- The probability of success of an attack is defined by:
    
    ```jsx
    10 + (attacker.Head - defender.Leg)
    ```
    
- The defending Gonster will receive damage defined by this formula rounded down to the nearest whole number
    
    ```jsx
    attacker.Arm - (attacker.Arm * (defender.Torso / 100))
    ```
    
- The minimum total received damage is 0
- The attacking Gonster’s computed Total Damage is deducted from the health of the defending Gonster.

 

**Gonster Knockout**

Given a set of Gonster, the program will run a blind draw single-elimination competition. Every Gonster will fight against one competitor in the Gonster Battle Simulator and the winning Gonster advances to the next match. A defeated Gonster is removed from the competition after one loss.

**Auto-battle server**

A single web server for **Gonster Knockout** that is capable of serving 10k concurrent users with a max latency of 150ms per API request. Users can generate 3 Gonsters for their roster and submit it as entry for competition with other users Gonster roster.