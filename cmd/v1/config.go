package main

type Config struct {
	DatabaseHostname string `toml:"database_hostname"`
	DatabaseUsername string `toml:"database_username"`
	DatabasePassword string `toml:"database_password"`
	DatabaseName     string `toml:"database_name"`

	Repack string         `toml:"repack"`
	Manips []Manipulation `toml:"manipulation"`
}
