package validator

type Compare int

const (
	Less Compare = iota
	Equal
	More
)

func (r Compare) String() string {
	switch r {
	case Less:
		return "<"
	case Equal:
		return "="
	case More:
		return ">"
	default:
		return "?"
	}
}

type Order int

const (
	Ascending Order = iota
	None
	Descending
)

func (o Order) String() string {
	switch o {
	case Ascending:
		return "asc"
	case None:
		return "none"
	case Descending:
		return "desc"
	default:
		return "?"
	}
}

type Parity int

const (
	Even Parity = iota
	Odd
)

func (p Parity) String() string {
	switch p {
	case Even:
		return "even"
	case Odd:
		return "odd"
	default:
		return "?"
	}
}
