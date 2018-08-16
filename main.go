package main

import (
	"log"
	"github.com/gitu/TravelTimeCalculator/timeCalculator"
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

	fmt.Printf("%-20s; %-20s; %-20s;%8s;%8s;%8s;%8s;%8s;%8s;%8s;%8s;%8s\n",
		"HOME", "TARGET", "WORKPLACE", "HT", "HW", "MORN", "MORN_R", "TH", "WH", "EVEN", "EVEN_R", "TWOWAY")

	for _, workPlace := range workPlaces {
		for _, home := range homes {
			for _, target := range targets {
				QuickCalc(home, target, workPlace)
			}
		}
	}
}

func QuickCalc(home string, target string, workPlace string) {
	_, _, err := timeCalculator.Calculate(home, target, workPlace)
	if err != nil {
		log.Fatalln(err)
	}
}
