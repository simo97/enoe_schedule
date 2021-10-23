package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gocolly/colly"
)

type ObservationType struct {
	Observations     string `json:observations`
	Prog_date        string `json:prog_date`
	Prog_heure_debut string `json:prog_heure_debut`
	Prog_heure_fin   string `json:prog_heure_fin`
	Region           string `json:region`
	Ville            string `json:ville`
	Quartier         string `json:quartier`
}

type ENEOReponse struct {
	Status int               `json:status`
	Data   []ObservationType `json:data`
}

func _getCity(id string) []ObservationType {

	endpoint := "https://alert.eneo.cm/ajaxOutage.php"
	data := url.Values{
		"region": {id},
	}

	resp, err := http.PostForm(endpoint, data)

	if err != nil {
		log.Fatal(err)
	}

	var res ENEOReponse
	json.NewDecoder(resp.Body).Decode(&res)

	return res.Data

}

func getCitySchdule(regionsChannel chan string, cities *[]ObservationType) {
	for {
		select {
		case id := <-regionsChannel:
			*cities = append(*cities, _getCity(id)...)
		}
	}
}

func main() {
	a := make(chan string)

	c := colly.NewCollector()
	t := time.Now()
	fName := "scheduled_interuptions." + t.Format("2006-01-02 15:04:05") + ".json"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()

	cities := make([]ObservationType, 0)

	defer func() {
		res, err := json.Marshal(cities)
		if err != nil {
			log.Fatal(err)
		}
		file.Write(res)
	}()

	go getCitySchdule(a, &cities)

	c.OnHTML("option", func(e *colly.HTMLElement) {
		id := e.Attr("value")
		if id != "" {
			a <- id
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)

	})

	c.Visit("https://alert.eneo.cm/?header=no")

}
