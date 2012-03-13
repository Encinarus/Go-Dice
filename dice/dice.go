package dice

import (
  "fmt"
  "math"
)

func chooseHelper(n int, k int, cache map[string]int64) int64 {
  // chooseHelper implements the recursive method of determining binomial
  // coefficients with memoization, as indicated at
  // http://en.wikipedia.org/wiki/Binomial_coefficient

  // if n==k we'll reduce to the k = 0 case and don't need to traverse
  if n == k {
    return 1
  }
  // if n < k, we'll reduce to the n = 0 case and don't need to traverse
  if n < k {
    return 0
  }
  if n == 0 {
    return 0
  }
  if k == 0 {
    return 1
  }

  // take advantage of symetry of pascals triangle and normalize to the
  // "left half" of the tree
  if k > (n / 2) {
    k = n - k
  }
  nCk := string(n) + "C" + string(k)

  value, ok := cache[nCk]
  if ok {
    return value
  }

  value = chooseHelper(n-1, k-1, cache) + chooseHelper(n-1, k, cache)
  cache[nCk] = value
  return value
}

func choose(n, k int) int64 {
  cache := make(map[string]int64)
  return chooseHelper(n, k, cache)
}

func probability(p, n, s int, cache map[string]int64) float64 {
  totalRolls := math.Pow(float64(s), float64(n))
  count := possibilityCount(p, n, s, cache)

  return float64(count) / totalRolls
}

func possibilityCount(p, n, s int, cache map[string]int64) int64 {
  // See http://mathworld.wolfram.com/Dice.html for an explanation of this
  kMax := int(math.Floor(float64(p - n) / float64(s)))
  
  total := int64(0)
  sign := int64(-1)
  for k := int(0); k <= kMax; k++ {
    sign *= int64(-1)
    left := chooseHelper(n, k, cache)
    right := chooseHelper((p - s*k - 1), n - 1, cache)

    term := sign * left * right
    total += term
  }

  return total
}

type DieRolls struct {
  die int
  rolls int
  min int
  max int
  quartile25 int
  quartile50 int
  quartile75 int
}

// Calculates the min, max and quartile values that can be generated by
// rolling a die rolls times.
//
func calculateProbabilities(rolls, die int) *DieRolls {
  dieRolls := new(DieRolls)
  dieRolls.rolls = rolls
  dieRolls.die = die
  dieRolls.min = rolls
  dieRolls.max = die * rolls

  cumilativeProbability := 0.0
  cache := make(map[string]int64)
  for value := dieRolls.min; value <= dieRolls.max; value++ {
    pOfValue := probability(value, rolls, die, cache)
    if dieRolls.quartile25 == 0 && (pOfValue + cumilativeProbability) > .25 {
      dieRolls.quartile25 = value
    } else if dieRolls.quartile50 == 0 && (pOfValue + cumilativeProbability) > .5 {
      dieRolls.quartile50 = value
    } else if dieRolls.quartile75 == 0 && (pOfValue + cumilativeProbability) > .75 {
      dieRolls.quartile75 = value
    }
    cumilativeProbability += pOfValue
  }

  return dieRolls
}

