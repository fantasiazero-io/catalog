package valueObjects

type Description struct {
	Short string
	Full  string
}

func NewDescription(short string, full string) Description {
	return Description{
		Short: short,
		Full:  full,
	}
}

func (d Description) EqualTo(other Description) bool {
	return d.Short == other.Short && d.Full == other.Full
}
