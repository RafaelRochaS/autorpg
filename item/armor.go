package item

import "fmt"

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

func (a ArmorImpl) String() string {
	return fmt.Sprintf(`%s
Defense: %v
Weight: %d
Str Req: %d
Dex Req: %d
Int Req: %d
`, a.GetName(), a.GetDefense(), a.GetWeight(), a.GetStrReq(), a.GetDexReq(), a.GetIntReq())
}
