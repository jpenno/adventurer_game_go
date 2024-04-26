package Color

type Color int8

const (
	Red Color = iota
	Green
	Yellow
	Blue
	Mangenta
	Cyan
	White
	Defalut
)

func (c Color) String() string {
	switch c {
	case Red:
		return "\033[1;31m"
	case Green:
		return "\033[1;32m"
	case Yellow:
		return "\033[1;33m"
	case Blue:
		return "\033[1;34m"
	case Mangenta:
		return "\033[1;35m"
	case Cyan:
		return "\033[1;36m"
	case White:
		return "\033[1;37m"
	case Defalut:
		return "\033[1;39m"
	}
	return "\033[1;39m"
}
