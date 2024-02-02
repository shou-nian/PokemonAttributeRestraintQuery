package config

import "github.com/pokemon/AttributeRestraintQuery/app/model"

func BugAttributeRestraint() model.ResultType {
	restraint := make(map[float32][]string)

	restraint[0.5] = []string{Fighting, Grass, Ground}
	restraint[2] = []string{Fire, Rock, Flying}

	return map[string]map[float32][]string{Bug: restraint}
}

func DarkAttributeRestraint() model.ResultType {
	restraint := make(map[float32][]string)

	restraint[0] = []string{Psychic}
	restraint[0.5] = []string{Ghost, Dark}
	restraint[2] = []string{Fighting, Bug, Fairy}

	return map[string]map[float32][]string{Dark: restraint}
}

func DragonAttributeRestraint() model.ResultType {
	restraint := make(map[float32][]string)

	restraint[0.5] = []string{Fire, Water, Grass, Electric}
	restraint[2] = []string{Ice, Dragon, Fairy}

	return map[string]map[float32][]string{Dragon: restraint}
}

func ElectricAttributeRestraint() model.ResultType {
	restraint := make(map[float32][]string)

	restraint[0.5] = []string{Flying, Steel, Electric}
	restraint[2] = []string{Ground}

	return map[string]map[float32][]string{Dragon: restraint}
}

func FairyAttributeRestraint() model.ResultType {
	restraint := make(map[float32][]string)

	restraint[0] = []string{Dragon}
	restraint[0.5] = []string{Bug, Dragon, Fighting}
	restraint[2] = []string{Poison, Steel}

	return map[string]map[float32][]string{Fairy: restraint}
}

func FightingAttributeRestraint() model.ResultType {
	restraint := make(map[float32][]string)

	restraint[0.5] = []string{Dark, Bug, Rock}
	restraint[2] = []string{Flying, Psychic, Fairy}

	return map[string]map[float32][]string{Fairy: restraint}
}

func FireAttributeRestraint() model.ResultType {
	restraint := make(map[float32][]string)

	restraint[0.5] = []string{Bug, Ice, Steel, Fire, Grass, Fairy}
	restraint[2] = []string{Water, Rock, Ground}

	return map[string]map[float32][]string{Fairy: restraint}
}

func FlyingAttributeRestraint() model.ResultType {
	restraint := make(map[float32][]string)

	restraint[0] = []string{Ground}
	restraint[0.5] = []string{Bug, Grass, Fighting}
	restraint[2] = []string{Rock, Electric, Ice}

	return map[string]map[float32][]string{Fairy: restraint}
}

func GhostAttributeRestraint() model.ResultType {
	restraint := make(map[float32][]string)

	restraint[0] = []string{Normal, Fighting}
	restraint[0.5] = []string{Bug, Poison}
	restraint[2] = []string{Ghost, Dark}

	return map[string]map[float32][]string{Fairy: restraint}
}
