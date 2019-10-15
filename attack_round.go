package main

func attackRound() (Character, bool) {

	// PLAYER ATTACK
	attackRoll, attackDamage := handleClass(playerCharacter, enemy, playerCharacter.Class)

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
	attackRoll, attackDamage = handleClass(enemy, playerCharacter, enemy.Class)

	if attackRoll >= playerCharacter.ArmourClass {
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

func handleClass(self Character, target Character, class string) (int, int) {
	attackRoll := 0
	attackDamage := 0

	switch class {
	case "Paladin":
		attackRoll, attackDamage = handlePaladin(self, target)
	default:
		attackRoll, attackDamage, _, _ = rollWeapon(self)
	}

	return attackRoll, attackDamage
}

func handlePaladin(self Character, target Character) (int, int) {
	attackRoll := 0
	attackDamage := 0

	decideAttack, _ := decideAttackPaladin(self)
	if decideAttack == "divineSmite" {
		attackRoll, attackDamage, _, _ = rollWeapon(self)
		if attackRoll >= target.ArmourClass {
			self.SpellSlots.FirstLevel--
		}
		attackDamage += castDivineSmite(self)

	}
	if decideAttack == "attack" {
		attackRoll, attackDamage, _, _ = rollWeapon(self)
	}

	return attackRoll, attackDamage
}