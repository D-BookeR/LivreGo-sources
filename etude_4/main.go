package main

import (
	"fmt"

	"github.com/D-BookeR/LivreGo-sources/etude_4/elections/helper"
	"github.com/D-BookeR/LivreGo-sources/etude_4/elections/model"
)

func main() {

	// Creating the model struct
	m := model.FromMongo{Server: "localhost", DbName: "elections", PoliticiansCollection: "politicians", VotesCollection: "votes"}

	// Getting all votes
	allVotes, err := m.AllVotes()
	if err != nil {
		fmt.Println(err)
		return
	}

	//
	r := helper.ComputeRound(allVotes)

	// Delete blanc
	delete(r, 0)

	//
	w, err := r.Winner(&m)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		return
	}

	fmt.Printf("Le gagnant est %s!\n", w)

	fmt.Println(len(allVotes))

	// Getting the winner directly
	p, err := m.Winner()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		return
	}
	fmt.Printf("Le gagnant est %s!\n", p)

}
