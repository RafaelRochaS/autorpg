package item

type ItemImpl struct {
	Name   string
	Level  int
	Type   ItemType
	StrReq int
	IntReq int
	DexReq int
}

func (i ItemImpl) GetLevel() int {
	return i.Level
}

func (i ItemImpl) GetName() string {
	return i.Name
}

func (i ItemImpl) GetType() ItemType {
	return i.Type
}

func (i ItemImpl) GetStrReq() int {
	return i.StrReq
}

func (i ItemImpl) GetIntReq() int {
	return i.IntReq
}

func (i ItemImpl) GetDexReq() int {
	return i.DexReq
}
