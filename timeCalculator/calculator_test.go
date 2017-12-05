package timeCalculator_test

import (
	"testing"
	"github.com/gitu/travelTimeCalculator/timeCalculator"
	"time"
)

func calculate(t *testing.T, home string, target string, workPlace string, expectedDuration string) {
	duration, err := timeCalculator.Calculate(home, target, workPlace)
	if err != nil {
		t.Log("Error with query: ", err)
		t.Fail()
	}
	expected, err := time.ParseDuration(expectedDuration)
	if err != nil {
		t.Fatal("Error parsing your duration", expectedDuration, err)
		t.FailNow()
	}
	if duration != expected {
		t.Log("Duration for ", home, " to ", target, " workplace ", workPlace, " expected: ", expected, " actual ", duration)
		t.Fail()
	}
}

func TestCalculate(t *testing.T) {
	calculate(t, "Rotkreuz", "Zürich", "Zug", "22m0s")
	calculate(t, "Rotkreuz", "Zürich", "Zürich", "0m")
	calculate(t, "Rotkreuz", "Luzern", "Zürich", "15m0s")
	calculate(t, "Rotkreuz", "Bendern", "Zug", "2h9m0s")
	calculate(t, "Zürich", "Zürich", "Zug", "0m")
	calculate(t, "Zürich", "St Gallen", "Zug", "1h2m0s")
	calculate(t, "Bern", "Zürich", "Zug", "0m")
	calculate(t, "Bern", "St Gallen", "Zug", "35m0s")
	calculate(t, "Bern", "Aarau", "Zug", "17m0s")
	calculate(t, "Bern", "Aarau", "Zürich", "0s")
	calculate(t, "Thalwil", "Zürich", "Zug", "9m0s")
	calculate(t, "Reinach AG", "Luzern", "Zug", "0m")
	calculate(t, "Reinach AG", "Liestal", "Zug", "13m0s")
	calculate(t, "Aarau", "Luzern", "Zug", "0m")
	calculate(t, "Aarau", "Liestal", "Zug", "25m0s")
	calculate(t, "Aarau", "St. Gallen", "Zug", "46m0s")
	calculate(t, "Wintherthur", "St. Gallen", "Zug", "36m")
	calculate(t, "Zug", "Zug", "Zürich", "0s")
	calculate(t, "Zug", "Zug", "Zug", "0s")
	calculate(t, "Zürich", "Zug", "Zürich", "21m0s")
}

func round(t *testing.T, in string, expected string) {
    di, err := time.ParseDuration(in)
	if err != nil {
		t.Log("Error with query: ", err)
		t.Fail()
	}
    de, err := time.ParseDuration(expected)
    if err != nil {
		t.Log("Error with query: ", err)
		t.Fail()
	}
	if de != timeCalculator.Round(di) {
		t.Log(di, " was not rounded correctly ",timeCalculator.Round(di), " it should be ", de)
		t.Fail()
	}
}

func TestRound(t *testing.T) {
	round(t, "13m", "15m")
	round(t, "12m10s", "0m")
	round(t, "33m", "30m")
	round(t, "38m", "30m")
	round(t, "1h33m", "1h30m")
	round(t, "1h33m", "1h30m")
	round(t, "1h30m", "1h30m")
}
