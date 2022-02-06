package handlers

import (
	"golang-fifa-world-cup-web-service/data"
	"net/http"
)

// RootHandler returns an empty body status code
func RootHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNoContent) // send the headers with a 204 response code.
}

// ListWinners returns winners from the list
func ListWinners(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var year = req.URL.Query().Get("year")
	if year == "" {
		var winners, err = data.ListAllJSON()
		if err != nil {
			// handle the error
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
		res.Write(winners)

	} else {
		var filteredWinners, err = data.ListAllByYear(year)
		if err != nil {
			// handle the error
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		res.Write(filteredWinners)
	}

}

// AddNewWinner adds new winner to the list
func AddNewWinner(res http.ResponseWriter, req *http.Request) {

}

// WinnersHandler is the dispatcher for all /winners URL
func WinnersHandler(res http.ResponseWriter, req *http.Request) {

}
