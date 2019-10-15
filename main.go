package main

import "fmt"

var (
	playerCharacter Character
	enemy           Character
	playerWinCount  int
	enemyWinCount   int
	playerDPT       []float64
	enemyDPT        []float64
)

const (
	iterations = 10000
)

func main() {
	var err error

	for i := 0; i < iterations; i++ {
		playerCharacter, err = getCharacter("./sheets/test_paladin.json")
		if err != nil {
			fmt.Println(err)
		}

		enemy, err = getCharacter("./sheets/test_enemy.json")
		if err != nil {
			fmt.Println(err)
		}

		endCombat := false
		winner := Character{}

		for !endCombat {
			winner, endCombat = attackRound()
		}

		if playerCharacter.Name == winner.Name {
			playerWinCount++
		} else if enemy.Name == winner.Name {
			enemyWinCount++
		}
	}

	playerAverageDPT := calcAverage(playerDPT)
	enemyAverageDPT := calcAverage(enemyDPT)

	if playerWinCount > enemyWinCount {
		fmt.Printf("Likely winner: %v\n", playerCharacter.Name)
		fmt.Printf("Winrate: %v percent\n", float64(playerWinCount)/float64(iterations)*100)
	} else if playerWinCount < enemyWinCount {
		fmt.Printf("Likely winner: %v\n", enemy.Name)
		fmt.Printf("Winrate: %v percent\n", float64(enemyWinCount)/float64(iterations)*100)
	}

	fmt.Printf("Player Avg DPT: %v\n", playerAverageDPT)
	fmt.Printf("Enemy Avg DPT: %v\n", enemyAverageDPT)
}

func calcAverage(dataSet []float64) float64 {
	var total float64 = 0
	for _, value := range dataSet {
		total += value
	}
	return total / float64(len(dataSet))
}
