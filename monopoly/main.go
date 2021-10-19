package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rollDice() (int, bool) {
	dieOne := rand.Intn(5) + 1
	dieTwo := rand.Intn(5) + 1
	return dieOne + dieTwo, dieOne == dieTwo
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var board [BOARD_SIZE]int
	var chancePile [NUMBER_CHANCE_CARDS]int
	m := Monopoly{
		board:      board,
		chancePile: chancePile,
	}
	m.initialize()
	positionToNumLandings := make(map[int]int64)
	for position := 0; position < BOARD_SIZE; position++ {
		positionToNumLandings[position] = 0
	}
	for iteration := 0; iteration < 10000; iteration++ {
		m.shuffleChancePile()
		currPosition := 0
		inJail := false
		turnsInJail := 0
		for turn := 0; turn < 1000; turn++ {
			rollAgain := true
			numDoubles := 0
			for rollAgain {
				rollAgain = false
				rolled, rolledDoubles := rollDice()
				if rolledDoubles {
					numDoubles++
					inJail = false
					rollAgain = true
				} else if inJail && turnsInJail < 3 {
					turnsInJail++
					continue
				}
				if numDoubles == 3 {
					// rolling three doubles sends you to jail and ends the turn
					currPosition = JAIL
					inJail = true
					rollAgain = false
					continue
				}

				nextPosition := (currPosition + rolled) % BOARD_SIZE
				currPosition, inJail = m.adjustPositionBasedOnNextPosition(nextPosition)
				if inJail {
					// if you end up in jail, even after rolling doubles, end the turn
					rollAgain = false
				}
				positionToNumLandings[currPosition]++
			}
		}
	}
	fmt.Println(positionToNumLandings)
}
