package model

import (
	"reflect"
	"testing"
)

func TestNewAttribute(t *testing.T) {
	r := NewAttribute("fire")
	exp := &RestraintRelational{"fire": {}}

	if !reflect.DeepEqual(r, exp) {
		t.Errorf("expected %#v, got %#v", exp, r)
	}
}
