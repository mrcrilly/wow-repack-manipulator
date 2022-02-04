package catav15

import (
	"fmt"

	"gorm.io/gorm"
)

type CreatureByZoneAndAreaManipulator struct {
	ZoneID int
	AreaID int

	// Generic part:
	Key   string
	Value interface{}
}

func (m *CreatureByZoneAndAreaManipulator) SetFlag(name string, value interface{}) error {
	switch name {
	case "zoneid":
		m.ZoneID = value.(int)
	case "areaid":
		m.AreaID = value.(int)
	case "key":
		m.Key = value.(string)
	case "value":
		m.Value = value
	default:
		return fmt.Errorf("unknown flag given: %s", name)
	}

	return nil
}

func (m CreatureByZoneAndAreaManipulator) Execute(db *gorm.DB) error {
	var creatures []Creature
	db.Where("zone = ?", m.ZoneID).Where("area = ?", m.AreaID).Find(&creatures)

	if len(creatures) == 0 {
		return fmt.Errorf("no creatures found in zone %d with area %d", m.ZoneID, m.AreaID)
	}

	// Cannot just pass a []int to GORM here as the list can be too large
	// resulting in operahand error
	for _, c := range creatures {
		db.Model(&CreatureTemplate{}).Where("entry = ?", c.Id).UpdateColumn(m.Key, m.Value)
	}

	return nil
}

type RemoveNPCByGUIDManipulator struct {
	GUID []int
}

func (m *RemoveNPCByGUIDManipulator) SetFlag(name string, value interface{}) error {
	switch name {
	case "guid":
		m.GUID = value.([]int)
	default:
		return fmt.Errorf("unknown flag given: %s", name)
	}

	return nil
}

func (m RemoveNPCByGUIDManipulator) Execute(db *gorm.DB) error {
	var creatures []Creature

	// Cannot use built in []int feature of GORM as it causes problems
	// with MySQL
	for _, guid := range m.GUID {
		var c Creature
		db.Where("guid = ?", guid).Find(&c)
		creatures = append(creatures, c)
	}

	if len(creatures) == 0 {
		return fmt.Errorf("no creatures found for given guid(s): %v", m.GUID)
	}

	// Cannot just pass a []int to GORM here as the list can be too large
	// resulting in operahand error
	for _, c := range creatures {
		db.Where("guid = ?", c.Guid).Delete(&c)
	}

	return nil
}
