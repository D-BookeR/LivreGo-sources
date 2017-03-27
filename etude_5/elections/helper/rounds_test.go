package helper

import (
	"errors"
	"reflect"
	"testing"

	"github.com/D-BookeR/LivreGo-sources/etude_5/elections/model"
)

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
		t.Errorf("Round is not as computed. Computed: %v. Expected: %v.", computedRound, expectedRound)
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
			t.Errorf("Round is not as computed. Computed: %v. Expected: %v.", computedRound, testCase.out)
		}
	}
}

/* Now testing Winner(), with mocking and table-driven test with types that were defined */

// Creating useful structures for the test cases
type winnerTestOut struct {
	politicianID int
	err          error
}
type winnerTest struct {
	in  Round
	out winnerTestOut
}

// Test cases
var winnerTests = []winnerTest{
	winnerTest{
		in: Round{1: 3, 2: 2, 3: 1},
		out: winnerTestOut{
			politicianID: 1,
			err:          nil,
		},
	},
	winnerTest{
		in: Round{},
		out: winnerTestOut{
			politicianID: 0,
			err:          errors.New("No vote seems to have been registered yet."),
		},
	},
	winnerTest{
		in: Round{1: 2, 2: 2, 3: 1},
		out: winnerTestOut{
			politicianID: 0,
			err:          errors.New("Two candidates are tied! John Doe, of \"GOP\" and John Doe, of \"GOP\" both have 2 votes"),
		},
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
		if errorsUnequal(err, testCase.out.err) {
			t.Errorf("Error status unexpected. Computed: %v. Expected: %v.", err, testCase.out.err)
		}

		// Comparing the value of the politician
		if computedPolitician.ID != testCase.out.politicianID {
			t.Errorf("Unexpected winning politician. Computed: %v. Expected: %v.", computedPolitician.ID, testCase.out.politicianID)
		}
	}
}

// Utility function to compare errors
func errorsUnequal(e1 error, e2 error) bool {
	return (e1 != nil || e2 != nil) && ((e1 == nil || e2 == nil) || (e1.Error() != e2.Error()))
}