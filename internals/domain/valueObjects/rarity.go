package valueObjects

type Rarity int

const (
	Platinum Rarity = iota + 1
	Gold
	Silver
	Bronze
)
