package main

import (
	"fmt"
	"github.com/briandowns/spinner"
	"sync"
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

	s := spinner.New(spinner.CharSets[36], time.Duration(iterations/50)*time.Millisecond)
	s.Start()

	var wg sync.WaitGroup
	wg.Add(iterations)
	start := time.Now()
	for i := 0; i < iterations; i++ {
		go simEncounter(&wg)
	}

	wg.Wait()
	s.Stop()

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Finished after %v!\n\nResults:\n---------------------\n", elapsed)

	playerAverageDPT := calcAverage(playerDPT)
	enemyAverageDPT := calcAverage(enemyDPT)

	fmt.Printf("Player Win Count: %v\nEnemy Win Count: %v\n", playerWinCount, enemyWinCount)

	if playerWinCount >= enemyWinCount {
		fmt.Printf("Likely winner: %v\n", playerCharacter.Name)
		fmt.Printf("Winrate: %.2f%%\n", float64(playerWinCount)/float64(iterations)*100)
	} else if playerWinCount < enemyWinCount {
		fmt.Printf("Likely winner: %v\n", enemy.Name)
		fmt.Printf("Winrate: %.2f percent\n", float64(enemyWinCount)/float64(iterations)*100)
	}

	fmt.Printf("Player Avg DPT: %.2f (%.2f%%)\n", playerAverageDPT, (playerAverageDPT/float64(enemy.Health))*100)
	fmt.Printf("Enemy Avg DPT: %.2f (%.2f%%)\n", enemyAverageDPT, (enemyAverageDPT/float64(playerCharacter.Health))*100)
}

func simEncounter(wg *sync.WaitGroup) {
	iterablePlayer := playerCharacter
	iterableEnemy := enemy

	endCombat := false
	winner := Character{}

	firstAttack := calcFirstAttack(iterablePlayer, iterableEnemy)

	for !endCombat {
		winner, endCombat = attackRound(&iterablePlayer, &iterableEnemy, firstAttack)
	}

	if playerCharacter.Name == winner.Name {
		playerWinCount++
	} else if enemy.Name == winner.Name {
		enemyWinCount++
	}

	wg.Done()
}

func calcFirstAttack(iterablePlayer, iterableEnemy Character) string {
	playerInitiative := rollInitiative(iterablePlayer)
	enemyInitiative := rollInitiative(iterableEnemy)

	if playerInitiative >= enemyInitiative {
		return "player"
	} else {
		return "enemy"
	}
}

func calcAverage(dataSet []float64) float64 {
	var total float64 = 0
	for _, value := range dataSet {
		total += value
	}
	return total / float64(len(dataSet))
}
