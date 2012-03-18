package main

import (
  "dice"
  "fmt"
)


func printProbabilities(rolls *dice.DieRolls) {
  fmt.Printf("Rolling %s, min: %d max: %d\n", rolls.Roll, rolls.Min, rolls.Max)

  for index, probability := range rolls.Rolls {
    fmt.Printf("  %2d  %8.5f%%\n", index + rolls.Min, probability * 100)
  }
}

func printAllRolls() {
  for rolls := 1; rolls <= 10; rolls++ {
    for die := 4; die <= 12; die += 2 {
      printProbabilities(dice.CalculateProbabilities(rolls, die))
    }

    printProbabilities(dice.CalculateProbabilities(rolls, 20))
  }
}

func dieRolls_firstTry(sides, rolls, dieSum int) float64 {
  if rolls == 1 {
    if dieSum >= 1 && dieSum <= sides {
      return 1.0 / float64(sides)
    } else {
      return 0
    }
  }

  sum := 0.0

  for i := 1; i <= dieSum - rolls + 1; i++ {
    sum += dieRolls_firstTry(sides, 1, i) *
        dieRolls_firstTry(sides, rolls - 1, dieSum - i)
  }
  return sum
}

func printRolls_firstTry(sides, rolls int) {
  fmt.Printf("%dd%d rolls: \n", rolls, sides)

  maxValue := rolls * sides
  for value := rolls; value <= maxValue; value++ {
    fmt.Printf("  %2d  %8.5f%%\n", value,
       dieRolls_firstTry(sides, rolls, value) * 100)
  }
}

func printDiceRolls(sides, rolls int) {
  printProbabilities(dice.CalculateProbabilities(rolls, sides))
}

func compareRollMethods(sides, rolls int) {
  printDiceRolls(sides, rolls)
  printRolls_firstTry(sides, rolls)
}

func main() {
  compareRollMethods(6, 3)
}

