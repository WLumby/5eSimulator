package main

import "fmt"

func main() {
	playerCharacter, err := getCharacter("./sheets/test_paladin.json")
	if err != nil {
		fmt.Println(err)
	}
	enemy, err := getCharacter("./sheets/test_enemy.json")
	if err != nil {
		fmt.Println(err)
	}

	dead := false

	for !dead {
		attackRoll, attackDamage, damageType, crit := rollWeapon(playerCharacter)
		if crit {
			fmt.Printf("%v --- Attack Roll: %v (CRIT!), Damage: %v %v\n", playerCharacter.Name, attackRoll, attackDamage, damageType)
		} else {
			fmt.Printf("%v --- Attack Roll: %v, Damage: %v %v\n", playerCharacter.Name, attackRoll, attackDamage, damageType)
		}

		if attackRoll >= int(enemy.ArmourClass) {
			fmt.Printf("HIT on %v for %v %v damage!\n\n", enemy.Name, attackDamage, damageType)
			enemy.Health = int32(int(enemy.Health) - attackDamage)
			if enemy.Health <= 0 {
				dead = true
				fmt.Printf("%v dies!", enemy.Name)
				break
			}
		} else {
			fmt.Printf("MISS on %v!\n\n", enemy.Name)
		}

		attackRoll, attackDamage, damageType, crit = rollWeapon(enemy)
		if crit {
			fmt.Printf("%v --- Attack Roll: %v (CRIT!), Damage: %v %v\n", enemy.Name, attackRoll, attackDamage, damageType)
		} else {
			fmt.Printf("%v --- Attack Roll: %v, Damage: %v %v\n", enemy.Name, attackRoll, attackDamage, damageType)
		}

		if attackRoll >= int(playerCharacter.ArmourClass) {
			fmt.Printf("HIT on %v for %v %v damage!\n\n", playerCharacter.Name, attackDamage, damageType)
			playerCharacter.Health = int32(int(playerCharacter.Health) - attackDamage)
			if playerCharacter.Health <= 0 {
				dead = true
				fmt.Printf("%v dies!", playerCharacter.Name)
				break
			}
		} else {
			fmt.Printf("MISS on %v!\n\n", playerCharacter.Name)
		}

		fmt.Printf("\nPlayer Health: ")
		for i := 0; i < int(playerCharacter.Health); i++ {
			fmt.Printf("-")
		}
		fmt.Println()

		fmt.Printf("Enemy  Health: ")
		for i := 0; i < int(enemy.Health); i++ {
			fmt.Printf("-")
		}
		fmt.Println("\n")
	}
}
