package catav15

import (
	"fmt"

	"gorm.io/gorm"
)

type CreatureTemplate struct {
	Entry              int     `gorm:"column:entry"`
	DifficultyEntry1   int     `gorm:"column:difficulty_entry_1"`
	DifficultyEntry2   int     `gorm:"column:difficulty_entry_2"`
	DifficultyEntry3   int     `gorm:"column:difficulty_entry_3"`
	KillCredit1        int     `gorm:"column:KillCredit1"`
	KillCredit2        int     `gorm:"column:KillCredit2"`
	Modelid1           int     `gorm:"column:modelid1"`
	Modelid2           int     `gorm:"column:modelid2"`
	Modelid3           int     `gorm:"column:modelid3"`
	Modelid4           int     `gorm:"column:modelid4"`
	Name               string  `gorm:"column:name"`
	Subname            string  `gorm:"column:subname"`
	IconName           string  `gorm:"column:IconName"`
	GossipMenuId       int     `gorm:"column:gossip_menu_id"`
	Minlevel           int     `gorm:"column:minlevel"`
	Maxlevel           int     `gorm:"column:maxlevel"`
	Exp                int     `gorm:"column:exp"`
	ExpUnk             int     `gorm:"column:exp_unk"`
	FactionA           int     `gorm:"column:faction_a"`
	FactionH           int     `gorm:"column:faction_h"`
	Npcflag            int     `gorm:"column:npcflag"`
	SpeedWalk          float64 `gorm:"column:speed_walk"`
	SpeedRun           float64 `gorm:"column:speed_run"`
	SpeedSwim          float64 `gorm:"column:speed_swim"`
	SpeedFly           float64 `gorm:"column:speed_fly"`
	Scale              float64 `gorm:"column:scale"`
	Rank               int     `gorm:"column:rank"`
	Mindmg             float64 `gorm:"column:mindmg"`
	Maxdmg             float64 `gorm:"column:maxdmg"`
	Dmgschool          int     `gorm:"column:dmgschool"`
	Attackpower        int     `gorm:"column:attackpower"`
	DmgMultiplier      float64 `gorm:"column:dmg_multiplier"`
	Baseattacktime     int     `gorm:"column:baseattacktime"`
	Rangeattacktime    int     `gorm:"column:rangeattacktime"`
	UnitClass          int     `gorm:"column:unit_class"`
	UnitFlags          int     `gorm:"column:unit_flags"`
	UnitFlags2         int     `gorm:"column:unit_flags2"`
	Dynamicflags       int     `gorm:"column:dynamicflags"`
	Family             int     `gorm:"column:family"`
	TrainerType        int     `gorm:"column:trainer_type"`
	TrainerSpell       int     `gorm:"column:trainer_spell"`
	TrainerClass       int     `gorm:"column:trainer_class"`
	TrainerRace        int     `gorm:"column:trainer_race"`
	Minrangedmg        float64 `gorm:"column:minrangedmg"`
	Maxrangedmg        float64 `gorm:"column:maxrangedmg"`
	Rangedattackpower  int     `gorm:"column:rangedattackpower"`
	Type               int     `gorm:"column:type"`
	TypeFlags          int     `gorm:"column:type_flags"`
	TypeFlags2         int     `gorm:"column:type_flags2"`
	Lootid             int     `gorm:"column:lootid"`
	Pickpocketloot     int     `gorm:"column:pickpocketloot"`
	Skinloot           int     `gorm:"column:skinloot"`
	Resistance1        int     `gorm:"column:resistance1"`
	Resistance2        int     `gorm:"column:resistance2"`
	Resistance3        int     `gorm:"column:resistance3"`
	Resistance4        int     `gorm:"column:resistance4"`
	Resistance5        int     `gorm:"column:resistance5"`
	Resistance6        int     `gorm:"column:resistance6"`
	Spell1             int     `gorm:"column:spell1"`
	Spell2             int     `gorm:"column:spell2"`
	Spell3             int     `gorm:"column:spell3"`
	Spell4             int     `gorm:"column:spell4"`
	Spell5             int     `gorm:"column:spell5"`
	Spell6             int     `gorm:"column:spell6"`
	Spell7             int     `gorm:"column:spell7"`
	Spell8             int     `gorm:"column:spell8"`
	PetSpellDataId     int     `gorm:"column:PetSpellDataId"`
	VehicleId          int     `gorm:"column:VehicleId"`
	Mingold            int     `gorm:"column:mingold"`
	Maxgold            int     `gorm:"column:maxgold"`
	AIName             string  `gorm:"column:AIName"`
	MovementType       int     `gorm:"column:MovementType"`
	InhabitType        int     `gorm:"column:InhabitType"`
	HoverHeight        float64 `gorm:"column:HoverHeight"`
	HealthMod          float64 `gorm:"column:Health_mod"`
	ManaMod            float64 `gorm:"column:Mana_mod"`
	ManaModExtra       float64 `gorm:"column:Mana_mod_extra"`
	ArmorMod           float64 `gorm:"column:Armor_mod"`
	RacialLeader       int     `gorm:"column:RacialLeader"`
	QuestItem1         int     `gorm:"column:questItem1"`
	QuestItem2         int     `gorm:"column:questItem2"`
	QuestItem3         int     `gorm:"column:questItem3"`
	QuestItem4         int     `gorm:"column:questItem4"`
	QuestItem5         int     `gorm:"column:questItem5"`
	QuestItem6         int     `gorm:"column:questItem6"`
	MovementId         int     `gorm:"column:movementId"`
	RegenHealth        int     `gorm:"column:RegenHealth"`
	EquipmentId        int     `gorm:"column:equipment_id"`
	MechanicImmuneMask int     `gorm:"column:mechanic_immune_mask"`
	FlagsExtra         int     `gorm:"column:flags_extra"`
	ScriptName         string  `gorm:"column:ScriptName"`
	WDBVerified        int     `gorm:"column:WDBVerified"`
	ReactState         int     `gorm:"column:ReactState"`
}

func (CreatureTemplate) TableName() string {
	return "creature_template"
}

type GenericCreatureTemplateManipulator struct {
	Column string
	ID     []int
	Key    string
	Value  interface{}
}

func (m *GenericCreatureTemplateManipulator) SetFlag(name string, value interface{}) error {
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

func (m GenericCreatureTemplateManipulator) Execute(db *gorm.DB) error {
	model := CreatureTemplate{}
	db.Model(&model).Where(fmt.Sprintf("%s = ?", m.Column), m.ID).UpdateColumn(m.Key, m.Value)
	return nil
}
