package helper

import (
	"errors"
	"fmt"

	"github.com/D-BookeR/LivreGo-sources/etude_2/elections/model"
)

// Round is a count, for each politician, of the number of votes they got
type Round map[int]int

// ComputeRound computes the summary of the round
func ComputeRound(vs model.Votes) Round {
	r := make(Round)
	for _, v := range vs {
		val, exists := r[v.PoliticianID]
		if exists {
			r[v.PoliticianID] = val + 1
			continue
		}
		r[v.PoliticianID] = 1
	}
	return r
}

// Winner finds the winner from a round
func (r *Round) Winner(m model.Reader) (model.Politician, error) {
	var currentMaxScore int
	var secondMaxScore int
	var currentWinner int
	var secondToWinner int

	for p, s := range *r {
		if s >= currentMaxScore {
			secondMaxScore = currentMaxScore
			currentMaxScore = s
			secondToWinner = currentWinner
			currentWinner = p
		}
	}

	if currentMaxScore == 0 {
		return model.Politician{}, errors.New("il ne semble y avoir aucun vote enregistré pour le moment.")
	}

	if currentMaxScore == secondMaxScore {
		currentWinnerPolitician, err := m.PoliticianFromID(currentWinner)
		if err != nil {
			return model.Politician{}, err
		}
		secondToWinnerPolitician, err := m.PoliticianFromID(secondToWinner)
		if err != nil {
			return model.Politician{}, err
		}
		return model.Politician{}, fmt.Errorf("deux candidats sont à égalité ! %s et %s ont tous deux %d votes", currentWinnerPolitician.String(), secondToWinnerPolitician.String(), currentMaxScore)
	}

	currentWinnerPolitician, err := m.PoliticianFromID(currentWinner)
	if err != nil {
		return model.Politician{}, err
	}
	return currentWinnerPolitician, nil
}
