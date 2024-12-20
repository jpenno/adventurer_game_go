package Player

import (
	"encoding/json"
	"fmt"
	"os"
)

type Controlls struct {
	Up     string `json:"Up"`
	Down   string `json:"Down"`
	Left   string `json:"Left"`
	Right  string `json:"Right"`
	Quit   string `json:"Quit"`
	Attack string `json:"Attack"`
	Pickup string `json:"Pickup"`
}

func NewControlls() *Controlls {

	fileData, err := os.ReadFile(".config/controlls.json")
	if err != nil {
		fmt.Println(err)

		return &Controlls{
			Up:     "up",
			Down:   "down",
			Left:   "left",
			Right:  "right",
			Quit:   "q",
			Attack: "attack",
			Pickup: "pickup",
		}
	}

	var controlls Controlls
	json.Unmarshal(fileData, &controlls)

	return &controlls
}
