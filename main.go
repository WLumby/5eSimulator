package main

import "fmt"

var (
	playerCharacter Character
	enemy           Character
)

const (
	iterations = 10000
)

func main() {
	var err error

	playerCharacter, err = getCharacter("./sheets/test_paladin.json")
	if err != nil {
		fmt.Println(err)
	}

	enemy, err = getCharacter("./sheets/test_enemy.json")
	if err != nil {
		fmt.Println(err)
	}



	for i := 0; i < iterations; i++ {
		endCombat := false
		winner := Character{}

		for !endCombat {
			winner, endCombat = attackRound()
		}

		if playerCharacter.Name == winner.Name {
			playerCharacter.WinCount++
		} else if enemy.Name == winner.Name {
			enemy.WinCount++
		}
	}

	if playerCharacter.WinCount > enemy.WinCount {
		fmt.Printf("Likely winner: %v\n", playerCharacter.Name)
		fmt.Printf("Winrate: %v percent\n", float64(playerCharacter.WinCount)/float64(iterations)*100)
	} else if playerCharacter.WinCount < enemy.WinCount {
		fmt.Printf("Likely winner: %v\n", enemy.Name)
		fmt.Printf("Winrate: %v percent\n", float64(enemy.WinCount)/float64(iterations)*100)

	}
}

func attackRound() (Character, bool) {

	// PLAYER ATTACK
	attackRoll, attackDamage, _, _ := rollWeapon(playerCharacter)

	if attackRoll >= int(enemy.ArmourClass) {
		enemy.Health = int32(int(enemy.Health) - attackDamage)
		if enemy.Health <= 0 {
			return playerCharacter, true
		}
	}

	// ENEMY ATTACK
	attackRoll, attackDamage, _, _ = rollWeapon(enemy)

	if attackRoll >= int(playerCharacter.ArmourClass) {
		playerCharacter.Health = int32(int(playerCharacter.Health) - attackDamage)
		if playerCharacter.Health <= 0 {
			return enemy, true
		}
	}

	return Character{}, false
}
