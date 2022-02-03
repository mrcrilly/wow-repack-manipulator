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
			err := manipulator(db, m.CreatureID, p.Key, p.Value)

			if err != nil {
				ErrorLogger.Fatal(err)
			} else {
				InfoLogger.Printf("updated created %d: '%s' = '%v'", m.CreatureID, p.Key, p.Value)
			}

		}
	}
}
