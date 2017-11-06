package main

import (
	"errors"
	"fmt"
	"log"
)

type politician struct {
	Name  string
	ID    int
	Party string
}

type voter struct {
	Name string
	ID   int
}

func (p *politician) String() string {
	return fmt.Sprintf("%s, de \"%s\"", (*p).Name, (*p).Party)
}

func (v *voter) String() string {
	return (*v).Name
}

type votes map[voter]*politician

type round map[politician]int

// Computes the summary of the round
func (v *votes) computeRound() round {
	r := make(round)
	for _, p := range *v {
		val, exists := r[*p]
		if exists {
			r[*p] = val + 1
			continue
		}
		r[*p] = 1
	}
	return r
}

func (r *round) winner() (politician, error) {
	var currentMaxScore int
	var secondMaxScore int
	var currentWinner politician
	var secondToWinner politician

	for p, s := range *r {
		if s >= currentMaxScore {
			secondMaxScore = currentMaxScore
			currentMaxScore = s
			secondToWinner = currentWinner
			currentWinner = p
		}
	}

	if currentMaxScore == 0 {
		return politician{}, errors.New("Il ne semble y avoir aucun vote enregistré pour le moment.")
	}

	if currentMaxScore == secondMaxScore {
		return politician{}, fmt.Errorf("deux candidats sont à égalité ! %s et %s ont tous deux %d votes", currentWinner.String(), secondToWinner.String(), currentMaxScore)
	}

	return currentWinner, nil
}

func main() {

	// Creating data
	blanc := politician{Name: "Vote blanc"}
	rouleau := politician{Name: "Chantal Rouleau", ID: 1, Party: "Ensemble demain"}
	lagace := politician{Name: "Loraine Lagace", ID: 2, Party: "Écologie"}
	dorion := politician{Name: "Jean-Baptiste Dorion", ID: 3, Party: "Par et pour le peuple"}

	v := make(votes)

	v[voter{Name: "Henri Lepic", ID: 0}] = &rouleau
	v[voter{Name: "Mathilde Meilichzon", ID: 1}] = &dorion
	v[voter{Name: "Mathieu Coupe", ID: 2}] = &dorion
	v[voter{Name: "Ramzi Sayal", ID: 3}] = &lagace
	v[voter{Name: "Rober Menard", ID: 4}] = &blanc
	v[voter{Name: "Benjamin Deon", ID: 5}] = &dorion
	v[voter{Name: "Mireille Gaillard", ID: 6}] = &lagace
	v[voter{Name: "Olivier Vaselin", ID: 7}] = &lagace
	v[voter{Name: "Antoine Tarre", ID: 8}] = &dorion

	r := v.computeRound()

	delete(r, blanc)

	w, err := r.winner()
	if err != nil {
		log.Printf("error : %s \n", err)
		return
	}

	log.Printf("Le gagnant est %s!\n", w.String())
}
