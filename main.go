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

	attackRoll, attackDamage, damageType, crit := rollWeapon(playerCharacter)
	if crit {
		fmt.Printf("Attack Roll: %v (CRIT!), Damage: %v %v\n", attackRoll, attackDamage, damageType)
	} else {
		fmt.Printf("Attack Roll: %v, Damage: %v %v\n", attackRoll, attackDamage, damageType)
	}

	attackRoll, attackDamage, damageType, crit = rollWeapon(enemy)
	if crit {
		fmt.Printf("Attack Roll: %v (CRIT!), Damage: %v %v\n", attackRoll, attackDamage, damageType)
	} else {
		fmt.Printf("Attack Roll: %v, Damage: %v %v\n", attackRoll, attackDamage, damageType)
	}}
