package elections

import (
	"errors"
	"fmt"
	model "votes/3/elections/model"
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
		} else {
			r[v.PoliticianID] = 1
		}
	}
	return r
}

// Winner finds the winner from a round
func (r Round) Winner(m model.Reader) (model.Politician, error) {
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

	if currentMaxScore == 0 {
		return model.Politician{}, errors.New("No vote seems to have been registered yet.")
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
		errString := fmt.Sprintf("Two candidates are tied! %s and %s both have %d votes", currentWinnerPolitician, secondToWinnerPolitician, currentMaxScore)
		return model.Politician{}, errors.New(errString)
	}

	currentWinnerPolitician, err := m.PoliticianFromID(currentWinner)
	if err != nil {
		return model.Politician{}, err
	}
	return currentWinnerPolitician, nil
}
