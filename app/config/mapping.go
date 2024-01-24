package config

type funcMap func() map[string][]string

var FuncMapping = map[string]funcMap{
	Bug: BugAttributeRestraint,
}
