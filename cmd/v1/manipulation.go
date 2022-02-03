package main

import "gorm.io/gorm"

type Manipulator func(*gorm.DB, int, string, interface{}) error

type Manipulation struct {
	Type       string                  `toml:"type"`
	CreatureID int                     `toml:"entry"`
	Pairs      []ManipulationFieldPair `toml:"modifiers"`
}

type ManipulationFieldPair struct {
	Key   string      `toml:"key"`
	Value interface{} `toml:"value"`
}
