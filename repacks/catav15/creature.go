package catav15

import (
	"fmt"

	"gorm.io/gorm"
)

type Creature struct {
	Guid            int
	Id              int
	Map             int
	Zone            int
	Area            int
	SpawnMask       int
	PhaseMask       int
	Modelid         int
	EquipmentId     int
	PositionX       float64
	PositionY       float64
	PositionZ       float64
	Orientation     float64
	Spawntimesecs   int
	Spawndist       float64
	Currentwaypoint int
	Curhealth       int
	Curmana         int
	MovementType    int
	Npcflag         int
	UnitFlags       int
	Dynamicflags    int
	Walkmode        float64
	Saiscriptflag   float64
}

func (Creature) TableName() string {
	return "creature"
}

type GenericCreatureManipulator struct {
	Column string
	ID     []int
	Key    string
	Value  interface{}
}

func (m *GenericCreatureManipulator) SetFlag(name string, value interface{}) error {
	switch name {
	case "column":
		m.Column = value.(string)
	case "id":
		m.ID = value.([]int)
	case "key":
		m.Key = value.(string)
	case "value":
		m.Value = value
	default:
		return fmt.Errorf("unknown flag given: %s", name)
	}

	return nil
}

func (m *GenericCreatureManipulator) Execute(db *gorm.DB) error {
	model := Creature{}
	db.Model(&model).Where(fmt.Sprintf("%s = ?", m.Column), m.ID).UpdateColumn(m.Key, m.Value)
	return nil
}
