package catav15

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
