package main

type Config struct {
	Repack string         `toml:"repack"`
	Manips []Manipulation `toml:"manipulation"`
}
