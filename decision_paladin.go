package main

func decideAttackPaladin(character Character) (string, int) {
	if character.SpellSlots.FirstLevel > 0 {
		return "divineSmite", 0
	}

	return "attack", -1
}
