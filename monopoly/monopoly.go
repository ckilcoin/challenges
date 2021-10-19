package main

import (
	"fmt"
	"math/rand"
)

// Chance cards
const (
	ADVANCE_TO_NEAREST_RAILROAD = iota
	ADVANCE_TO_NEAREST_UTILITY
	ADVANCE_TO_GO
	ADVANCE_TO_BOARDWALK
	ADVANCE_TO_ILLINOIS
	ADVANCE_TO_ST_CHARLES
	ADVANCE_TO_READING
	GO_BACK_THREE_SPACES
	GO_TO_JAIL_CARD

	// no-ops

	CHAIRMAN_OF_THE_BOARD
	GET_OUT_OF_JAIL_FREE
	BANK_PAYS_DIVIDEND
	POOR_TAX
	BUILDING_AND_LOAN_MATURES
	GENERAL_REPAIRS
)

// Special Spaces
const (
	GO                     = 0
	FIRST_COMMUNITY_CHEST  = 2
	READING_RAILROAD       = 5
	FIRST_CHANCE           = 7
	JAIL                   = 10
	ST_CHARLES             = 11
	ELECTRIC_COMPANY       = 12
	SECOND_COMMUNITY_CHEST = 17
	SECOND_CHANCE          = 22
	ILLINOIS               = 24
	WATERWORKS             = 28
	GO_TO_JAIL_POSITION    = 30
	THIRD_COMMUNITY_CHEST  = 33
	THIRD_CHANCE           = 36
	BOARDWALK              = 39
)

const BOARD_SIZE = 40
const NUMBER_CHANCE_CARDS = 16

type Monopoly struct {
	board      [BOARD_SIZE]int
	chancePile [NUMBER_CHANCE_CARDS]int
}

func (m *Monopoly) initialize() {
	m.initializeBoard()
	m.initializeChancePile()
}

// Creates the board
func (m *Monopoly) initializeBoard() {
	for i := 0; i < BOARD_SIZE; i++ {
		m.board[i] = i
	}
}

// Creates the chance pile, which will be shuffled and then returned in order in the games
func (m *Monopoly) initializeChancePile() {
	for i := 0; i < NUMBER_CHANCE_CARDS-1; i++ {
		m.chancePile[i] = i
	}
	// Chance pile has 2 cards for advancing to nearest railroad
	m.chancePile[NUMBER_CHANCE_CARDS-1] = ADVANCE_TO_NEAREST_RAILROAD
}

func (m *Monopoly) shuffleChancePile() {
	rand.Shuffle(NUMBER_CHANCE_CARDS, func(i, j int) {
		m.chancePile[i], m.chancePile[j] = m.chancePile[j], m.chancePile[i]
	})
}

// Based on Community Chest / Chance position, advance to nearest railroad and get position
func (m *Monopoly) getNearestRailroad(currPosition int) int {
	switch currPosition {
	case FIRST_CHANCE:
		return 15
	case SECOND_CHANCE:
		return 25
	case THIRD_CHANCE:
		return READING_RAILROAD
	default:
		fmt.Println("Should not be ending up here, should only be called from certain positions")
		return 0
	}
}

// Based on Community Chest / Chance position, advance to nearest utility and get position
func (m *Monopoly) getNearestUtility(currPosition int) int {
	switch currPosition {
	case FIRST_CHANCE:
		return ELECTRIC_COMPANY
	case SECOND_CHANCE:
		return WATERWORKS
	// could add in case from FIRST_CHANCE, but putting in separate to make it easier to read pictorally
	case THIRD_CHANCE:
		return ELECTRIC_COMPANY
	default:
		fmt.Println("Should not be ending up here, should only be called from certain positions")
		return 0
	}
}

// Draws a card and adjusts the position from landing on Chance
// then puts the card at the bottom of the pile
func (m *Monopoly) getNextPositionFromDrawingChanceCard(currPosition int) (int, bool) {
	// fmt.Printf("Chance pile before drawing: %v\n", m.chancePile)
	chanceCard := m.chancePile[0]
	nextPosition := currPosition
	inJail := false
	switch chanceCard {
	case ADVANCE_TO_NEAREST_RAILROAD:
		nextPosition = m.getNearestRailroad(currPosition)
	case ADVANCE_TO_NEAREST_UTILITY:
		nextPosition = m.getNearestUtility(currPosition)
	case ADVANCE_TO_GO:
		nextPosition = GO
	case ADVANCE_TO_BOARDWALK:
		nextPosition = BOARDWALK
	case ADVANCE_TO_ILLINOIS:
		nextPosition = ILLINOIS
	case ADVANCE_TO_ST_CHARLES:
		nextPosition = ST_CHARLES
	case ADVANCE_TO_READING:
		nextPosition = READING_RAILROAD
	case GO_BACK_THREE_SPACES:
		nextPosition = (currPosition - 3) % BOARD_SIZE
		nextPosition, inJail = m.adjustPositionBasedOnNextPosition(nextPosition)
	case GO_TO_JAIL_CARD:
		nextPosition = JAIL
		inJail = true
		// All cards above and including CHAIRMAN_OF_THE_BOARD are no ops
	}

	// move up cards in chance pile
	for i := 1; i < NUMBER_CHANCE_CARDS; i++ {
		m.chancePile[i-1] = m.chancePile[i]
	}
	m.chancePile[NUMBER_CHANCE_CARDS-1] = chanceCard
	// fmt.Printf("Chance pile after drawing: %v\n", m.chancePile)
	return nextPosition, inJail
}

func (m *Monopoly) getNextPositionFromDrawingCommunityChestCard(currPosition int) (int, bool) {
	// TODO: implement community chest
	return currPosition, false
}

func (m *Monopoly) adjustPositionBasedOnNextPosition(nextPosition int) (int, bool) {
	adjustedPosition := nextPosition
	adjustToJail := false
	switch nextPosition {
	case FIRST_COMMUNITY_CHEST:
	case SECOND_COMMUNITY_CHEST:
	case THIRD_COMMUNITY_CHEST:
		// fmt.Println("Adjusting the community chest, TODO")
		adjustedPosition, adjustToJail = m.getNextPositionFromDrawingCommunityChestCard(nextPosition)
	case FIRST_CHANCE:
	case SECOND_CHANCE:
	case THIRD_CHANCE:
		adjustedPosition, adjustToJail = m.getNextPositionFromDrawingChanceCard(nextPosition)
	case GO_TO_JAIL_POSITION:
		adjustToJail = true
	}
	if adjustToJail {
		adjustedPosition = JAIL
	}
	// if nextPosition != adjustedPosition {
	// 	fmt.Printf("Updated position %v to adjustedPosition %v. In jail? %v \n", nextPosition, adjustedPosition, adjustToJail)
	// }
	return adjustedPosition, adjustToJail
}
