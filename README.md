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

Given 2 Gonster entering a simulation, the simulator defines the actions a Gonster can do for every turn. The simulator is composed of multiple rounds of battle, wherein, a round is composed of 2 turns. On the first turn of the round, the first Gonster is allowed to be in the offensive the second Gonster is in the defensive. On the second, Gonster trade places.  The simulator will run an indefinite amount of rounds until a Gonster is left with zero health. The Gonster with remaining positive health is the victor of the simulation. The simulator will log every action and result of each turn.

 

**Gonster Knockout**

Given a set of Gonster, the program will run a blind draw single-elimination competition. Every Gonster will fight against one competitor in the Gonster Battle Simulator and the winning Gonster advances to the next match. A defeated Gonster is removed from the competition after one loss.

**Auto-battle server**

A single web server for **Gonster Knockout** that is capable of serving 10k concurrent users with a max latency of 150ms per API request. Users can generate 3 Gonsters for their roster and submit it as entry for competition with other users Gonster roster.