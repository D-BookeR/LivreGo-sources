package main

import (
	"fmt"
	"os/user"

	"github.com/D-BookeR/LivreGo-sources/etude_2/elections/helper"
	"github.com/D-BookeR/LivreGo-sources/etude_2/elections/model"
)

func main() {

	// Building a slice of the name of each vote data file, so that we can pass it in the model
	votesFileNames := []string{}
	for i := 0; i < 100; i++ {
		votesFileNames = append(votesFileNames, fmt.Sprintf("votes_%d.json", i+1))
	}

	// A cross-platform way to get the user's home directory, so that we can pass it in the model
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Creating the model struct
	m := model.FromFiles{DirPath: usr.HomeDir + "/LivreGo-data", PoliticiansFileName: "politicians.json", VotesFileNames: votesFileNames}

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
}
