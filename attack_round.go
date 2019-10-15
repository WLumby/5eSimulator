package main

func attackRound(playerC, enemyC *Character) (Character, bool) {

	// PLAYER ATTACK
	attackRoll, attackDamage := handleClass(*playerC, *enemyC, playerC.Class)

	if attackRoll >= enemyC.ArmourClass {
		enemyC.Health = enemyC.Health - attackDamage
		playerDPT = append(playerDPT, float64(attackDamage))
		if enemyC.Health <= 0 {
			return *playerC, true
		}
	} else {
		playerDPT = append(playerDPT, float64(0))
	}

	// ENEMY ATTACK
	attackRoll, attackDamage = handleClass(*enemyC, *playerC, enemyC.Class)

	if attackRoll >= playerC.ArmourClass {
		playerC.Health = playerC.Health - attackDamage
		enemyDPT = append(enemyDPT, float64(attackDamage))
		if playerC.Health <= 0 {
			return enemy, true
		}
	} else {
		enemyDPT = append(enemyDPT, float64(0))
	}

	return Character{}, false
}

func handleClass(self, target Character, class string) (int, int) {
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

func handlePaladin(self, target Character) (int, int) {
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