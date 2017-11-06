package controller

import (
	"net/http"
	"strconv"

	"github.com/D-BookeR/LivreGo-sources/etude_6/elections/model"
)

// RegisterVote , provided all params "name", "id" and "politician_id" are provided, adds that vote to the database
func RegisterVote(w http.ResponseWriter, r *http.Request) {

	// Preparing the params and making sure they're there
	name := r.URL.Query().Get("name")
	id := r.URL.Query().Get("id")
	politicianID := r.URL.Query().Get("politician_id")
	if name == "" || id == "" || politicianID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Not all 3 expected parameters, name, id, and politician_id, were present"))
		return
	}

	// Turning the string param politician_id into an int, complaining if impossible
	politicianIDInt, err := strconv.Atoi(politicianID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("The politician_id parameter is not a valid number - "))
		w.Write([]byte(err.Error()))
		return
	}

	// Saving vote into the MongoDB database
	m := model.FromMongo{Server: "localhost", DbName: "elections", PoliticiansCollection: "politicians", VotesCollection: "votes"}
	vote := model.Vote{Name: name, ID: id, PoliticianID: politicianIDInt}
	if err = m.SaveVote(vote); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Success! \o/
	w.Write([]byte("Vote was successfully saved"))
}
