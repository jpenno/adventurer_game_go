package Player

type Controlls struct {
	Up    string
	Down  string
	Left  string
	Right string
	Quit  string
}

func NewControlls() Controlls {
	return Controlls{
		Up:    "up",
		Down:  "down",
		Left:  "left",
		Right: "right",
		Quit:  "q",
	}
}