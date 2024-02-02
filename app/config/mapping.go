package config

import "github.com/pokemon/AttributeRestraintQuery/app/model"

type FuncMap func() model.ResultType

var FuncMapping = map[string]FuncMap{
	Bug:      BugAttributeRestraint,
	Dark:     DarkAttributeRestraint,
	Dragon:   DragonAttributeRestraint,
	Electric: ElectricAttributeRestraint,
	Fairy:    FairyAttributeRestraint,
	Fighting: FightingAttributeRestraint,
	Fire:     FireAttributeRestraint,
	Flying:   FlyingAttributeRestraint,
	Ghost:    GhostAttributeRestraint,
}
