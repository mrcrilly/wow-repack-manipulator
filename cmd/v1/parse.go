package main

import (
	"tcdm/repacks/catav15"
)

func parseCataV15Repack(manips []Manipulation) {
	for _, m := range manips {
		var manipulator Manipulator

		switch m.Type {
		case "creature_template":
			manipulator = catav15.ManipulateCreatureTemplate
		case "creature":
			manipulator = catav15.ManipulateCreature
		}

		for _, p := range m.Pairs {
			err := manipulator(db, m.Column, m.UniqueIDs, p.Key, p.Value)

			if err != nil {
				ErrorLogger.Fatal(err)
			} else {
				InfoLogger.Printf("%s: updated id '%v': '%s' = '%v'", m.Type, m.UniqueIDs, p.Key, p.Value)
			}

		}
	}
}
