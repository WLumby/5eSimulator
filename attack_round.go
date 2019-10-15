package main

import "fmt"

func attackRound() (Character, bool) {

	// PLAYER ATTACK
	attackRoll, attackDamage := handleClass(playerCharacter.Class)

	if attackRoll >= enemy.ArmourClass {
		enemy.Health = enemy.Health - attackDamage
		fmt.Printf("PLAYER HIT -- Damage: %v\n\n", attackDamage)
		playerDPT = append(playerDPT, float64(attackDamage))
		if enemy.Health <= 0 {
			fmt.Printf("Enemy DIES\n\n\n")
			return playerCharacter, true
		}
	} else {
		fmt.Println("PLAYER MISS\n")
		playerDPT = append(playerDPT, float64(0))
	}

	// ENEMY ATTACK
	attackRoll, attackDamage, _, _ = rollWeapon(enemy)

	if attackRoll >= playerCharacter.ArmourClass {
		playerCharacter.Health = playerCharacter.Health - attackDamage
		fmt.Printf("ENEMY HIT -- Damage: %v\n\n", attackDamage)
		enemyDPT = append(enemyDPT, float64(attackDamage))
		if playerCharacter.Health <= 0 {
			fmt.Printf("Player DIES\n\n\n")
			return enemy, true
		}
	} else {
		fmt.Println("ENEMY MISS\n")
		enemyDPT = append(enemyDPT, float64(0))
	}

	return Character{}, false
}

func handleClass(class string) (int, int) {
	attackRoll := 0
	attackDamage := 0

	switch class {
	case "Paladin":
		attackRoll, attackDamage = handlePaladin()
	}

	return attackRoll, attackDamage
}

func handlePaladin() (int, int) {
	attackRoll := 0
	attackDamage := 0

	decideAttack, _ := decideAttackPaladin(playerCharacter)
	if decideAttack == "divineSmite" {
		attackRoll, attackDamage, _, _ = rollWeapon(playerCharacter)
		if attackRoll >= enemy.ArmourClass {
			playerCharacter.SpellSlots.FirstLevel--
		}
		attackDamage += castDivineSmite(playerCharacter)

	}
	if decideAttack == "attack" {
		attackRoll, attackDamage, _, _ = rollWeapon(playerCharacter)
	}

	return attackRoll, attackDamage
}