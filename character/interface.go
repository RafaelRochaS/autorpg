package character

import "autorpg/item"

type Character interface {
	Create()
	SetName(string)
	SetStats()
	SetClass(Class)
	AttachItem(item.Item)
	RemoveItem(item.Item)
	CheckLevelItem(item.Item) bool
	CheckStatRequirementsItem(item.Item) bool
}
