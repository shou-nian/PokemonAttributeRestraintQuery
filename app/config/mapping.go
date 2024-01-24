package config

type funcMap func() map[float32][]string

var FuncMapping = map[string]funcMap{
	Bug:      BugAttributeRestraint,
	Dark:     DarkAttributeRestraint,
	Dragon:   DragonAttributeRestraint,
	Electric: ElectricAttributeRestraint,
	Fairy:    FairyAttributeRestraint,
}
