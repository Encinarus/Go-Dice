package main

import (
	"fmt"
	"github.com/Encinarus/Go-Dice/dice"
	"math"
	"time"
)

func printProbabilities(rolls *dice.DieRolls) {
	fmt.Printf("Rolling %s, min: %d max: %d\n", rolls.Roll, rolls.Min, rolls.Max)

	for index, probability := range rolls.Rolls {
		fmt.Printf("  %2d  %8.5f%%\n", index+rolls.Min, probability*100)
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

func dieRolls_firstTry(sides, rolls, dieSum int) int64 {
	if rolls == 1 {
		if dieSum >= 1 && dieSum <= sides {
			return 1.0
		} else {
			return 0
		}
	}

	sum := int64(0)

	for i := 1; i <= dieSum-rolls+1; i++ {
		sum += dieRolls_firstTry(sides, 1, i) *
			dieRolls_firstTry(sides, rolls-1, dieSum-i)
	}
	return sum
}

func dieRolls_secondTryWithCache(sides, rolls, dieSum int,
	cache map[string]int64) int64 {
	cacheKey := fmt.Sprintf("%dd%d:%d", rolls, sides, dieSum)
	value, ok := cache[cacheKey]
	if ok {
		return value
	}

	if rolls == 1 {
		if dieSum >= 1 && dieSum <= sides {
			return 1
		} else {
			return 0
		}
	}

	sum := int64(0)

	for i := 1; i <= dieSum-rolls+1; i++ {
		sum += dieRolls_secondTryWithCache(sides, 1, i, cache) *
			dieRolls_secondTryWithCache(sides, rolls-1, dieSum-i, cache)
	}
	cache[cacheKey] = sum
	return sum
}

func printRolls_firstTry(sides, rolls int) {
	fmt.Printf("%dd%d rolls: \n", rolls, sides)

	maxValue := rolls * sides
	rollSequenceCount := math.Pow(float64(sides), float64(rolls))
	for value := rolls; value <= maxValue; value++ {
		valueRolls := float64(dieRolls_firstTry(sides, rolls, value))
		fmt.Printf(" %3d  %8.5f%%\n", value,
			valueRolls/rollSequenceCount*100)
	}
}

func printRolls_secondTry(sides, rolls int) {
	fmt.Printf("%dd%d rolls: \n", rolls, sides)

	maxValue := rolls * sides
	rollSequenceCount := math.Pow(float64(sides), float64(rolls))
	cache := make(map[string]int64)
	for value := rolls; value <= maxValue; value++ {
		valueRolls := float64(dieRolls_secondTryWithCache(sides, rolls, value, cache))
		fmt.Printf(" %3d  %8.5f%%\n", value,
			valueRolls/rollSequenceCount*100)
	}
}

func printDiceRolls(sides, rolls int) {
	printProbabilities(dice.CalculateProbabilities(rolls, sides))
}

func compareRollMethods(sides, rolls int) {
	start1 := time.Now()
	printDiceRolls(sides, rolls)
	start2 := time.Now()
	printRolls_secondTry(sides, rolls)
	end := time.Now()

	binomialMillis := (start2.Sub(start1)) / 1000000
	cachedMillis := (end.Sub(start2)) / 1000000
	fmt.Printf("Binomial took %d millis\n", binomialMillis)
	fmt.Printf("Cached Recursive took %d millis\n", cachedMillis)
	fmt.Printf("Binomial was %5.2f times faster\n", float64(cachedMillis)/float64(binomialMillis))

	//printRolls_firstTry(sides, rolls)
}

func main() {
	compareRollMethods(12, 15)
}
