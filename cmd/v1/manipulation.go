package main

import "gorm.io/gorm"

type Manipulator interface {
	Execute(*gorm.DB) error
	SetFlag(string, interface{}) error
}

type Manipulation struct {
	Type      string                  `toml:"type"`
	ZoneID    int                     `toml:"zone_id"`
	AreaID    int                     `toml:"area_id"`
	UniqueIDs []int                   `toml:"id"`
	Column    string                  `toml:"column"`
	Pairs     []ManipulationFieldPair `toml:"modifiers"`
}

type ManipulationFieldPair struct {
	Key   string      `toml:"key"`
	Value interface{} `toml:"value"`
}
