package config

func BugAttributeRestraint() map[string][]string {
	restraint := make(map[string][]string)

	restraint["0.5"] = []string{Fighting, Grass, Ground}
	restraint["2"] = []string{Fire, Rock, Flying}

	return restraint
}
