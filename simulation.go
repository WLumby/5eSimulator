package main

import (
	"fmt"
	"time"
)

var (
	playerCharacter Character
	enemy           Character
	playerWinCount  int
	enemyWinCount   int
	playerDPT       []float64
	enemyDPT        []float64
)

func runSim(playerPath, enemyPath string, iterations int) {
	var err error

	playerCharacter, err = getCharacter(playerPath)
	if err != nil {
		fmt.Println(err)
	}

	enemy, err = getCharacter(enemyPath)
	if err != nil {
		fmt.Println(err)
	}

	start := time.Now()
	for i := 0; i < iterations; i++ {
		simEncounter()
	}

	playerAverageDPT := calcAverage(playerDPT)
	enemyAverageDPT := calcAverage(enemyDPT)

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Finished after %v!\n\nResults:\n---------------------\n", elapsed)


	fmt.Printf("Player Win Count: %v\nEnemy Win Count: %v\n", playerWinCount, enemyWinCount)

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

func simEncounter() {
	iterablePlayer := playerCharacter
	iterableEnemy := enemy

	endCombat := false
	winner := Character{}

	for !endCombat {
		winner, endCombat = attackRound(&iterablePlayer, &iterableEnemy)
	}

	if playerCharacter.Name == winner.Name {
		playerWinCount++
	} else if enemy.Name == winner.Name {
		enemyWinCount++
	}
}

func calcAverage(dataSet []float64) float64 {
	var total float64 = 0
	for _, value := range dataSet {
		total += value
	}
	return total / float64(len(dataSet))
}
