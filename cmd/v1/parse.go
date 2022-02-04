package main

import (
	"tcdm/repacks/catav15"
)

func parseCataV15Repack(manips []Manipulation) {
	for _, m := range manips {
		switch m.Type {

		case "creature_template":
			manipulator := &catav15.GenericCreatureTemplateManipulator{}
			manipulator.SetFlag("column", "entry")
			manipulator.SetFlag("id", m.UniqueIDs)

			for _, p := range m.Pairs {
				manipulator.SetFlag("key", p.Key)
				manipulator.SetFlag("value", p.Value)

				err := manipulator.Execute(db)

				if err != nil {
					ErrorLogger.Fatal(err)
				} else {
					InfoLogger.Printf("%s: updated id '%v': '%s' = '%v'", m.Type, m.UniqueIDs, p.Key, p.Value)
				}
			}

		case "scale_creature_in_zone_area":
			manipulator := &catav15.CreatureByZoneAndAreaManipulator{}
			manipulator.SetFlag("zoneid", m.ZoneID)
			manipulator.SetFlag("areaid", m.AreaID)

			for _, p := range m.Pairs {
				manipulator.SetFlag("key", p.Key)
				manipulator.SetFlag("value", p.Value)

				err := manipulator.Execute(db)

				if err != nil {
					ErrorLogger.Fatal(err)
				} else {
					InfoLogger.Printf("%s: updated zone id '%d' and area id '%d': '%s' = '%v'", m.Type, m.ZoneID, m.AreaID, p.Key, p.Value)
				}
			}
		}

	}
}
