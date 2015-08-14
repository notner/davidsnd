package main

import (
		"net/http"
		"fmt"
        "github.com/julienschmidt/httprouter"
        )

type Dice struct {
	name string
	value int
	additionalValue int
}

var (
	availableDice = make(map[string]*Dice) 
)

func initaliseDice() {
	availableDice["D20"] = &Dice{"d20", 20, 0}

}

func GetDice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello")

}

// Roll will return a randomly generater number 
// 1 and Dice.value
func (c *Dice) Roll() (int) {
	return 0

}