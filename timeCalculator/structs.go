package timeCalculator


type ConnectionsResult struct {
	Connections []struct {
		From struct {
			Station struct {
				ID    string      `json:"id"`
				Name  string      `json:"name"`
				Score interface{} `json:"score"`
				Coordinate struct {
					Type string  `json:"type"`
					X    float64 `json:"x"`
					Y    float64 `json:"y"`
				} `json:"coordinate"`
				Distance interface{} `json:"distance"`
			} `json:"station"`
			Arrival            interface{} `json:"arrival"`
			ArrivalTimestamp   interface{} `json:"arrivalTimestamp"`
			Departure          string      `json:"departure"`
			DepartureTimestamp int         `json:"departureTimestamp"`
			Delay              int         `json:"delay"`
			Platform           string      `json:"platform"`
			Prognosis struct {
				Platform    interface{} `json:"platform"`
				Arrival     interface{} `json:"arrival"`
				Departure   string      `json:"departure"`
				Capacity1St interface{} `json:"capacity1st"`
				Capacity2Nd interface{} `json:"capacity2nd"`
			} `json:"prognosis"`
			RealtimeAvailability interface{} `json:"realtimeAvailability"`
			Location struct {
				ID    string      `json:"id"`
				Name  string      `json:"name"`
				Score interface{} `json:"score"`
				Coordinate struct {
					Type string  `json:"type"`
					X    float64 `json:"x"`
					Y    float64 `json:"y"`
				} `json:"coordinate"`
				Distance interface{} `json:"distance"`
			} `json:"location"`
		} `json:"from"`
		To struct {
			Station struct {
				ID    string      `json:"id"`
				Name  string      `json:"name"`
				Score interface{} `json:"score"`
				Coordinate struct {
					Type string  `json:"type"`
					X    float64 `json:"x"`
					Y    float64 `json:"y"`
				} `json:"coordinate"`
				Distance interface{} `json:"distance"`
			} `json:"station"`
			Arrival            string      `json:"arrival"`
			ArrivalTimestamp   int         `json:"arrivalTimestamp"`
			Departure          interface{} `json:"departure"`
			DepartureTimestamp interface{} `json:"departureTimestamp"`
			Delay              interface{} `json:"delay"`
			Platform           string      `json:"platform"`
			Prognosis struct {
				Platform    interface{} `json:"platform"`
				Arrival     interface{} `json:"arrival"`
				Departure   interface{} `json:"departure"`
				Capacity1St interface{} `json:"capacity1st"`
				Capacity2Nd interface{} `json:"capacity2nd"`
			} `json:"prognosis"`
			RealtimeAvailability interface{} `json:"realtimeAvailability"`
			Location struct {
				ID    string      `json:"id"`
				Name  string      `json:"name"`
				Score interface{} `json:"score"`
				Coordinate struct {
					Type string  `json:"type"`
					X    float64 `json:"x"`
					Y    float64 `json:"y"`
				} `json:"coordinate"`
				Distance interface{} `json:"distance"`
			} `json:"location"`
		} `json:"to"`
		Duration    string      `json:"duration"`
		Transfers   int         `json:"transfers"`
		Service     interface{} `json:"service"`
		Products    []string    `json:"products"`
		Capacity1St interface{} `json:"capacity1st"`
		Capacity2Nd interface{} `json:"capacity2nd"`
		Sections []struct {
			Journey struct {
				Name         string      `json:"name"`
				Category     string      `json:"category"`
				Subcategory  interface{} `json:"subcategory"`
				CategoryCode interface{} `json:"categoryCode"`
				Number       string      `json:"number"`
				Operator     string      `json:"operator"`
				To           string      `json:"to"`
				PassList []struct {
					Station struct {
						ID    string      `json:"id"`
						Name  string      `json:"name"`
						Score interface{} `json:"score"`
						Coordinate struct {
							Type string  `json:"type"`
							X    float64 `json:"x"`
							Y    float64 `json:"y"`
						} `json:"coordinate"`
						Distance interface{} `json:"distance"`
					} `json:"station"`
					Arrival            interface{} `json:"arrival"`
					ArrivalTimestamp   interface{} `json:"arrivalTimestamp"`
					Departure          string      `json:"departure"`
					DepartureTimestamp int         `json:"departureTimestamp"`
					Delay              int         `json:"delay"`
					Platform           string      `json:"platform"`
					Prognosis struct {
						Platform    interface{} `json:"platform"`
						Arrival     interface{} `json:"arrival"`
						Departure   string      `json:"departure"`
						Capacity1St interface{} `json:"capacity1st"`
						Capacity2Nd interface{} `json:"capacity2nd"`
					} `json:"prognosis"`
					RealtimeAvailability interface{} `json:"realtimeAvailability"`
					Location struct {
						ID    string      `json:"id"`
						Name  string      `json:"name"`
						Score interface{} `json:"score"`
						Coordinate struct {
							Type string  `json:"type"`
							X    float64 `json:"x"`
							Y    float64 `json:"y"`
						} `json:"coordinate"`
						Distance interface{} `json:"distance"`
					} `json:"location"`
				} `json:"passList"`
				Capacity1St interface{} `json:"capacity1st"`
				Capacity2Nd interface{} `json:"capacity2nd"`
			} `json:"journey"`
			Walk interface{} `json:"walk"`
			Departure struct {
				Station struct {
					ID    string      `json:"id"`
					Name  string      `json:"name"`
					Score interface{} `json:"score"`
					Coordinate struct {
						Type string  `json:"type"`
						X    float64 `json:"x"`
						Y    float64 `json:"y"`
					} `json:"coordinate"`
					Distance interface{} `json:"distance"`
				} `json:"station"`
				Arrival            interface{} `json:"arrival"`
				ArrivalTimestamp   interface{} `json:"arrivalTimestamp"`
				Departure          string      `json:"departure"`
				DepartureTimestamp int         `json:"departureTimestamp"`
				Delay              int         `json:"delay"`
				Platform           string      `json:"platform"`
				Prognosis struct {
					Platform    interface{} `json:"platform"`
					Arrival     interface{} `json:"arrival"`
					Departure   string      `json:"departure"`
					Capacity1St interface{} `json:"capacity1st"`
					Capacity2Nd interface{} `json:"capacity2nd"`
				} `json:"prognosis"`
				RealtimeAvailability interface{} `json:"realtimeAvailability"`
				Location struct {
					ID    string      `json:"id"`
					Name  string      `json:"name"`
					Score interface{} `json:"score"`
					Coordinate struct {
						Type string  `json:"type"`
						X    float64 `json:"x"`
						Y    float64 `json:"y"`
					} `json:"coordinate"`
					Distance interface{} `json:"distance"`
				} `json:"location"`
			} `json:"departure"`
			Arrival struct {
				Station struct {
					ID    string      `json:"id"`
					Name  string      `json:"name"`
					Score interface{} `json:"score"`
					Coordinate struct {
						Type string  `json:"type"`
						X    float64 `json:"x"`
						Y    float64 `json:"y"`
					} `json:"coordinate"`
					Distance interface{} `json:"distance"`
				} `json:"station"`
				Arrival            string      `json:"arrival"`
				ArrivalTimestamp   int         `json:"arrivalTimestamp"`
				Departure          interface{} `json:"departure"`
				DepartureTimestamp interface{} `json:"departureTimestamp"`
				Delay              int         `json:"delay"`
				Platform           string      `json:"platform"`
				Prognosis struct {
					Platform    interface{} `json:"platform"`
					Arrival     string      `json:"arrival"`
					Departure   interface{} `json:"departure"`
					Capacity1St interface{} `json:"capacity1st"`
					Capacity2Nd interface{} `json:"capacity2nd"`
				} `json:"prognosis"`
				RealtimeAvailability interface{} `json:"realtimeAvailability"`
				Location struct {
					ID    string      `json:"id"`
					Name  string      `json:"name"`
					Score interface{} `json:"score"`
					Coordinate struct {
						Type string  `json:"type"`
						X    float64 `json:"x"`
						Y    float64 `json:"y"`
					} `json:"coordinate"`
					Distance interface{} `json:"distance"`
				} `json:"location"`
			} `json:"arrival"`
		} `json:"sections"`
	} `json:"connections"`
	From struct {
		ID    string      `json:"id"`
		Name  string      `json:"name"`
		Score interface{} `json:"score"`
		Coordinate struct {
			Type string  `json:"type"`
			X    float64 `json:"x"`
			Y    float64 `json:"y"`
		} `json:"coordinate"`
		Distance interface{} `json:"distance"`
	} `json:"from"`
	To struct {
		ID    string      `json:"id"`
		Name  string      `json:"name"`
		Score interface{} `json:"score"`
		Coordinate struct {
			Type string  `json:"type"`
			X    float64 `json:"x"`
			Y    float64 `json:"y"`
		} `json:"coordinate"`
		Distance interface{} `json:"distance"`
	} `json:"to"`
	Stations struct {
		From []struct {
			ID    string      `json:"id"`
			Name  string      `json:"name"`
			Score interface{} `json:"score"`
			Coordinate struct {
				Type string  `json:"type"`
				X    float64 `json:"x"`
				Y    float64 `json:"y"`
			} `json:"coordinate"`
			Distance interface{} `json:"distance"`
		} `json:"from"`
		To []struct {
			ID    string      `json:"id"`
			Name  string      `json:"name"`
			Score interface{} `json:"score"`
			Coordinate struct {
				Type string  `json:"type"`
				X    float64 `json:"x"`
				Y    float64 `json:"y"`
			} `json:"coordinate"`
			Distance interface{} `json:"distance"`
		} `json:"to"`
	} `json:"stations"`
}
