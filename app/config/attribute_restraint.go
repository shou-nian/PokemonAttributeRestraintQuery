package config

func BugAttributeRestraint() map[float32][]string {
	restraint := make(map[float32][]string)

	restraint[0.5] = []string{Fighting, Grass, Ground}
	restraint[2] = []string{Fire, Rock, Flying}

	return restraint
}

func DarkAttributeRestraint() map[float32][]string {
	restraint := make(map[float32][]string)

	restraint[0] = []string{Psychic}
	restraint[0.5] = []string{Ghost, Dark}
	restraint[2] = []string{Fighting, Bug, Fairy}

	return restraint
}

func DragonAttributeRestraint() map[float32][]string {
	restraint := make(map[float32][]string)

	restraint[0.5] = []string{Fire, Water, Grass, Electric}
	restraint[2] = []string{Ice, Dragon, Fairy}

	return restraint
}

func ElectricAttributeRestraint() map[float32][]string {
	restraint := make(map[float32][]string)

	restraint[0.5] = []string{Flying, Steel, Electric}
	restraint[2] = []string{Ground}

	return restraint
}

func FairyAttributeRestraint() map[float32][]string {
	restraint := make(map[float32][]string)

	restraint[0] = []string{Dragon}
	restraint[0.5] = []string{Bug, Dragon, Fighting}
	restraint[2] = []string{Poison, Steel}

	return restraint
}
