package main

import "gorm.io/gorm"

type Manipulator func(*gorm.DB, string, []int, string, interface{}) error

type Manipulation struct {
	Type      string                  `toml:"type"`
	UniqueIDs []int                   `toml:"id"`
	Column    string                  `toml:"column"`
	Pairs     []ManipulationFieldPair `toml:"modifiers"`
}

type ManipulationFieldPair struct {
	Key   string      `toml:"key"`
	Value interface{} `toml:"value"`
}
