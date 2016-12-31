package main

import (
	"encoding/json"
	"fmt"
)

type Card struct {
	Uuid    string `json:"uuid"`
	NfcUUID string `json:"nfcUUID"`
}

func main() {
	bBool, _ := json.Marshal(true)
	fmt.Println(string(bBool))

	c := &Card{
		Uuid:    "20",
		NfcUUID: "weesddsds"}
	bCard, _ := json.Marshal(c)

	fmt.Println(string(bCard))

	s := `{"uuid": "1", "nfcUUID": "ewwee1"}`

	var newCard Card
	json.Unmarshal([]byte(s), &newCard)

	fmt.Println(newCard)

}
