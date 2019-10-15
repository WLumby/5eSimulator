package main

func castDivineSmite(character Character) int {
	spellDamage := roll(character.Spells[0].DamageRoll)

	return spellDamage
}
