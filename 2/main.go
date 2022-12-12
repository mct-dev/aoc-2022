package main

import (
	"fmt"
	"os"
	"strings"
)

const L_ROCK = "A"
const L_PAPER = "B"
const L_SCISSORS = "C"

const R_ROCK = "X"
const R_PAPER = "Y"
const R_SCISSORS = "Z"

const ROUND_LOSE = "X"
const ROUND_DRAW = "Y"
const ROUND_WIN = "Z"

func getShapeScore(shape string) int {
	switch shape {
	case L_ROCK:
		return 1
	case L_PAPER:
		return 2
	case L_SCISSORS:
		return 3
	case R_ROCK:
		return 1
	case R_PAPER:
		return 2
	case R_SCISSORS:
		return 3
	}

	return 0
}

func isWinningPlay(mine int, theirs int) bool {
	if mine == theirs {
		return false
	}

	switch mine {
	case getShapeScore(L_ROCK):
		return theirs == getShapeScore(R_SCISSORS)
	case getShapeScore(L_PAPER):
		return theirs == getShapeScore(R_ROCK)
	case getShapeScore(L_SCISSORS):
		return theirs == getShapeScore(R_PAPER)
	}

	return false
}

func isDraw(mine int, theirs int) bool {
	return mine == theirs
}

func getInputLines() []string {
	bytes, err := os.ReadFile("2/input.txt")
	if err != nil {
		panic(err)
	}

	st := string(bytes)

	return strings.Split(st, "\n")
}

func part1() {
	lines := getInputLines()
	winCount := 0
	totalPoints := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		moves := strings.Split(line, " ")
		lMove := moves[0]
		rMove := moves[1]

		lScore := getShapeScore(lMove)
		rScore := getShapeScore(rMove)

		isDraw := isDraw(rScore, lScore)
		isWinningPlay := isWinningPlay(rScore, lScore)

		roundScore := 0

		if isDraw {
			roundScore = 3 + rScore
		} else if isWinningPlay {
			roundScore = 6 + rScore
			winCount++
		} else {
			roundScore = 0 + rScore
		}

		totalPoints += roundScore
	}

	fmt.Println("win count:", winCount)
	fmt.Println("total points:", totalPoints)
}

func getRequiredMove(theirMove string, requiredRoundEnding string) string {
	if requiredRoundEnding == ROUND_DRAW {
		return theirMove
	}

	switch theirMove {
	case L_ROCK:
		if requiredRoundEnding == ROUND_WIN {
			return R_PAPER
		} else if requiredRoundEnding == ROUND_LOSE {
			return R_SCISSORS
		}
	case L_PAPER:
		if requiredRoundEnding == ROUND_WIN {
			return R_SCISSORS
		} else if requiredRoundEnding == ROUND_LOSE {
			return R_ROCK
		}
	case L_SCISSORS:
		if requiredRoundEnding == ROUND_WIN {
			return R_ROCK
		} else if requiredRoundEnding == ROUND_LOSE {
			return R_PAPER
		}
	}

	return theirMove
}

func part2() {
	lines := getInputLines()
	winCount := 0
	totalPoints := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		moves := strings.Split(line, " ")
		theirMove := moves[0]
		requiredRoundEnding := moves[1]

		theirScore := getShapeScore(theirMove)

		myMove := getRequiredMove(theirMove, requiredRoundEnding)
		myScore := getShapeScore(myMove)

		isDraw := isDraw(myScore, theirScore)
		isWin := isWinningPlay(myScore, theirScore)

		roundScore := 0
		if isDraw {
			roundScore = 3 + myScore
		} else if isWin {
			roundScore = 6 + myScore
			winCount++
		} else {
			roundScore = 0 + myScore
		}

		totalPoints += roundScore
	}

	fmt.Println("win count:", winCount)
	fmt.Println("total points:", totalPoints)
}

func main() {
	part2()
}
