package main

func attackRound(playerC, enemyC *Character, firstAttack string) (Character, bool) {

	if firstAttack == "player" {
		if playerAttack(playerC, enemyC) {
			return *playerC, true
		}
		if enemyAttack(playerC, enemyC) {
			return *enemyC, true
		}
	} else if firstAttack == "enemy" {
		if enemyAttack(playerC, enemyC) {
			return *enemyC, true
		}
		if playerAttack(playerC, enemyC) {
			return *playerC, true
		}
	}
	return Character{}, false
}

func playerAttack(playerC, enemyC *Character) bool {
	// PLAYER ATTACK
	attackRoll, attackDamage := handleClass(*playerC, *enemyC, playerC.Class)
	if attackRoll >= enemyC.ArmourClass {
		enemyC.Health = enemyC.Health - attackDamage
		playerDPT = append(playerDPT, float64(attackDamage))
		if enemyC.Health <= 0 {
			return true
		}
	} else {
		playerDPT = append(playerDPT, float64(0))
	}

	return false
}

func enemyAttack(playerC, enemyC *Character) bool {
	// ENEMY ATTACK
	attackRoll, attackDamage := handleClass(*enemyC, *playerC, enemyC.Class)

	if attackRoll >= playerC.ArmourClass {
		playerC.Health = playerC.Health - attackDamage
		enemyDPT = append(enemyDPT, float64(attackDamage))
		if playerC.Health <= 0 {
			return true
		}
	} else {
		enemyDPT = append(enemyDPT, float64(0))
	}

	return false
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
