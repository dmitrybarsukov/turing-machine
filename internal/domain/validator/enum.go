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
	Unordered
	Descending
)

func (o Order) String() string {
	switch o {
	case Ascending:
		return "ascending"
	case Unordered:
		return "unordered"
	case Descending:
		return "descending"
	default:
		return "?"
	}
}

type Parity int

const (
	Even Parity = iota
	Odd
)

func (p Parity) not() Parity {
	return 1 - p
}

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

type Count int

const (
	Zero Count = iota
	One
	Two
	Three
)
