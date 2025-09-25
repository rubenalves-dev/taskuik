package culture

type Culture map[Locale]string

type Locale int

const (
	PT Locale = iota
	EN
)
