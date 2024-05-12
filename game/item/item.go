package item

type Item interface {
	GetIcon() string
	GetName() string
	GetStats() map[string]uint32
}

type baseItem struct {
	icon  string
	name  string
	stats map[string]uint32
}

func (bi baseItem) GetIcon() string {
	return bi.icon
}

func (bi baseItem) GetName() string {
	return bi.name
}

func (bi baseItem) GetStats() map[string]uint32 {
	return bi.stats
}
