package main

import (
	"math"
)

func rollWeapon(character Character) (int, int, string, bool) {

	attackRoll, crit := rollAttack(character)
	weaponDamage := rollWeaponDamage(character, crit)
	damageType := character.Weapon.DamageType

	return attackRoll, weaponDamage, damageType, crit
}

func rollCastSpell(character Character, spell Spell) (int, int, string, bool) {
	attackRoll, crit := rollSpell(character, spell)
	spellDamage := roll(spell.DamageRoll)
	spellDamageType := spell.DamageType

	return attackRoll, spellDamage, spellDamageType, crit
}

func rollWeaponDamage(character Character, crit bool) int {
	critDamageRoll := 0

	damageRoll := roll(character.Weapon.DamageRoll)
	if crit {
		critDamageRoll = roll(character.Weapon.DamageRoll)
	}
	return damageRoll + critDamageRoll + calculateAttributeMod(character, character.Weapon.Attribute)
}

func rollAttack(character Character) (int, bool) {
	var crit bool

	d20Roll := roll("1d20")

	if d20Roll == 20 {
		crit = true
	}

	damageMod := calculateAttributeMod(character, character.Weapon.Attribute)
	proficiency := calcProficiency(character)
	return d20Roll + damageMod + proficiency, crit
}

func rollSpell(character Character, spell Spell) (int, bool) {
	var crit bool

	d20Roll := roll("1d20")

	if d20Roll == 20 {
		crit = true
	}

	damageMod := calculateAttributeMod(character, spell.Attribute)
	proficiency := calcProficiency(character)
	return d20Roll + damageMod + proficiency, crit
}

func calcProficiency(character Character) int {
	// TODO: Add proficiency map
	return 2
}

func calculateAttributeMod(character Character, attribute string) int {
	attributeValue := getAttribute(character, attribute)
	damageMod := int(math.Floor(float64(attributeValue-10) / 2))

	return damageMod
}

func getAttribute(character Character, attribute string) int {
	switch attribute {
	case "STR":
		return character.Attributes.STR
	case "DEX":
		return character.Attributes.DEX
	case "CON":
		return character.Attributes.CON
	case "INT":
		return character.Attributes.INT
	case "WIS":
		return character.Attributes.WIS
	case "CHA":
		return character.Attributes.CHA
	}

	return -1
}
