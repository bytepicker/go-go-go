package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var dayNames = [...]string{"Sweetmorn", "Boomtime", "Pungenday", "Prickle-Prickle", "Setting Orange"}
var seasonNames = [...]string{"Chaos", "Discord", "Confusion", "Bureaucracy", "The Aftermath"}
var holydaysFive = [...]string{"Mungday", "Mojoday", "Syaday", "Zaraday", "Maladay"}
var holydaysFifty = [...]string{"Chaoflux", "Discoflux", "Confuflux", "Bureflux", "Afflux"}

const (
	daysInYear         = 365
	initialYearPadding = 1166
	hailEris           = daysInYear / len(seasonNames)
)

type discordianDate struct {
	Day     string `json:"day"`
	Number  int    `json:"number"`
	Season  string `json:"season"`
	Year    int    `json:"year"`
	Holyday string `json:"holyday"`
}

func getCurrentDate(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	result := convert(now)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func convert(now time.Time) discordianDate {
	ddate := discordianDate{}

	ddate.Year = now.Year() + initialYearPadding
	ddate.Season = seasonNames[(now.YearDay()-1)/hailEris]
	ddate.Day = dayNames[(now.YearDay()-1)%hailEris%5]
	ddate.Number = now.YearDay() % 73

	if ddate.Number == 5 {
		ddate.Holyday = holydaysFive[now.YearDay()%5]
	}

	if ddate.Number == 50 {
		ddate.Holyday = holydaysFifty[now.YearDay()%5]
	}

	if int(now.Month()) == 2 && now.Day() == 29 {
		ddate.Holyday = "St. Tib's Day"
	}

	return ddate
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", getCurrentDate).Methods("GET")

	port := 8080
	fmt.Printf("Server is running on :%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
