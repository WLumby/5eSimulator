package main

type Character struct {
	Name        string     `json:"name"`
	Race        string     `json:"race"`
	Class       string     `json:"class"`
	Level       float64    `json:"level"`
	Attributes  Attributes `json:"attributes"`
	ArmourClass int16      `json:"armourClass"`
	Health      int32      `json:"health"`
	Weapon      Weapon     `json:"weapon"`
	WinCount    int
}

type Attributes struct {
	STR int16 `json:"STR"`
	DEX int16 `json:"DEX"`
	CON int16 `json:"CON"`
	INT int16 `json:"INT"`
	WIS int16 `json:"WIS"`
	CHA int16 `json:"CHA"`
}

type Weapon struct {
	Name       string `json:"name"`
	Attribute  string `json:"attribute"`
	DamageType string `json:"damageType"`
	DamageRoll string `json:"damageRoll"`
}
