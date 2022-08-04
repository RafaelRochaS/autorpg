package item

type ArmorImpl struct {
	Item
	Defense int
	Weight  int
}

func (a ArmorImpl) GetDefense() int {
	return a.Defense
}

func (a ArmorImpl) GetWeight() int {
	return a.Weight
}
