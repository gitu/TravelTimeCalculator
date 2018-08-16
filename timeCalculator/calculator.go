package timeCalculator

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
	"crypto/sha1"
	"os"
	"io"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func Round(duration time.Duration) time.Duration {
	quarters := (int64(duration/time.Minute) + 14) / 15
	return time.Minute * time.Duration(quarters*15)
}

func getJson(url string, target interface{}) error {

	sum := sha1.Sum([]byte(url))

	fileName := fmt.Sprintf("store-%x.json", sum)

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		r, err := myClient.Get(url)
		if err != nil {
			return err
		}
		defer r.Body.Close()

		outFile, err := os.Create(fileName)
		defer outFile.Close()
		_, err = io.Copy(outFile, r.Body)

		if err != nil {
			return err
		}

	}

	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewDecoder(f).Decode(target)
}

func ParseDuration(s string) (time.Duration, error) {
	return time.ParseDuration(s[3:5] + "h" + s[6:8] + "m")
}
func CalculateAverage(from string, to string, departure string) (time.Duration, error) {
	conn := new(ConnectionsResult) // or &Foo{}
	myUrl := fmt.Sprintf("http://transport.opendata.ch/v1/connections?from=%s&to=%s&time=%s&limit=8",
		template.URLQueryEscaper(from), template.URLQueryEscaper(to), departure)

	err := getJson(myUrl, conn)
	if err != nil {
		log.Printf("error while reading result from url (%s) %s", myUrl, err)
		return 0, err
	}

	//total := 0
	min := 0
	for i := range conn.Connections {
		c := conn.Connections[i]
		duration, err := ParseDuration(c.Duration)
		if err != nil {
			log.Printf("error while parsing Duration (%s) %s", c.Duration, err)
			return 0, err
		}
		//fmt.Printf("%s - %s - %s\n", c.From.Departure, c.Duration, duration)
		if i == 0 {
			min = int(duration / time.Second)
		} else if min > int(duration/time.Second) {
			min = int(duration / time.Second)
		}
		//total += int(duration / time.Second)
	}
	//if len(conn.Connections) > 0 {
	//	total /= len(conn.Connections)
	//}

	//averageDuration := time.Duration(total) * time.Second
	averageDuration := time.Duration(min) * time.Second
	//fmt.Printf("Average %s\n", averageDuration)

	return averageDuration, nil
}

func Calculate(home, target, workPlace string) (time.Duration, time.Duration, error) {
	morningHT, err := CalculateAverage(home, target, "06:30")
	if err != nil {
		return 0, 0, err
	}
	morningHW, err := CalculateAverage(home, workPlace, "06:30")
	if err != nil {
		return 0, 0, err
	}

	morningResult := morningHT - morningHW
	if morningResult < 0 {
		morningResult = 0
	}

	eveningTH, err := CalculateAverage(target, home, "17:00")
	if err != nil {
		return 0, 0, err
	}
	eveningWH, err := CalculateAverage(workPlace, home, "17:00")
	if err != nil {
		return 0, 0, err
	}

	eveningResult := eveningTH - eveningWH
	if eveningResult < 0 {
		eveningResult = 0
	}

	fmt.Printf("%-20s; %-20s; %-20s;%8s;%8s;%8s;%8s;%8s;%8s;%8s;%8s;%8s\n", home, target, workPlace, morningHT, morningHW, morningResult, Round(morningResult), eveningTH, eveningWH, eveningResult, Round(eveningResult), Round(morningResult+eveningResult))

	return morningResult, eveningResult, nil
}
