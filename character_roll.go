package main

import (
	"math"
)

func rollWeaponDamage(character Character) int {
	damageRoll := roll(character.Weapon.DamageRoll)
	//fmt.Printf("damageRoll: %v\n", damageRoll)
	return damageRoll + calculateAttributeMod(character, character.Weapon.Attribute)

}

func calculateAttributeMod(character Character, attribute string) int {
	attributeValue := getAttribute(character, attribute)
	//fmt.Printf("attributeValue: %v\n", attributeValue)

	damageMod := int(math.Floor(float64(attributeValue-10) / 2))
	//fmt.Printf("damageMod: %v\n", damageMod)

	return damageMod
}

func getAttribute(character Character, attribute string) int {
	switch attribute {
	case "STR":
		return int(character.Attributes.STR)
	case "DEX":
		return int(character.Attributes.DEX)
	case "CON":
		return int(character.Attributes.CON)
	case "INT":
		return int(character.Attributes.INT)
	case "WIS":
		return int(character.Attributes.WIS)
	case "CHA":
		return int(character.Attributes.CHA)
	}

	return -1
}
