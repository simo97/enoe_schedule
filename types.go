package main

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
