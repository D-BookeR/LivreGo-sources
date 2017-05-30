package helper

import (
	"errors"
	"reflect"
	"testing"

	"github.com/D-BookeR/LivreGo-sources/etude_5/elections/model"
)

/* A test that never fails! */
func TestAlwaysPasses(t *testing.T) {

}

/* Simple test for the ComputeRound() function, made of one test case */
func TestComputeRound(t *testing.T) {
	// Computed
	computedRound := ComputeRound(model.Votes{
		model.Vote{Name: "John Doe", ID: "a", PoliticianID: 1},
		model.Vote{Name: "John Doe", ID: "b", PoliticianID: 0},
		model.Vote{Name: "John Doe", ID: "c", PoliticianID: 1},
		model.Vote{Name: "John Doe", ID: "d", PoliticianID: 2},
		model.Vote{Name: "John Doe", ID: "e", PoliticianID: 2},
		model.Vote{Name: "John Doe", ID: "f", PoliticianID: 2},
	})

	// Expected
	expectedRound := Round{0: 1, 1: 2, 2: 3}

	// Assertions
	if !reflect.DeepEqual(computedRound, expectedRound) {
		t.Errorf("Round is not as expected. Computed: %v. Expected: %v.", computedRound, expectedRound)
	}
}

/* Same test on the ComputeRound() function, but this time with several use cases, using table-driven tests */

// Test cases
var computeRoundsTests = []struct {
	in  model.Votes
	out Round
}{
	{
		model.Votes{
			model.Vote{Name: "John Doe", ID: "a", PoliticianID: 1},
			model.Vote{Name: "John Doe", ID: "b", PoliticianID: 0},
			model.Vote{Name: "John Doe", ID: "c", PoliticianID: 1},
			model.Vote{Name: "John Doe", ID: "d", PoliticianID: 2},
			model.Vote{Name: "John Doe", ID: "e", PoliticianID: 2},
			model.Vote{Name: "John Doe", ID: "f", PoliticianID: 2},
		},
		Round{0: 1, 1: 2, 2: 3},
	},
	{
		model.Votes{
			model.Vote{Name: "John Doe", ID: "a", PoliticianID: 1},
			model.Vote{Name: "John Doe", ID: "b", PoliticianID: 0},
			model.Vote{Name: "John Doe", ID: "c", PoliticianID: 1},
			model.Vote{Name: "John Doe", ID: "d", PoliticianID: 2},
			model.Vote{Name: "John Doe", ID: "e", PoliticianID: 2},
			model.Vote{Name: "John Doe", ID: "f", PoliticianID: 3},
		},
		Round{0: 1, 1: 2, 2: 2, 3: 1},
	},
	{
		model.Votes{},
		Round{},
	},
}

// Test function
func TestComputeRounds(t *testing.T) {
	for _, testCase := range computeRoundsTests {
		computedRound := ComputeRound(testCase.in)

		if !reflect.DeepEqual(computedRound, testCase.out) {
			t.Errorf("Round is not as expected. Computed: %v. Expected: %v.", computedRound, testCase.out)
		}
	}
}

/* Now testing Winner(), with mocking and table-driven test with types that were defined */

// Test cases
var winnerTests = []struct {
	in              Round
	outPoliticianID int
	outErr          error
}{
	{
		in:              Round{1: 3, 2: 2, 3: 1},
		outPoliticianID: 1,
		outErr:          nil,
	},
	{
		in:              Round{},
		outPoliticianID: 0,
		outErr:          errors.New("Il ne semble y avoir aucun vote enregistré pour le moment."),
	},
	{
		in:              Round{1: 2, 2: 2, 3: 1},
		outPoliticianID: 0,
		outErr:          errors.New("Deux candidats sont à égalité ! John Doe, de \"GOP\" et John Doe, de \"GOP\" ont tous deux 2 votes."),
	},
}

// And now, the structure that will mock a model.Reader, whose PoliticianFromID() is the only function that interests us.
// PoliticianFromID() will simply return from the ID a politician with dummy data, but the right ID
type mockedReader struct{}

func (r mockedReader) AllPoliticians() (model.Politicians, error) {
	return nil, nil
}
func (r mockedReader) AllVotes() (model.Votes, error) {
	return nil, nil
}
func (r mockedReader) PoliticianFromID(ID int) (model.Politician, error) {
	return model.Politician{ID: ID, Name: "John Doe", Party: "GOP"}, nil
}

// Test function
func TestWinner(t *testing.T) {
	for _, testCase := range winnerTests {
		computedPolitician, err := testCase.in.Winner(mockedReader{})

		// Comparing the value of the error
		if errorsUnequal(err, testCase.outErr) {
			t.Errorf("Error status unexpected. Computed: %v. Expected: %v.", err, testCase.outErr)
		}

		// Comparing the value of the politician
		if computedPolitician.ID != testCase.outPoliticianID {
			t.Errorf("Unexpected winning politician. Computed: %v. Expected: %v.", computedPolitician.ID, testCase.outPoliticianID)
		}
	}
}

// Utility function to compare errors
func errorsUnequal(e1 error, e2 error) bool {
	return (e1 != nil || e2 != nil) && ((e1 == nil || e2 == nil) || (e1.Error() != e2.Error()))
}
