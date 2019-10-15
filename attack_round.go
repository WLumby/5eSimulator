package main

func attackRound() (Character, bool) {

	// PLAYER ATTACK
	attackRoll, attackDamage, _, _ := rollWeapon(playerCharacter)

	if attackRoll >= enemy.ArmourClass {
		enemy.Health = enemy.Health - attackDamage
		playerDPT = append(playerDPT, float64(attackDamage))
		if enemy.Health <= 0 {
			return playerCharacter, true
		}
	} else {
		playerDPT = append(playerDPT, float64(0))
	}

	// ENEMY ATTACK
	attackRoll, attackDamage, _, _ = rollWeapon(enemy)

	if attackRoll >= int(playerCharacter.ArmourClass) {
		playerCharacter.Health = playerCharacter.Health - attackDamage
		enemyDPT = append(enemyDPT, float64(attackDamage))
		if playerCharacter.Health <= 0 {
			return enemy, true
		}
	} else {
		enemyDPT = append(enemyDPT, float64(0))
	}

	return Character{}, false
}