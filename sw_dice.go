// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type dieFace struct {
	distance int
	surges   int
	damage   int
}

type die struct {
	faces []dieFace
}

func makeFace(distance int, surges int, damage int) dieFace {
	return dieFace{distance: distance, surges: surges, damage: damage}
}

func addFaces(face1 dieFace, face2 dieFace) dieFace {
 	return dieFace{
		distance: face1.distance + face2.distance,
		surges: face1.surges + face2.surges,
		damage: face1.damage + face2.damage,
		}
}



var redDie = die{
	faces: []dieFace{
		makeFace(0, 0, 1),
		makeFace(0, 0, 2),
		makeFace(0, 0, 2),
		makeFace(0, 1, 2),
		makeFace(0, 0, 3),
		makeFace(0, 0, 3),
	},
}
var blueDie = die{
	faces: []dieFace{
		makeFace(2, 0, 1),
		makeFace(2, 1, 0),
		makeFace(3, 0, 2),
		makeFace(3, 1, 1),
		makeFace(4, 0, 2),
		makeFace(5, 0, 1),
	},
}
var greenDie = die{
	faces: []dieFace{
		makeFace(1, 1, 0),
		makeFace(1, 1, 1),
		makeFace(1, 0, 2),
		makeFace(2, 0, 2),
		makeFace(2, 1, 1),
		makeFace(3, 0, 2),
	},
}
var yellowDie = die{
	faces: []dieFace{
		makeFace(0, 1, 0),
		makeFace(0, 2, 1),
		makeFace(1, 1, 1),
		makeFace(1, 0, 2),
		makeFace(2, 1, 0),
		makeFace(2, 0, 1),
	},
}

func largestValue(die []int) int {
	max := 0
	for i := range die {
		if die[i] > max {
			max = die[i]
		}
	}
	return max
}

func sum(hist []int) int {
	sum := 0
	for i := range hist {
		sum += hist[i]
	}
	return sum
}

func makeHistogram(die []int) []int {
	hist := make([]int, largestValue(die)+1)
	for i := range die {
		hist[die[i]] += 1
	}
	return hist
}

func histogram(dice ...[]int) []int {
	hist := makeHistogram(dice[0])
	rest := dice[1:]
	for i := range rest {
		hist = hist_expand(hist, rest[i])
	}
	return hist
}

func printHist(hist []int) {
	totalCount := sum(hist)
	totalSeen := 0
	over_odds := 100.0
	fmt.Println("  #:  % =  ( % <= / % >= )")
	for i := range hist {
		totalSeen += hist[i]
		individual_odds := 100.0 * float64(hist[i]) / float64(totalCount)
		under_odds := 100.0 * float64(totalSeen) / float64(totalCount)
		fmt.Printf("%3d: %4.1f  (%5.1f / %5.1f)\n", i, individual_odds, under_odds, over_odds)
		over_odds -= individual_odds
	}
}

func hist_expand(hist []int, newDie []int) []int {
	newDie_max := largestValue(newDie)

	new_hist := make([]int, newDie_max+len(hist))
	for i := range newDie {
		for histValue := range hist {
			sum := newDie[i] + histValue
			histCount := hist[histValue]
			new_hist[sum] = new_hist[sum] + histCount
		}
	}

	return new_hist
}

func runAll() {
	redDamage := []int{1, 2, 2, 2, 3, 3}
	blueDamage := []int{1, 0, 2, 1, 2, 1}
	greenDamage := []int{0, 1, 2, 2, 1, 2}
	yellowDamage := []int{0, 1, 1, 2, 0, 1}

	histogram(redDamage, blueDamage, greenDamage, yellowDamage)

	redSurge := []int{0, 0, 0, 1, 0, 0}
	blueSurge := []int{0, 1, 0, 1, 0, 0}
	greenSurge := []int{1, 1, 0, 0, 1, 0}
	yellowSurge := []int{1, 2, 1, 0, 1, 0}
	histogram(redSurge, blueSurge, greenSurge, yellowSurge)

	fmt.Println("Force pike damage")
	printHist(histogram(redDamage, yellowDamage, yellowDamage))

	fmt.Println("Diala perception")
	printHist(histogram(blueSurge, greenSurge, yellowSurge))
}
