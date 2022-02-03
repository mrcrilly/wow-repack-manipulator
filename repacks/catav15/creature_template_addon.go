package catav15

type CreatureTemplateAddon struct {
	Entry                  int
	PathId                 int
	Mount                  int
	Bytes1                 int
	Bytes2                 int
	Emote                  int
	Auras                  string
	VisibilityDistanceType int
}

func (CreatureTemplateAddon) TableName() string {
	return "creature_template_addon"
}
