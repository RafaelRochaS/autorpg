package item

type ArmorImpl struct {
	Item
	Defense int
}

func (a ArmorImpl) GetDefense() int {
	return a.Defense
}
