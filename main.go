package main

import (
  "dice"
)

func main() {
  rolls := calculateProbabilities(10, 12)
  fmt.Printf("Rolling %dd%d : %d %d %d %d %d\n",
    rolls.rolls, rolls.die,
    rolls.min,
    rolls.quartile25, rolls.quartile50, rolls.quartile75,
    rolls.max)
}

