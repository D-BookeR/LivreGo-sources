package main

import (
	"fmt"
	"votes/etude_2/elections/helper"
	"votes/etude_2/elections/model"
)

func main() {

	votesFileNames := []string{}
	for i := 0; i < 100; i++ {
		votesFileNames = append(votesFileNames, fmt.Sprintf("votes_%d.json", i+1))
	}

	m := model.ModelFiles{DirPath: "2_files", PoliticiansFileName: "politicians.json", VotesFileNames: votesFileNames}

	// fmt.Println(m.GenerateAndWriteVotes(100, 10000, 3))

	allVotes, err := m.AllVotes()
	if err != nil {
		fmt.Println(err)
		return
	}

	r := helper.ComputeRound(allVotes)

	// fmt.Println(r)

	delete(r, 0) // delete blanc

	w, err := r.Winner(&m)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		return
	}

	fmt.Printf("The winner is %s!\n", w)
}
