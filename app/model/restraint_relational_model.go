package model

type RestraintRelational map[string][]string

func NewAttribute(attr string) *RestraintRelational {
	return &RestraintRelational{attr: {}}
}
