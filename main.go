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
			
		}

		fmt.Printf("Health: ")
		for i := 0; i < int(playerCharacter.Health); i++ {
			fmt.Printf("-")
		}
		fmt.Println("\n")

		attackRoll, attackDamage, damageType, crit = rollWeapon(enemy)
		if crit {
			fmt.Printf("%v --- Attack Roll: %v (CRIT!), Damage: %v %v\n", enemy.Name, attackRoll, attackDamage, damageType)
		} else {
			fmt.Printf("%v --- Attack Roll: %v, Damage: %v %v\n", enemy.Name, attackRoll, attackDamage, damageType)
		}

		fmt.Printf("Health: ")
		for i := 0; i < int(enemy.Health); i++ {
			fmt.Printf("-")
		}
		fmt.Println("\n")

		dead = true
	}
}
