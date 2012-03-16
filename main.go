package main

import (
  "dice"
  "fmt"
)

func printProbabilities(rolls *dice.DieRolls) {
  fmt.Printf("%dd%d\tAvg: %d\tLikely Spread: %d - %d\tCrit: %d\n",
    rolls.Rolls, rolls.Die,
    (rolls.Max + rolls.Min) / 2,
    rolls.Quartile25, rolls.Quartile75,
    rolls.Max)
}

func main() {
  for rolls := 1; rolls <= 10; rolls++ {
    for die := 4; die <= 12; die += 2 {
      printProbabilities(dice.CalculateProbabilities(rolls, die))
    }

    printProbabilities(dice.CalculateProbabilities(rolls, 20))
  }
}

