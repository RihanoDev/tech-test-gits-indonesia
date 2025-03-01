package main

import "fmt"

func calculateRanks(leaderboard []int, playerScores []int) []int {
	filteredScores := []int{}
	for _, val := range leaderboard {
		if len(filteredScores) == 0 || filteredScores[len(filteredScores)-1] != val {
			filteredScores = append(filteredScores, val)
		}
	}

	playerRanks := make([]int, len(playerScores))

	for idx, pScore := range playerScores {
		pos := 1
		for _, lScore := range filteredScores {
			if pScore >= lScore {
				break
			}
			pos++
		}
		playerRanks[idx] = pos
	}

	return playerRanks
}

func main() {
	leaderboardScores := []int{100, 100, 50, 40, 40, 20, 10}
	gitsScores := []int{5, 25, 50, 120}

	finalRanks := calculateRanks(leaderboardScores, gitsScores)

	for _, rank := range finalRanks {
		fmt.Print(rank, " ")
	}
}
