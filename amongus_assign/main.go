package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var playersToImposters = map[int]int{
	5:  1,
	6:  1,
	7:  2,
	8:  2,
	9:  2,
	10: 2,
}

var CREW = 'C'
var IMPOSTER = 'I'
var NO_PREFERENCE = 'N'

func assign_users(preferences string) (crew []int, imposters []int, assignmentScore int) {
	numPlayers := len(preferences)
	if numImposters, ok := playersToImposters[numPlayers]; ok {
		numCrew := numPlayers - numImposters
		preferCrew := []int{}
		preferImposter := []int{}
		noPreferences := []int{}
		for player, preference := range preferences {
			switch preference {
			case CREW:
				preferCrew = append(preferCrew, player)
			case IMPOSTER:
				preferImposter = append(preferImposter, player)
			case NO_PREFERENCE:
				noPreferences = append(noPreferences, player)
			default:
				return nil, nil, -1
			}
		}
		imposters := []int{}
		crew := []int{}
		numPreferCrew := len(preferCrew)
		numPreferImposter := len(preferImposter)
		assignmentScore := 0
		if numCrew >= numPreferCrew && numImposters >= numPreferImposter {
			// if everyone can get their preference, give them it
			imposters = preferImposter
			crew = preferCrew
			// randomly assign no preference players
			rand.Seed(time.Now().UnixNano())
			rand.Shuffle(len(noPreferences), func(i, j int) {
				noPreferences[i], noPreferences[j] = noPreferences[j], noPreferences[i]
			})
			numAssignedImposter := len(imposters)
			for _, player := range noPreferences {
				if numAssignedImposter < numImposters {
					numAssignedImposter += 1
					imposters = append(imposters, player)
				} else {
					crew = append(crew, player)
				}
			}
		} else if numCrew >= numPreferCrew {
			crew = preferCrew
			// if more people want to be imposter, all noPreferences will be crew
			// and some preferImposter will be crew as well.
			crew = append(crew, noPreferences...)
			rand.Seed(time.Now().UnixNano())
			rand.Shuffle(len(preferImposter), func(i, j int) {
				preferImposter[i], preferImposter[j] = preferImposter[j], preferImposter[i]
			})
			numAssignedImposter := 0
			for _, player := range preferImposter {
				if numAssignedImposter == numImposters {
					assignmentScore += 1
					crew = append(crew, player)
				} else {
					imposters = append(imposters, player)
					numAssignedImposter += 1
				}
			}
		} else if numImposters >= numPreferImposter {
			// not enough people want to be imposter
			// since numCrew > numPreferCrew, cannot have len(noPrefences) + len(preferImposter)
			// be greater than numImposters, so all noPreferences become imposters
			imposters = preferImposter
			imposters = append(imposters, noPreferences...)
			numAssignedImposter := len(imposters)
			rand.Seed(time.Now().UnixNano())
			rand.Shuffle(len(preferCrew), func(i, j int) {
				preferCrew[i], preferCrew[j] = preferCrew[j], preferCrew[i]
			})
			for _, player := range preferCrew {
				if numAssignedImposter == numImposters {
					crew = append(crew, player)
				} else {
					imposters = append(imposters, player)
					numAssignedImposter += 1
					assignmentScore += 1
				}
			}
		}
		sort.Ints(crew)
		sort.Ints(imposters)
		// cannot have a case where preferences for both are larger than number of roles
		return crew, imposters, assignmentScore
	} else {
		return nil, nil, -1
	}
}

func main() {
	var preferencesToExpectedScore = map[string]int{
		"CCCCC":      1,
		"NCCCC":      0,
		"CCICC":      0,
		"IIIII":      4,
		"IIINI":      3,
		"IICCC":      1,
		"IINNN":      1,
		"CCCCCCCCCC": 2,
		"CCCIICCCCC": 0,
		"ICCICCCCCI": 1,
		"IIIIIIIIII": 8,
		"IIIIIIIINI": 7,
		"NNNNNNNNNN": 0,
	}
	tests := make([]string, 0)
	for preference, _ := range preferencesToExpectedScore {
		tests = append(tests, preference)
	}
	sort.Strings(tests)
	for _, preference := range tests {
		score := preferencesToExpectedScore[preference]
		crew, imposters, assignmentScore := assign_users(preference)
		if assignmentScore != score {
			fmt.Print(
				fmt.Sprintln("Failed test ", preference, ", got assignments for crew: ", crew, "\t imposters: ", imposters),
			)
		} else {
			fmt.Print(
				fmt.Sprintln("PASSED test!"),
			)
		}
	}
}
