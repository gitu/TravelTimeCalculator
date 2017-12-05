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
	return duration
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
func CalculateAverage(from string, to string, via string) (time.Duration, error) {
	conn := new(ConnectionsResult) // or &Foo{}
	myUrl := fmt.Sprintf("http://transport.opendata.ch/v1/connections?from=%s&to=%s&time=06:00&limit=6",
		template.URLQueryEscaper(from), template.URLQueryEscaper(to))

	if !(via == from || via == to) {
		myUrl += "&via=" + via
	}

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

func Calculate(home string, target string, workPlace string) (time.Duration, error) {
	HT, err := CalculateAverage(home, target, target)
	if err != nil {
		return 0, err
	}
	WT, err := CalculateAverage(workPlace, target, target)
	if err != nil {
		return 0, err
	}
	HW, err := CalculateAverage(home, workPlace, workPlace)
	if err != nil {
		return 0, err
	}

	result := time.Duration(0)
	if HT < WT {
		if HW > WT {
			result = WT - HT
		} else {
			result = HT
		}
	}
	if HT > WT && HT-HW > 0 {
		result = HT - HW
	}

	if result > WT {
		result = WT
	}
	if result > HT {
		result = HT
	}


	fmt.Printf("%-15s; %-15s; %-15s;%8s;%8s;%8s;%8s\n", home, target, workPlace, HT, WT, HW, result)

	return result, nil
}
