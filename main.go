package main

import (
	"log"
	"github.com/gitu/travelTimeCalculator/timeCalculator"
	"fmt"
)

func main() {
	workPlaces := []string{"Zug", "Zürich",}
	homes := []string{
		"Zürich",
	}
	targets := []string{
		"Zürich",
	}

	fmt.Printf("%s;%s;%s;%s;%s;%s;%s\n", "HOME", "TARGET", "WORKPLACE", "HT", "WT", "HW", "RES")

	for _, workPlace := range workPlaces {
		for _, home := range homes {
			for _, target := range targets {
				QuickCalc(home, target, workPlace)
			}
		}
	}
}

func QuickCalc(home string, target string, workPlace string) {
	_, err := timeCalculator.Calculate(home, target, workPlace)
	if err != nil {
		log.Fatalln(err)
	}
}
