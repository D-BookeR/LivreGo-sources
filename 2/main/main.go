package main

import (
	"errors"
	"fmt"
	"votes/2/elections"
)

type politiciansMap map[int]elections.Politician

type round map[int]int

// Computes the summary of the round
func computeRound(vs elections.Votes) round {
	r := make(round)
	for _, v := range vs {
		val, exists := r[v.PoliticianID]
		if exists {
			r[v.PoliticianID] = val + 1
		} else {
			r[v.PoliticianID] = 1
		}
	}
	return r
}

// Notice: no getWinner
func (r round) winner(m *elections.ModelFiles) (elections.Politician, error) {
	currentMaxScore := 0
	secondMaxScore := 0
	var currentWinner int
	var secondToWinner int

	for p, s := range r {
		if s >= currentMaxScore {
			secondMaxScore = currentMaxScore
			currentMaxScore = s
			secondToWinner = currentWinner
			currentWinner = p
		}
	}

	if currentMaxScore == secondMaxScore {
		currentWinnerPolitician, err := m.PoliticianFromID(currentWinner)
		if err != nil {
			return elections.Politician{}, err
		}
		secondToWinnerPolitician, err := m.PoliticianFromID(secondToWinner)
		if err != nil {
			return elections.Politician{}, err
		}
		errString := fmt.Sprintf("Two candidates are tied! %s and %s both have %d votes", currentWinnerPolitician, secondToWinnerPolitician, currentMaxScore)
		return elections.Politician{}, errors.New(errString)
	}

	currentWinnerPolitician, err := m.PoliticianFromID(currentWinner)
	if err != nil {
		return elections.Politician{}, err
	}
	return currentWinnerPolitician, nil
}

func main() {

	votesFileNames := []string{}
	for i := 0; i < 100; i++ {
		votesFileNames = append(votesFileNames, fmt.Sprintf("votes_%d.json", i+1))
	}

	m := elections.ModelFiles{DirPath: "2_files", PoliticiansFileName: "politicians.json", VotesFileNames: votesFileNames}

	// fmt.Println(m.GenerateAndWriteVotes(100, 10000, 3))

	allVotes, err := m.AllVotes()
	if err != nil {
		fmt.Println(err)
		return
	}

	r := computeRound(allVotes)

	// fmt.Println(r)

	delete(r, 0) // delete blanc

	w, err := r.winner(&m)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		return
	}

	fmt.Printf("The winner is %s!\n", w)

	fmt.Println("******")
	fmt.Println(len(allVotes))
}
