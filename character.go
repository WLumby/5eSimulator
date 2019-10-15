package main

type Character struct {
	Name        string     `json:"name"`
	Race        string     `json:"race"`
	Class       string     `json:"class"`
	Level       float64    `json:"level"`
	Attributes  Attributes `json:"attributes"`
	ArmourClass int        `json:"armourClass"`
	Health      int        `json:"health"`
	Weapon      Weapon     `json:"weapon"`
	Spells      []Spell    `json:"spells"`
	SpellSlots  SpellSlots `json:"spellSlots"`
}

type Attributes struct {
	STR int `json:"STR"`
	DEX int `json:"DEX"`
	CON int `json:"CON"`
	INT int `json:"INT"`
	WIS int `json:"WIS"`
	CHA int `json:"CHA"`
}

type Weapon struct {
	Name       string `json:"name"`
	Attribute  string `json:"attribute"`
	DamageType string `json:"damageType"`
	DamageRoll string `json:"damageRoll"`
}

type SpellSlots struct {
	FirstLevel int `json:"firstLevel"`
}

type Spell struct {
	Name                    string `json:"name"`
	Attribute               string `json:"attribute"`
	CastType                string `json:"castType"`
	DamageType              string `json:"damageType"`
	DamageRoll              string `json:"damageRoll"`
	HigherCastingDamageRoll string `json:"higherCastingDamageRoll"`
	SpellSlotLevel          int    `json:"spellSlotLevel"`
}
