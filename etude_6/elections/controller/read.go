package controller

import (
	"fmt"
	"net/http"

	"github.com/D-BookeR/LivreGo-sources/etude_6/elections/model"
)

// Winner is an HTTP handler to write in the response the winning politician
func Winner(w http.ResponseWriter, r *http.Request) {
	m := model.FromMongo{Server: "localhost", DbName: "elections", PoliticiansCollection: "politicians", VotesCollection: "votes"}

	winner, err := m.Winner()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Fprint(w, winner)
}

// Count writes to the response the number of votes currently in the database
func Count(w http.ResponseWriter, r *http.Request) {
	m := model.FromMongo{Server: "localhost", DbName: "elections", PoliticiansCollection: "politicians", VotesCollection: "votes"}

	votes, err := m.AllVotes()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Fprint(w, len(votes))
}
