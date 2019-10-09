package main

import "fmt"

func main() {
	character, err := getCharacter("./sheets/test_paladin.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v %v\n", rollWeaponDamage(character), character.Weapon.DamageType)
}
